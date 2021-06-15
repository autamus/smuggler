package web

import (
	"log"
	"net/http"
)

// NotFound returns an error message for all invalid requests.
// For example, if any unknown URI is used this message is returned.
func NotFound(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Found", 404)
}

// ReportError reports an encountered error back to the user.
func reportError(w http.ResponseWriter, r *http.Request, location string, err error) {
	report := "[ERROR Encountered]: " +
		"[Location: " + location + "] " +
		"[Error: " + err.Error() + "]"
	http.Error(w, report, http.StatusInternalServerError)
	log.Println(report)
}
