FROM golang:1.17-alpine AS base
WORKDIR /app

ENV GO111MODULE="on"
ENV GOOS="linux"
ENV CGO_ENABLED=0

# System dependencies
RUN apk update \
    && apk add --no-cache \
    ca-certificates \
    git \
    && update-ca-certificates

### Development with hot reload and debugger
FROM base AS dev
WORKDIR /app

# Hot reloading mod
RUN go get -u github.com/cosmtrek/air
EXPOSE 8080

ENTRYPOINT ["air"]

### Executable builder
FROM base AS builder
WORKDIR /app

# Application dependencies
COPY ../groupie-tracker /app
RUN go mod download \
    && go mod verify

RUN go build -o my-great-program -a .

### Production
FROM alpine:latest

RUN apk update \
    && apk add --no-cache \
    ca-certificates \
    curl \
    tzdata \
    && update-ca-certificates

# Copy executable
COPY --from=builder /app/my-great-program /usr/local/bin/my-great-program
EXPOSE 8080

ENTRYPOINT ["/usr/local/bin/my-great-program"]