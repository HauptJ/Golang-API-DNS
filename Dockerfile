############################################################
# Dockerfile for Github DNS API
# Based on golang latest
############################################################

# Set the base image to Golang Latest
FROM golang:latest

# File Author / Maintainer
MAINTAINER joshua@hauptj.com

################## Begin Installation ######################

# Install Go Dependencies
RUN go get github.com/gorilla/context \
    github.com/gorilla/mux \
    golang.org/x/net/context/ctxhttp \
    github.com/google/go-cmp/cmp


# Copy Golang source code
RUN mkdir /DNSAPI
ADD API /DNSAPI/

WORKDIR /DNSAPI

# Build the API
RUN go build -o main .

# Run the API
CMD ["/DNSAPI/main"]

# Healthcheck to restart API if it crashes
HEALTHCHECK CMD curl --fail http://localhost:8080/mx/hauptj.com || exit 1

