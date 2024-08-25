package model

type Secret struct {
	Id             string `json:"id"`
	Description    string `json:"description,omitempty"`
	CreatedAt      string `json:"createdAt,omitempty"`
	CreatedBy      string `json:"createdBy,omitempty"`
	LastModifiedAt string `json:"lastModifiedAt,omitempty"`
	LastModifiedBy string `json:"lastModifiedBy,omitempty"`
}
