FROM golang:alpine

WORKDIR /EmailAliasManager/
COPY . ./
RUN apk add build-base
RUN CGO_ENABLED=1 go build ./server.go

ENV CRYPTO_PEPPER=""
ENV SQLITE_DATABASE_PATH="/database.db"
ENV SERVER_ADDRESS=":4041"

CMD ["/EmailAliasManager/server"]
