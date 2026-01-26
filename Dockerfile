FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /service ./cmd/service/main.go

RUN go build -o /migrate ./cmd/migrate/main.go

FROM alpine:3.18


WORKDIR /app

COPY --from=builder /service /app/service
COPY --from=builder /migrate /app/migrate

COPY ./migrations /app/migrations

EXPOSE 8080

CMD ["/app/service"]
