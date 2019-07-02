package types

type JsonResponse struct {
	Code    int64                  `json:"code"`
	Data    map[string]interface{} `json:"data"`
	Message string                 `json:"message"`
}

type ErrorResponse JsonResponse

var (
	RequestValidationException = ErrorResponse{
		400,
		map[string]interface{}{},
		"Request Validation Exception",
	}
)
