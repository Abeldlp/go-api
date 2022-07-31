## Go api simple project

This project was only to learn how to make a crud api from scratch with the go language.

Structure of the code is as follows

```
cmd
    |_ main.go                  => This is the entry point server
    |_ cmd                      => The compiled server
pkg
    |_ config
        |_ app.go               => This is the configuration for the database
    |_ controller
        |_ book-controller.go   => The controller for the book Model
    |_ model
        |_ book.go              => The model holding the structure for a book
    |_ routes
        |_ bookstore-routes.go  => The routes pointing to the book-controller
    |_ utils
        |_ utils.go             => Body parser for request json
docker-compose.yml              => docker compose file for the database

```

To run the project locally follow these steps

1. Clone the project `git clone https://github.com/Abeldlp/go-api.git`
2. Install dependencies `go mod tidy`
3. Run a docker container with the database `docker-compose up -d`
4. Run the server `cd cmd && go run main.go`

Those steps assume that you have docker / docker-compose installed alongside go
