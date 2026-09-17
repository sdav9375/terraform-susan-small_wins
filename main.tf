terraform {
  required_version = ">=1.16.3"

  required_providers {
    bufo = {
      source = "austinvalle/bufo"
    }

    https-pipedream = {
      source  = "sdav9375/terraform-provider-https-pipedream"
      version = "~> 1.0"
    }
  }
}

resource "terraform_data" "bufo_trigger" {
  input = var.color

  lifecycle {
    action_trigger {
      events  = [after_create]
      actions = [action.bufo_print.hello]
    }
  }
}

action "bufo_print" "hello" {
  config {
    color = var.color
  }
}

action "pipedream_request" "smoke_test" {
  config {
    url    = "https://eobzplaeaw1eev8.m.pipedream.net"
    method = "POST"
    payload = jsonencode({
      message = "Hello from a Terraform action"
    })
  }
}
