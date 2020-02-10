package models

// AccessTokenBody : Model for Reddit AccessBody Response
type AccessTokenBody struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   uint   `json:"expires_in"`
	Scope       string `json:"scope"`
}
