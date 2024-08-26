package pkg

import (
	"fmt"
	"github.com/previder/vault-cli/pkg/model"
)

func (v *VaultClient) GetEnvironments() ([]model.Environment, error) {
	var result []model.Environment
	err := v.request("GET", "/environment", nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (v *VaultClient) GetEnvironment(id string) (*model.Environment, error) {
	var result model.Environment
	err := v.request("GET", fmt.Sprintf("/environment/%v", id), nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (v *VaultClient) CreateEnvironment(create model.EnvironmentCreate) (*model.EnvironmentCreateResponse, error) {
	var result model.EnvironmentCreateResponse
	err := v.request("POST", "/environment", create, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (v *VaultClient) DeleteEnvironment(id string) error {
	err := v.request("DELETE", fmt.Sprintf("/environment/%v", id), nil, nil)
	if err != nil {
		return err
	}
	return nil
}
