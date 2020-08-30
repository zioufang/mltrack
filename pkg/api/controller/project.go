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

// CreateProject creates the entity in the database
func (s *Server) CreateProject(w http.ResponseWriter, r *http.Request) {
	m := model.Project{}
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

// GetAllProjects gets all the projects from the database
func (s *Server) GetAllProjects(w http.ResponseWriter, r *http.Request) {
	m := model.Project{}
	projects, err := m.GetAll(s.DB)
	if err != nil {
		apiutil.RespError(w, http.StatusInternalServerError, err)
		return
	}
	apiutil.RespSuccess(w, projects)
}

// GetProjectByID gets one project given an ID from the database
func (s *Server) GetProjectByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		apiutil.RespError(w, http.StatusBadRequest, err)
		return
	}
	m := model.Project{}
	projectGet, err := m.GetByID(s.DB, id)
	if err != nil {
		apiutil.RespError(w, http.StatusBadRequest, err)
		return
	}
	apiutil.RespSuccess(w, projectGet)
}

// GetProjectByParam gets one project from the database, expects 'id' or 'name' from url param
func (s *Server) GetProjectByParam(w http.ResponseWriter, r *http.Request) {
	m := model.Project{}
	var projectGet *model.Project
	var err error

	// ?id= prioritized over ?name= if both are provided in the url parameter
	if idParam := r.URL.Query().Get("id"); idParam != "" {
		var id uint64
		id, err = strconv.ParseUint(idParam, 10, 64)
		if err != nil {
			apiutil.RespError(w, http.StatusBadRequest, err)
			return
		}
		projectGet, err = m.GetByID(s.DB, id)
	} else if name := r.URL.Query().Get("name"); name != "" {
		projectGet, err = m.GetByName(s.DB, name)
	} else {
		apiutil.RespError(w, http.StatusUnprocessableEntity, errors.New("Need to provide the parameter 'id' or 'name'"))
		return
	}

	// if no error from retrieving projectGet
	if err != nil {
		apiutil.RespError(w, http.StatusBadRequest, err)
		return
	}
	apiutil.RespSuccess(w, projectGet)

}

// UpdateProjectByID updates the project name by ID
func (s *Server) UpdateProjectByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		apiutil.RespError(w, http.StatusBadRequest, err)
		return
	}
	m := model.Project{}
	// TODO validation requires name field to be non-empty, not suitable for update
	err = apiutil.ReadReqBodyWithoutValidate(w, r, s.DB, &m)
	if err != nil {
		apiutil.RespError(w, http.StatusUnprocessableEntity, err)
		return
	}
	var updated *model.Project
	updated, err = m.UpdateByID(s.DB, id)
	if err != nil {
		apiutil.RespError(w, http.StatusInternalServerError, err)
		return
	}
	apiutil.RespSuccess(w, updated)
}

// DeleteProjectByID deletes a project by ID from database
func (s *Server) DeleteProjectByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		apiutil.RespError(w, http.StatusBadRequest, err)
		return
	}
	m := model.Project{}
	err = m.DeleteByID(s.DB, id)
	if err != nil {
		apiutil.RespError(w, http.StatusBadRequest, err)
		return
	}
	apiutil.RespSuccessWithMessage(w, fmt.Sprintf("id %d deleted", id), "")
}
