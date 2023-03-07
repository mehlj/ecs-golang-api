data "aws_iam_policy_document" "instance-assume-role-policy" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["ec2.amazonaws.com"]
    }
  }
}

resource "aws_iam_role" "ecsinstance_role" {
  name               = "ecsinstance_role"
  assume_role_policy = data.aws_iam_policy_document.instance-assume-role-policy.json
}

resource "aws_iam_role_policy" "ecspolicy" {
  name = "ecspolicy"
  role = aws_iam_role.ecsinstance_role.id

  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Effect" : "Allow",
        "Action" : [
          "ec2:DescribeTags",
          "ecs:CreateCluster",
          "ecs:DeregisterContainerInstance",
          "ecs:DiscoverPollEndpoint",
          "ecs:Poll",
          "ecs:RegisterContainerInstance",
          "ecs:StartTelemetrySession",
          "ecs:UpdateContainerInstancesState",
          "ecs:Submit*",
          "ecr:GetAuthorizationToken",
          "ecr:BatchCheckLayerAvailability",
          "ecr:GetDownloadUrlForLayer",
          "ecr:BatchGetImage",
          "logs:CreateLogStream",
          "logs:PutLogEvents"
        ],
        "Resource" : "*"
      }
    ]
  })
}

resource "aws_iam_instance_profile" "ecs_role" {
  name = "ecs_role"
  role = aws_iam_role.ecsinstance_role.name
}

resource "aws_security_group" "ecs_ec2_sg" {
  name        = "ecs_ec2_sg"
  description = "ECS and EC2-backing firewall rules"
  vpc_id      = aws_vpc.ecs_vpc.id

  ingress {
    description = "SSH to container instances for debugging"
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_instance" "ecs_cluster_member" {
  # aws ssm get-parameters --names /aws/service/ecs/optimized-ami/amazon-linux-2/recommended/image_id --region us-east-1
  ami           = "ami-0ac7415dd546fb485" # latest ECS-optimized AMI
  instance_type = "t2.micro"

  user_data = templatefile("userdata.sh",
    {
      clustername = aws_ecs_cluster.ecs_cluster.name
      region      = "us-east-1"
  })

  iam_instance_profile = aws_iam_instance_profile.ecs_role.name
  security_groups      = [aws_security_group.ecs_ec2_sg.id]
}