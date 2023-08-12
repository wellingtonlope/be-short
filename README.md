# Be Short

Documentation(Local): http://localhost:8080/swagger/index.html

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