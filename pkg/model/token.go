package model

type Token struct {
	Id            string `json:"id"`
	Description   string `json:"description,omitempty"`
	EnvironmentId string `json:"environmentId,omitempty"`
	CreatedAt     string `json:"createdAt,omitempty"`
	CreatedBy     string `json:"createdBy,omitempty"`
	ExpiresAt     string `json:"expiresAt,omitempty"`
	TokenType     string `json:"tokenType,omitempty"`
}
