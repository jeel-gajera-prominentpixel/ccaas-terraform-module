module "secrets-manager" {
  source                  = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-secret-manager?ref=main"
  prefix_company          = "jb"
  lob                     = "itsd"
  prefix_region           = "usw2"
  application             = "recordings"
  env                     = "sandbox"
  name_prefix             = "jb-secret-manager"
  description             = "Example Secrets Manager secret"
  recovery_window_in_days = 0
  replica = {
    us-east-1 = {}
    another = {
      region = "us-west-2"
    }
  }
  create_policy       = true
  block_public_policy = true
  policy_statements = {
    read = {
      sid = "AllowAccountRead"
      principals = [{
        type        = "AWS"
        identifiers = ["arn:aws:iam::${data.aws_caller_identity.current.account_id}:root"]
      }]
      actions   = ["secretsmanager:GetSecretValue"]
      resources = ["*"]
    }
  }

  create_random_password = true
  secret_string = jsonencode({
    engine   = "mariadb",
    host     = "mydb.cluster-123456789012.us-east-1.rds.amazonaws.com",
    username = "Bill",
    password = "ThisIsMySuperSecretString12356!"
    dbname   = "mydb",
    port     = 3306
  })
  random_password_length = 64
  ignore_secret_changes  = true
  enable_rotation        = true
  tags                   = local.tags
}
