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
type modelSingle struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    model.Model `json:"data"`
}

// response body with multiple models
type modelMulti struct {
	Success bool          `json:"success"`
	Message string        `json:"message"`
	Data    []model.Model `json:"data"`
}

func TestCreateModel(t *testing.T) {
	projects := seedProjectTable()
	testCases := []struct {
		input          string
		expName        string
		expProjectID   uint64
		expStatus      string
		expDescription string
		expSuccess     bool
		statusCode     int
		errMsg         string
	}{
		{
			input:          fmt.Sprintf(`{"name":"shoes_1065  ", "project_id":%d, "status":"prod", "description":"this is a model for shoes"}`, projects[0].ID),
			expName:        "shoes_1065",
			expProjectID:   projects[0].ID,
			expStatus:      "prod",
			expDescription: "this is a model for shoes",
			expSuccess:     true,
			statusCode:     http.StatusOK,
			errMsg:         "",
		},
		{
			input:      fmt.Sprintf(`{"name":"shoes_1065  ", "project_id":%d, "status":"prod", "description":"this is a model for shoes"}`, projects[0].ID),
			expSuccess: false,
			statusCode: http.StatusUnprocessableEntity,
			errMsg:     "Model Name already exists",
		},
		{
			input:      fmt.Sprintf(`{"name":"   ", "project_id":%d, "status":"prod", "description":"this is a model for shoes"}`, projects[0].ID),
			expSuccess: false,
			statusCode: http.StatusUnprocessableEntity,
			errMsg:     "Model Name cannot be empty",
		},
		{
			input:      `{"name":"somename", "status":"prod", "description":"this is a model for shoes"}`,
			expSuccess: false,
			statusCode: http.StatusUnprocessableEntity,
			errMsg:     "Project ID cannot be empty",
		},
		{
			input:      `{"name":"somename", "project_id":12345, "status":"prod", "description":"this is a model for shoes"}`,
			expSuccess: false,
			statusCode: http.StatusUnprocessableEntity,
			errMsg:     "Project with id 12345 doesn't exist",
		},
	}

	for _, c := range testCases {
		req, _ := http.NewRequest("POST", "/models", bytes.NewBufferString(c.input))
		resp := execRequest(req)
		var respMap modelSingle
		json.Unmarshal([]byte(resp.Body.String()), &respMap)

		// compare response with expected
		fmt.Println("Testing: " + req.Method + " " + req.URL.String())
		assert.Equal(t, resp.Code, c.statusCode)
		if c.statusCode == http.StatusOK {
			assert.Equal(t, c.expName, respMap.Data.Name)
			assert.Equal(t, c.expProjectID, respMap.Data.ProjectID)
			assert.Equal(t, c.expStatus, respMap.Data.Status)
			assert.Equal(t, c.expDescription, respMap.Data.Description)
		} else {
			assert.Equal(t, c.errMsg, respMap.Message)
		}
	}

}

