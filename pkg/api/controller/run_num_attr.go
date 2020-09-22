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

// GetRunNumAttrListByParam taks Param in list to get a list of RunNumAttrs
func (s *Server) GetRunNumAttrListByParam(w http.ResponseWriter, r *http.Request) {
	var RunNumAttrsGet *[]model.RunNumAttr
	var err error
	m := model.RunNumAttr{}

	// url param in list
	// https://golang.org/pkg/net/url/#Values
	// e.g. friend=Jess&friend=Sarah&friend=Zoe
	modelRunIDsParam := r.URL.Query()["model_run_id"]
	if len(modelRunIDsParam) == 0 {
		apiutil.RespError(w, http.StatusUnprocessableEntity, errors.New("Need to provide the parameter 'model_run_id'"))
		return
	}
	modelRunIDs := make([]uint64, len(modelRunIDsParam))
	for i := range modelRunIDsParam {
		modelRunIDs[i], err = strconv.ParseUint(modelRunIDsParam[i], 10, 64)
		if err != nil {
			apiutil.RespError(w, http.StatusBadRequest, err)
			return
		}
	}

	names := r.URL.Query()["name"]
	categories := r.URL.Query()["category"]
	RunNumAttrsGet, err = m.Get(s.DB, modelRunIDs, names, categories)
	if err != nil {
		apiutil.RespError(w, http.StatusBadRequest, err)
		return
	}
	apiutil.RespSuccess(w, RunNumAttrsGet)
}
