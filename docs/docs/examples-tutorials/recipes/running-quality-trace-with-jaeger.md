# Running Qualitytrace with Jaeger

:::note
[Check out the source code on GitHub here.](https://github.com/intelops/quality-trace/tree/main/examples/quick-start-jaeger-nodejs)
:::

[Qualitytrace]<!--(https://tracetest.io/)-->is a testing tool based on [OpenTelemetry](https://opentelemetry.io/) that allows you to test your distributed application. It allows you to use data from distributed traces generated by OpenTelemetry to validate and assert if your application has the desired behavior defined by your test definitions.

[Jaeger](https://www.jaegertracing.io/) is an open-source, end-to-end distributed tracing solution. It allows you to monitor and troubleshoot transactions in complex distributed systems. It was developed and then open sourced by Uber Technologies. Jaeger provides a distributed tracing solution to enable transactions across multiple heterogeneous systems or microservices to be tracked and displayed as a cascading series of spans.

## Sample Node.js App with Jaeger, OpenTelemetry and Qualitytrace

This is a simple quick start on how to configure a Node.js app to use OpenTelemetry instrumentation with traces and Qualitytrace for enhancing your E2E and integration tests with trace-based testing. The infrastructure will use Jaeger as the trace data store, and OpenTelemetry Collector to receive traces from the Node.js app and send them to Jaeger.

## Prerequisites

You will need [Docker](https://docs.docker.com/get-docker/) and [Docker Compose](https://docs.docker.com/compose/install/) installed on your machine to run this quick start app!

## Project Structure

The project is built with Docker Compose. It contains two distinct `docker-compose.yaml` files.

### 1. Node.js App

The `docker-compose.yaml` file and `Dockerfile` in the root directory are for the Node.js app.

### 2. Qualitytrace

The `docker-compose.yaml` file, `collector.config.yaml`, `quality-trace-provision.yaml`, and `quality-trace-config.yaml` in the `quality-trace` directory are for the setting up Qualitytrace, Jaeger, and the OpenTelemetry Collector.

The `quality-trace` directory is self-contained and will run all the prerequisites for enabling OpenTelemetry traces and trace-based testing with Qualitytrace.

### Docker Compose Network

All `services` in the `docker-compose.yaml` are on the same network and will be reachable by hostname from within other services. For example, `jaeger:14250` in the `collector.config.yaml` will map to the `jaeger` service, where the port `14250` is the port where Jaeger accepts traces. And, `jaeger:16685` in the `quality-trace-provision.yaml` will map to the `jaeger` service and port `16685` where Qualitytrace will fetch trace data from Jaeger.

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

The `docker-compose.yaml` in the `quality-trace` directory is configured with four services.

- **Postgres** - Postgres is a prerequisite for Qualitytrace to work. It stores trace data when running the trace-based tests.
- [**OpenTelemetry Collector**](https://opentelemetry.io/docs/collector/) - A vendor-agnostic implementation of how to receive, process and export telemetry data.
- [**Jaeger**](https://www.jaegertracing.io/) - A trace data store.
- [**Qualitytrace**]<!--(https://tracetest.io/) -->- Trace-based testing that generates end-to-end tests automatically from traces.

```yaml
version: "3"
services:
  quality-trace:
    image: intelops/quality-trace:${TAG:-latest}
    platform: linux/amd64
    volumes:
      - type: bind
        source: ./quality-trace/quality-trace.config.yaml
        target: /app/quality-trace.yaml
      - type: bind
        source: quality-trace/quality-trace-provision.yaml
        target: /app/provision.yaml
    command: --provisioning-file /app/provision.yaml
    ports:
      - 11633:11633
    depends_on:
      postgres:
        condition: service_healthy
      jaeger:
        condition: service_started
      otel-collector:
        condition: service_started
    healthcheck:
      test: ["CMD", "wget", "--spider", "localhost:11633"]
      interval: 1s
      timeout: 3s
      retries: 60
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
    image: otel/opentelemetry-collector-contrib:0.59.0
    command:
      - "--config"
      - "/otel-local-config.yaml"
    volumes:
      - ./quality-trace/collector.config.yaml:/otel-local-config.yaml

  jaeger:
    image: jaegertracing/all-in-one:latest
    restart: unless-stopped
    ports:
      - "16686:16686"
    healthcheck:
      test: ["CMD", "wget", "--spider", "localhost:16686"]
      interval: 1s
      timeout: 3s
      retries: 60
```

Qualitytrace depends on Postgres, Jaeger and the OpenTelemetry Collector. Both Qualitytrace and the OpenTelemetry Collector require config files to be loaded via a volume. The volumes are mapped from the root directory into the `quality-trace` directory and the respective config files.

To start both the Node.js app and Qualitytrace we will run this command:

```bash
docker-compose -f docker-compose.yaml -f quality-trace/docker-compose.yaml up # add --build if the images are not built already
```

The `quality-trace-config.yaml` file contains the basic setup of connecting Qualitytrace to the Postgres instance.

```yaml
postgres:
  host: postgres
  user: postgres
  password: postgres
  port: 5432
  dbname: postgres
  params: sslmode=disable

telemetry:
  exporters:
    collector:
      serviceName: quality-trace
      sampling: 100 # 100%
      exporter:
        type: collector
        collector:
          endpoint: otel-collector:4317

server:
  telemetry:
    exporter: collector
```

The `quality-trace.provision.yaml` file defines the trace data store, set to Jaeger, meaning the traces will be stored in Jaeger and Qualitytrace will fetch them from Jaeger when running tests.

But how does Qualitytrace fetch traces?

Qualitytrace uses `jaeger:16685` to connect to Jaeger and fetch trace data.

```yaml
---
type: PollingProfile
spec:
  name: Default
  strategy: periodic
  default: true
  periodic:
    retryDelay: 5s
    timeout: 10m

---
type: DataStore
spec:
  name: Jaeger
  type: jaeger
  default: true
  jaeger:
    endpoint: jaeger:16685
    tls:
      insecure: true
```

How do traces reach Jaeger?

The `collector.config.yaml` explains that. It receives traces via either `grpc` or `http`. Then, exports them to Jaegers's `model.proto` endpoint `jaeger:14250`.

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
    loglevel: debug
  jaeger:
    endpoint: jaeger:14250
    tls:
      insecure: true

service:
  pipelines:
    traces/1:
      receivers: [otlp]
      processors: [batch]
      exporters: [jaeger]
```

## Run Both the Node.js App and Qualitytrace

To start both the Node.js app and Qualitytrace, we will run this command:

```bash
docker-compose -f docker-compose.yaml -f quality-trace/docker-compose.yaml up # add --build if the images are not built already
```

This will start your Qualitytrace instance on `http://localhost:11633/`.

Open the URL and start creating tests! Make sure to use the `http://app:8080/` URL in your test creation, because your Node.js app and Qualitytrace are in the same network.

## Learn More

Feel free to check out our [examples in GitHub](https://github.com/intelops/quality-trace/tree/main/examples) <!--  and join our [Discord Community](https://discord.gg/8MtcMrQNbX) -->for more info!