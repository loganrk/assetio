package port

import "net/http"

type Response interface {
	SetError(errCode string, errMsg string)
	SetStatus(status int)
	SetData(data any)
	Send(w http.ResponseWriter)
}
