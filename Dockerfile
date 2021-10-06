FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on

COPY go.mod /build
COPY go.sum /build/

RUN cd /build/ && git clone https://github.com/NaddiNadja/grpc101.git
RUN cd /build/grpc101/server && go build ./...

EXPOSE 9080

ENTRYPOINT [ "/build/grpc101/server/server" ]