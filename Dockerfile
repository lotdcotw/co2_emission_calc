FROM golang:1.14.3 AS builder
ADD . /co2e
WORKDIR /co2e
RUN go mod download && go mod vendor
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app ./cmd/co2e

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /co2e/.env .
COPY --from=builder /co2e/data/co2e.ini /root/data/co2e.ini
COPY --from=builder /co2e/app .
CMD ["./app"]
