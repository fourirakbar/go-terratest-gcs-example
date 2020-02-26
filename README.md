# Terratest Example

## Description

Simple terratest example with GCS modules. By the way special thanks to [@romen](https://www.linkedin.com/in/rohman/?originalSubdomain=id) for teaching me about terratest.

## Feature(s)

What you can do on top of this repository:

- [x] Create a test for your module terraform
- [x] Test create a real env on your google project, then destroy it

## Prerequisite(s)

To smoothly interact with this repository please install following application in your local device:

- Linux only (`apt install build-essential`)
- Golang (tested on go1.13.1)
- Terraform (tested on v0.12.19)
- [pre-commit](https://pre-commit.com/) (tested on v1.17.0)
- Text editor (Vim, Visual Studio Code, Intellij Idea, etc.)

## How to Use This Repo
### Clone this repository
First thing first, make sure that all prerequisite(s) installed in your device, then clone this repository (make sure you have access to it)
```$bash
~ git clone git@github.com:fourirakbar/go-terratest-example.git
```

### Setup .pre-commit

```$bash
~ cd go-terratest-example
~ make setup
```

### Test your code using terratest
```$bash
~ cd test
~ go test -test.v -run TestTerraformGcpExample
```

## To Do

- [ ] Integrate with jenkins
- [ ] Integrate with gitlab-ci
- [ ] Add more complex testing
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
