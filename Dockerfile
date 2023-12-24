FROM golang:1.21.3

WORKDIR /app

# download packages
COPY go.mod go.sum ./
RUN go mod download
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# build application
COPY . .
RUN go build -o go-backend ./cmd/main.go

# expose 8080 port
EXPOSE 8080

# run application
CMD ["./go-backend"]
