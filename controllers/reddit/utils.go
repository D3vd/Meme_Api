package reddit

import "encoding/base64"

// EncodeCredentials : Return base64 Encoded client ID and Secret required for authentication
func (r *Reddit) EncodeCredentials() (encodedCredentials string) {
	data := r.ClientID + ":" + r.ClientSecret
	encodedCredentials = base64.StdEncoding.EncodeToString([]byte(data))
	return
}
