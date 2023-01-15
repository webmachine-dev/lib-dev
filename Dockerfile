FROM golang:latest

RUN go install github.com/webmachine-dev/lib-dev@latest

ENTRYPOINT [ "lib-dev" ]
