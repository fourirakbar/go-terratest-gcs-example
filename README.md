# Terratest Example

## Description

Simple terratest example with GCS modules
<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|:----:|:-----:|:-----:|
| force\_destroy | Force destroy when deleting bucket | string | n/a | yes |
| location | Location of the bucket | string | n/a | yes |
| name | The name of the bucket | string | n/a | yes |
| project | Project of the bucket | string | n/a | yes |
| storage\_class | Storage class of the bucket | string | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| bucket\_url | Bucket URL of the bucket |
| force\_destroy | Force destroy of the bucket |
| location | Location of the bucker |
| name | Name of the bucket |
| project | Project of the bucker |
| storage\_class | Storage class of the bucket |

<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK --
