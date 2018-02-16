package dto

// TokenResponse represents the positive response of token request
type TokenResponse struct {
	Name       string `json:"name"`
	Token      string `json:"token"`
	LastAccess string `json:"last_access"`
}
