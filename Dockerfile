FROM golang:1.16.5-stretch

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"
ENV CGO_ENABLED=0

RUN go mod init github.com/andersonribeir0/config-server
RUN go mod tidy

CMD ["tail", "-f", "/dev/null"]