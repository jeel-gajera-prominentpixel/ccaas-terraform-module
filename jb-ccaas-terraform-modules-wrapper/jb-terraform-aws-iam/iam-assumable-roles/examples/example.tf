module "iam_role" {
  source           = "../"
  prefix_company   = "jb"
  lob              = "connect"
  prefix_region    = "usw2"
  application      = "internal"
  env              = "sandbox"
  create_role      = true
  role_name        = "test-role"
  role_description = "This is a test role"
}
