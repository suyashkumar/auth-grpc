FROM golang:1.9

WORKDIR /go/src/app
COPY . .  
RUN wget https://github.com/google/protobuf/releases/download/v3.5.1/protoc-3.5.1-linux-x86_64.zip
RUN apt-get update
RUN apt-get install unzip
RUN unzip protoc-3.5.1-linux-x86_64.zip 
RUN cp bin/protoc /usr/local/bin/protoc 
RUN go get -u github.com/golang/dep/cmd/dep
RUN go get -u github.com/golang/protobuf/protoc-gen-go
RUN make
