# From base image
FROM golang:1.12 AS builder

#RUN go get -u github.com/Masterminds/glide
#setting up working directory
WORKDIR /go/src/github.com/vm
ADD . /go/src/github.com/vm

#RUN glide install
RUN go build -o server
RUN cp server /server

EXPOSE 8000
ENTRYPOINT ["/server"]