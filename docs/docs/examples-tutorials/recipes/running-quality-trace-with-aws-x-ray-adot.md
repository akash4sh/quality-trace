# Running Qualitytrace with AWS X-Ray (AWS X-Ray Node.js SDK & AWS Distro for OpenTelemetry)

<!--:::note
[Check out the source code on GitHub here.](https://github.com/intelops/quality-trace/tree/main/examples/quality-trace-amazon-x-ray-adot)
:::-->

[Qualitytrace]<!--(https://quality-trace.io/) -->is a testing tool based on [OpenTelemetry](https://opentelemetry.io/) that allows you to test your distributed application. It allows you to use data from distributed traces generated by OpenTelemetry to validate and assert if your application has the desired behavior defined by your test definitions.

[AWS X-Ray](https://aws.amazon.com/xray/) provides a complete view of requests as they travel through your application and filters visual data across payloads, functions, traces, services, APIs and more with no-code and low-code motions.

[AWS Distro for OpenTelemetry (ADOT)](https://aws-otel.github.io/docs/getting-started/collector) is a secure, production-ready, AWS-supported distribution of the OpenTelemetry project. Part of the Cloud Native Computing Foundation, OpenTelemetry provides open source APIs, libraries and agents to collect distributed traces and metrics for application monitoring.

## Simple Node.js API with AWS X-Ray and Qualitytrace

This is a simple quick start guide on how to configure a Node.js app to use instrumentation with traces and Qualitytrace for enhancing your E2E and integration tests with trace-based testing. The infrastructure will use AWS X-Ray as the trace data store, the ADOT as a middleware and a Node.js app to generate the telemetry data.

## Prerequisites

You will need [Docker](https://docs.docker.com/get-docker/) and [Docker Compose](https://docs.docker.com/compose/install/) installed on your machine to run this quick start app!

## Project Structure

The project is built with Docker Compose.

### 1. Node.js App

The `Dockerfile` in the root directory is for the Node.js app.

### 2. Qualitytrace

The `docker-compose.yaml` file, `quality-trace.provision.yaml`, and `quality-trace-config.yaml` in the `root` direforctory are for the setting up the Node.js App, Qualitytrace and ADOT Collector.

### Docker Compose Network

All `services` in the `docker-compose.yaml` are on the same network and will be reachable by hostname from within other services. For example, `adot-collector:2000` in the `src/index.js` will map to the `adot-collector` service, where port `2000` is the port where the X-Ray Daemon accepts telemetry data.

## Node.js App

The Node.js app is a simple Express app, contained in the `src/index.js` file.

It is instrumented using [AWS X-Ray SDK](https://docs.aws.amazon.com/xray/latest/devguide/xray-sdk-nodejs.html), sending the data to the ADOT collector that will push the telemetry information to both the AWS service and the Qualitytrace OLTP endpoint.

The key instrumentation section from the `src/index.js` file.

```js
const AWSXRay = require("aws-xray-sdk");
const XRayExpress = AWSXRay.express;
const express = require("express");

AWSXRay.setDaemonAddress("adot-collector:2000");

// Capture all AWS clients we create
const AWS = AWSXRay.captureAWS(require("aws-sdk"));
AWS.config.update({
  region: process.env.AWS_REGION || "us-west-2",
});

// Capture all outgoing https requests
AWSXRay.captureHTTPsGlobal(require("https"));
const https = require("https");

const app = express();
const port = 3000;

app.use(XRayExpress.openSegment("Qualitytrace"));
```

To start the server run this command.

```bash
npm start
```

As you can see the `Dockerfile` uses the command above.

```Dockerfile
FROM node:slim
WORKDIR /usr/src/app

COPY ./src/package*.json ./

RUN npm install
COPY ./src .

EXPOSE 3000
CMD [ "npm", "start" ]
```

## Qualitytrace

The `docker-compose.yaml` includes three other services.

- **Postgres** - Postgres is a prerequisite for Qualitytrace to work. It stores trace data when running the trace-based tests.
- [**AWS Distro for OpenTelemetry (ADOT)**](https://aws-otel.github.io/docs/getting-started/collector) - Software application that listens for traffic on UDP port 2000, gathers raw segment data and relays it to the AWS X-Ray API. The daemon works in conjunction with the AWS X-Ray SDKs and must be running so that data sent by the SDKs can reach the X-Ray service.
- [**Qualitytrace**] <!--(https://tracetest.io/) -->- Trace-based testing that generates end-to-end tests automatically from traces.

```yaml
version: "3"

services:
  quality-trace:
    image: intelops/quality-trace:${TAG:-latest}
    platform: linux/amd64
    volumes:
      - type: bind
        source: ./quality-trace-config.yaml
        target: /app/quality-trace.yaml
      - type: bind
        source: ./quality-trace.provision.yaml
        target: /app/provisioning.yaml
    ports:
      - 11633:11633
    command: --provisioning-file /app/provisioning.yaml
    extra_hosts:
      - "host.docker.internal:host-gateway"
    depends_on:
      postgres:
        condition: service_healthy
      adot-collector:
        condition: service_started
    healthcheck:
      test: ["CMD", "wget", "--spider", "localhost:11633"]
      interval: 1s
      timeout: 3s
      retries: 60

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
    ports:
      - 5432:5432

  adot-collector:
    image: amazon/aws-otel-collector:latest
    command:
      - "--config"
      - "/otel-local-config.yaml"
    volumes:
      - ./collector.config.yaml:/otel-local-config.yaml
    environment:
      AWS_ACCESS_KEY_ID: ${AWS_ACCESS_KEY_ID}
      AWS_SECRET_ACCESS_KEY: ${AWS_SECRET_ACCESS_KEY}
      AWS_SESSION_TOKEN: ${AWS_SESSION_TOKEN}
      AWS_REGION: ${AWS_REGION}
    ports:
      - 4317:4317
      - 2000:2000
```

Qualitytrace depends on Postgres and the ADOT collector. Qualitytrace requires config files to be loaded via a volume. The volumes are mapped from the root directory into the `root` directory and the respective config files.

The `quality-trace.config.yaml` file contains the basic setup of connecting Qualitytrace to the Postgres instance.

```yaml
postgres:
  host: postgres
  user: postgres
  password: postgres
  port: 5432
  dbname: postgres
  params: sslmode=disable
```

The `quality-trace.provision.yaml` file defines the trace data store, set to the OTel Collector, meaning the traces will be sent to the ADOT instance where later on will be pushed to the AWX X-Ray service and the OTLP Qualitytrace endpoint.

```yaml
---
type: DataStore
spec:
  name: otlp
  type: otlp

```

But how does Qualitytrace fetch traces?

The ADOT collector is configured with the `awsxray` receiver and exporting the OpenTelemetry format directly intro Qualitytrace.

```yaml
receivers:
  awsxray:
    transport: udp

processors:
  batch:

exporters:
  awsxray:
    region: ${AWS_REGION}
  otlp/quality-trace:
    endpoint: quality-trace:4317
    tls:
      insecure: true

service:
  pipelines:
    traces/quality-trace:
      receivers: [awsxray]
      processors: [batch]
      exporters: [otlp/quality-trace]
    traces/awsxray:
      receivers: [awsxray]
      exporters: [awsxray]
```

How do traces reach AWS X-Ray?

The application code in the `src/index.js` file uses the native AWS SDK X-Ray library which sends telemetry data to the ADOT Collector to be processed and then sent to the configured AWS X-Ray SaaS and Qualitytrace.

## Run Both the Node.js App and Qualitytrace

To start both the Node.js app and Qualitytrace, run this command:

```bash
docker-compose up
```

This will start your Qualitytrace instance on `http://localhost:11633/`. Open the instance and start creating tests!
Make sure to use the `http://app:3000/` url in your test creation, because your Node.js app and Qualitytrace are in the same network.

## Learn More

Please visit our [examples in GitHub](https://github.com/intelops/quality-trace/tree/main/examples) <!-- and join our [Discord Community](https://discord.gg/8MtcMrQNbX) --> for more info!