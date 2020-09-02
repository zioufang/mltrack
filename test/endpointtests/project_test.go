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

// response body with single project
type projectSingle struct {
	Success bool          `json:"success"`
	Message string        `json:"message"`
	Data    model.Project `json:"data"`
}

// response body with multiple projects
type projectMulti struct {
	Success bool            `json:"success"`
	Message string          `json:"message"`
	Data    []model.Project `json:"data"`
}

func TestCreateProject(t *testing.T) {
	testCases := []struct {
		input          string
		expName        string
		expDescription string
		expSuccess     bool
		statusCode     int
		errMsg         string
	}{
		{
			input:          `{"name":"shoes_1065  ", "description":"this is a project for shoes"}`,
			expName:        "shoes_1065",
			expDescription: "this is a project for shoes",
			expSuccess:     true,
			statusCode:     http.StatusOK,
			errMsg:         "",
		},
		{
			input:      `{"name":"shoes_1065", "description":"this is a project for shoes"}`,
			expSuccess: false,
			statusCode: http.StatusUnprocessableEntity,
			errMsg:     "Project Name already exists",
		},
		{
			input:      `{"name":"  ", "description":"nothingness"}`,
			expSuccess: false,
			statusCode: http.StatusUnprocessableEntity,
			errMsg:     "Project Name cannot be empty",
		},
	}

	resetTables()

	for _, c := range testCases {
		req, _ := http.NewRequest("POST", "/projects", bytes.NewBufferString(c.input))
		resp := execRequest(req)
		var respMap projectSingle
		json.Unmarshal([]byte(resp.Body.String()), &respMap)

		// compare response with expected
		fmt.Println("Testing: " + req.Method + " " + req.URL.String())
		assert.Equal(t, resp.Code, c.statusCode)
		if c.statusCode == http.StatusOK {
			assert.Equal(t, c.expName, respMap.Data.Name)
			assert.Equal(t, c.expDescription, respMap.Data.Description)
		} else {
			assert.Equal(t, c.errMsg, respMap.Message)
		}
	}

}

func TestGetProjectByID(t *testing.T) {
	projects := seedProjectTable()
	testCases := []struct {
		inID       string
		expName    string
		expSuccess bool
		statusCode int
	}{
		{
			inID:       strconv.Itoa(int(projects[0].ID)),
			expName:    projects[0].Name,
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
		req, _ := http.NewRequest("GET", "/projects/"+c.inID, nil)
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

func TestGetProjectByParam(t *testing.T) {
	projects := seedProjectTable()
	testCases := []struct {
		inID       string
		inName     string
		expID      uint64
		expSuccess bool
		statusCode int
	}{
		{
			inID:       fmt.Sprint(projects[0].ID),
			expID:      projects[0].ID,
			expSuccess: true,
			statusCode: http.StatusOK,
		},
		{
			inName:     projects[1].Name,
			expID:      projects[1].ID,
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
		req, _ := http.NewRequest("GET", "/projects", nil)
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
		var respMap projectSingle
		json.Unmarshal([]byte(resp.Body.String()), &respMap)

		assert.Equal(t, resp.Code, c.statusCode)
		assert.Equal(t, c.expSuccess, respMap.Success)
		if c.statusCode == http.StatusOK {
			assert.Equal(t, c.expID, respMap.Data.ID)
		}
	}
}

func TestUpdateProjectByID(t *testing.T) {
	projects := seedProjectTable()
	testCases := []struct {
		inID           string
		updateJSON     string
		expName        string
		expDescription string
		expSuccess     bool
		statusCode     int
	}{
		{
			inID:           fmt.Sprint(projects[0].ID),
			updateJSON:     `{"name":"daoud_2"}`,
			expName:        "daoud_2",
			expDescription: projects[0].Description,
			expSuccess:     true,
			statusCode:     http.StatusOK,
		},
		{
			inID:           fmt.Sprint(projects[1].ID),
			updateJSON:     `{"description":"this is project estobar 2"}`,
			expName:        projects[1].Name,
			expDescription: "this is project estobar 2",
			expSuccess:     true,
			statusCode:     http.StatusOK,
		},
		{
			inID:           fmt.Sprint(projects[1].ID),
			updateJSON:     `{"name":"berwick", "description":"this is project berwick"}`,
			expName:        "berwick",
			expDescription: "this is project berwick",
			expSuccess:     true,
			statusCode:     http.StatusOK,
		},
		{
			inID:       "12345",
			updateJSON: `{"name":"berwick", "description":"this is project berwick"}`,
			expSuccess: false,
			statusCode: http.StatusInternalServerError,
		},
		{
			inID:       fmt.Sprint(projects[1].ID),
			updateJSON: `{}`,
			expSuccess: false,
			statusCode: http.StatusInternalServerError,
		},
	}

	for _, c := range testCases {
		req, _ := http.NewRequest("PUT", "/projects/"+c.inID, bytes.NewBufferString(c.updateJSON))
		resp := execRequest(req)
		var respMap projectSingle
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

func TestDeleteProjectByID(t *testing.T) {
	projects := seedProjectTable()
	testCases := []struct {
		inID       string
		expSuccess bool
		statusCode int
	}{
		{
			inID:       fmt.Sprint(projects[0].ID),
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
		req, _ := http.NewRequest("DELETE", "/projects/"+c.inID, nil)
		resp := execRequest(req)
		var respMap projectSingle
		json.Unmarshal([]byte(resp.Body.String()), &respMap)

		fmt.Println("Testing: " + req.Method + " " + req.URL.String())
		assert.Equal(t, respMap.Success, c.expSuccess)
		assert.Equal(t, resp.Code, c.statusCode)
	}
}

func TestGetAllProjects(t *testing.T) {
	seedProjectTable()
	req, _ := http.NewRequest("GET", "/projects/all", nil)
	resp := execRequest(req)
	var respMap projectMulti
	json.Unmarshal([]byte(resp.Body.String()), &respMap)

	fmt.Println("Testing: " + req.Method + " " + req.URL.String())
	assert.Equal(t, resp.Code, http.StatusOK)
	assert.Equal(t, respMap.Success, true)
	assert.Equal(t, len(respMap.Data), 2)
}
