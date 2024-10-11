package response

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
