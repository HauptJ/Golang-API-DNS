package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"

	"./MXLookup"
	"./AddrLookup"
	"./CNAMELookup"
	"./HostLookup"
)


func MXLookupEndPt(writer http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	mxRecord, err := MXLookup.GetMXRecord(params["host"])
	if err != nil {
		respondWithError(writer, http.StatusBadRequest, err.Error())
	} else {
		respondWithJson(writer, http.StatusOK, mxRecord)
	}
}


func AddrLookupEndPt(writer http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	hostnames, err := AddrLookup.GetAddrHost(params["host"])
	if err != nil {
		respondWithError(writer, http.StatusBadRequest, err.Error())
	} else {
		respondWithJson(writer, http.StatusOK, hostnames)
  }
}


func CNAMELookupEndPt(writer http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	cnameRecord, err := CNAMELookup.GetCNAMERecord(params["host"])
	if err != nil {
		respondWithError(writer, http.StatusBadRequest, err.Error())
	} else {
		respondWithJson(writer, http.StatusOK, cnameRecord)
	}
}

func HostLookupEndPt(writer http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	cnameRecord, err := HostLookup.GetHostRecord(params["host"])
	if err != nil {
		respondWithError(writer, http.StatusBadRequest, err.Error())
	} else {
		respondWithJson(writer, http.StatusOK, cnameRecord)
	}
}


func respondWithError(writer http.ResponseWriter, code int, msg string) {
	respondWithJson(writer, code, map[string]string{"ERROR": msg})
}


func respondWithJson(writer http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Problem generating JSON response %v\n", err)
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(code)
	writer.Write(response)
}


func main() {
	router := mux.NewRouter()

	// Endpoints
	router.HandleFunc("/addr/{host}", AddrLookupEndPt).Methods("GET")
	router.HandleFunc("/mx/{host}", MXLookupEndPt).Methods("GET")
	router.HandleFunc("/cname/{host}", CNAMELookupEndPt).Methods("GET")
	router.HandleFunc("/host/{host}", HostLookupEndPt).Methods("GET")

  
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}