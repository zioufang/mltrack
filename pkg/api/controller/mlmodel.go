package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/zioufang/mltrackapi/pkg/api/apiutil"
	"github.com/zioufang/mltrackapi/pkg/api/model"
)

// CreateModel creates the entity in the database
func (s *Server) CreateModel(w http.ResponseWriter, r *http.Request) {
	m := model.Model{}
	err := apiutil.ReadReqBody(w, r, s.db, &m)
	if err != nil {
		apiutil.RespError(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = m.Create(s.db)
	if err != nil {
		apiutil.RespError(w, http.StatusInternalServerError, err)
		return
	}
	apiutil.RespSuccess(w, m)
}

// GetAllModels gets all the models from the database
func (s *Server) GetAllModels(w http.ResponseWriter, r *http.Request) {
	m := model.Model{}
	models, err := m.GetAll(s.db)
	if err != nil {
		apiutil.RespError(w, http.StatusInternalServerError, err)
		return
	}
	log.Println(models)
	apiutil.RespSuccess(w, models)
}

// GetModelByID gets one model given an ID from the database
func (s *Server) GetModelByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		apiutil.RespError(w, http.StatusBadRequest, err)
		return
	}
	m := model.Model{}
	modelGet, err := m.GetByID(s.db, id)
	if err != nil {
		apiutil.RespError(w, http.StatusBadRequest, err)
		return
	}
	apiutil.RespSuccess(w, modelGet)
}

// DeleteModelByID deletes one model given an ID from the database
func (s *Server) DeleteModelByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		apiutil.RespError(w, http.StatusBadRequest, err)
		return
	}
	m := model.Model{}
	err = m.DeleteByID(s.db, id)
	if err != nil {
		apiutil.RespError(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", id))
	w.WriteHeader(http.StatusNoContent)
	apiutil.RespSuccess(w, "")
}
