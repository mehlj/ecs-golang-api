resource "aws_vpc" "ecs_vpc" {
  cidr_block = "10.0.0.0/16"
}

resource "aws_subnet" "private" {
  vpc_id            = aws_vpc.ecs_vpc.id
  cidr_block        = "10.0.1.0/24"
  availability_zone = "us-east-1"
}

resource "aws_subnet" "public" {
  vpc_id            = aws_vpc.ecs_vpc.id
  cidr_block        = "10.0.2.0/24"
  availability_zone = "us-east-1"

  # instances launched in subnet get assigned a public IP
  # TODO: how does this work with a predefined CIDR block?
  map_public_ip_on_launch = true
}

resource "aws_internet_gateway" "gw" {
  depends_on = [
    aws_vpc.ecs_vpc
  ]

  vpc_id = aws_vpc.ecs_vpc.id
}

# default gateway for "public" subnet is the Internet Gateway
resource "aws_route_table" "public" {
  depends_on = [
    aws_internet_gateway.gw
  ]

  vpc_id = aws_vpc.ecs_vpc.id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.gw.id
  }
}

resource "aws_route_table" "private" {
  vpc_id = aws_vpc.ecs_vpc.id
}

resource "aws_route_table_association" "public" {
  subnet_id      = aws_subnet.public.id
  route_table_id = aws_route_table.public.id
}

resource "aws_route_table_association" "private" {
  subnet_id      = aws_subnet.private.id
  route_table_id = aws_route_table.private.id
}

resource "aws_eip" "ecs_eip" {
  depends_on = [
    aws_internet_gateway.gw
  ]

  vpc = true
}

resource "aws_nat_gateway" "ecs_nat" {
  depends_on = [
    aws_subnet.private
  ]

  allocation_id = aws_eip.ecs_eip.id
  subnet_id     = aws_subnet.public.id
}

# default gateway for "private" subnet is the NAT Gateway
resource "aws_route" "nat_route" {
  route_table_id         = aws_route_table.private.id
  nat_gateway_id         = aws_nat_gateway.ecs_nat.id
  destination_cidr_block = "0.0.0.0/0"
}