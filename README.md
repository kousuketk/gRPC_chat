# gRPC_chat
```
- start server
go run server/main.go

- createMessage
grpc_cli call localhost:9090 CreateMessage "message: 'test_message'"

- getMessages
grpc_cli call localhost:9090 GetMessages ""
```