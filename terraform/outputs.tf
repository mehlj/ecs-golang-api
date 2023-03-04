output "repository_url" {
  description = "URL for the ECR repository."

  value = aws_ecr_repository.ecr.repository_url
}