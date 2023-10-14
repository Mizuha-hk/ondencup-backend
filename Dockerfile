FROM golang:1.21.1

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /onden-backend

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o main ./cmd/main.go

RUN chmod 755 ./main

EXPOSE 8080

CMD [ "./main" ]