package endpointtests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/zioufang/mltrackapi/pkg/api/model"
	"gopkg.in/go-playground/assert.v1"
)

// response body with single attrs
type runNumAttrSingle struct {
	Success bool             `json:"success"`
	Message string           `json:"message"`
	Data    model.RunNumAttr `json:"data"`
}

// response body with multiple attrs
type runNumAttrMulti struct {
	Success bool               `json:"success"`
	Message string             `json:"message"`
	Data    []model.RunNumAttr `json:"data"`
}

func TestCreateRunNumAttr(t *testing.T) {
	runs := seedModelRunTable()
	testCases := []struct {
		input         string
		expName       string
		expModelRunID uint64
		expCategory   string
		expValue      float32
		expSuccess    bool
		statusCode    int
		errMsg        string
	}{
		{
			input:         fmt.Sprintf(`{"name":"shoes_1065   ", "model_run_id":%d, "category":"metric", "value":0.1}`, runs[0].ID),
			expName:       "shoes_1065",
			expModelRunID: runs[0].ID,
			expCategory:   "metric",
			expValue:      0.1,
			expSuccess:    true,
			statusCode:    http.StatusUnprocessableEntity,
			errMsg:        "",
		},
		{
			input:      fmt.Sprintf(`{"name":"shoes_1065   ", "model_run_id":%d, "category":"metric", "value":1}`, runs[0].ID),
			expSuccess: false,
			statusCode: http.StatusOK,
			errMsg:     "Name shoes_1065 already exists for this model run",
		},
		{
			input:      fmt.Sprintf(`{"name":"", "model_run_id":%d, "category":"metric", "value":1}`, runs[0].ID),
			expSuccess: false,
			statusCode: http.StatusUnprocessableEntity,
			errMsg:     "Name cannot be empty",
		},
		{
			input:      fmt.Sprintf(`{"name":"something", "model_run_id":12345, "category":"metric", "value":1}`),
			expSuccess: false,
			statusCode: http.StatusUnprocessableEntity,
			errMsg:     "Model Run with id 12345 doesn't exist",
		},
	}
	for _, c := range testCases {
		req, _ := http.NewRequest("POST", "/num_attrs", bytes.NewBufferString(c.input))
		resp := execRequest(req)
		var respMap runNumAttrSingle
		json.Unmarshal([]byte(resp.Body.String()), &respMap)

		// compare response with expected
		fmt.Println("Testing: " + req.Method + " " + req.URL.String())
		if respMap.Message != "" {
			fmt.Println("Msg: " + respMap.Message)
		}
		assert.Equal(t, resp.Code, c.statusCode)
		assert.Equal(t, c.expSuccess, respMap.Success)
		if c.statusCode == http.StatusOK {
			assert.Equal(t, c.expName, respMap.Data.Name)
			assert.Equal(t, c.expModelRunID, respMap.Data.ModelRunID)
			assert.Equal(t, c.expCategory, respMap.Data.Category)
			assert.Equal(t, c.expValue, respMap.Data.Value)
		} else {
			assert.Equal(t, c.errMsg, respMap.Message)
		}

	}
}

func TestGetRunNumAttrListByParam(t *testing.T) {
	attrs := seedRunNumAttrTable()

	testCases := []struct {
		inModelRunIDs []string
		inNames       []string
		inCategories  []string
		expCount      int
		expSuccess    bool
		statusCode    int
	}{
		{
			inModelRunIDs: []string{fmt.Sprint(attrs[0].ModelRunID)},
			inNames:       []string{"metric_1"},
			inCategories:  []string{"metric"},
			expCount:      1,
			expSuccess:    true,
			statusCode:    http.StatusOK,
		},
		{
			inModelRunIDs: []string{fmt.Sprint(attrs[0].ModelRunID)},
			expCount:      3,
			expSuccess:    true,
			statusCode:    http.StatusOK,
		},
		// TODO better way of getting two different model_run_id
		{
			inModelRunIDs: []string{fmt.Sprint(attrs[0].ModelRunID), fmt.Sprint(attrs[4].ModelRunID)},
			inNames:       []string{"metric_1"},
			expCount:      2,
			expSuccess:    true,
			statusCode:    http.StatusOK,
		},
	}

	for _, c := range testCases {
		req, _ := http.NewRequest("GET", "/num_attrs/list", nil)
		q := req.URL.Query()
		if len(c.inModelRunIDs) > 0 {
			for i := range c.inModelRunIDs {
				q.Add("model_run_id", c.inModelRunIDs[i])

			}
		}
		if len(c.inNames) > 0 {
			for i := range c.inNames {
				q.Add("name", c.inNames[i])
			}
		}
		if len(c.inCategories) > 0 {
			for i := range c.inCategories {
				q.Add("category", c.inCategories[i])
			}
		}
		req.URL.RawQuery = q.Encode()
		fmt.Println("Testing: " + req.Method + " " + req.URL.String())
		resp := execRequest(req)
		var respMap runNumAttrMulti
		json.Unmarshal([]byte(resp.Body.String()), &respMap)

		if respMap.Message != "" {
			fmt.Println("Msg: " + respMap.Message)
		}
		assert.Equal(t, resp.Code, c.statusCode)
		assert.Equal(t, c.expSuccess, respMap.Success)
		if c.statusCode == http.StatusOK {
			assert.Equal(t, c.expCount, len(respMap.Data))
		}
	}
}
