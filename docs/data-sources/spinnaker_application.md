# spinnaker_application data source

Use this data source to get information about the spinnaker application.

### Example Usage
```terraform
data "spinnaker_application" "sample-app" {
  application = "sample-app"
}
```

### Argument Reference
The following arguments are supported:
* `application` - (Required) The name of application use to find the spinnaker application in Spin GATE api.

### Attribute Reference
* `application` - The name of the Spinnaker application.
* `id` - The unique identifier of Spinnaker application.
