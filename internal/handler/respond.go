package handler

import (
	"encoding/json"
	"net/http"
)

type errResp struct {
	Error string `json:"error"`
}

func writeJson(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}

func badRequest(w http.ResponseWriter, msg string) {
	writeJson(w, http.StatusBadRequest, errResp{Error: msg})
}
func serverErr(w http.ResponseWriter, err error) {
	writeJson(w, http.StatusInternalServerError, errResp{Error: err.Error()})
}
