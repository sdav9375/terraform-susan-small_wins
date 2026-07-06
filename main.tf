terraform {
  required_version = "~>1.15.7"

  required_providers {
    bufo = {
      source = "austinvalle/bufo"
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