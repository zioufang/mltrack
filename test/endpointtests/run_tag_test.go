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
type runTagSingle struct {
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Data    model.RunTag `json:"data"`
}

// response body with multiple attrs
type runTagMulti struct {
	Success bool           `json:"success"`
	Message string         `json:"message"`
	Data    []model.RunTag `json:"data"`
}

func TestCreateRunTag(t *testing.T) {
	runs := seedModelRunTable()
	testCases := []struct {
		input         string
		expKey        string
		expModelRunID uint64
		expValue      string
		expSuccess    bool
		statusCode    int
		errMsg        string
	}{
		{
			input:         fmt.Sprintf(`{"key":"shoes_1065   ", "model_run_id":%d, "value":"lace"}`, runs[0].ID),
			expKey:        "shoes_1065",
			expModelRunID: runs[0].ID,
			expValue:      "lace",
			expSuccess:    true,
			statusCode:    http.StatusOK,
			errMsg:        "",
		},
		{
			input:      fmt.Sprintf(`{"key":"shoes_1065   ", "model_run_id":%d, "value":"lace"}`, runs[0].ID),
			expSuccess: false,
			statusCode: http.StatusUnprocessableEntity,
			errMsg:     "Key shoes_1065 already exists for this model run",
		},
		{
			input:      fmt.Sprintf(`{"key":"", "model_run_id":%d, "value":"lace"}`, runs[0].ID),
			expSuccess: false,
			statusCode: http.StatusUnprocessableEntity,
			errMsg:     "Key cannot be empty",
		},
		{
			input:      fmt.Sprintf(`{"key":"something", "model_run_id":12345,  "value":"lace"}`),
			expSuccess: false,
			statusCode: http.StatusUnprocessableEntity,
			errMsg:     "Model Run with id 12345 doesn't exist",
		},
	}
	for _, c := range testCases {
		req, _ := http.NewRequest("POST", "/tags", bytes.NewBufferString(c.input))
		resp := execRequest(req)
		var respMap runTagSingle
		json.Unmarshal([]byte(resp.Body.String()), &respMap)

		// compare response with expected
		fmt.Println("Testing: " + req.Method + " " + req.URL.String())
		if respMap.Message != "" {
			fmt.Println("Msg: " + respMap.Message)
		}
		assert.Equal(t, resp.Code, c.statusCode)
		assert.Equal(t, c.expSuccess, respMap.Success)
		if c.statusCode == http.StatusOK {
			assert.Equal(t, c.expKey, respMap.Data.Key)
			assert.Equal(t, c.expModelRunID, respMap.Data.ModelRunID)
			assert.Equal(t, c.expValue, respMap.Data.Value)
		} else {
			assert.Equal(t, c.errMsg, respMap.Message)
		}

	}
}

func TestGetRunTagListByParam(t *testing.T) {
	tags := seedRunTagTable()

	testCases := []struct {
		inModelRunIDs []string
		inKeys        []string
		expCount      int
		expSuccess    bool
		statusCode    int
	}{
		{
			inModelRunIDs: []string{fmt.Sprint(tags[0].ModelRunID)},
			inKeys:        []string{"git_hash"},
			expCount:      1,
			expSuccess:    true,
			statusCode:    http.StatusOK,
		},
		{
			inModelRunIDs: []string{fmt.Sprint(tags[0].ModelRunID)},
			expCount:      2,
			expSuccess:    true,
			statusCode:    http.StatusOK,
		},
		// TODO better way of getting two different model_run_id
		{
			inModelRunIDs: []string{fmt.Sprint(tags[0].ModelRunID), fmt.Sprint(tags[4].ModelRunID)},
			inKeys:        []string{"git_hash"},
			expCount:      2,
			expSuccess:    true,
			statusCode:    http.StatusOK,
		},
	}

	for _, c := range testCases {
		req, _ := http.NewRequest("GET", "/tags/list", nil)
		q := req.URL.Query()
		if len(c.inModelRunIDs) > 0 {
			for i := range c.inModelRunIDs {
				q.Add("model_run_id", c.inModelRunIDs[i])

			}
		}
		if len(c.inKeys) > 0 {
			for i := range c.inKeys {
				q.Add("key", c.inKeys[i])
			}
		}
		req.URL.RawQuery = q.Encode()
		fmt.Println("Testing: " + req.Method + " " + req.URL.String())
		resp := execRequest(req)
		var respMap runTagMulti
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
