package response

type SuccessResponse struct {
	Status     string      `json:"status"`
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Timestamp  string      `json:"timestamp"`
	Payload    interface{} `json:"payload"`
}
