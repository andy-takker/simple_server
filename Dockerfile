FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY . /app/
RUN go mod download
RUN go build -o server.bin cmd/app/main.go
RUN go build -o migrate.bin cmd/migrations/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/server.bin . 
COPY --from=builder /app/migrate.bin .
COPY ./migrations /app/migrations
EXPOSE 8000

CMD ["/app/server.bin"]
