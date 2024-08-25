package model

type Environment struct {
	Id             string `json:"id"`
	Name           string `json:"name,omitempty"`
	Contact        string `json:"contact,omitempty"`
	Active         bool   `json:"active,omitempty"`
	CreatedAt      string `json:"createdAt,omitempty"`
	CreatedBy      string `json:"createdBy,omitempty"`
	LastModifiedAt string `json:"lastModifiedAt,omitempty"`
	LastModifiedBy string `json:"lastModifiedBy,omitempty"`
}
