package v1

type ErrorResponse struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}
