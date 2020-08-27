package controllertests

import (
	"net/http"
	"testing"
)

func clearModelTable() {
	server.DB.Exec("DELETE FROM model;")
}

func seedModelTable(){
    clearModelTable()
    db := server.DB
    db.Exec(sql string, values ...interface{})
}

func TestGetModel(t *testing.T) {
	clearModelTable()
	testSamples := []struct {
		inID          string
		inName        string
		expID         uint64
		expName       string
		expStatusCode int
		expErrMsg     string
	}{
		{
			inID:          "1",
			inName:        "",
			expID:         1,
			expName:       "testmodel",
			expStatusCode: 200,
			expErrMsg:     "",
		},
		// {
		// inID         :,
		// inName       :,
		// expID         :,
		// expName       :,
		// expStatusCode :,
		// expErrMsg     :,
		// },
	}
	req, _ := http.NewRequest("GET", "/models", nil)
	q := req.URL.Query()

	for _, v := range testSamples {
		if v.id != "" {
			q.Add("id", v.id)
		}
		if v.name != "" {
			q.Add("name", v.name)
		}
		req.URL.RawQuery = q.Encode()

		execRequest(req)
	}
}
