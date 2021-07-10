## Step to run

1. Start Zipkin server with Docker
```
$docker container run -d -p 9411:9411 openzipkin/zipkin
```

Open url http://localhost:9411

2. Start service 1
```
$cd service01
$go run main_with_tracing.go
```

Open url http://localhost:8080/users/1

3. Start service 2
```
$cd service02
$go run main.go
```

Open url http://localhost:8080/users/123