FROM golang:alpine3.18 as builder
WORKDIR /golang-sqlite-crud
RUN apk add --no-cache gcc musl-dev sqlite-dev
COPY . .
RUN CGO_ENABLED=1 GOOS=linux GOARCH=arm64 go build -ldflags="-w -s" .
FROM golang:alpine3.18
WORKDIR /golang-sqlite-crud
COPY --from=builder /golang-sqlite-crud /usr/bin/
ENTRYPOINT ["golang-sqlite-crud"]
