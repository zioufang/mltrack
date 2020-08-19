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

type respBody struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

func (rb *respBody) parseResp(success bool, data interface{}) {
	rb.Success = success
	rb.Data = data
}

type errorMsg struct {
	Message string `json:"message"`
}

// responseJSON wraps the response in {success: "", data: {}} format
func respJSON(w http.ResponseWriter, success bool, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	resp := &respBody{}
	resp.parseResp(success, data)
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

// RespSuccess sets the success to true and add the data payload in response
func RespSuccess(w http.ResponseWriter, data interface{}) {
	respJSON(w, true, data)
}

// RespError sets status code, sets the success to false and add an error message to data paylaod in response
func RespError(w http.ResponseWriter, statusCode int, err error) {
	w.WriteHeader(statusCode)
	errMsg := errorMsg{Message: err.Error()}
	respJSON(w, false, errMsg)
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
