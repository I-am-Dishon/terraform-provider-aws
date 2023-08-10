---
subcategory: "Lightsail"
layout: "aws"
page_title: "AWS: aws_lightsail_bucket"
description: |-
  Provides a lightsail bucket
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_lightsail_bucket

Provides a lightsail bucket.

## Example Usage

```python
# Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.lightsail_bucket import LightsailBucket
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        LightsailBucket(self, "test",
            bundle_id="small_1_0",
            name="mytestbucket"
        )
```

## Argument Reference

This resource supports the following arguments:

* `name` - (Required) The name for the bucket.
* `bundle_id` - (Required) - The ID of the bundle to use for the bucket. A bucket bundle specifies the monthly cost, storage space, and data transfer quota for a bucket. Use the [get-bucket-bundles](https://docs.aws.amazon.com/cli/latest/reference/lightsail/get-bucket-bundles.html) cli command to get a list of bundle IDs that you can specify.
* `tags` - (Optional) A map of tags to assign to the resource. To create a key-only tag, use an empty string as the value. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - The name used for this bucket (matches `name`).
* `arn` - The ARN of the lightsail bucket.
* `availability_zone` - The resource Availability Zone. Follows the format us-east-2a (case-sensitive).
* `created_at` - The timestamp when the bucket was created.
* `region` - The Amazon Web Services Region name.
* `support_code` - The support code for the resource. Include this code in your email to support when you have questions about a resource in Lightsail. This code enables our support team to look up your Lightsail information more easily.
* `tags_all` - A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import `aws_lightsail_bucket` using the `name` attribute. For example:

```python
# Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
```

Using `terraform import`, import `aws_lightsail_bucket` using the `name` attribute. For example:

```console
% terraform import aws_lightsail_bucket.test example-bucket
```

<!-- cache-key: cdktf-0.17.1 input-31ef4d7cd5556521a3c48b3bae0b34c347181d340f923da9040d9741794675d4 -->