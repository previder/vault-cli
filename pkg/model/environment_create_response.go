package model

type EnvironmentCreateResponse struct {
	Environment
	Secret string `json:"secret"`
}
