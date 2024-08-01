FROM golang:1.22-alpine3.20

WORKDIR /app

RUN apk --no-cache add protobuf protobuf-dev

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

RUN wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.59.1

ADD . /app

RUN chmod +x scripts/test.sh

CMD ["./scripts/test.sh"]
