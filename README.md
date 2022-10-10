log-socket
==========
`log-socket` is a drop-in replacement for Go's `log` package that allows for streaming of logs via WebSockets.

## Installation
To install the library:
`go get github.com/taigrr/log-socket`

## Running
To run a demo of this library:
`go run main.go`

This demo will do a sample of every log type and push results to `0.0.0.0:8080`. Once running, you can open a browser and navigate to 
`0.0.0.0:8080` to see an example implementation of how logs are streamed.

