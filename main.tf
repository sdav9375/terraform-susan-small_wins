terraform {
  required_version = "~> 1.15.7"

  required_providers {
    bufo = {
      source = "austinvalle/bufo"
    }

    httpaction = {
      source = "terraform.local/local/httpaction"
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

action "httpaction_request" "smoke_test" {
  config {
    url    = "https://eobzplaeaw1eev8.m.pipedream.net"
    method = "POST"
    payload = jsonencode({
      message = "Hello from a Terraform action"
    })
  }
}
