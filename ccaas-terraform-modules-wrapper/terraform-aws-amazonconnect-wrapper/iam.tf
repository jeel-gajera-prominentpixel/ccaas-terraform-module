module "connect_federation_policy" {
  count       = var.sso_integration ? 1 : 0
  source      = "../../ccaas-terraform-modules/terraform-aws-iam/modules/iam-policy"
  name        = "${local.amazon_connect_iam_name}-federation-pol"
  description = "IAM policy to get federation token and list for Amazon Connect instance"
  policy      = data.aws_iam_policy_document.connect_federation_policy_document.json
}

data "aws_iam_policy_document" "connect_federation_policy_document" {
  statement {
    effect = "Allow"
    actions = [
      "connect:GetFederationToken"
    ]
    resources = [
      "*"
    ]
    condition {
      test     = "StringEquals"
      variable = "connect:InstanceId"
      values   = [module.amazon_connect.instance_id]
    }
  }
}

module "connect_azure_ad_policy" {
  count       = var.sso_integration ? 1 : 0
  source      = "../../ccaas-terraform-modules/terraform-aws-iam/modules/iam-policy"
  name        = "${local.amazon_connect_iam_name}-azure-pol"
  description = "Azure AD IAM access policy"
  policy      = data.aws_iam_policy_document.connect_azure_ad_policy_document.json
}


data "aws_iam_policy_document" "connect_azure_ad_policy_document" {
  statement {
    effect = "Allow"
    actions = [
      "iam:ListRoles",
      "iam:ListAccountAliases"
    ]
    resources = ["*"]
  }
}

resource "aws_iam_saml_provider" "AWSSSO_admin_Connect" {
  count                  = var.sso_integration ? 1 : 0
  name                   = "${format("connect-%s-%s", var.region_prefix, var.env)}-saml-pvr"
  saml_metadata_document = file(var.connect_admin_sso_xml_filepath)
}

resource "aws_iam_saml_provider" "AWSSSO_agent_Connect" {
  count                  = var.sso_integration ? 1 : 0
  name                   = "${format("connect-%s", var.region_prefix)}-agsaml-pvr"
  saml_metadata_document = file(var.connect_agent_sso_xml_filepath)
}

data "aws_iam_policy_document" "AWSConnectTrustPolicy" {
  count = var.sso_integration ? 1 : 0
  statement {
    effect = "Allow"
    actions = [
      "sts:TagSession",
      "sts:AssumeRoleWithSAML"
    ]
    principals {
      type = "Federated"
      identifiers = [
        aws_iam_saml_provider.AWSSSO_admin_Connect[count.index].arn,
        aws_iam_saml_provider.AWSSSO_agent_Connect[count.index].arn
      ]
    }
    condition {
      test     = "StringLike"
      variable = "SAML:aud"

      values = [
        "https://${module.amazon_connect.instance_id}.${data.aws_region.current.name}.sign-in.connect.aws/saml*"
      ]
    }
  }
}

resource "aws_iam_role" "AWSSSO_connect_role" {
  count                = var.sso_integration ? 1 : 0
  name                 = "${local.amazon_connect_iam_name}-sso-role"
  assume_role_policy   = data.aws_iam_policy_document.AWSConnectTrustPolicy[count.index].json
  description          = "IAM role to access Connect Application"
  permissions_boundary = var.permissions_boundary_arn
}

resource "aws_iam_role_policy_attachment" "AWSConnectPolicyAttachment1" {
  count      = var.sso_integration ? 1 : 0
  role       = aws_iam_role.AWSSSO_connect_role[count.index].name
  policy_arn = module.connect_federation_policy[count.index].arn
}

resource "aws_iam_role_policy_attachment" "AWSConnectPolicyAttachment2" {
  count      = var.sso_integration ? 1 : 0
  role       = aws_iam_role.AWSSSO_connect_role[count.index].name
  policy_arn = module.connect_azure_ad_policy[count.index].arn
}
