package controller

import (
	"log"
	"net/http"

	"github.com/zioufang/mltrackapi/pkg/api/apiutil"
	"github.com/zioufang/mltrackapi/pkg/api/model"
)

// CreateModel creates the entity in the database
func (s *Server) CreateModel(w http.ResponseWriter, r *http.Request) {
	m := model.Model{}
	apiutil.ReadReqBody(w, r, &m)
	log.Println(m)
	err := m.Create(s.db)
	if err != nil {
		apiutil.HTTPError(w, http.StatusInternalServerError, err)
	}
	apiutil.ResponseJSON(w, m)
}
