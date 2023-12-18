FROM --platform=linux/amd64 golang:latest as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o main

FROM --platform=arm64 ubuntu:latest as runner

WORKDIR /app

RUN set -x
RUN apt-get update
RUN DEBIAN_FRONTEND=noninteractive apt-get install -y ca-certificates
RUN rm -rf /var/lib/apt/lists/*

RUN mkdir -p /home/runner
RUN groupadd -r runner && useradd -r -g runner runner

COPY --from=builder /app/main .
# COPY --from=builder /app/.env .

EXPOSE 8000
ENV PORT 8000

USER runner

CMD [ "./main", "start"]
