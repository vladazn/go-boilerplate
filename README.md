
migrate the database
```bash
go run main.go db migrate
```

run the grpc server
```bash
go run main.go server start
```

run the api 
```bash
go run main.go gateway start
```