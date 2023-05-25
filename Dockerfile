FROM golang:alpine
WORKDIR /EmailAliasManager/
COPY . ./
RUN apk add build-base
RUN CGO_ENABLED=1 go build ./server.go
CMD ["/EmailAliasManager/server"]