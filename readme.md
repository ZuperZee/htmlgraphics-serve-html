# Serve html

A simple http server to serve html files related to this issue <https://github.com/gapitio/gapit-htmlgraphics-panel/issues/167>.

## Usage

### Go

Prerequisite <https://go.dev/>

Start the server on port 8080

```bash
go run main.go
```

### Docker

Prerequisite <https://www.docker.com/>

Build the docker image

```bash
docker build . -t serve-html
```

Change the network to the same network as Grafana (or something Grafana can reach).

Start the docker container

```bash
docker compose up -d
```
