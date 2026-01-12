module "iam_user" {
  source                        = "../"
  prefix_company                = "jb"
  lob                           = "connect"
  prefix_region                 = "usw2"
  application                   = "internal"
  env                           = "sandbox"
  name                          = "test-user"
  create_user                   = true
  create_iam_user_login_profile = true
}
