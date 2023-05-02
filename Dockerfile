FROM golang:1.19.7-alpine3.17 as builder
WORKDIR /app

RUN apk update
RUN apk add libc6-compat git build-base

COPY go.mod go.sum ./

RUN go mod download

COPY ./cmd ./cmd
COPY ./config ./config
COPY ./docs ./docs
COPY ./internal ./internal
COPY ./pkg ./pkg
COPY example.env .

ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0

RUN go build -o main ./cmd/main.go

FROM alpine:latest as runner
WORKDIR /app

RUN apk add libc6-compat

EXPOSE 4420
EXPOSE 4421

COPY --from=builder /app/main .
COPY --from=builder /app/example.env .

ENTRYPOINT [ "/app/main" ]
