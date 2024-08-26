package cmd

import (
	"encoding/json"
	"github.com/olekukonko/tablewriter"
	"github.com/previder/vault-cli/pkg"
	"github.com/spf13/cobra"
	"log"
	"os"
	"reflect"
	"strconv"
)

var (
	baseUri          string
	token            string
	outputType       string
	verbose          bool
	vaultClient      *pkg.VaultClient
	validOutputTypes = []string{"json", "pretty"}
)

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

var rootCmd = &cobra.Command{
	Use:   "vault-cli",
	Short: "A client for the Previder Vault",
	Long:  `vault.previder.io`,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&baseUri, "baseuri", "u", pkg.DefaultBaseUri, "The URL to call the Previder Vault")
	rootCmd.PersistentFlags().StringVarP(&token, "token", "t", "", "The authentication token")
	rootCmd.PersistentFlags().StringVarP(&outputType, "output", "o", "json", "Output format [pretty / json]")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output")
}

func setupClient(cmd *cobra.Command, args []string) {
	if !contains(validOutputTypes, outputType) {
		log.Fatal("Invalid output type")
	}

	if token == "" {
		if os.Getenv("VAULT_TOKEN") != "" {
			token = os.Getenv("VAULT_TOKEN")
		} else {
			log.Fatal("No token found")
		}
	}

	client, err := pkg.NewVaultClient(baseUri, token)
	if err != nil {
		log.Println("Error while instancing client", err.Error())
		return
	}
	vaultClient = client
}

func printJson(input any) {
	marshal, _ := json.MarshalIndent(input, "", "  ")
	println(string(marshal))
}

func printTable(headers []string, content []interface{}) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(headers)

	for _, row := range content {
		var rowValues []string
		rowVal := reflect.ValueOf(row)
		for _, header := range headers {
			field := rowVal.FieldByName(header)
			if !field.IsValid() {
				rowValues = append(rowValues, "")
				continue
			}

			var value string
			switch field.Kind() {
			case reflect.Bool:
				value = strconv.FormatBool(field.Bool())
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				value = strconv.FormatInt(field.Int(), 10)
			case reflect.Float32, reflect.Float64:
				value = strconv.FormatFloat(field.Float(), 'f', -1, 64)
			case reflect.String:
				value = field.String()
			default:
				value = "Unsupported Type"
			}
			rowValues = append(rowValues, value)
		}
		table.Append(rowValues)
	}
	table.Render()
}
