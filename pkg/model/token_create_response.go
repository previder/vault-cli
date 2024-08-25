package model

type TokenCreateResponse struct {
	Token
	Secret string `json:"secret"`
}
