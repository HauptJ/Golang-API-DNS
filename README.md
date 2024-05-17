# Golang-API-DNS
[![Build Status](https://travis-ci.org/HauptJ/Golang-API-DNS.svg?branch=master)](https://travis-ci.org/HauptJ/Golang-API-DNS)

A simple Golang REST API for DNS queries.

### To Build and Run:
- **NOTE:** Port `8080` or the one specified in `main.go` must be available.

1. Install Golang if you want to run the API on your host machine.
  - [Golang Download and installation instructions](https://golang.org/dl/)


2. Install Golang dependencies if you want to run the API on your host machine. If you wish to run the API in a Docker container, you can skip this step.
  - `go get -u github.com/gorilla/mux`
  - `go get -u github.com/google/go-cmp/cmp`
  - `go get -u github.com/HauptJ/Golang-API-DNS/MXLookup`
  - `go get -u github.com/HauptJ/Golang-API-DNS/AddrLookup`
  - `go get -u github.com/HauptJ/Golang-API-DNS/CNAMELookup`
  - `go get -u github.com/HauptJ/Golang-API-DNS/HostLookup`

3. To run on host system
  1. `go build main.go`
  2. `./main` or `.\main.exe`

### Docker

To Build and Run:

    NOTE 1: Port 8880 or the one specified in docker-compose.yml and main.go must be available.
    NOTE 2: On Windows 10 Pro or greater? You can use Chocolatey to install Docker Desktop for Windows and everything else listed below.

    Install Docker and Docker Compose.

    Docker
        Chocolatey for Windows
    Docker Compose
        Chocolatey for Windows

    Build and Run the container:
        docker-compose build
        docker-compose up

    To stop
        docker-compose down



### Endpoints:

#### /mx
- **TYPE:** GET
- Returns the list of MX records for the specified host
- `/mx/{host}`
- **Example:** `/mx/hauptj.com`
- Returns:
```
[
    {
        "Host": "string",
        "Pref": int
    },
...
]
```
- Example:
```
[
    {
        "Host": "aspmx.l.google.com.",
        "Pref": 10
    },
    {
        "Host": "alt1.aspmx.l.google.com.",
        "Pref": 20
    },
    {
        "Host": "alt2.aspmx.l.google.com.",
        "Pref": 30
    },
    {
        "Host": "alt3.aspmx.l.google.com.",
        "Pref": 40
    },
    {
        "Host": "alt4.aspmx.l.google.com.",
        "Pref": 50
    }
]
```

### /addr
- **TYPE:** GET
- Returns a the hostname assigned to an IP address
- `/addr/{host}`
- **Example:** `/addr/1.1.1.1`
- Returns:
[
    "string"
]
- Example:
```
[
    "one.one.one.one."
]
```

### /cname
- **TYPE:** GET
- Returns the cname record for the specified host
- **Example:** `/cname/research.swtch.com`
- Returns:
`string`
- Example:
```
"ghs.google.com."
```

### /host
- **TYPE:** GET
- Returns the A and AAAA records for a given host
- **Example:** `/host/localhost`
- Returns:
```
[
    "string",
    "string"
]
```
- Example:
```
[
    "::1",
    "127.0.0.1"
]
```

TODO:
- Fix Rest endpoint test for MX endpoint.
- Implement Rest endpoint tests for the other endpoints.
- Implement, IP, NX, SRV, and TXT endpoints.
- Swagger Docs