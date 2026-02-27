package model

import "fmt"

type Secret struct {
	Id             string `json:"id"`
	Description    string `json:"description,omitempty"`
	CreatedAt      string `json:"createdAt,omitempty"`
	CreatedBy      string `json:"createdBy,omitempty"`
	LastModifiedAt string `json:"lastModifiedAt,omitempty"`
	LastModifiedBy string `json:"lastModifiedBy,omitempty"`
}

func (s Secret) String() string {
	return fmt.Sprintf(
		"Id: %v\nDescription: %v\nCreatedAt: %v\nCreatedBy: %v\nLastModifiedAt: %v\nLastModifiedBy: %v",
		s.Id,
		s.Description,
		s.CreatedAt,
		s.CreatedBy,
		s.LastModifiedAt,
		s.LastModifiedBy,
	)
}
