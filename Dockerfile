# Base build image
FROM golang:alpine
# Install some dependencies needed to build the project
RUN apk add --no-cache git
RUN apk add --no-cache bash
RUN apk add --no-cache gcc
RUN mkdir /build
ADD . /build/
ADD ./vendor /build/vendor
WORKDIR /build
# Force the go compiler to use modules
ENV GO111MODULE=on

#COPY secrets.json /build/secrets.json
RUN CGO_ENABLED=0 GOOS=linux go build -mod vendor -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .
CMD ["./main"]