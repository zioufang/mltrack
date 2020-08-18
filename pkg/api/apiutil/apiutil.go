package apiutil

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Entity is an abstract interface for a DB model used in API
type Entity interface {
	FormatAndValidate() error
}

// ResponseJSON wraps the payload data in JSON format
func ResponseJSON(w http.ResponseWriter, data interface{}) {
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

// ResponseError sets status code and error message to the ResponseWriter
func ResponseError(w http.ResponseWriter, statusCode int, err error) {
	w.WriteHeader(statusCode)
	w.Write([]byte(err.Error()))
}

// ReadReqBody reads the request body into a specified Struct and validate it
func ReadReqBody(w http.ResponseWriter, r *http.Request, e Entity) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &e)
	if err != nil {
		return err
	}
	err = e.FormatAndValidate()
	if err != nil {
		return err
	}
	return nil
}
