---
subcategory: "API Gateway"
layout: "aws"
page_title: "AWS: aws_api_gateway_api_key"
description: |-
  Provides an API Gateway API Key.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_api_gateway_api_key

Provides an API Gateway API Key.

~> **NOTE:** Since the API Gateway usage plans feature was launched on August 11, 2016, usage plans are now **required** to associate an API key with an API stage.

## Example Usage

```python
# Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.api_gateway_api_key import ApiGatewayApiKey
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        ApiGatewayApiKey(self, "MyDemoApiKey",
            name="demo"
        )
```

## Argument Reference

This resource supports the following arguments:

* `name` - (Required) Name of the API key.
* `description` - (Optional) API key description. Defaults to "Managed by Terraform".
* `enabled` - (Optional) Whether the API key can be used by callers. Defaults to `true`.
* `value` - (Optional) Value of the API key. If specified, the value must be an alphanumeric string between 20 and 128 characters. If not specified, it will be automatically generated by AWS on creation.
* `tags` - (Optional) Key-value map of resource tags. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - ID of the API key
* `created_date` - Creation date of the API key
* `last_updated_date` - Last update date of the API key
* `value` - Value of the API key
* `arn` - ARN
* `tags_all` - Map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import API Gateway Keys using the `id`. For example:

```python
# Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
```

Using `terraform import`, import API Gateway Keys using the `id`. For example:

```console
% terraform import aws_api_gateway_api_key.my_demo_key 8bklk8bl1k3sB38D9B3l0enyWT8c09B30lkq0blk
```

<!-- cache-key: cdktf-0.17.1 input-c096e0bf3eef1372bbc03ef43fd91d12e9bce137de455aa87618afa795050ecf -->