resource "aws_s3_bucket" "default" {
  count = module.this.enabled ? 1 : 0

  bucket = "${module.this.id}-logs"
}

module "tfstate_backend" {
  source = "cloudposse/tfstate-backend/aws"

  force_destroy = true

  bucket_enabled   = true
  dynamodb_enabled = true

  logging = [
    {
      target_bucket = one(aws_s3_bucket.default[*].id)
      target_prefix = "tfstate/"
    }
  ]

  bucket_ownership_enforced_enabled = true

  context = module.this.context
}

output "s3_bucket_id" {
  value       = module.tfstate_backend.s3_bucket_id
  description = "S3 bucket ID"
}

output "s3_replication_role_arn" {
  value       = module.tfstate_backend.s3_replication_role_arn
  description = "The ARN of the IAM Role created for replication, if enabled."
}

output "dynamodb_table_name" {
  value       = module.tfstate_backend.dynamodb_table_name
  description = "DynamoDB table name"
}

output "dynamodb_table_id" {
  value       = module.tfstate_backend.dynamodb_table_id
  description = "DynamoDB table ID"
}