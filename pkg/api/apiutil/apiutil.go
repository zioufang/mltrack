package apiutil

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Entity is an abstract interface for a DB model used in API
type Entity interface {
	FormatAndValidate() error
}

// HTTPError sets status code and error message to the ResponseWriter
func HTTPError(w http.ResponseWriter, statusCode int, err error) {
	w.WriteHeader(statusCode)
	w.Write([]byte(err.Error()))
}

// ReadReqBody reads the request body into a specified Struct and validate it
func ReadReqBody(w http.ResponseWriter, r *http.Request, e Entity) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		HTTPError(w, http.StatusUnprocessableEntity, err)
	}
	err = json.Unmarshal(body, &e)
	if err != nil {
		HTTPError(w, http.StatusUnprocessableEntity, err)
	}
	err = e.FormatAndValidate()
	if err != nil {
		HTTPError(w, http.StatusUnprocessableEntity, err)
	}
}
