package controllertests

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/zioufang/mltrackapi/pkg/api/controller"
)

const testDBName string = "test.db"

var server = controller.Server{}

func TestMain(m *testing.M) {
	os.Exit(testMainWrapper(m))
}

func testMainWrapper(m *testing.M) int {
	server.Init("sqlite3", testDBName)
	defer os.Remove(testDBName)
	testRun := m.Run()
	return testRun

}

func execRequest(req *http.Request) *httptest.ResponseRecorder {
	rec := httptest.NewRecorder()
	server.Router.ServeHTTP(rec, req)
	return rec
}
