FROM golang:1.21-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main ./cmd/

FROM alpine:1.21-alpine

RUN addgroup -g 1001 -S appgroup && \
    adduser -u 1001 -S appuser -G appgroup

COPY --from=builder /build/main /app/main

USER appuser

EXPOSE 8080

ENTRYPOINT ["/app/main"]