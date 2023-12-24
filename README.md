# go-clean-architecture
Implementing Go clean architecture, simple CRUD API with gin, gorm, and PostgreSQL

### Prerequisite
- Docker and docker-compose installed
- Basic knowledge of Docker container

### Run Project
1. Clone the repository
```shell
git clone https://github.com/cjr96720/go-clean-architecture.git
cd go-clean-architecture
```
2. Add a new `.env`
3. Docker Compse
```shell
# create docker network
docker network create "go-shop"

# on Mac
docker-compose up -d --bulid

# check if the containers are up and running
❯ docker ps -f "network=go-shop"
CONTAINER ID   IMAGE                              COMMAND                  CREATED          STATUS                    PORTS                                       NAMES
b5dea899d905   go-clean-architecture-go-backend   "./go-backend"           50 seconds ago   Up 40 seconds             0.0.0.0:8080->8080/tcp, :::8080->8080/tcp   go-backend
89ff8fc1ac9a   postgres:15.4                      "docker-entrypoint.s…"   50 seconds ago   Up 50 seconds (healthy)   0.0.0.0:5433->5432/tcp, :::5433->5432/tcp   db-postgres
```

4. Visit Swagger UI on `http://127.0.0.1:8080/docs/index.html`
