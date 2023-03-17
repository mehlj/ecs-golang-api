variable "pg_dsn" {
  description = "Connection string for Postgres RDS instance"
  type        = string
  sensitive   = true
  default     = "placeholder"
}

variable "image_tag" {
  description = "Image tag that allows terraform to update ECS task definition every pipeline run"
  type        = string
  default     = "placeholder"
}