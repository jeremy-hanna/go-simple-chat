# simple-chat

Simple chat server for publishing messages to netcat connected clients.

## Use
Start the server in a shell
```
  go run server.go
```

Connect to the port via netcat in another tab and type messages
```
nc localhost 7896
```

Other connections from additional tabs will receive messages broadcast to them
