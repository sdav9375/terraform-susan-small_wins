output "color_enabled" {
  description = "The color setting used when the bufo_print action was configured."
  value       = var.color
}

output "trigger_resource_id" {
  description = "ID of the terraform_data resource that triggered the bufo_print action."
  value       = terraform_data.bufo_trigger.id
}