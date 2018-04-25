package api

import (
	"net/http"
	"encoding/json"
)

func testv1(w http.ResponseWriter, r *http.Reques) {
	json.NewEncoder.Encode("echo")
	w.WriteHeader(http.StatusOK)
	return
}