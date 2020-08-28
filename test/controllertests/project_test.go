package controllertests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/zioufang/mltrackapi/pkg/api/model"
	"gopkg.in/go-playground/assert.v1"
)

type projectRespBody struct {
	Success bool          `json:"success"`
	Message string        `json:"message"`
	Data    model.Project `json:"data"`
}

func TestCreateProject(t *testing.T) {
	clearProjectTable()
	testCases := []struct {
		input          string
		expName        string
		expDescription string
		statusCode     int
		errMsg         string
	}{
		{
			input:          `{"name":"shoes_1065  ", "description":"this is a project for shoes"}`,
			expName:        "shoes_1065",
			expDescription: "this is a project for shoes",
			statusCode:     200,
			errMsg:         "",
		},
		{
			input:      `{"name":"shoes_1065", "description":"this is a project for shoes"}`,
			statusCode: 422,
			errMsg:     "Project Name already exists",
		},
		{
			input:      `{"name":"  ", "description":"nothingness"}`,
			statusCode: 422,
			errMsg:     "Project Name cannot be empty",
		},
	}

	for _, c := range testCases {
		req, _ := http.NewRequest("POST", "/projects", bytes.NewBufferString(c.input))
		resp := execRequest(req)
		var respMap projectRespBody
		json.Unmarshal([]byte(resp.Body.String()), &respMap)

		// compare response with expected
		assert.Equal(t, resp.Code, c.statusCode)
		if c.statusCode == 200 {
			assert.Equal(t, c.expName, respMap.Data.Name)
			assert.Equal(t, c.expDescription, respMap.Data.Description)
		} else {
			assert.Equal(t, c.errMsg, respMap.Message)
		}
	}

}

// func TestGetAllProjects(t *testing.T) {
// 	seedProjectTable()
// 	req, _ := http.NewRequest("POST", "/projects/all", nil)
// 	resp := execRequest(req)
// 	var respMap []projectRespBody
// 	json.Unmarshal([]byte(resp.Body.String()), &respMap)

// 	assert.Equal(t, resp.Code, http.StatusOK)
// 	assert.Equal(t, len(respMap), 2)
// }
