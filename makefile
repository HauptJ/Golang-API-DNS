# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=main
BINARY_UNIX=$(BINARY_NAME)_unix

all: test build
build:
        $(GOBUILD) -o $(BINARY_NAME) -v
test: 
        $(GOTEST) -v ./...
clean: 
        $(GOCLEAN)
        rm -f $(BINARY_NAME)
        rm -f $(BINARY_UNIX)
run:
        $(GOBUILD) -o $(BINARY_NAME) -v ./...
        ./$(BINARY_NAME)
deps:
        $(GOGET) github.com/gorilla/mux
        $(GOGET) github.com/google/go-cmp/cmp
        $(GOGET) github.com/HauptJ/Golang-API-DNS/MXLookup
        $(GOGET) github.com/HauptJ/Golang-API-DNS/AddrLookup
        $(GOGET) github.com/HauptJ/Golang-API-DNS/CNAMELookup
        $(GOGET) github.com/HauptJ/Golang-API-DNS/HostLookup