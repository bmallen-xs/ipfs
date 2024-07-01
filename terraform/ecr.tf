resource "aws_ecr_repository" "this" {
  name                 = "blockparty/metadata"
  image_tag_mutability = "MUTABLE"

  force_delete = true

  image_scanning_configuration {
    scan_on_push = true
  }
}

output "ecr" {
  value = aws_ecr_repository.this
}