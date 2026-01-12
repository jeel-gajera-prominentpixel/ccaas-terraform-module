resource "aws_wafv2_ip_set" "dynamic_ip_set" {
  for_each           = { for ipset in var.waf_ip_sets : ipset.name => ipset }
  name               = each.value.name
  scope              = var.scope # Assuming CloudFront scope; change if needed
  ip_address_version = each.value.ip_address_version
  addresses          = each.value.addresses_list

  description = var.description
  tags = {
    Name = each.value.name
  }
}


resource "aws_wafv2_web_acl" "this" {
  name        = var.name
  description = var.description
  scope       = var.scope

  default_action {
    dynamic "allow" {
      for_each = var.default_action == "allow" ? [1] : []
      content {}
    }
    dynamic "block" {
      for_each = var.default_action == "block" ? [1] : []
      content {}
    }
  }

  dynamic "rule" {
    for_each = var.waf_rules
    content {
      name     = rule.value.name
      priority = rule.value.priority

      action {
        dynamic "allow" {
          for_each = rule.value.action == "allow" ? [1] : []
          content {}
        }
        dynamic "block" {
          for_each = rule.value.action == "block" ? [1] : []
          content {}
        }
      }

      statement {
        ip_set_reference_statement {
          arn = aws_wafv2_ip_set.dynamic_ip_set[rule.value.ip_set_name].arn
        }
      }

      visibility_config {
        sampled_requests_enabled   = rule.value.sampled_requests_enabled
        cloudwatch_metrics_enabled = rule.value.cloudwatch_metrics_enabled
        metric_name                = rule.value.metric_name
      }
    }
  }

  visibility_config {
    cloudwatch_metrics_enabled = var.visibility_config.cloudwatch_metrics_enabled
    metric_name                = var.visibility_config.metric_name
    sampled_requests_enabled   = var.visibility_config.sampled_requests_enabled
  }
  tags = var.tags
}


resource "aws_cloudwatch_log_group" "waf_log_group" {
  count = var.enabled_logging_configuration ? 1 : 0
  name  = "aws-waf-logs-${var.name}"
  tags  = var.tags
}


resource "aws_wafv2_web_acl_logging_configuration" "this" {
  count = var.enabled_logging_configuration ? 1 : 0

  log_destination_configs = [aws_cloudwatch_log_group.waf_log_group[0].arn]
  resource_arn            = aws_wafv2_web_acl.this.arn

  dynamic "redacted_fields" {
    for_each = var.redacted_fields

    content {
      dynamic "method" {
        for_each = redacted_fields.value.method ? [true] : []
        content {}
      }

      dynamic "query_string" {
        for_each = redacted_fields.value.query_string ? [true] : []
        content {}
      }

      dynamic "uri_path" {
        for_each = redacted_fields.value.uri_path ? [true] : []
        content {}
      }

      dynamic "single_header" {
        for_each = lookup(redacted_fields.value, "single_header", null) != null ? toset(redacted_fields.value.single_header) : []
        content {
          name = single_header.value
        }
      }
    }
  }


  dynamic "logging_filter" {
    for_each = var.logging_filter == null ? [] : [var.logging_filter]
    content {
      default_behavior = lookup(logging_filter.value, "default_behavior")

      dynamic "filter" {
        for_each = lookup(logging_filter.value, "filter")
        iterator = filter
        content {
          behavior    = lookup(filter.value, "behavior")
          requirement = lookup(filter.value, "requirement")

          dynamic "condition" {
            for_each = lookup(filter.value, "condition")
            content {
              dynamic "action_condition" {
                for_each = lookup(condition.value, "action_condition", null) == null ? {} : lookup(condition.value, "action_condition")
                iterator = action_condition
                content {
                  action = action_condition.value
                }
              }

              dynamic "label_name_condition" {
                for_each = lookup(condition.value, "label_name_condition", null) == null ? {} : lookup(condition.value, "label_name_condition")
                iterator = label_name_condition
                content {
                  label_name = label_name_condition.value
                }
              }
            }
          }
        }
      }
    }
  }
}

