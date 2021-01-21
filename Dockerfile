FROM golang:1.12-alpine AS builder
RUN apk add --no-cache git

WORKDIR /app
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .
RUN go build -o main .

FROM alpine:3.7 AS app
WORKDIR /app
EXPOSE 80
ENV PORT 80
RUN apk add --no-cache ca-certificates

COPY --from=builder /app/main .
HEALTHCHECK CMD curl --fail http://localhost:80/health || exit 1
CMD ["./main"]