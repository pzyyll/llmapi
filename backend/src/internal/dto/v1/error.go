package v1

type SuccessResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}
