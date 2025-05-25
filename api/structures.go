package api

type APIError struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}
