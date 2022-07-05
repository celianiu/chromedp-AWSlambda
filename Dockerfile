FROM golang:1.18-alpine AS builder


WORKDIR /app

RUN apk add go git
RUN go env -w GOPROXY=direct

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main





FROM chromedp/headless-shell:latest

RUN apt-get update; apt-get install -y dumb-init musl-dev ca-certificates
COPY --from=builder /app/main .

ENTRYPOINT ["./main"]