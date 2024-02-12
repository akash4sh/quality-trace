# Running a Python app with OpenTelemetry manual instrumention

:::note
[Check out the source code on GitHub here.](https://github.com/intelops/quality-trace/tree/main/examples/quick-start-python)
:::

[Qualitytrace]<!--(https://quality-trace.io/)--> is a testing tool based on [OpenTelemetry](https://opentelemetry.io/) that allows you to test your distributed application. It allows you to use data from distributed traces generated by OpenTelemetry to validate and assert if your application has the desired behavior defined by your test definitions.

## Sample Python app with OpenTelemetry Collector and Qualitytrace

This is a simple quick start on how to configure a Python app to use OpenTelemetry instrumentation with traces, and Qualitytrace for enhancing your e2e and integration tests with trace-based testing.

## Prerequisites

You will need [Docker](https://docs.docker.com/get-docker/) and [Docker Compose](https://docs.docker.com/compose/install/) installed on your machine to run this quick start app!

## Project structure

The project is built with Docker Compose. It contains two distinct `docker-compose.yaml` files.

### 1. Python app

The `docker-compose.yaml` file and `Dockerfile` in the root directory are for the Python app.

### 2. Qualitytrace

The `docker-compose.yaml` file, `collector.config.yaml`, `quality-trace-provision.yaml`, and `quality-trace-config.yaml` in the `quality-trace` directory are for setting up Qualitytrace and the OpenTelemetry Collector.

The `quality-trace` directory is self-contained and will run all the prerequisites for enabling OpenTelemetry traces and trace-based testing with Qualitytrace.

### Docker Compose Network

All `services` in the `docker-compose.yaml` are on the same network and will be reachable by hostname from within other services. E.g. `quality-trace:4317` in the `collector.config.yaml` will map to the `quality-trace` service, where the port `4317` is the port where Qualitytrace accepts traces.

## Python app

The Python app is a simple Flask app, contained in the `app.py` file.

The code below imports all the Flask, and OpenTelemetry libraries and configures both manual and automatic OpenTelemetry instrumentation.

```python
from flask import Flask, request
import json

from opentelemetry import trace
from opentelemetry.sdk.resources import Resource
from opentelemetry.sdk.trace import TracerProvider
from opentelemetry.sdk.trace.export import BatchSpanProcessor
from opentelemetry.sdk.trace.export import ConsoleSpanExporter

provider = TracerProvider()
processor = BatchSpanProcessor(ConsoleSpanExporter())
provider.add_span_processor(processor)
trace.set_tracer_provider(provider)
tracer = trace.get_tracer(__name__)
```

There are 3 endpoints in the Flask app. For seeing manual instrumentation trigger the `"/manual"` endpoint. For seeing the automatic instrumentation trigger the `"/automatic"` endpoint respectively. 

```python
app = Flask(__name__)

@app.route("/manual")
def manual():
    with tracer.start_as_current_span(
        "manual", 
        attributes={ "endpoint": "/manual", "foo": "bar" } 
    ):
        return "App works with a manual instrumentation."

@app.route('/automatic')
def automatic():
    return "App works with automatic instrumentation."

@app.route("/")
def home():
    return "App works."
```

The `Dockerfile` includes bootstrapping the needed OpenTelemetry packages. As you can see it does not have the `CMD` command. Instead, it's configured in the `docker-compose.yaml` below.

```Dockerfile
FROM python:3.10.1-slim
WORKDIR /opt/app
COPY . .
RUN pip install --no-cache-dir -r requirements.txt
RUN opentelemetry-bootstrap -a install
EXPOSE 8080
```

The `docker-compose.yaml` contains just one service for the Python app. The service is stared with the `command` parameter.

```yaml
version: '3'
services:
  app:
    image: quick-start-python
    platform: linux/amd64
    extra_hosts:
      - "host.docker.internal:host-gateway"
    build: .
    ports:
      - "8080:8080"
    # using the command here instead of the Dockerfile
    command: opentelemetry-instrument --traces_exporter otlp --service_name app --exporter_otlp_endpoint otel-collector:4317 --exporter_otlp_insecure true flask run --host=0.0.0.0 --port=8080
    depends_on:
      quality-trace:
        condition: service_started
```

To start it, run this command:

```bash
docker compose build # optional if you haven't already built the image
docker compose up
```

This will start the Python app. But, you're not sending the traces anywhere.

Let's fix this by configuring Qualitytrace and OpenTelemetry Collector.

## Qualitytrace

The `docker-compose.yaml` in the `quality-trace` directory is configured with three services.

- **Postgres** - Postgres is a prerequisite for Qualitytrace to work. It stores trace data when running the trace-based tests.
- [**OpenTelemetry Collector**](https://opentelemetry.io/docs/collector/) - A vendor-agnostic implementation of how to receive, process and export telemetry data.
- [**Qualitytrace**](https://quality-trace.io/) - Trace-based testing that generates end-to-end tests automatically from traces.

```yaml
version: "3"
services:
  quality-trace:
    image: intelops/quality-trace:latest
    platform: linux/amd64
    volumes:
      - type: bind
        source: ./quality-trace/quality-trace-config.yaml
        target: /app/quality-trace.yaml
      - type: bind
        source: ./quality-trace/quality-trace-provision.yaml
        target: /app/provisioning.yaml
    ports:
      - 11633:11633
    command: --provisioning-file /app/provisioning.yaml
    depends_on:
      postgres:
        condition: service_healthy
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

```

Qualitytrace depends on both Postgres and the OpenTelemetry Collector. Both Qualitytrace and the OpenTelemetry Collector require config files to be loaded via a volume. The volumes are mapped from the root directory into the `quality-trace` directory and the respective config files.

The `quality-trace-config.yaml` file contains the basic setup of connecting Qualitytrace to the Postgres instance.

```yaml
postgres:
  host: postgres
  user: postgres
  password: postgres
  port: 5432
  dbname: postgres
  params: sslmode=disable
```

The `quality-trace-provision.yaml` file provisions the trace data store and polling to store in the Postgres database. The data store is set to OTLP meaning the traces will be stored in Qualitytrace itself.

```yaml
---
type: DataStore
spec:
  name: OpenTelemetry Collector
  type: otlp
  isdefault: true
```
  
But how are traces sent to Qualitytrace?

The `collector.config.yaml` explains that. It receives traces via either `grpc` or `http`. Then, exports them to Qualitytrace's otlp endpoint `quality-trace:4317`.

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
  otlp/1:
    endpoint: quality-trace:4317
    # Send traces to Qualitytrace.
    # Read more in docs here: https://docs.tracetest.io/configuration/connecting-to-data-stores/opentelemetry-collector
    tls:
      insecure: true

service:
  pipelines:
    traces/1:
      receivers: [otlp]
      processors: [batch]
      exporters: [otlp/1]

```

## Run both the Python app and Qualitytrace

To start both the Python app and Qualitytrace we will run this command:

```bash
docker-compose -f docker-compose.yaml -f quality-trace/docker-compose.yaml up # add --build if the images are not built already
```

This will start your Qualitytrace instance on `http://localhost:11633/`. Go ahead and open it up.

Start creating tests! Make sure to use the `http://app:8080/` url in your test creation, because your Python app and Qualitytrace are in the same network.

## Learn more

Feel free to check out our [examples in GitHub](https://github.com/intelops/quality-trace/tree/main/examples)
<!--  and join our [Discord Community](https://discord.gg/8MtcMrQNbX) for more info! -->