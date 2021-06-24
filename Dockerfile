FROM golang:1.16.5-stretch

WORKDIR /go/src/config-server
COPY . .

RUN CGO_ENABLED=0 go get -d -v ./...
RUN CGO_ENABLED=0 go install -v ./...

CMD ["config-server", "http"]