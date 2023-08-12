# Be Short

Documentation(Local): http://localhost:8080/swagger/index.html

![Build](https://github.com/wellingtonlope/ticket-api/actions/workflows/build.yaml/badge.svg)
![Coverage](https://img.shields.io/badge/Coverage-0%25-red)

## Usage
You can use .env to set up the environment and start the http server with the following command:
```bash
go run cmd/api/main.go
```

Also, you can start the server with docker and docker-compose, just run the following command:
```bash
# by default the server will listen on port 8080
docker-compose up --build
```

## Tests
To run the tests, run the following command:
```bash
go test ./...
```