package controllertests

import (
	"log"
	"os"
	"testing"

	"github.com/zioufang/mltrackapi/pkg/api/controller"
)

const testDBName string = "test.db"

var s controller.Server

func TestMain(m *testing.M) {
	s = controller.Server{}
	s.Init("sqlite3", testDBName)
	testRun := m.Run()
	err := os.Remove(testDBName)
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(testRun)
}
