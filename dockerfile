FROM golang:1.25.4 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o api ./cmd/server/main.go

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/api .

EXPOSE 8080

CMD ["./api"]
