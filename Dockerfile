FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o holiday-api ./cmd/api

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/holiday-api .
COPY --from=builder /app/docs ./docs

EXPOSE 8080

CMD ["./holiday-api"]