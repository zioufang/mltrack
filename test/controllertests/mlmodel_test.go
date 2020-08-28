package controllertests

// func TestGetModel(t *testing.T) {
// 	testSamples := []struct {
// 		inID          string
// 		inName        string
// 		expID         uint64
// 		expName       string
// 		expStatusCode int
// 		expErrMsg     string
// 	}{
// 		{
// 			inID:          "1",
// 			inName:        "",
// 			expID:         1,
// 			expName:       "testmodel",
// 			expStatusCode: 200,
// 			expErrMsg:     "",
// 		},
// 		// {
// 		// inID         :,
// 		// inName       :,
// 		// expID         :,
// 		// expName       :,
// 		// expStatusCode :,
// 		// expErrMsg     :,
// 		// },
// 	}
// 	req, _ := http.NewRequest("GET", "/models", nil)
// 	q := req.URL.Query()

// 	for _, v := range testSamples {
// 		if v.inID != "" {
// 			q.Add("id", v.inID)
// 		}
// 		if v.inName != "" {
// 			q.Add("name", v.inName)
// 		}
// 		req.URL.RawQuery = q.Encode()

// 		execRequest(req)
// 	}
// }
