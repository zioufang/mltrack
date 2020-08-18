package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/zioufang/mltrackapi/pkg/api/apiutil"
	"github.com/zioufang/mltrackapi/pkg/api/model"
)

// CreateModel creates the entity in the database
func (s *Server) CreateModel(w http.ResponseWriter, r *http.Request) {
	m := model.Model{}
	err := apiutil.ReadReqBody(w, r, &m)
	if err != nil {
		apiutil.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = m.CreateModel(s.db)
	if err != nil {
		apiutil.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	apiutil.ResponseJSON(w, m)
}

// GetAllModels gets all the models from the database
func (s *Server) GetAllModels(w http.ResponseWriter, r *http.Request) {
	m := model.Model{}
	models, err := m.GetAllModels(s.db)
	if err != nil {
		apiutil.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	log.Println(models)
	apiutil.ResponseJSON(w, models)
}

// GetModelByID gets one model given an ID from the database
func (s *Server) GetModelByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		apiutil.ResponseError(w, http.StatusBadRequest, err)
		return
	}
	m := model.Model{}
	modelGet, err := m.GetModelByID(s.db, uid)
	if err != nil {
		apiutil.ResponseError(w, http.StatusBadRequest, err)
		return
	}
	apiutil.ResponseJSON(w, modelGet)
}

// DeleteModelByID deletes one model given an ID from the database
func (s *Server) DeleteModelByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		apiutil.ResponseError(w, http.StatusBadRequest, err)
		return
	}
	m := model.Model{}
	err = m.DeleteModelByID(s.db, uid)
	if err != nil {
		apiutil.ResponseError(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", uid))
	w.WriteHeader(http.StatusNoContent)
	apiutil.ResponseJSON(w, "")
}
