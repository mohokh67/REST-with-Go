package handlers

import (
	"net/http"
	"strings"

	"github.com/google/uuid"
)

// AccountRouter handles the accounts route
func AccountRouter(w http.ResponseWriter, r *http.Request) {
	// fmt.Printf("[%s] %q\n", r.Method, r.URL.String())
	q := r.URL.Query()

	path := strings.TrimSuffix(r.URL.Path, "/")
	if path == "/organisation/accounts" {
		switch r.Method {
		case http.MethodGet:
			accountsGetAll(w, r, q)
			return
		case http.MethodPost:
			accountsPostOne(w, r)
			return
		case http.MethodHead:
			accountsGetAll(w, r, q)
			return
		case http.MethodOptions:
			postOptionResponse(w, []string{http.MethodGet, http.MethodPost, http.MethodHead, http.MethodOptions}, nil)
			return
		default:
			postError(w, http.StatusMethodNotAllowed)
			return
		}
	}

	path = strings.TrimPrefix(path, "/organisation/accounts/")
	id, err := uuid.Parse(path)
	if err != nil {
		postError(w, http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		accountsGetOne(w, r, id)
		return
	case http.MethodDelete:
		accountsDeleteOne(w, r, id)
		return
	case http.MethodHead:
		accountsGetOne(w, r, id)
		return
	case http.MethodOptions:
		postOptionResponse(w, []string{http.MethodGet, http.MethodDelete, http.MethodHead, http.MethodOptions}, nil)
		return
	default:
		postError(w, http.StatusMethodNotAllowed)
		return
	}
}
