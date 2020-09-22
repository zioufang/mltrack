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

// CreateModelRun creates the entity in the database
func (s *Server) CreateModelRun(w http.ResponseWriter, r *http.Request) {
	m := model.ModelRun{}
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

// GetModelRunByID gets one model given an ID from the database
func (s *Server) GetModelRunByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		apiutil.RespError(w, http.StatusBadRequest, err)
		return
	}
	m := model.ModelRun{}
	modelGet, err := m.GetByID(s.DB, id)
	if err != nil {
		apiutil.RespError(w, http.StatusBadRequest, err)
		return
	}
	apiutil.RespSuccess(w, modelGet)
}

// DeleteModelRunByID deletes one model run given an ID from the database
func (s *Server) DeleteModelRunByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		apiutil.RespError(w, http.StatusBadRequest, err)
		return
	}
	m := model.ModelRun{}
	err = m.DeleteByID(s.DB, id)
	if err != nil {
		apiutil.RespError(w, http.StatusBadRequest, err)
		return
	}
	apiutil.RespSuccessWithMessage(w, fmt.Sprintf("id %d deleted", id), "")
}

// GetModelRunListByParam gets a list of model runs by the model_id
func (s *Server) GetModelRunListByParam(w http.ResponseWriter, r *http.Request) {
	var modelRunsGet *[]model.ModelRun
	m := model.ModelRun{}

	modelIDParam := r.URL.Query().Get("model_id")
	if modelIDParam == "" {
		apiutil.RespError(w, http.StatusUnprocessableEntity, errors.New("Need to provide the parameter 'model_id'"))
		return
	}
	modelID, err := strconv.ParseUint(modelIDParam, 10, 64)
	if err != nil {
		apiutil.RespError(w, http.StatusBadRequest, err)
		return
	}
	modelRunsGet, err = m.GetByModelID(s.DB, modelID)
	if err != nil {
		apiutil.RespError(w, http.StatusBadRequest, err)
		return
	}
	apiutil.RespSuccess(w, modelRunsGet)
}
