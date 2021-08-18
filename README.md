# gRPC_chat
```
- start server
$ go run server/main.go
2021/08/18 18:19:18 Received: test1
2021/08/18 18:19:18 Sent: test1: 2021-08-18 18:19:18
2021/08/18 18:19:22 Received: test2
2021/08/18 18:19:22 Sent: test2: 2021-08-18 18:19:22
2021/08/18 18:19:23 Received: test3
2021/08/18 18:19:23 Sent: test3: 2021-08-18 18:19:23
2021/08/18 18:19:25 Received: test4
2021/08/18 18:19:25 Sent: test4: 2021-08-18 18:19:25

- createMessage
$ grpc_cli call localhost:9090 CreateMessage "message: 'test1'"          
connecting to localhost:9090
message: "test1"
Rpc succeeded with OK status

- getMessages
$ grpc_cli call localhost:9090 GetMessages ""
connecting to localhost:9090
message: "test1: 2021-08-18 18:19:18"
message: "test2: 2021-08-18 18:19:22"
message: "test3: 2021-08-18 18:19:23"
message: "test4: 2021-08-18 18:19:25"
```