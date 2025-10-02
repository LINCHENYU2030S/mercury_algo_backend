# Dockerfile
FROM golang:1.22-bullseye AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/mercury_algo ./main.go

FROM gcr.io/distroless/base-debian12
WORKDIR /app
COPY --from=build /out/trading-bot-service .
ENV KITEX_ADDR=:9000 \
    DB_HOST=db \
    DB_PORT=3306 \
    DB_USER=root \
    DB_PASSWORD=secret \
    DB_NAME=mercuryalgodb
EXPOSE 9000
ENTRYPOINT ["/api/mercury_algo"]