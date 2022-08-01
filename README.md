## Go api simple project

This project was only to learn how to make a crud api from scratch with the go language.

Structure of the code is as follows

```
.
├── cmd
│   ├── cmd
│   └── main.go
├── docker-compose.yml
├── go.mod
├── go.sum
├── pkg
│   ├── config
│   │   └── app.go
│   ├── controller
│   │   └── book-controller.go
│   ├── model
│   │   └── book.go
│   ├── routes
│   │   └── booksstore-routes.go
│   └── utils
│       └── utils.go
└── README.md

```

To run the project locally follow these steps

1. Clone the project `git clone https://github.com/Abeldlp/go-api.git`
2. Install dependencies `go mod tidy`
3. Run a docker container with the database `docker-compose up -d`
4. Run the server `cd cmd && go run main.go`

Those steps assume that you have docker / docker-compose installed alongside go
