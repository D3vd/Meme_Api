package reddit

import "encoding/base64"

func (r *Reddit) EncodeCredentials() (encodedCredentials string) {
	data := r.ClientID + ":" + r.ClientSecret
	encodedCredentials = base64.StdEncoding.EncodeToString([]byte(data))
	return
}
