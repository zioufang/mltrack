package apiutil

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jinzhu/gorm"
)

// Entity is an abstract interface for a DB model used in API
type Entity interface {
	FormatAndValidate(*gorm.DB) error
}

// respBody is the struct for response body from http request
type respBody struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (rb *respBody) parseResp(success bool, message string, data interface{}) {
	rb.Success = success
	rb.Message = message
	rb.Data = data
}

// responseJSON wraps the response in {success: "", data: {}} format
func respJSON(w http.ResponseWriter, success bool, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	resp := &respBody{}
	resp.parseResp(success, message, data)
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

// RespSuccessWithMessage sets the success to true and add the data payload in response
func RespSuccessWithMessage(w http.ResponseWriter, message string, data interface{}) {
	respJSON(w, true, message, data)
}

// RespSuccess sets the success to true and add the data payload in response without the message
func RespSuccess(w http.ResponseWriter, data interface{}) {
	RespSuccessWithMessage(w, "", data)
}

// RespError sets status code, sets the success to false and add an error message to data paylaod in response
func RespError(w http.ResponseWriter, statusCode int, err error) {
	w.WriteHeader(statusCode)
	respJSON(w, false, err.Error(), "")
}

// ReadReqBody reads the request body into a specified Struct and validate it
func ReadReqBody(w http.ResponseWriter, r *http.Request, db *gorm.DB, e Entity) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &e)
	if err != nil {
		return err
	}
	err = e.FormatAndValidate(db)
	if err != nil {
		return err
	}
	return nil
}
