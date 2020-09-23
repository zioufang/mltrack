package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/zioufang/mltrackapi/pkg/api/apiutil"
	"github.com/zioufang/mltrackapi/pkg/api/model"
)

// CreateRunTag creates the entity in the database
func (s *Server) CreateRunTag(w http.ResponseWriter, r *http.Request) {
	m := model.RunTag{}
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

// GetRunTagListByParam taks Param in list to get a list of RunTags
func (s *Server) GetRunTagListByParam(w http.ResponseWriter, r *http.Request) {
	var RunTagsGet *[]model.RunTag
	var err error
	m := model.RunTag{}

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

	keys := r.URL.Query()["key"]
	RunTagsGet, err = m.Get(s.DB, modelRunIDs, keys)
	if err != nil {
		apiutil.RespError(w, http.StatusBadRequest, err)
		return
	}
	apiutil.RespSuccess(w, RunTagsGet)
}
