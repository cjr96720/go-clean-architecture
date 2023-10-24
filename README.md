# go-clean-architecture
Implementing Go clean architecture, simple CRUD API with gin, gorm, and PostgreSQL

### Folder Structure


### Run Project
```shell
git clone https://github.com/cjr96720/go-clean-architecture.git
cd go-clean-architecture
```
#### Docker Compse
```shell
# create docker network
docker network create "go-shop"

# on Mac
docker-compose up -d --bulid
```

#### Migration
```shell
# [Optional] create migration file
migrate create -ext sql -dir internal/database/postgres/migration -seq create_tables

# migrate
migrate -database ${POSTGRES_URL} -path internal/database/postgres/migration up
```

#### Run API server
```shell
go run cmd/main.go
```