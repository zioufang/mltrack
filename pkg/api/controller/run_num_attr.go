package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/zioufang/mltrackapi/pkg/api/apiutil"
	"github.com/zioufang/mltrackapi/pkg/api/model"
)

// CreateRunNumAttr creates the entity in the database
func (s *Server) CreateRunNumAttr(w http.ResponseWriter, r *http.Request) {
	m := model.RunNumAttr{}
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

// GetRunNumAttrListByParam gets a list of RunNumAttrs
func (s *Server) GetRunNumAttrListByParam(w http.ResponseWriter, r *http.Request) {
	var RunNumAttrsGet *[]model.RunNumAttr
	m := model.RunNumAttr{}

	modelRunIDParam := r.URL.Query().Get("model_run_id")
	if modelRunIDParam == "" {
		apiutil.RespError(w, http.StatusUnprocessableEntity, errors.New("Need to provide the parameter 'model_run_id'"))
		return
	}
	modelRunID, err := strconv.ParseUint(modelRunIDParam, 10, 64)
	if err != nil {
		apiutil.RespError(w, http.StatusBadRequest, err)
		return
	}

	name := r.URL.Query().Get("name")
	category := r.URL.Query().Get("category")
	RunNumAttrsGet, err = m.Get(s.DB, modelRunID, name, category)
	if err != nil {
		apiutil.RespError(w, http.StatusBadRequest, err)
		return
	}
	apiutil.RespSuccess(w, RunNumAttrsGet)
}
