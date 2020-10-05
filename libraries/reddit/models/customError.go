package models

// CustomRedditError : Custom Reddit Error
type CustomRedditError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
