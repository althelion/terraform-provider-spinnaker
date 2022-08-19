# spinnaker_application data source

Use this data source to get information about the spinnaker application.

### Example Usage
```terraform
data "spinnaker_application" "sample-app" {
  name = "sample-app"
}
```

### Argument Reference
The following arguments are supported:
* `name` - (Required) The name of applicate use to find the spinnaker application in Spin GATE api.

### Attribute Reference
* `name` - The name of the Spinnaker application.
* `id` - The unique identifier of Spinnaker application.
