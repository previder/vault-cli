<div style="text-align: center">
<img src="https://vault.pnck.nl/cdn/logo.png" width="250" alt=""/>
</div>

Vault-cli is a project to have a light-weight, secure and multi-tenant solution for encrypted password storage. Is uses a the Vault Rest API where you can manage your environments, tokens and secrets.

**Release:**

[![Release Version](https://img.shields.io/github/v/release/gkwmiddelkamp/vault?label=vault)](https://github.com/previder/vault-cli/releases/latest)

**Last build:**

![Last build](https://github.com/previder/vault-cli/actions/workflows/go.yml/badge.svg)

**Last release:**

![Last publish](https://github.com/previder/vault-cli/actions/workflows/goreleaser.yml/badge.svg)

# Environments
Security is key in the project. You can create separate environments for your projects or customers. All environments use unique encryption keys, which are never stored in the database and are only available to the customer.

The MasterAdmin token can create an Environment. As a response to this call an EnvironmentAdmin token is returned once. This type of token can be used to create ReadWrite or ReadOnly tokens. Read the section [Tokens](#Tokens) for more detailed view of the different token types.

# Tokens
There are 4 types of tokens, each having its own purpose.

|                                  | MasterAdmin   	 | EnvironmentAdmin  	 | ReadWrite  	 | ReadOnly   	 |
|----------------------------------|-----------------|---------------------|--------------|--------------|
| Create MasterAdmin token	        | 	     ✅         | 	                   | 	            | 	            |
| Create EnvironmentAdmin token	   | 	 ✅              | 	                   | 	            | 	            |
| Create ReadWrite/ReadOnly token	 | 	               | 	      ✅            | 	            | 	            |
| Manage environments              | 	     ✅          | 	                   | 	            | 	            |
| Manage secrets	                  | 	               | 	                   | 	  ✅          | 	            |
| Get decrypted secret             | 	               | 	                   | 	   ✅         | 	    ✅        |


# Getting started
Vault-cli is a stand-alone binary to use with the Vault API. 

To see all usages, run
```shell
./vault-cli --help
```

## Token
Use the token directly from the command-line or define the VAULT_TOKEN environment variable.

## Usage example
```shell
./vault-cli -t <insert-token> secret list
```
Will print all secrets in the Vault environment

```shell
export VAULT_TOKEN="insert-token"
./vault-cli secret decode <yoursecret>
```
To get the decrypted secret back to use in an application.

## Output
The default output format is `json`. Lists of environments, tokens and secrets can also be pretty-printed with the `-o pretty` parameter.
