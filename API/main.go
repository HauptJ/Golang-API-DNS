package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"

	"./MXLookup"
	"./CNAMELookup"
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

func CNAMELookupEndPt(writer http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	cnameRecord, err := CNAMELookup.GetCNAMERecord(params["host"])
	if err != nil {
		respondWithError(writer, http.StatusBadRequest, err.Error())
	} else {
		respondWithError(writer, http.StatusOK, cnameRecord)
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
	router.HandleFunc("/cname/{host}", CNAMELookupEndPt).Methods("GET")
	if err := http.ListenAndServe(":8880", router); err != nil {
		log.Fatal(err)
	}
}