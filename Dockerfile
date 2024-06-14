FROM golang:1.22-bullseye as builder
ENV APP_HOME /go/src/app

WORKDIR "$APP_HOME"
COPY go.* .

RUN go mod download
RUN go mod verify
COPY . .
# RUN go build -o .
RUN go build \
  -ldflags="-linkmode external -extldflags -static" \
  -tags netgo \
  -o .

FROM busybox
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY .env .env
COPY --from=builder /go/src/app/AAS .

EXPOSE 8081
CMD ["./AAS"]


