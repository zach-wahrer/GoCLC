# GoCLC
GoCLC is a command line, multi-user chat server and client written in, you guessed it, Go!

## To Run Your Own Server
Clone the repo, then execute `go run .` in the project's root directory. The server will begin running, listening on `localhost:8000`.

Since GoCLC is currently in the early stages of development, the server cannot do much more than greet a user and offer rudimentary commands. However, this will change rapidly, so keep an eye on this repo!

## To Run the Test Suite
From the root of the project:
* `go test` will run integration tests for starting/running the server.
* `go test server` will run all unit tests related to the server code.

## Planned Features
* General chat
* \#Channel based chat
* Private messaging between clients
* Encrypted message passing between client-server-client
