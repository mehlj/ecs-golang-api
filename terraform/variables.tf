variable "pg_dsn" {
  description = "Connection string for Postgres RDS instance"
  type        = string
  sensitive   = true
}