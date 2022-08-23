FROM golang:1.19.0-bullseye AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN go build .

FROM debian:bullseye-slim as app
CMD ["/go-images"]
EXPOSE 8080
ENV MONGO_ADDR="mongo"
COPY --from=builder /app/go-images /
USER 9999:9999
