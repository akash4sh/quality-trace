# Running Qualitytrace With Honeycomb

<!-->:::note
[Check out the source code on GitHub here.](https://github.com/kubeshop/tracetest/tree/main/examples/tracetest-honeycomb)
:::
-->
[Qualitytrace]<!--(https://tracetest.io/)--> is a testing tool based on [OpenTelemetry](https://opentelemetry.io/) that allows you to test your distributed application. It allows you to use data from distributed traces generated by OpenTelemetry to validate and assert if your application has the desired behavior defined by your test definitions.

[Honeycomb](https://honeycomb.io/) is an observability solution that shows you the patterns and outliers of how users experience your code in complex and unpredictable environments.

## Prerequisites

You will need [Docker](https://docs.docker.com/get-docker/) and [Docker Compose](https://docs.docker.com/compose/install/) installed on your machine to run this sample app! Additionally, you will need a Honeycomb account and api key. Sign up to use Honeycomb [here](https://ui.honeycomb.io/signup).

## Project Structure

The project is built with Docker Compose. It contains two distinct `docker-compose.yaml` files.

### 1. Node.js App

The `docker-compose.yaml` file and `Dockerfile` in the root directory are for the Node.js app.

### 2. Qualitytrace

The `docker-compose.yaml` file, `collector.config.yaml`, `quality-trace-provision.yaml`, and `quality-trace-config.yaml` in the root directory are for the setting up Qualitytrace and the OpenTelemetry Collector.

The `root` directory is self-contained and will run all the prerequisites for enabling OpenTelemetry traces and trace-based testing with Qualitytrace, as well as routing all traces the Node.js App generates to Honeycomb.

### Docker Compose Network

All `services` in the `docker-compose.yaml` are on the same network and will be reachable by hostname from within other services. E.g. `quality-trace:4317` in the `collector.config.yaml` will map to the `quality-trace` service, where the port `4317` is the port where Qualitytrace accepts traces.

## Node.js App

The Node.js app is a simple Express app contained in the `app.js` file.

The OpenTelemetry tracing is contained in the `tracing.otel.grpc.js` or `tracing.otel.http.js` files.
Traces will be sent to the OpenTelemetry Collector.

Here's the content of the `tracing.otel.grpc.js` file:

```js
const opentelemetry = require("@opentelemetry/sdk-node");
const {
  getNodeAutoInstrumentations,
} = require("@opentelemetry/auto-instrumentations-node");
const {
  OTLPTraceExporter,
} = require("@opentelemetry/exporter-trace-otlp-grpc");

const sdk = new opentelemetry.NodeSDK({
  traceExporter: new OTLPTraceExporter({ url: "http://otel-collector:4317" }),
  instrumentations: [getNodeAutoInstrumentations()],
});
sdk.start();
```

Depending on which of these you choose, traces will be sent to either the `grpc` or `http` endpoint.

The hostnames and ports for these are:

- GRPC: `http://otel-collector:4317`
- HTTP: `http://otel-collector:4318/v1/traces`

Enabling the tracer is done by preloading the trace file.

```bash
node -r ./tracing.otel.grpc.js app.js
```

In the `package.json` you will see two npm scripts for running the respective tracers alongside the `app.js`.

```json
"scripts": {
  "with-grpc-tracer":"node -r ./tracing.otel.grpc.js app.js",
  "with-http-tracer":"node -r ./tracing.otel.http.js app.js"
},
```

To start the server, run this command:

```bash
npm run with-grpc-tracer
# or
npm run with-http-tracer
```

As you can see the `Dockerfile` uses the command above.

```Dockerfile
FROM node:slim
WORKDIR /usr/src/app
COPY package*.json ./
RUN npm install
COPY . .
EXPOSE 8080
CMD [ "npm", "run", "with-grpc-tracer" ]
```

And, the `docker-compose.yaml` contains just one service for the Node.js app.

```yaml
version: "3"
services:
  app:
    image: quick-start-nodejs
    build: .
    ports:
      - "8080:8080"
```

To start it, run this command:

```bash
docker compose build # optional if you haven't already built the image
docker compose up
```

This will start the Node.js app. But, you're not sending the traces anywhere.

Let's fix this by configuring Qualitytrace and OpenTelemetry Collector.

## Qualitytrace

The `docker-compose.yaml` in the `quality-trace` directory is configured with three services.

- **Postgres** - Postgres is a prerequisite for Qualitytrace to work. It stores trace data when running the trace-based tests.
- [**OpenTelemetry Collector**](https://opentelemetry.io/docs/collector/) - A vendor-agnostic implementation of how to receive, process and export telemetry data.
- [**Qualitytrace**]<!--(https://tracetest.io/)--> - Trace-based testing that generates end-to-end tests automatically from traces.

```yaml
version: "3.2"
services:
  quality-trace:
    restart: unless-stopped
    image: intelops/quality-trace:${TAG:-latest}
    platform: linux/amd64
    ports:
      - 11633:11633
    extra_hosts:
      - "host.docker.internal:host-gateway"
    volumes:
      - type: bind
        source: ./quality-trace/quality-trace-config.yaml
        target: /app/quality-trace.yaml
      - type: bind
        source: ./quality-trace/quality-trace-provision.yaml
        target: /app/provision.yaml
    command: --provisioning-file /app/provision.yaml
    healthcheck:
      test: ["CMD", "wget", "--spider", "localhost:11633"]
      interval: 1s
      timeout: 3s
      retries: 60
    depends_on:
      postgres:
        condition: service_healthy
      otel-collector:
        condition: service_started
    environment:
      QUALITYTRACE_DEV: ${QUALITYTRACE_DEV}

  postgres:
    image: postgres:14
    environment:
      POSTGRES_PASSWORD: postgres
      POSTGRES_USER: postgres
    healthcheck:
      test: pg_isready -U "$$POSTGRES_USER" -d "$$POSTGRES_DB"
      interval: 1s
      timeout: 5s
      retries: 60

  otel-collector:
    image: otel/opentelemetry-collector-contrib:0.68.0
    restart: unless-stopped
    command:
      - "--config"
      - "/otel-local-config.yaml"
    volumes:
      - ./quality-trace/collector.config.yaml:/otel-local-config.yaml

```

Qualitytrace depends on both Postgres and the OpenTelemetry Collector. Both Qualitytrace and the OpenTelemetry Collector require config files to be loaded via a volume. The volumes are mapped from the root directory into the `quality-trace` directory and the respective config files.

To start both the Node.js App and Qualitytrace, run this command:

```bash
docker-compose up # add --build if the images are not built already
```

The `quality-trace-config.yaml` file contains the basic setup of connecting Qualitytrace to the Postgres instance.

```yaml
# quality-trace-config.yaml

---
postgres:
  host: postgres
  user: postgres
  password: postgres
  port: 5432
  dbname: postgres
  params: sslmode=disable

```

The `quality-trace-provision.yaml` file contains the data store setup. The data store is set to `Honeycomb` meaning the traces will be received by Qualitytrace OTLP API and stored in Qualitytrace itself.

```yaml
# quality-trace-provision.yaml
---
type: DataStore
spec:
  name: Honeycomb
  type: honeycomb
```

**How to Send Traces to Qualitytrace and Honeycomb**

The `collector.config.yaml` explains that. It receives traces via either `grpc` or `http`. Then, exports them to Qualitytrace's OTLP endpoint `quality-trace:4317` in one pipeline, and to Honeycomb in another.

Make sure to add your Honeycomb access token in the headers of the `otlp/ls` exporter.

```yaml
receivers:
  otlp:
    protocols:
      grpc:
      http:

processors:
  batch:
    timeout: 100ms

exporters:
  logging:
    logLevel: debug
  # OTLP for Qualitytrace
  otlp/quality-trace:
    endpoint: quality-trace:4317 # Send traces to Qualitytrace. Read more in docs here:  https://docs.tracetest.io/configuration/connecting-to-data-stores/opentelemetry-collector
    tls:
      insecure: true
  # OTLP for Honeycomb
  otlp/honeycomb:
    endpoint: "api.honeycomb.io:443"
    headers:
      "x-honeycomb-team": <HONEYCOMB_API_KEY>
      # Read more in docs here: https://docs.honeycomb.io/getting-data-in/otel-collector/

service:
  pipelines:
    traces/quality-trace:
      receivers: [otlp]
      processors: [batch]
      exporters: [otlp/quality-trace]
    traces/honeycomb:
      receivers: [otlp]
      processors: [batch]
      exporters: [logging, otlp/honeycomb]
```

## Run Both the Node.js App and Qualitytrace

To start both the Node.js App and Qualitytrace, run this command:

```bash
docker-compose up # add --build if the images are not built already
```

This will start your Qualitytrace instance on `http://localhost:11633/`.

Open the URL and start creating tests! Make sure to use the `http://app:8080/` URL in your test creation, because your Node.js app and Qualitytrace are in the same network.

## Run Qualitytrace Tests with the Qualitytrace CLI

First, [install the CLI]<!--(https://docs.tracetest.io/getting-started/installation#install-the-tracetest-cli).-->
Then, configure the CLI:

```bash
quality-trace configure --endpoint http://localhost:11633
```

Once configured, you can run a test against the Qualitytrace instance via the terminal.

Check out the `test-api.yaml` file.

```yaml
# test-api.yaml

type: Test
spec:
  id: W656Q0c4g
  name: http://app:8080
  description: akadlkasjdf
  trigger:
    type: http
    httpRequest:
      url: http://app:8080
      method: GET
      headers:
        - key: Content-Type
          value: application/json
  specs:
    - selector: span[quality-trace.span.type="http" name="GET /" http.target="/" http.method="GET"]
      assertions:
        - attr:http.status_code  =  200
        - attr:quality-trace.span.duration  <  500ms

```

This file defines a test the same way you would through the Web UI.

To run the test, run this command in the terminal:

```bash
quality-trace run test -f ./test-api.yaml
```

```bash
✔ http://app:8080 (http://localhost:11633/test/W656Q0c4g/run/2/test)
	✔ span[quality-trace.span.type="http" name="GET /" http.target="/" http.method="GET"]
```

## View Trace Spans Over Time in Honeycomb

To access a historical overview of all the trace spans the Node.js App generates, jump over to your Honeycomb account.

![Honeycomb trace overview](https://res.cloudinary.com/djwdcmwdz/image/upload/v1683042900/Blogposts/Docs/honeycomb_trace_kbjdl4.png)

You can also drill down into a particular trace.

![Honeycomb trace drilldown](https://res.cloudinary.com/djwdcmwdz/image/upload/v1683042900/Blogposts/Docs/honeycome_dashboard_gyisdg.png)

With Honeycomb and Qualitytrace, you get the best of both worlds. You can run trace-based tests and automate running E2E and integration tests against real trace data. And, use Honeycomb to get a historical overview of all traces your distributed application generates.

## Learn More

Feel free to check out our [examples in GitHub](https://github.com/intelops/quality-trace/tree/main/examples) <!-- and join our [Discord Community](https://discord.gg/8MtcMrQNbX) -->
for more info!