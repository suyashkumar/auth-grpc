FROM golang

WORKDIR /go/src/app
COPY . .
RUN go get github.com/golang/dep/cmd/dep
RUN go get -u github.com/golang/protobuf/protoc-gen-go
RUN wget https://github.com/google/protobuf/releases/download/v3.4.0/protoc-3.4.0-linux-x86_64.zip
RUN apt-get update
RUN apt-get install unzip
RUN unzip protoc-3.4.0-linux-x86_64.zip
RUN cp bin/protoc /usr/bin
RUN make
CMD ./auth-grpc
