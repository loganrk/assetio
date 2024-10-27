package response

import (
	"assetio/internal/domain"
	"encoding/json"
	"net/http"
)

type response struct {
	Status  int        `json:"status"`
	Success bool       `json:"success"`
	Err     []errorMsg `json:"error,omitempty"`
	Data    any        `json:"data,omitempty"`
}

type errorMsg struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

func New() domain.Response {
	return &response{}
}

func (r *response) SetError(errCode string, errMsg string) {
	r.Err = append(r.Err, errorMsg{
		Code: errCode,
		Msg:  errMsg,
	})
}

func (r *response) SetStatus(status int) {
	r.Status = status
}

func (r *response) SetData(data any) {
	if data != nil {
		r.Data = data
	} else {
		r.Data = struct{}{}
	}
}

func (r *response) Send(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	if len(r.Err) > 0 {
		w.WriteHeader(r.Status)
		r.Success = false
		r.Data = struct{}{}
	} else {
		w.WriteHeader(http.StatusOK)
		r.Status = http.StatusOK
		r.Success = true
		r.Err = make([]errorMsg, 0)
	}

	json.NewEncoder(w).Encode(r)
}
