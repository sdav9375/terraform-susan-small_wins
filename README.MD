# terraform-local-bufo

A minimal, zero-infrastructure module for testing HCP Terraform's `action` blocks and no-code provisioning workflow. It invokes the [`austinvalle/bufo`](https://registry.terraform.io/providers/austinvalle/bufo) provider's `bufo_print` action, which prints an ASCII frog to run output. No cloud resources, credentials, or real infrastructure are involved.

## Usage

```hcl
module "bufo" {
  source  = "app.terraform.io/<YOUR_ORG>/bufo/local"
  version = "1.0.0"

  color = true
}
```

## What it does

- Declares the `bufo` provider as a required provider.
- Creates a `terraform_data` resource whose `lifecycle.action_trigger` invokes the `bufo_print` action automatically after the resource is created.
- No real infrastructure is provisioned — `terraform_data` exists only to give the action something to hang off of and to give this module a resource that shows up in state.

## Inputs

| Name    | Description                                          | Type   | Default | Required |
|---------|-------------------------------------------------------|--------|---------|----------|
| `color` | Whether the bufo frog prints in color or plain ASCII. | `bool` | `true`  | no       |

## Outputs

| Name                   | Description                                                      |
|------------------------|--------------------------------------------------------------------|
| `color_enabled`        | Echoes back the `color` value used to configure the action.        |
| `trigger_resource_id`  | ID of the `terraform_data` resource that triggered the action.     |

## Requirements

| Name      | Version   |
|-----------|-----------|
| terraform | ~>1.15.7  |
| bufo      | n/a       |

## Notes

- This module is intentionally trivial — it exists to validate the actions and no-code publishing mechanics, not as a production pattern.
- Actions have no state of their own (no create/update/delete lifecycle), so `terraform_data` is the only resource Terraform tracks; the action itself only appears in run output and the workspace's Actions tab.
- Safe to apply and destroy repeatedly — nothing persists outside of Terraform state.