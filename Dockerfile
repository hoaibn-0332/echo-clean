# Builder
FROM golang:1.23.0-alpine3.19 as builder

RUN apk update && apk upgrade && \
    apk --update add git make bash build-base

WORKDIR /app

COPY . .

RUN make build

# Distribution
FROM alpine:latest

RUN apk update && apk upgrade && \
    apk --update --no-cache add tzdata && \
    mkdir /app

WORKDIR /app

EXPOSE 8080

COPY --from=builder /app/engine /app/.env /app/
COPY --from=builder /app/migrations /app/migrations

CMD /app/engine