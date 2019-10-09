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