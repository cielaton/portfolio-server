FROM golang:1.22.5

WORKDIR /app

EXPOSE 1323
EXPOSE 5432

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN cat .env_docker > .env
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -o /server

CMD ["/server"]
