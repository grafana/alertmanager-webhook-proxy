# Alertmanager Webhook Proxy

[![CircleCI](https://circleci.com/gh/grafana/alertmanager-webhook-proxy/tree/main.svg?style=svg)](https://circleci.com/gh/grafana/alertmanager-webhook-proxy/tree/main)

A reverse proxy for the Prometheus AlertManager webhook notifier. This proxy service can take in a GoLang template to modify the payload for consumers that do not support the AlertManager webhook payload.

## Build

```sh
$ go build -o awp ./cmd/main.go
```

## Usage

```sh
$ ./awp -address=":8080" -target="http://my.consumer.io" -template="./template.txt"
```

| Command Flag | Default | Description |
| ------------ | ------- | ----------- |
| `-address`   | `:8080` | Server bind address |
| `-target`    | `http://localhost:8090` | Target URL |
| `-template`  | `/tmp/awp/template.txt` | Path to payload transformation template |

## Templating

The proxy loads a GoLang template of your choice to transform the alertmanager
payload.

For GoLang templating syntax, check out the [documentation](https://pkg.go.dev/text/template).
The AlertManager webhook payload data format is detailed [here](https://prometheus.io/docs/alerting/latest/configuration/#webhook_config).
Example templates can be found in the `templater` [test data directory](pkg/templater/testdata).
test
