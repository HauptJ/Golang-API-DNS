package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"

	"./MXLookup"
	"./AddrLookup"
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
	hostnames, err := AddrLookup.GetMXRecord(params["host"])
	if err != null {
		respondWithError(writer, http.StatusBadRequest, err.Error())
	} else {
		respondWithJson(writer, http.StatusOK, hostnames)
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
	router.HandleFunc("/mxlookup/{host}", MXLookupEndPt).Methods("GET")
	router.HandleFunc("/addrlookup/{host}", AddrLookupEndPt).Methods("GET")
	
	if err := http.ListenAndServe(":8880", router); err != nil {
		log.Fatal(err)
	}
}