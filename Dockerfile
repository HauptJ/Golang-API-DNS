############################################################
# Dockerfile for Github Golang API
# Based on golang latest
############################################################

# Set the base image to Golang Latest
FROM golang:1.22.0

# File Author / Maintainer
MAINTAINER joshua@hauptj.com


################## Begin Installation ######################

# Install Go Dependencies


# Copy Golang source code
RUN mkdir /DNSAPI
ADD . /DNSAPI/

WORKDIR /DNSAPI

# build API
RUN go build -o main .

# Run API
CMD ["/DNSAPI/main"]
