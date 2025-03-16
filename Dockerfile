# --- Build Stage ---
FROM golang:1.24.1-alpine AS builder

RUN apk add --no-cache git

ENV GOCACHE=/root/.cache/go-build

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN --mount=type=cache,target="/root/.cache/go-build" CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o uuidserver .

# --- Final Stage ---
FROM alpine:latest

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /src/uuidserver .

RUN chown appuser:appgroup /app/uuidserver

ENV GIN_MODE=release

EXPOSE 80

USER appuser

ENTRYPOINT ["./uuidserver"]
