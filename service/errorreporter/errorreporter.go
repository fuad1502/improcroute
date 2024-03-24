package errorreporter

import (
	"encoding/json"
	"log"
	"net/http"
)

// ErrorReporter is a struct that handles both writing error information to the
// HTTP response and logging the Error. `FuncName` gives the log output more
// context on which function it failed."
type ErrorReporter struct {
	FuncName string
}

// Report writes an HTTP response with status code `status` and a JSON payload
// body {Error: `err`}. It also writes the error information to log, with an
// additional function context.
func (reporter ErrorReporter) Report(w http.ResponseWriter, status int, err error) {
	errorStruct := struct{ Error string }{Error: err.Error()}
	b, _ := json.Marshal(&errorStruct)
	http.Error(w, string(b), http.StatusBadRequest)
	log.Printf("%v: %v\n", reporter.FuncName, err)
}
