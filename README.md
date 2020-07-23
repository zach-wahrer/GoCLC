# GoCLC
GoCLC is a command line, multi-user chat server and client written in, you guessed it, Go!

## To Run Your Own Server
Clone the repo, then execute `go run . -server=true` in the project's root directory. The server will begin running, listening on `localhost:8000`. Optionally, you can specify your own address and port with the `-address` and `-port` command flags, respectively.

To connect to an already running server as a client, use the command `go run . -address=[server IP] -port=[server port]`. Similarly to the server command, if you omit the `-address` and `-port` flags, the client will attempt to connect to a server at `localhost:8000`.

Since GoCLC is currently in the early stages of development, the server cannot do much more than greet a user and offer rudimentary commands. However, this will change rapidly, so keep an eye on this repo!

## To Run the Test Suite
From the root of the project:
* `go test` will run integration tests for starting/running the server.
* `go test server` will run all unit tests related to the server code.
* `go test client` will run all unit tests related to the client code.

## Implemented Features
* General chat
* Comprehensive server logs

## Planned Features
* \#Channel based chat
* Private messaging between clients
* Encrypted message passing between client-server-client

## Maybe Features
* GUI client
* Web client
* User registration / authentication
