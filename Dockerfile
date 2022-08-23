FROM golang:1.19.0-bullseye
CMD ["/app/go-images"]
EXPOSE 8080
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN go build .

USER 9999:9999
