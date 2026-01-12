module "cloudfront" {
  source = "../../jb-ccaas-terraform-modules/terraform-aws-cloudfront"

  aliases                              = var.aliases
  comment                              = var.comment
  enabled                              = var.enabled
  staging                              = var.staging
  http_version                         = var.http_version
  is_ipv6_enabled                      = var.is_ipv6_enabled
  price_class                          = var.price_class
  retain_on_delete                     = var.retain_on_delete
  wait_for_deployment                  = var.wait_for_deployment
  create_origin_access_identity        = var.create_origin_access_identity
  origin_access_identities             = var.origin_access_identities
  create_origin_access_control         = var.create_origin_access_control
  origin_access_control                = var.origin_access_control
  logging_config                       = var.logging_config
  origin                               = var.origin
  origin_group                         = var.origin_group
  default_cache_behavior               = var.default_cache_behavior
  ordered_cache_behavior               = var.ordered_cache_behavior
  custom_error_response                = var.custom_error_response
  create_monitoring_subscription       = var.create_monitoring_subscription
  realtime_metrics_subscription_status = var.realtime_metrics_subscription_status
  geo_restriction                      = var.geo_restriction
  viewer_certificate                   = var.viewer_certificate
  web_acl_id                           = var.web_acl_id
  default_root_object                  = var.default_root_object
  continuous_deployment_policy_id      = var.continuous_deployment_policy_id
  create_distribution                  = var.create_distribution
  tags = merge(local.tags, {
    Name = var.name == "" ? local.cloudfront_name : var.name
  })
}



resource "aws_cloudfront_origin_request_policy" "custom" {
  count   = var.create_request_policy && var.custom_origin_request_policy != null ? 1 : 0
  name    = var.custom_origin_request_policy.name
  comment = var.custom_origin_request_policy.comment

  cookies_config {
    cookie_behavior = var.custom_origin_request_policy.cookies_config.cookie_behavior

    dynamic "cookies" {
      for_each = var.custom_origin_request_policy.cookies_config.cookie_behavior == "whitelist" ? [var.custom_origin_request_policy.cookies_config.cookies] : []
      content {
        items = cookies.value
      }
    }
  }

  headers_config {
    header_behavior = var.custom_origin_request_policy.headers_config.header_behavior

    dynamic "headers" {
      for_each = var.custom_origin_request_policy.headers_config.header_behavior != "none" ? [var.custom_origin_request_policy.headers_config.headers] : []
      content {
        items = headers.value
      }
    }
  }

  query_strings_config {
    query_string_behavior = var.custom_origin_request_policy.query_strings_config.query_string_behavior

    dynamic "query_strings" {
      for_each = var.custom_origin_request_policy.query_strings_config.query_string_behavior == "whitelist" ? [var.custom_origin_request_policy.query_strings_config.query_strings] : []
      content {
        items = query_strings.value
      }
    }
  }
}
