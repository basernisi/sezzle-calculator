package auth

type TokenRequest struct {
	ClientID     string
	ClientSecret string
}

type TokenResponse struct {
	AccessToken string
	TokenType   string
	ExpiresIn   int64
}
