package pkg

import (
	"fmt"
	"github.com/previder/vault-cli/pkg/model"
)

func (v *VaultClient) GetSecrets() ([]model.Secret, error) {
	var result []model.Secret
	err := v.request("GET", "/secret", nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (v *VaultClient) GetSecret(id string) (*model.Secret, error) {
	var result model.Secret
	err := v.request("GET", fmt.Sprintf("/secret/%v", id), nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (v *VaultClient) CreateSecret(create model.SecretCreate) (*model.Secret, error) {
	var result model.Secret
	err := v.request("POST", "/secret", create, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (v *VaultClient) DecryptSecret(id string) (*model.SecretDecrypt, error) {
	var result model.SecretDecrypt
	err := v.request("GET", fmt.Sprintf("/secret/%v/decrypt", id), nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
