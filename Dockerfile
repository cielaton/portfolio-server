FROM golang:1.22.5

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -o /server

CMD ["/server"]