func TestGetModelByID(t *testing.T) {
	models := seedModelTable()
	testCases := []struct {
		inID       string
		expName    string
		expSuccess bool
		statusCode int
	}{
		{
			inID:       strconv.Itoa(int(models[0].ID)),
			expName:    models[0].Name,
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
		req, _ := http.NewRequest("GET", "/models/"+c.inID, nil)
		resp := execRequest(req)
		var respMap projectSingle
		json.Unmarshal([]byte(resp.Body.String()), &respMap)

		fmt.Println("Testing: " + req.Method + " " + req.URL.String())
		assert.Equal(t, resp.Code, c.statusCode)
		assert.Equal(t, c.expSuccess, respMap.Success)
		if c.statusCode == http.StatusOK {
			assert.Equal(t, c.expName, respMap.Data.Name)
		}
	}
}

func TestGetModelByParam(t *testing.T) {
	models := seedModelTable()
	testCases := []struct {
		inID       string
		inName     string
		expID      uint64
		expSuccess bool
		statusCode int
	}{
		{
			inID:       fmt.Sprint(models[0].ID),
			expID:      models[0].ID,
			expSuccess: true,
			statusCode: http.StatusOK,
		},
		{
			inName:     models[1].Name,
			expID:      models[1].ID,
			expSuccess: true,
			statusCode: http.StatusOK,
		},
		{
			inID:       "123456",
			expSuccess: false,
			statusCode: http.StatusBadRequest,
		},
		{
			inName:     "wrongname",
			expSuccess: false,
			statusCode: http.StatusBadRequest,
		},
	}

	for _, c := range testCases {
		req, _ := http.NewRequest("GET", "/models", nil)
		q := req.URL.Query()
		if c.inID != "" {
			q.Add("id", c.inID)
		}
		if c.inName != "" {
			q.Add("name", c.inName)
		}
		req.URL.RawQuery = q.Encode()
		fmt.Println("Testing: " + req.Method + " " + req.URL.String())
		resp := execRequest(req)
		var respMap modelSingle
		json.Unmarshal([]byte(resp.Body.String()), &respMap)

		fmt.Println("Resp Message: " + respMap.Message)
		assert.Equal(t, resp.Code, c.statusCode)
		assert.Equal(t, c.expSuccess, respMap.Success)
		if c.statusCode == http.StatusOK {
			assert.Equal(t, c.expID, respMap.Data.ID)
		}
	}
}

func TestUpdateModelByID(t *testing.T) {
	models := seedModelTable()
	testCases := []struct {
		inID           string
		updateJSON     string
		expName        string
		expStatus      string
		expDescription string
		expSuccess     bool
		statusCode     int
	}{
		{
			inID:           fmt.Sprint(models[0].ID),
			updateJSON:     `{"name":"daoud_2"}`,
			expName:        "daoud_2",
			expStatus:      models[0].Status,
			expDescription: models[0].Description,
			expSuccess:     true,
			statusCode:     http.StatusOK,
		},
		{
			inID:       fmt.Sprint(models[0].ID),
			updateJSON: `{"name":"berwick","project_id":123, "description":"this is model berwick"}`,
			expSuccess: false,
			statusCode: http.StatusInternalServerError,
		},
		{
			inID:           fmt.Sprint(models[1].ID),
			updateJSON:     `{"status":"newstatus", "description":"this is model estobar 2"}`,
			expName:        models[1].Name,
			expStatus:      "newstatus",
			expDescription: "this is model estobar 2",
			expSuccess:     true,
			statusCode:     http.StatusOK,
		},
		{
			inID:           fmt.Sprint(models[1].ID),
			updateJSON:     `{"name":"berwick", "description":"this is model berwick"}`,
			expName:        "berwick",
			expStatus:      models[1].Status,
			expDescription: "this is model berwick",
			expSuccess:     true,
			statusCode:     http.StatusOK,
		},
		{
			inID:       "12345",
			updateJSON: `{"name":"berwick", "description":"this is model berwick"}`,
			expSuccess: false,
			statusCode: http.StatusInternalServerError,
		},
		{
			inID:       fmt.Sprint(models[1].ID),
			updateJSON: `{}`,
			expSuccess: false,
			statusCode: http.StatusInternalServerError,
		},
	}

	for _, c := range testCases {
		req, _ := http.NewRequest("PUT", "/models/"+c.inID, bytes.NewBufferString(c.updateJSON))
		resp := execRequest(req)
		var respMap modelSingle
		json.Unmarshal([]byte(resp.Body.String()), &respMap)

		fmt.Println("Testing: " + req.Method + " " + req.URL.String())
		fmt.Println(respMap.Message)
		assert.Equal(t, resp.Code, c.statusCode)
		assert.Equal(t, respMap.Success, c.expSuccess)
		if c.statusCode == http.StatusOK {
			assert.Equal(t, c.expName, respMap.Data.Name)
			assert.Equal(t, c.expDescription, respMap.Data.Description)
		}
	}
}

func TestDeleteModelByID(t *testing.T) {
	models := seedModelTable()
	testCases := []struct {
		inID       string
		expSuccess bool
		statusCode int
	}{
		{
			inID:       fmt.Sprint(models[0].ID),
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
		req, _ := http.NewRequest("DELETE", "/models/"+c.inID, nil)
		resp := execRequest(req)
		var respMap modelSingle
		json.Unmarshal([]byte(resp.Body.String()), &respMap)

		fmt.Println("Testing: " + req.Method + " " + req.URL.String())
		assert.Equal(t, respMap.Success, c.expSuccess)
		assert.Equal(t, resp.Code, c.statusCode)
	}
}

func TestGetAllModels(t *testing.T) {
	models := seedModelTable()
	req, _ := http.NewRequest("GET", "/models/all", nil)
	resp := execRequest(req)
	var respMap modelMulti
	json.Unmarshal([]byte(resp.Body.String()), &respMap)

	fmt.Println("Testing: " + req.Method + " " + req.URL.String())
	assert.Equal(t, resp.Code, http.StatusOK)
	assert.Equal(t, respMap.Success, true)
	assert.Equal(t, len(respMap.Data), len(models))
}

func TestGetModelListByParam(t *testing.T) {
	models := seedModelTable()
	// TODO fix the hard coded expCount
	testCases := []struct {
		inProjectID string
		expCount    int
		expSuccess  bool
		statusCode  int
	}{
		{
			inProjectID: fmt.Sprint(models[0].ProjectID),
			expCount:    2,
			expSuccess:  true,
			statusCode:  http.StatusOK,
		},
		{
			inProjectID: fmt.Sprint(models[2].ProjectID),
			expCount:    1,
			expSuccess:  true,
			statusCode:  http.StatusOK,
		},
		{
			inProjectID: "123456",
			expSuccess:  false,
			statusCode:  http.StatusBadRequest,
		},
	}

	for _, c := range testCases {
		req, _ := http.NewRequest("GET", "/models/list", nil)
		q := req.URL.Query()
		if c.inProjectID != "" {
			q.Add("project_id", c.inProjectID)
		}
		req.URL.RawQuery = q.Encode()
		fmt.Println("Testing: " + req.Method + " " + req.URL.String())
		resp := execRequest(req)
		var respMap modelMulti
		json.Unmarshal([]byte(resp.Body.String()), &respMap)

		fmt.Println("Resp Message: " + respMap.Message)
		assert.Equal(t, resp.Code, c.statusCode)
		assert.Equal(t, c.expSuccess, respMap.Success)
		if c.statusCode == http.StatusOK {
			assert.Equal(t, c.expCount, len(respMap.Data))
		}
	}

}
