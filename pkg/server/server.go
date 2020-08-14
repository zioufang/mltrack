package server

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type Server struct {
	db     *gorm.DB
	router *mux.Router
}
