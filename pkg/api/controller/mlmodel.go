package controller

import (
	"errors"
	"log"
	"net/http"
	"strconv"

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

// GetModel gets one model from the database, expects 'id' or 'name' form url param
func (s *Server) GetModel(w http.ResponseWriter, r *http.Request) {
	m := model.Model{}
	var modelGet *model.Model
	var err error

	// ?id= prioritized over ?name= if both are provided in the url parameter
	if idParam := r.URL.Query().Get("id"); idParam != "" {
		var id uint64
		id, err = strconv.ParseUint(idParam, 10, 64)
		if err != nil {
			apiutil.RespError(w, http.StatusBadRequest, err)
			return
		}
		modelGet, err = m.GetByID(s.db, id)
	} else if name := r.URL.Query().Get("name"); name != "" {
		modelGet, err = m.GetByName(s.db, name)
	} else {
		apiutil.RespError(w, http.StatusUnprocessableEntity, errors.New("Need to provide the parameter 'id' or 'name'"))
		return
	}

	// if no error from retrieving modelGet
	if err != nil {
		apiutil.RespError(w, http.StatusBadRequest, err)
		return
	}
	apiutil.RespSuccess(w, modelGet)

}

// DeleteModel deletes a model from the database
func (s *Server) DeleteModel(w http.ResponseWriter, r *http.Request) {
	m := model.Model{}
	var err error

	// ?id= prioritized over ?name= if both are provided in the url parameter
	if idParam := r.URL.Query().Get("id"); idParam != "" {
		var id uint64
		id, err = strconv.ParseUint(idParam, 10, 64)
		if err != nil {
			apiutil.RespError(w, http.StatusBadRequest, err)
			return
		}
		err = m.DeleteByID(s.db, id)
	} else if name := r.URL.Query().Get("name"); name != "" {
		err = m.DeleteByName(s.db, name)
	} else {
		apiutil.RespError(w, http.StatusUnprocessableEntity, errors.New("Need to provide the parameter 'id' or 'name'"))
		return
	}

	if err != nil {
		apiutil.RespError(w, http.StatusBadRequest, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	apiutil.RespSuccess(w, "")
}
