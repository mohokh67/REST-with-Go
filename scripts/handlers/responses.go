package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
)

type jsonResponse map[string]interface{}

func postError(w http.ResponseWriter, code int) {
	http.Error(w, http.StatusText(code), code)
}

func postBodyResponse(w http.ResponseWriter, code int, content jsonResponse) {
	if content != nil {
		json, err := json.Marshal(content)
		if err != nil {
			postError(w, http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(code)
		w.Write(json)
		return
	}
	w.WriteHeader(code)
	w.Write([]byte(http.StatusText(code)))

}

func postOptionResponse(w http.ResponseWriter, methods []string, content jsonResponse) {
	w.Header().Set("Allow", strings.Join(methods, ","))
	postBodyResponse(w, http.StatusOK, content)
}
