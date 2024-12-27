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

// SetError appends an error message to the response.
// The error is represented by an errorMsg struct containing a code and a message.
func (r *response) SetError(errCode string, errMsg string) {
	// Append the error code and message to the Err field of the response
	r.Err = append(r.Err, errorMsg{
		Code: errCode, // Error code
		Msg:  errMsg,  // Error message
	})
}

// SetStatus sets the HTTP status code for the response.
// The status code is used when sending the response to indicate success or failure.
func (r *response) SetStatus(status int) {
	// Set the HTTP status code
	r.Status = status
}

// SetData sets the response data.
// If the data is non-nil, it is assigned to the Data field of the response.
// If the data is nil, an empty struct is assigned to represent no data.
func (r *response) SetData(data any) {

	if data != nil {
		// Assign the provided data to the Data field
		r.Data = data
	} else {
		// If no data is provided, assign an empty struct to the Data field
		r.Data = struct{}{}
	}
}

// Send writes the response to the HTTP ResponseWriter.
// It sets appropriate headers, status code, and encodes the response as JSON.
func (r *response) Send(w http.ResponseWriter) {
	// Set the content type of the response to "application/json"
	w.Header().Set("Content-Type", "application/json")

	// If there are errors in the response, set the status to the error status and prepare the response.
	if len(r.Err) > 0 {
		// Set the status to the previously set status code (error scenario)
		w.WriteHeader(r.Status)
		r.Success = false   // Mark the response as failed
		r.Data = struct{}{} // Clear the Data field in case of an error
	} else {
		// If no errors, send a successful response
		w.WriteHeader(http.StatusOK) // HTTP 200 OK status
		r.Status = http.StatusOK     // Set the status field to HTTP 200
		r.Success = true             // Mark the response as successful
		r.Err = make([]errorMsg, 0)  // Clear any previous error messages
	}
	// Encode the response object as JSON and send it to the client
	json.NewEncoder(w).Encode(r)
}
