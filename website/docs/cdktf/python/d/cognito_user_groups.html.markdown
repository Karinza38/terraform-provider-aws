---
subcategory: "Cognito IDP (Identity Provider)"
layout: "aws"
page_title: "AWS: aws_cognito_user_groups"
description: |-
  Terraform data source for managing AWS Cognito IDP (Identity Provider) User Groups.
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_cognito_user_groups

Terraform data source for managing AWS Cognito IDP (Identity Provider) User Groups.

## Example Usage

### Basic Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_cognito_user_groups import DataAwsCognitoUserGroups
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        DataAwsCognitoUserGroups(self, "example",
            user_pool_id="us-west-2_aaaaaaaaa"
        )
```

## Argument Reference

The following arguments are required:

* `user_pool_id` - (Required) User pool the client belongs to.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `id` - User pool identifier.
* `groups` - List of groups. See [`groups`](#groups) below.

### groups

* `description` - Description of the user group.
* `group_name` - Name of the user group.
* `precedence` - Precedence of the user group.
* `role_arn` - ARN of the IAM role to be associated with the user group.

<!-- cache-key: cdktf-0.20.8 input-b9ec4c71ef4f31f9a750655ce24ceb46439e2d99ccaf3e06f390319d638dd1c3 -->