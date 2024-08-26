package cmd

import (
	"errors"
	"fmt"
	"github.com/previder/vault-cli/pkg/model"
	"github.com/spf13/cobra"
)

func init() {
	var environmentCmd = &cobra.Command{
		Use:   "environment",
		Short: "Environment commands",
	}
	rootCmd.AddCommand(environmentCmd)

	var cmdList = &cobra.Command{
		Use:    "list",
		Short:  "Get a list of environments",
		Args:   cobra.NoArgs,
		RunE:   listEnvironment,
		PreRun: setupClient,
	}
	environmentCmd.AddCommand(cmdList)

	var cmdGet = &cobra.Command{
		Use:        "get [id]",
		Short:      "Get an environment",
		Args:       cobra.ExactArgs(1),
		ArgAliases: []string{"id"},
		RunE:       getEnvironment,
		PreRun:     setupClient,
	}

	environmentCmd.AddCommand(cmdGet)

	var cmdCreate = &cobra.Command{
		Use:    "create",
		Short:  "Create a new environment and receive the environment token",
		RunE:   createEnvironment,
		PreRun: setupClient,
	}
	environmentCmd.AddCommand(cmdCreate)

	cmdCreate.Flags().StringP("name", "n", "", "Name of the environment")
	cmdCreate.Flags().StringP("contact", "c", "", "Contact email of the environment")
	cmdCreate.Flags().BoolP("active", "a", true, "Activate the environment")

	var cmdDelete = &cobra.Command{
		Use:        "delete [id]",
		Short:      "Delete an environment",
		Args:       cobra.ExactArgs(1),
		ArgAliases: []string{"id"},
		RunE:       deleteEnvironment,
		PreRun:     setupClient,
	}

	environmentCmd.AddCommand(cmdDelete)
}

func listEnvironment(cmd *cobra.Command, args []string) error {
	content, err := vaultClient.GetEnvironments()
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
		printTable([]string{"Id", "Name", "Contact", "Active"}, intContent)
	} else {
		printJson(content)
	}
	return nil
}

func getEnvironment(cmd *cobra.Command, args []string) error {
	content, err := vaultClient.GetEnvironment(args[0])
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

func createEnvironment(cmd *cobra.Command, args []string) error {
	create := model.EnvironmentCreate{}
	name, err := cmd.Flags().GetString("name")
	if err != nil {
		return err
	}
	if name == "" {
		return errors.New("name cannot be empty")
	}
	create.Name = name

	contact, err := cmd.Flags().GetString("contact")
	if err != nil {
		return err
	}
	if contact == "" {
		return errors.New("contact cannot be empty")
	}
	create.Contact = contact

	active, err := cmd.Flags().GetBool("active")
	if err != nil {
		return err
	}
	create.Active = active

	createResponse, err := vaultClient.CreateEnvironment(create)
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

func deleteEnvironment(cmd *cobra.Command, args []string) error {
	err := vaultClient.DeleteEnvironment(args[0])
	if err != nil {
		return err
	}

	return nil
}
