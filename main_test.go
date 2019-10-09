package main

import (
	"testing"
	/*"bytes"
	"net/http/httptest"
	"net/http"*/
)

func TestDummy(t* testing.T) {

}

// TODO: FIX Rest Endpoint test
/*
func TestMXLookupRestSuccess(t* testing.T) {
	req, err := http.NewRequest("GET", "/mx", bytes.NewBuffer([]byte("hauptj.com")))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(MXLookupEndPt)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: actual: %v expected: %v",
			status, http.StatusOK)
	}

	var expectedResult string  = '[{"Host":"aspmx.l.google.com.","Pref":10},{"Host":"alt1.aspmx.l.google.com.","Pref":20},{"Host":"alt2.aspmx.l.google.com.","Pref":30},{"Host":"alt3.aspmx.l.google.com.","Pref":40},{"Host":"alt4.aspmx.l.google.com.","Pref":50}]'

	if rr.Body.String() != expectedResult {
		t.Errorf("Handler returned wrong body: actual %v want %v",
			rr.Body.String(), expectedResult)
	}
}
*/