# go-client-server

[![Go Report Card](https://goreportcard.com/badge/github.com/thesampadilla/go-client-server)](https://goreportcard.com/report/github.com/thesampadilla/go-client-server)

> A Go Client-Server local application with Concurrency

`go-client-server` is a local application capable of implementing an http server listening for requests on `localhost:6969` or a concurrent client-server application with a RabbitMQ's external queue as an orchestrator/work queue.

The server implements an ordered map data structure to store and retrieve `key->value` pairs in order of addition.

Clients - either via RabbitMQ or http - can ask the server to perform any of the following operations:
- Add item
- Remove item
- Get item
- Get item by index
- Get all items
For documentation on how to structure those requests, see the [client README](client/README.md).

## Prequisites
1. Golang. See installation instructions for your OS [here](https://go.dev/doc/install).
2. Docker Engine. See installlation instructions for your OS [here](https://docs.docker.com/engine/install/).
3. RabbitMQ server running for non-http server. More information below.

## Install and Run
The repo comes with a pre-built binary `go_client_server` that is ready to run. Simply clone the repo to run it. The http server can run out of the box if the `http` flag is passed (see usage below), but for the concurrent client-server execution mode, the RabbitMQ server needs to be running.

Once Docker Engine is installed, simply run the script to start a container of the RabbitMQ local server:
```
./scripts/run_rabbitmq_server.sh
```

## Usage
To start the binary:
```
go_client_server [server_mode_flag] [sleep_flag] [non_sequential_flag]
```

### Flags and Defaults
`go_client_server` takes 3 optional flags:
- `-server-type`: If passed the value `http`, the application spins up an http server instead. Any other value (or the absence of the flag) defaults to a RabbitMQ client-server application. **If present and passed the value of `http`, all other flags are ignored.**
- `-sleep`: If present, workers will simulate a "heavy" task by randmoly waiting 1-5 seconds before executing the work. This delay will be applied to all methods **except** `getall`.
- `-non-sequential`: If present, the work queue will dispatch messages to workers without receiving acknowledgemnts. This may impact the sequentiality of operations passed to the data structure, espcially if `-sleep` is present.

## Logs and results
The http server does not write any logs. It just prints to stdout or to the request writer.

For the client-server with RabbitMQ application, errors are written to `stoud`. Information (`[INFO]`), warning (`[WARN]`), and results (`[RESULT]`) are written to the appropriate logfile under `logs/`. 

`go_client_server` will create and write to a different logfile based on the configurations passed. The options are:
- `logs/server.log` for default mode.
- `logs/sleep_nonSequential_server.log` for when the `-non-sequential` and `-sleep` flags are present.
- `logs/nonSequential_server.log` for when only the `-non-sequential` flag is present.
- `logs/sleep_server.log` for when only the `-sleep` flag is present.

Each log will contain information about the request, the Goroutine ID processing it, and how many goroutines are running at any time.
For a `sleep_nonSequential_server.log` example:
```
[INFO] 2023/03/26 03:24:58 handlers.go:82: Goroutine ID: 15
[RESULT] 2023/03/26 03:24:58 handlers.go:84: Got All Items: [map[C:c] map[D:d] map[X:x]]
[INFO] 2023/03/26 03:24:58 utils.go:36: Goroutines running at the end of GID 15 WORKER: 13

[INFO] 2023/03/26 03:24:58 handlers.go:16: Goroutine ID: 10
[RESULT] 2023/03/26 03:24:58 handlers.go:19: New orderedmap entry: A->a
[INFO] 2023/03/26 03:24:58 utils.go:36: Goroutines running at the end of GID 10 WORKER: 12

[INFO] 2023/03/26 03:24:58 handlers.go:32: Goroutine ID: 13
[RESULT] 2023/03/26 03:24:58 handlers.go:37: Removed D->d from ordered map
[INFO] 2023/03/26 03:24:58 utils.go:36: Goroutines running at the end of GID 13 WORKER: 11
```