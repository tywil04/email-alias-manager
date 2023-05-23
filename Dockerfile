FROM golang:alpine
WORKDIR /EmailAliasManager/
COPY . ./
RUN go build ./server.go
CMD ["/EmailAliasManager/server"]