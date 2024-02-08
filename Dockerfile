FROM golang:1.22.0-alpine AS builder

COPY . /source/
WORKDIR /source/

RUN go mod download
RUN go build -o ./bin/crud_server cmd/grpc_server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /source/bin/crud_server .

CMD ["./crud_server"]