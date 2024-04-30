package core_services_response

const (
	StatusError = "error"
)

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

func Error(message string) Response {
	return Response{
		Status: StatusError,
		Error:  message,
	}
}
