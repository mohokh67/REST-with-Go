package handlers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/asdine/storm"
	"github.com/google/uuid"

	"../account"
)

func bodyToAccount(r *http.Request, a *account.Account) error {
	if r.Body == nil {
		return errors.New("request body is empty")
	}
	if a == nil {
		return errors.New("an account is required")
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, a)

}

func getPaginationData(w http.ResponseWriter, query url.Values) (int, int) {
	pageNumber := 0
	pageSize := 100
	pn := query["page_number"]
	ps := query["page_size"]
	if len(pn) > 0 {
		pageNumber, _ = strconv.Atoi(pn[0])
	}

	if len(ps) > 0 {
		pageSize, _ = strconv.Atoi(ps[0])
	}
	return pageNumber, pageSize
}

func accountsGetAll(w http.ResponseWriter, r *http.Request, query url.Values) {
	pageNumber, pageSize := getPaginationData(w, query)

	accounts, err := account.All(pageNumber, pageSize)
	if err != nil {
		postError(w, http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodHead {
		postBodyResponse(w, http.StatusOK, jsonResponse{})
		return
	}

	postBodyResponse(w, http.StatusOK, jsonResponse{"data": accounts})
}

func accountsPostOne(w http.ResponseWriter, r *http.Request) {
	a := new(account.Account)
	err := bodyToAccount(r, a)
	if err != nil {
		postError(w, http.StatusBadRequest)
		return
	}

	// a.ID = bson.NewObjectId()
	a.ID = uuid.New()
	a.OrganisationID = uuid.New()

	err = a.Save()
	if err != nil {
		if err == account.ErrorRecordInvalid {
			postError(w, http.StatusBadRequest)
		} else {
			postError(w, http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Localtion", "/organisation/accounts/"+a.ID.String())
	w.WriteHeader(http.StatusCreated)
}

func accountsGetOne(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	account, err := account.Find(id)
	if err != nil {
		if err == storm.ErrNotFound {
			postError(w, http.StatusNotFound)
			return
		}
		postError(w, http.StatusInternalServerError)
		return
	}
	if r.Method == http.MethodHead {
		postBodyResponse(w, http.StatusOK, jsonResponse{})
		return
	}

	postBodyResponse(w, http.StatusOK, jsonResponse{"data": account})
}

func accountsDeleteOne(w http.ResponseWriter, _ *http.Request, id uuid.UUID) {
	err := account.Delete(id)
	if err != nil {
		if err == storm.ErrNotFound {
			postError(w, http.StatusNotFound)
			return
		}
		postError(w, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
