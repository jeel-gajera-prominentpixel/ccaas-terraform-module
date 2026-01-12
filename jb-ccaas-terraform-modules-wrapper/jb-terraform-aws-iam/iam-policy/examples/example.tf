module "iam_policy" {
  source         = "../"
  prefix_company = "jb"
  lob            = "connect"
  prefix_region  = "usw2"
  application    = "internal"
  env            = "sandbox"
  name           = "test-policy"
  description    = "My example policy"

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "ec2:Describe*"
      ],
      "Effect": "Allow",
      "Resource": "*"
    }
  ]
}
EOF
}
