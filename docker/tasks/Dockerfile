FROM golang:1.22.1-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o go-todo ./main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder ./app .

ENTRYPOINT ["go-todo"]
