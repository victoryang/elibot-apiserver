package main

import (
	"fmt"
	"context"
	"net/http"
	"time"
)

const (
	HealthPath = "/health"
	Metrics = "/metrics"
)

type Options struct {
	Headers   		map[string]string
	Hostname 		string
	Jar 			http.CookieJar
	Transport 		http.RoundTripper
	Interval  		time.Duration
}

type BackendConfig struct {
	Options
	url 		   	string
	requestTimeout 	time.Duration
}

type Collector struct {

}

type Backend struct {
	config 			*BackendConfig
	collectors		[]*Collector
	status 			bool
	err				error
	metrics			[][]byte
}

// this function adds additional http headers and hostname to http.request
func (b *Backend) addHeadersAndHost(req *http.Request) *http.Request {
	if b.config.Options.Hostname != "" {
		req.Host = b.config.Options.Hostname
	}

	for k, v := range b.config.Options.Headers {
		req.Header.Set(k, v)
	}
	return req
}

func NewBackendConfig(options Options, serverurl string, d time.Duration) *BackendConfig {
	return &BackendConfig {
			Options:        options,
			url:			serverurl,
			requestTimeout: d,
		}
}

func NewDefaultCollectors() []*Collector{
	return nil
}

// NewBackend Instantiate a new Backend
func NewBackend(config *BackendConfig) *Backend {
	return &Backend{
		config: 		config,
		collectors:		NewDefaultCollectors(),
		status:			false,
		metrics:		nil,
	}
}

// checkHealth returns a nil error in case it was successful and otherwise
// a non-nil error with a meaningful description why the health check failed.
func checkHealth(backend *Backend) {
	req, err := http.NewRequest(http.MethodGet, backend.config.url + HealthPath, nil)
	if err != nil {
		backend.err = fmt.Errorf("failed to create HTTP request: %s", err)
		backend.status = false
	}

	req = backend.addHeadersAndHost(req)

	client := http.Client{
		Timeout:    backend.config.requestTimeout * time.Second,
		Transport:  backend.config.Options.Transport,
		Jar:		backend.config.Options.Jar,
	}

	resp, err := client.Do(req)
	if err != nil {
		backend.err = fmt.Errorf("HTTP request failed: %s", err)
		backend.status = false
	}

	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest {
		backend.err = fmt.Errorf("received error status code: %v", resp.StatusCode)
		backend.status = false
		return
	}

	backend.status = true
	return
}

func (b *Backend) startcheckHealth(ctx context.Context) {
	ticker := time.NewTicker(b.config.Options.Interval*time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			go checkHealth(b)
		}
	}
}