# spinnaker_pipeline Data Source

Use this data source to get information about the spinnaker application pipeline.

## Example Usage
```terraform
data "spinnaker_pipeline" "sample-app-pipeline" {
  application = "sample-app"
  name        = "sample-pipeline"
}
```

## Argument Reference
The following arguments are supported:
- `application` - (Required) The application name use to find the Spinnaker application in Spin GATE api.
- `name` - (Required) The Spinnaker application pipeline name.

## Attribute Reference
* `application` - The name of the Spinnaker application.
* `id` - The Spinnaker application pipeline unique identifier.
* `pipeline` - The pipeline content in JSON format
* `pipeline_id` - The Spinnamer application pipeline unique identifier.
