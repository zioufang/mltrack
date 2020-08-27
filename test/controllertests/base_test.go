package controllertests

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/zioufang/mltrackapi/pkg/api/controller"
)

const testDBName string = "test.db"

var server = controller.Server{}

func TestMain(m *testing.M) {
	s.Init("sqlite3", testDBName)
	testRun := m.Run()
	err := os.Remove(testDBName)
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(testRun)

}

func execRequest(req *http.Request) *httptest.ResponseRecorder {
	rec := httptest.NewRecorder()
	s.Router.ServeHTTP(rec, req)
	return rec
}
