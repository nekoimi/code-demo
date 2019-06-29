package types

type JsonResponse struct {
	Code int64							`json:"code"`
	Data map[interface{}]interface{}	`json:"data"`
	Message string						`json:"message"`
}

type ErrorResponse JsonResponse

var (
	RequestValidationException = ErrorResponse{
		400,
		map[interface{}]interface{}{},
		"Request Validation Exception",
	}
)
