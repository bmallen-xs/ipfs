locals {
  region = "us-east-1"
  name   = "ipfs-metadata"

  vpc_cidr = "10.0.0.0/16"
  azs      = slice(data.aws_availability_zones.available.names, 0, 2)

  container_name = "ecsdemo-frontend"
  container_port = 8080

  tags = {
    Name       = local.name
    Example    = local.name
    Repository = "https://github.com/terraform-aws-modules/terraform-aws-ecs"
  }
}