package Interfaces

type HTTPResponse struct {
	Status			string			`json:"status"`
	Data			interface{} 	`json:"data,omitempty"`
	ErrorMessage	interface{}		`json:"errorMessage"`
}
func CreateHTTPResponsePayload(data interface{}, status string, errorMessage interface{}) *HTTPResponse {
	return &HTTPResponse{
		Status: 		status,
		Data: 			data,
		ErrorMessage: 	errorMessage,
	}
}