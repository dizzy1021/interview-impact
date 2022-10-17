package util

type APIResponse struct {
	Message *string			`json:"message"`
	Data	*interface{}	`json:"data"`
	Code	int				`json:"code"`
}

func NewAPIResponse(data *interface{}, message *string, code int) APIResponse {
	response := APIResponse{
		Data	: data,
		Message	: message,
		Code	: code,
	}

	return response
}