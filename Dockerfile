FROM golang:latest as build

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o ipfs-metadata *.go

FROM ubuntu:latest

RUN apt-get update && apt-get install ca-certificates -y && update-ca-certificates

COPY --from=build /app/ipfs-metadata /ipfs-metadata
COPY data /data

ENV POSTGRES_USER="user"
ENV POSTGRES_PASSWORD="password"
ENV POSTGRES_DB="user"
ENV POSTGRES_HOST="postgres"
ENV POSTGRES_PORT="5432"

CMD ["/ipfs-metadata"]