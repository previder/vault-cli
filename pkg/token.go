package pkg

import (
	"fmt"
	"github.com/previder/vault-cli/pkg/model"
)

func (v *VaultClient) GetTokens() ([]model.Token, error) {
	var result []model.Token
	err := v.request("GET", "/token", nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (v *VaultClient) GetToken(id string) (*model.Token, error) {
	var result model.Token
	err := v.request("GET", fmt.Sprintf("/token/%v", id), nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (v *VaultClient) CreateToken(create model.TokenCreate) (*model.TokenCreateResponse, error) {
	var result model.TokenCreateResponse
	err := v.request("POST", "/token", create, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
