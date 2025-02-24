FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY . /app/
RUN go mod download
RUN go build -o server.bin cmd/app/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/server.bin . 

EXPOSE 8000

CMD ["/app/server.bin"]
