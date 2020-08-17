package controller

import (
	"fmt"
	"log"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
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
	s.router.HandleFunc("/models", s.CreateModel).Methods("POST")
}
