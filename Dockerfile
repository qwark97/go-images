FROM golang:1.19.0-bullseye AS builder
WORKDIR /app
# disable cross-compiling (might cause some dynamic links to libc/libmusl; source: https://stackoverflow.com/a/55106860/14181841)
ENV CGO_ENABLED=0

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
# build the binary with debug information removed
RUN go build -ldflags="-w -s" -o go-images .

FROM scratch as app
CMD [ "./go-images" ]
EXPOSE 8080
ENV MONGO_ADDR="mongo"
COPY --from=builder /app/go-images go-images
USER 9999:9999
