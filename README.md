# Spinnaker Terraform Provider

## Spinnaker
https://spinnaker.hub.liveramp.com

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) 1.1.7
- [Go](https://golang.org/doc/install) 1.11 (to build the provider plugin)

## Building and Developing The Provider

```shell
git clone git@github.com:althelion/terraform-provider-spinnaker.git
cd terraform-provider-spinnaker/
go build
go test./...
```

## Building the provider
If you're building the provider, follow the instructions to [install it as a plugin.](https://www.terraform.io/docs/plugins/basics.html#installing-a-plugin) 

## Using the provider
After placing it into your plugins directory, initialize the your state file.
```shell
terraform init
```
