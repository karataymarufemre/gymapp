package model

type SaveResponse struct {
	Success bool   `json:"success"`
	ID      uint64 `json:"id"`
}

type TokenResponse struct {
	Token string `json:"token"`
}
