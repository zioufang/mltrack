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

// GetProjectWithQuery gets one project from the database, expects 'id' or 'name' form url param
func (s *Server) GetProjectWithQuery(w http.ResponseWriter, r *http.Request) {
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

// DeleteProject deletes a project from the database
// func (s *Server) DeleteProject(w http.ResponseWriter, r *http.Request) {
// 	m := model.Project{}
// 	var err error

// 	// ?id= prioritized over ?name= if both are provided in the url parameter
// 	if idParam := r.URL.Query().Get("id"); idParam != "" {
// 		var id uint64
// 		id, err = strconv.ParseUint(idParam, 10, 64)
// 		if err != nil {
// 			apiutil.RespError(w, http.StatusBadRequest, err)
// 			return
// 		}
// 		err = m.DeleteByID(s.DB, id)
// 	} else if name := r.URL.Query().Get("name"); name != "" {
// 		err = m.DeleteByName(s.DB, name)
// 	} else {
// 		apiutil.RespError(w, http.StatusUnprocessableEntity, errors.New("Need to provide the parameter 'id' or 'name'"))
// 		return
// 	}

// 	if err != nil {
// 		apiutil.RespError(w, http.StatusBadRequest, err)
// 		return
// 	}
// 	apiutil.RespSuccess(w, "record deleted")
// }
