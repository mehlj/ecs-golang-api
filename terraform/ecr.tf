resource "aws_ecr_repository" "ecr" {
  name = "mehlj-pipeline"
}

data "aws_iam_policy_document" "ecrpolicy" {
  statement {
    sid    = "Add full ECR access to PoC repository"
    effect = "Allow"

    principals {
      type        = "*"
      identifiers = ["*"]
    }

    actions = [
      "ecr:GetDownloadUrlForLayer",
      "ecr:BatchGetImage",
      "ecr:BatchCheckLayerAvailability",
      "ecr:PutImage",
      "ecr:InitiateLayerUpload",
      "ecr:UploadLayerPart",
      "ecr:CompleteLayerUpload",
      "ecr:DescribeRepositories",
      "ecr:GetRepositoryPolicy",
      "ecr:ListImages",
      "ecr:DeleteRepository",
      "ecr:BatchDeleteImage",
      "ecr:SetRepositoryPolicy",
      "ecr:DeleteRepositoryPolicy",
    ]
  }
}

resource "aws_ecr_repository_policy" "ecrpolicy" {
  repository = aws_ecr_repository.ecr.name
  policy     = data.aws_iam_policy_document.ecrpolicy.json
}