FROM golang:1.14-alpine3.11 AS builder

WORKDIR /go/src/github.com/PumpkinSeed/mysql-gcs-restore

COPY go.mod go.mod
COPY main.go main.go

RUN go get -d -v ./...

RUN go build -ldflags="-w -s" -o /go/bin/mysql-gcs-restore

CMD ["app"]

FROM alpine:3.11
RUN apk add mysql-client

COPY --from=builder /go/bin/mysql-gcs-restore /usr/local/bin/mysql-gcs-restore
COPY entrypoint.sh entrypoint.sh
RUN chmod +x entrypoint.sh

ENTRYPOINT ["./entrypoint.sh"]
