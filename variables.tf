variable "color" {
  description = "Whether the bufo frog prints in color (true) or plain ASCII (false)."
  type        = bool
  default     = true
}

variable "count" {
  description = "The number of bufos"
  type = number
}

variable "name" {
  description = "The name of the bufo"
  type = string
}