package cmd

import (
	"errors"
	"fmt"
	"github.com/gkwmiddelkamp/vault-cli/pkg/model"
	"github.com/spf13/cobra"
)

func init() {
	var secretCmd = &cobra.Command{
		Use:   "secret",
		Short: "Secret commands",
	}
	rootCmd.AddCommand(secretCmd)

	var cmdList = &cobra.Command{
		Use:    "list",
		Short:  "Get a list of secrets",
		Args:   cobra.NoArgs,
		RunE:   listSecret,
		PreRun: setupClient,
	}
	secretCmd.AddCommand(cmdList)

	var cmdGet = &cobra.Command{
		Use:        "get [id]",
		Short:      "Get a secret",
		Args:       cobra.ExactArgs(1),
		ArgAliases: []string{"id"},
		RunE:       getSecret,
		PreRun:     setupClient,
	}

	secretCmd.AddCommand(cmdGet)

	var cmdCreate = &cobra.Command{
		Use:    "create",
		Short:  "Create a new secret",
		RunE:   createSecret,
		PreRun: setupClient,
	}
	secretCmd.AddCommand(cmdCreate)

	cmdCreate.Flags().StringP("description", "d", "", "Description of the secret")
	cmdCreate.Flags().StringP("secret", "s", "", "The secret to store")

	var cmdDecrypt = &cobra.Command{
		Use:        "decrypt [id]",
		Short:      "Get the decrypted secret",
		Args:       cobra.ExactArgs(1),
		ArgAliases: []string{"id"},
		RunE:       decryptSecret,
		PreRun:     setupClient,
	}

	secretCmd.AddCommand(cmdDecrypt)
}

func listSecret(cmd *cobra.Command, args []string) error {
	content, err := vaultClient.GetSecrets()
	if err != nil {
		return err
	}
	if outputType == "pretty" {
		intContent := make([]interface{}, len(content))
		for idx, row := range content {
			intContent[idx] = row
		}
		if verbose {
			fmt.Printf("%+v\n", intContent)
		}
		printTable([]string{"Id", "Description"}, intContent)
	} else {
		printJson(content)
	}
	return nil
}

func getSecret(cmd *cobra.Command, args []string) error {
	content, err := vaultClient.GetSecret(args[0])
	if err != nil {
		return err
	}
	if outputType == "pretty" {
		fmt.Printf("%+v\n", content)
	} else {
		printJson(content)
	}
	return nil
}

func createSecret(cmd *cobra.Command, args []string) error {
	create := model.SecretCreate{}
	description, err := cmd.Flags().GetString("description")
	if err != nil {
		return err
	}
	if description == "" {
		return errors.New("description cannot be empty")
	}
	create.Description = description

	secret, err := cmd.Flags().GetString("secret")
	if err != nil {
		return err
	}
	if secret == "" {
		return errors.New("secret cannot be empty")
	}
	create.Secret = secret

	createResponse, err := vaultClient.CreateSecret(create)
	if err != nil {
		return err
	}

	if outputType == "pretty" {
		fmt.Printf("%+v\n", createResponse)
	} else {
		printJson(createResponse)
	}

	return nil
}

func decryptSecret(cmd *cobra.Command, args []string) error {
	content, err := vaultClient.DecryptSecret(args[0])
	if err != nil {
		return err
	}
	if outputType == "pretty" {
		fmt.Printf("%+v\n", content)
	} else {
		printJson(content)
	}
	return nil
}
