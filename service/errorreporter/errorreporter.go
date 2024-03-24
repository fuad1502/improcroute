package errorreporter

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorReporter struct {
	FuncName string
}

func (reporter ErrorReporter) Report(w http.ResponseWriter, status int, err error) {
	errorStruct := struct{ Error string }{Error: err.Error()}
	b, _ := json.Marshal(&errorStruct)
	http.Error(w, string(b), http.StatusBadRequest)
	log.Printf("%v: %v\n", reporter.FuncName, err)
}
