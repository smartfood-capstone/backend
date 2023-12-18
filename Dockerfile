FROM golang:latest as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o main

FROM ubuntu:latest as runner

WORKDIR /app

RUN set -x
RUN apt-get update
RUN DEBIAN_FRONTEND=noninteractive apt-get install -y ca-certificates
RUN rm -rf /var/lib/apt/lists/*

RUN mkdir -p /home/runner
RUN groupadd -r runner && useradd -r -g runner runner

COPY --from=builder /app/main .
# COPY --from=builder /app/.env .

ENV PORT=$PORT
ENV API_KEY=$API_KEY
ENV DB_HOST=$DB_HOST
ENV DB_PORT=$DB_PORT
ENV DB_USER=$DB_USER
ENV DB_PASSWORD=$DB_PASSWORD
ENV DB_NAME=$DB_NAME
ENV DB_SSL_MODE=$DB_SSL_MODE

EXPOSE $PORT

USER runner

CMD [ "./main", "start"]
