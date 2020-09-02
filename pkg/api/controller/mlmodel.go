package controller

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/zioufang/mltrackapi/pkg/api/apiutil"
	"github.com/zioufang/mltrackapi/pkg/api/model"
)

// CreateModel creates the entity in the database
func (s *Server) CreateModel(w http.ResponseWriter, r *http.Request) {
	m := model.Model{}
	err := apiutil.ReadReqBody(w, r, s.DB, &m)
	if err != nil {
		apiutil.RespError(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = m.Create(s.DB)
	if err != nil {
		apiutil.RespError(w, http.StatusInternalServerError, err)
		return
	}
	apiutil.RespSuccess(w, m)
}

// GetModelByID gets one model given an ID from the database
func (s *Server) GetModelByID(w http.ResponseWriter, r *http.Request) {
	m := model.Model{}
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		apiutil.RespError(w, http.StatusBadRequest, err)
		return
	}
	modelGet, err := m.GetByID(s.DB, id)
	if err != nil {
		apiutil.RespError(w, http.StatusBadRequest, err)
		return
	}
	apiutil.RespSuccess(w, modelGet)
}

// GetModelByParam gets one model from the database, expects 'id' or 'name' form url param
func (s *Server) GetModelByParam(w http.ResponseWriter, r *http.Request) {
	var modelGet *model.Model
	var err error
	m := model.Model{}

	// ?id= prioritized over ?name= if both are provided in the url parameter
	if idParam := r.URL.Query().Get("id"); idParam != "" {
		var id uint64
		id, err = strconv.ParseUint(idParam, 10, 64)
		if err != nil {
			apiutil.RespError(w, http.StatusBadRequest, err)
			return
		}
		modelGet, err = m.GetByID(s.DB, id)
	} else if name := r.URL.Query().Get("name"); name != "" {
		modelGet, err = m.GetByName(s.DB, name)
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

// UpdateModelByID updates the model instance by ID
func (s *Server) UpdateModelByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		apiutil.RespError(w, http.StatusBadRequest, err)
		return
	}
	m := model.Model{}
	// TODO validation requires name field to be non-empty, not suitable for update
	err = apiutil.ReadReqBodyWithoutValidate(w, r, s.DB, &m)
	if err != nil {
		apiutil.RespError(w, http.StatusUnprocessableEntity, err)
		return
	}
	var updated *model.Model
	updated, err = m.UpdateByID(s.DB, id)
	if err != nil {
		apiutil.RespError(w, http.StatusInternalServerError, err)
		return
	}
	apiutil.RespSuccess(w, updated)
}

// DeleteModelByID deletes a model by ID from database
func (s *Server) DeleteModelByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		apiutil.RespError(w, http.StatusBadRequest, err)
		return
	}
	m := model.Model{}
	err = m.DeleteByID(s.DB, id)
	if err != nil {
		apiutil.RespError(w, http.StatusBadRequest, err)
		return
	}
	apiutil.RespSuccessWithMessage(w, fmt.Sprintf("id %d deleted", id), "")
}

// GetAllModels gets all the models from the database
func (s *Server) GetAllModels(w http.ResponseWriter, r *http.Request) {
	m := model.Model{}
	models, err := m.GetAll(s.DB)
	if err != nil {
		apiutil.RespError(w, http.StatusInternalServerError, err)
		return
	}
	apiutil.RespSuccess(w, models)
}

// GetModelListByParam gets a list of models by the project_id
func (s *Server) GetModelListByParam(w http.ResponseWriter, r *http.Request) {
	var modelsGet *[]model.Model
	m := model.Model{}

	projectIDParam := r.URL.Query().Get("project_id")
	if projectIDParam == "" {
		apiutil.RespError(w, http.StatusUnprocessableEntity, errors.New("Need to provide the parameter 'project_id'"))
		return
	}
	projectID, err := strconv.ParseUint(projectIDParam, 10, 64)
	if err != nil {
		apiutil.RespError(w, http.StatusBadRequest, err)
		return
	}
	modelsGet, err = m.GetByProjectID(s.DB, projectID)
	if err != nil {
		apiutil.RespError(w, http.StatusBadRequest, err)
		return
	}
	apiutil.RespSuccess(w, modelsGet)
}
