# Dockerfile
FROM golang:1.23-bullseye AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/mercury_algo ./main.go

FROM gcr.io/distroless/base-debian12
WORKDIR /app
COPY --from=build /out/mercury_algo /app/mercury_algo
ENV KITEX_ADDR=:9000 HTTP_ADDR=:8888 \
    DB_HOST=mercuryalgodb.cho4u00yi8z9.us-east-2.rds.amazonaws.com \
    DB_PORT=3306 \
    DB_USER=SeekingAlphaMA \
    DB_PASSWORD=qoibLHkLR662CHPXsMwP \
    DB_NAME=mercuryalgo
EXPOSE 9000 8888
ENTRYPOINT ["/app/mercury_algo"]