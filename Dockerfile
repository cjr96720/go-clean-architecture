FROM golang:1.21.3

RUN apt-get update && apt-get install -y vim

# for migration
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# for swagger
RUN go install github.com/swaggo/swag/cmd/swag@latest

WORKDIR /root/go-clean-architecture
