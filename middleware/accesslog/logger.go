package accesslog

import (
	"bytes"

	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"
)

var LoggerDefaultName = "elibot-server "

// LoggerEntry is the structure
// passed to the template.
type LoggerEntry struct {
	StartTime string
	Status    int
	Duration  time.Duration
	Hostname  string
	Method    string
	Path      string
	Request   *http.Request
}

// LoggerDefaultFormat is the format
// logged used by the default Logger instance.
var LoggerDefaultFormat = "{{.StartTime}} | {{.Status}} | \t {{.Duration}} | {{.Hostname}} | {{.Method}} {{.Path}} \n"

// LoggerDefaultDateFormat is the
// format used for date by the
// default Logger instance.
var LoggerDefaultDateFormat = time.RFC3339

// ALogger interface
type ALogger interface {
	Println(v ...interface{})
	Printf(format string, v ...interface{})
}

// Logger is a middleware handler that logs the request as it goes in and the response as it goes out.
type Logger struct {
	// ALogger implements just enough log.Logger interface to be compatible with other implementations
	ALogger
	dateFormat string
	template   *template.Template
	file 	   *os.File
	filename   string
}

func openAccessLogFile(filePath string) (*os.File, error) {
	if ok := checkFileSizeExceededMax(filePath); ok {
		os.Rename(filePath, filePath + ".bak")
	}

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0664)
	if err != nil {
		return nil, fmt.Errorf("error opening file %s: %s", filePath, err)
	}

	return file, nil
}

func (l *Logger) Close() error {
	return l.file.Close()
}

// NewLogger returns a new Logger instance
func NewLogger(filename string) (*Logger, error) {
	file, err := openAccessLogFile(filename)
	if err!=nil {
		return nil, err
	}
	logger := &Logger{ALogger: log.New(file, LoggerDefaultName, 0), dateFormat: LoggerDefaultDateFormat}
	logger.file = file
	logger.filename = filename
	logger.SetFormat(LoggerDefaultFormat)
	return logger, nil
}

func (l *Logger) SetFormat(format string) {
	l.template = template.Must(template.New("elibot_parser").Parse(format))
}

func (l *Logger) SetDateFormat(format string) {
	l.dateFormat = format
}

func (l *Logger) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	start := time.Now()

	res := NewResponseWriter(rw)
	next(res, r)

	log := LoggerEntry{
		StartTime: start.Format(l.dateFormat),
		Status:    res.Status(),
		Duration:  time.Since(start),
		Hostname:  r.Host,
		Method:    r.Method,
		Path:      r.URL.Path,
		Request:   r,
	}

	buff := &bytes.Buffer{}
	l.template.Execute(buff, log)
	l.Printf(buff.String())
}
