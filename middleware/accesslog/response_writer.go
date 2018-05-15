package accesslog

import (
	"net/http"
)

// ResponseWriter is a wrapper around http.ResponseWriter that provides extra information about
// the response. It is recommended that middleware handlers use this construct to wrap a responsewriter
// if the functionality calls for it.
type ResponseWriter struct {
	http.ResponseWriter
	status int
}

func NewResponseWriter(res http.ResponseWriter) *ResponseWriter {
	// Default the status code to 200
	return &ResponseWriter{res}
}

// Give a way to get the status
func (w ResponseWriter) Status() int {
	return w.status
}

// Satisfy the http.ResponseWriter interface
func (w ResponseWriter) Header() http.Header {
	return w.ResponseWriter.Header()
}

func (w ResponseWriter) Write(data []byte) (int, error) {
	if w.status == 0 {
		w.status = http.StatusOK
	}
	return w.ResponseWriter.Write(data)
}

func (w ResponseWriter) WriteHeader(statusCode int) {
	// Store the status code
	w.status = statusCode

	// Write the status code onward.
	w.ResponseWriter.WriteHeader(statusCode)
}
