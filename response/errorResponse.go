package response

type ErrorResponse struct {
	Status     string      `json:"status"`
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Timestamp  string      `json:"timestamp"`
	Errors     interface{} `json:"errors"`
}
