# syntax=docker/dockerfile:1

FROM golang:1.25 AS builder
WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/server ./cmd/server

FROM alpine:3.20
RUN addgroup -S app && adduser -S app -G app
WORKDIR /app

COPY --from=builder /out/server /app/server

EXPOSE 8080
USER app

ENTRYPOINT ["/app/server"]
