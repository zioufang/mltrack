package endpointtests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"testing"

	"github.com/zioufang/mltrackapi/pkg/api/model"
	"gopkg.in/go-playground/assert.v1"
)

// response body with single model
type modelRunSingle struct {
	Success bool           `json:"success"`
	Message string         `json:"message"`
	Data    model.ModelRun `json:"data"`
}

// response body with multiple models
type modelRunMulti struct {
	Success bool             `json:"success"`
	Message string           `json:"message"`
	Data    []model.ModelRun `json:"data"`
}

func TestCreateModelRun(t *testing.T) {
	models := seedModelTable()
	testCases := []struct {
		input      string
		expName    string
		expModelID uint64
		expSuccess bool
		statusCode int
		errMsg     string
	}{
		{
			input:      fmt.Sprintf(`{"name":"shoes_1065  ", "model_id":%d}`, models[0].ID),
			expName:    "shoes_1065",
			expModelID: models[0].ID,
			expSuccess: true,
			statusCode: http.StatusOK,
			errMsg:     "",
		},
		{
			input:      fmt.Sprintf(`{"name":"shoes_1065  ", "model_id":%d}`, models[0].ID),
			expName:    "shoes_1065",
			expModelID: models[0].ID,
			expSuccess: true,
			statusCode: http.StatusOK,
			errMsg:     "",
		},
		{
			input:      fmt.Sprintf(`{"name":"   ", "model_id":%d}`, models[0].ID),
			expName:    "",
			expModelID: models[0].ID,
			expSuccess: true,
			statusCode: http.StatusOK,
			errMsg:     "",
		},
		{
			input:      `{"name":"somename"}`,
			expSuccess: false,
			statusCode: http.StatusUnprocessableEntity,
			errMsg:     "Model ID cannot be empty",
		},
		{
			input:      `{"name":"somename", "model_id":12345}`,
			expSuccess: false,
			statusCode: http.StatusUnprocessableEntity,
			errMsg:     "Model with id 12345 doesn't exist",
		},
	}

	for _, c := range testCases {
		req, _ := http.NewRequest("POST", "/runs", bytes.NewBufferString(c.input))
		resp := execRequest(req)
		var respMap modelRunSingle
		json.Unmarshal([]byte(resp.Body.String()), &respMap)

		// compare response with expected
		fmt.Println("Testing: " + req.Method + " " + req.URL.String())
		assert.Equal(t, resp.Code, c.statusCode)
		if c.statusCode == http.StatusOK {
			assert.Equal(t, c.expName, respMap.Data.Name)
			assert.Equal(t, c.expModelID, respMap.Data.ModelID)
		} else {
			assert.Equal(t, c.errMsg, respMap.Message)
		}
	}

}

func TestGetModelRunByID(t *testing.T) {
	runs := seedModelRunTable()
	testCases := []struct {
		inID       string
		expName    string
		expSuccess bool
		statusCode int
	}{
		{
			inID:       strconv.Itoa(int(runs[0].ID)),
			expName:    runs[0].Name,
			expSuccess: true,
			statusCode: http.StatusOK,
		},
		{
			inID:       "123456",
			expSuccess: false,
			statusCode: http.StatusBadRequest,
		},
	}

	for _, c := range testCases {
		req, _ := http.NewRequest("GET", "/runs/"+c.inID, nil)
		resp := execRequest(req)
		var respMap modelRunSingle
		json.Unmarshal([]byte(resp.Body.String()), &respMap)

		fmt.Println("Testing: " + req.Method + " " + req.URL.String())
		assert.Equal(t, resp.Code, c.statusCode)
		assert.Equal(t, c.expSuccess, respMap.Success)
		if c.statusCode == http.StatusOK {
			assert.Equal(t, c.expName, respMap.Data.Name)
		}
	}
}

func TestDeleteModelRunByID(t *testing.T) {
	runs := seedModelRunTable()
	testCases := []struct {
		inID       string
		expSuccess bool
		statusCode int
	}{
		{
			inID:       fmt.Sprint(runs[0].ID),
			expSuccess: true,
			statusCode: http.StatusOK,
		},
		{
			inID:       "123456",
			expSuccess: false,
			statusCode: http.StatusBadRequest,
		},
	}

	for _, c := range testCases {
		req, _ := http.NewRequest("DELETE", "/runs/"+c.inID, nil)
		resp := execRequest(req)
		var respMap modelRunSingle
		json.Unmarshal([]byte(resp.Body.String()), &respMap)

		fmt.Println("Testing: " + req.Method + " " + req.URL.String())
		assert.Equal(t, respMap.Success, c.expSuccess)
		assert.Equal(t, resp.Code, c.statusCode)
	}
}

func TestGetModelRunListByParam(t *testing.T) {
	runs := seedModelRunTable()
	// TODO fix the hard coded expCount
	testCases := []struct {
		inModelID  string
		expCount   int
		expSuccess bool
		statusCode int
	}{
		{
			inModelID:  fmt.Sprint(runs[0].ModelID),
			expCount:   2,
			expSuccess: true,
			statusCode: http.StatusOK,
		},
		{
			inModelID:  fmt.Sprint(runs[2].ModelID),
			expCount:   1,
			expSuccess: true,
			statusCode: http.StatusOK,
		},
		{
			inModelID:  "123456",
			expSuccess: false,
			statusCode: http.StatusBadRequest,
		},
	}

	for _, c := range testCases {
		req, _ := http.NewRequest("GET", "/runs/list", nil)
		q := req.URL.Query()
		if c.inModelID != "" {
			q.Add("model_id", c.inModelID)
		}
		req.URL.RawQuery = q.Encode()
		fmt.Println("Testing: " + req.Method + " " + req.URL.String())
		resp := execRequest(req)
		var respMap modelRunMulti
		json.Unmarshal([]byte(resp.Body.String()), &respMap)

		fmt.Println("Resp Message: " + respMap.Message)
		assert.Equal(t, resp.Code, c.statusCode)
		assert.Equal(t, c.expSuccess, respMap.Success)
		if c.statusCode == http.StatusOK {
			assert.Equal(t, c.expCount, len(respMap.Data))
		}
	}

}
