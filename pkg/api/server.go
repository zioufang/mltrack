package api

import (
	"github.com/zioufang/mltrackapi/pkg/api/controller"
)

// Run runs the server
func Run() {
	var server = controller.Server{}
	server.Init("sqlite3", "test.db")

	var port uint = 8000
	server.Run(port)
}
