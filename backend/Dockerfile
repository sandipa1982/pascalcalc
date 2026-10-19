# syntax=docker/dockerfile:1
FROM golang:1.23-alpine AS builder
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY cmd ./cmd
COPY internal ./internal
ARG SERVICE
RUN test "$SERVICE" = "arithmetic-api" || test "$SERVICE" = "math-api" || test "$SERVICE" = "expression-api"
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /service ./cmd/${SERVICE}

FROM alpine:3.21
RUN addgroup -S app && adduser -S -G app app
USER app
COPY --from=builder /service /service
ENTRYPOINT ["/service"]
