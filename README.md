# test2 

simple API for blog
featuring 2 APIS:
- localhost:8080/api/v1/posts - using Gin
- localhost:8080/api/v2/posts - using Gin + Huma
- localhost:8080/healthcheck
- localhost:8080/doc    - API doc
### Navigate to [http://localhost:8080/doc](http://localhost:8080/doc) to see API v2 spec/docs.
Also there you can download YAML/JSON schema in OpenAPI 3.1 format and make sample requests


## Fetching dependencies
```
make dep
```
or
```
go mod tidy
```

## Building & Running 
```
make run
```
or
```
go build -o bin/task2 cmd/main.go && ./bin/task2
```
### Testing 
```
make test
```
or
```
go test -v ./tests/unit/...
```

