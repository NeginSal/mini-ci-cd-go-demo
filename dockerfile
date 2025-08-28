FROM golang:1.24.5 AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o mini-ci-cd-go-demo ./main.go

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/mini-ci-cd-go-demo .

EXPOSE 8080

CMD ["./mini-ci-cd-go-demo"]
