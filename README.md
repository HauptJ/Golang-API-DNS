# Golang-API-DNS
A simple Golang REST API for DNS queries.

### To Build and Run:
- **NOTE 1:** Port `8080` or the one specified in `main.go` must be available.

1. Install Golang if you want to run the API on your host machine.
  - [Golang Download and installation instructions](https://golang.org/dl/)


2. Install Golang dependencies if you want to run the API on your host machine. If you wish to run the API in a Docker container, you can skip this step.
  - `go get github.com/gorilla/mux`

3. To run on host system
  1. `go build main.go`
  2. `./main` or `.\main.exe`


### Endpoints:

#### /mx
- **TYPE:** GET
- Returns the list of MX records for the specified host
- `/mx/{host}`
- **Example:** `/mx/hauptj.com`

### /addr

### /cname

### /host