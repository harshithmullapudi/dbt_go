package models

type Status struct {
	Success bool `json:"success"`
}

type APIResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}
