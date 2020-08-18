package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/zioufang/mltrackapi/pkg/api/middleware"
	"github.com/zioufang/mltrackapi/pkg/api/model"
)

// Server is the struct for the server
type Server struct {
	db     *gorm.DB
	router *mux.Router
}

// Init initialize the server with database and mux router
// TODO adds postgres, mysql support
func (s *Server) Init(DbDriver, DbName string) {
	var err error
	switch DbDriver {
	case "sqlite3":
		s.db, err = gorm.Open(DbDriver, DbName)
		if err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatal(fmt.Errorf("%s is not a supported database", DbDriver))
	}
	s.db.AutoMigrate(&model.Model{})
	s.router = mux.NewRouter()
	s.SetRoutes()
}

// SetRoutes sets the routs for the server
func (s *Server) SetRoutes() {
	s.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Mltrack\n")
	})
	s.router.HandleFunc("/models", middleware.SetJSONHeader(s.CreateModel)).Methods("POST")
}

// Run runs the server
func (s *Server) Run(port uint) {
	fmt.Printf("Listening to port %d\n", port)
	addr := ":" + fmt.Sprint(port)
	log.Fatal(http.ListenAndServe(addr, s.router))
}
