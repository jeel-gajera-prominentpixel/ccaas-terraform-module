################################################################################
# Cluster
################################################################################

module "ecs_cluster" {
  source = "../../jb-ccaas-terraform-modules/terraform-aws-ecs-fargate"

  create = var.create
  region = var.region

  # Cluster
  configuration            = var.cluster_configuration
  name                     = var.cluster_name
  service_connect_defaults = var.cluster_service_connect_defaults
  setting                  = var.cluster_setting

  # Cluster Cloudwatch log group
  create_cloudwatch_log_group            = var.create_cloudwatch_log_group
  cloudwatch_log_group_name              = var.cloudwatch_log_group_name
  cloudwatch_log_group_retention_in_days = var.cloudwatch_log_group_retention_in_days
  cloudwatch_log_group_kms_key_id        = var.cloudwatch_log_group_kms_key_id
  cloudwatch_log_group_class             = var.cloudwatch_log_group_class
  cloudwatch_log_group_tags              = var.cloudwatch_log_group_tags

  # Cluster capacity providers
  autoscaling_capacity_providers     = var.autoscaling_capacity_providers
  default_capacity_provider_strategy = var.default_capacity_provider_strategy

  # Task execution IAM role
  create_task_exec_iam_role               = var.create_task_exec_iam_role
  task_exec_iam_role_name                 = var.task_exec_iam_role_name
  task_exec_iam_role_use_name_prefix      = var.task_exec_iam_role_use_name_prefix
  task_exec_iam_role_path                 = var.task_exec_iam_role_path
  task_exec_iam_role_description          = var.task_exec_iam_role_description
  task_exec_iam_role_permissions_boundary = var.task_exec_iam_role_permissions_boundary
  task_exec_iam_role_tags                 = var.task_exec_iam_role_tags
  task_exec_iam_role_policies             = var.task_exec_iam_role_policies

  # Task execution IAM role policy
  create_task_exec_policy  = var.create_task_exec_policy
  task_exec_ssm_param_arns = var.task_exec_ssm_param_arns
  task_exec_secret_arns    = var.task_exec_secret_arns
  task_exec_iam_statements = var.task_exec_iam_statements

  tags = merge(var.tags, var.cluster_tags)
}

################################################################################
# Service(s)
################################################################################

module "ecs-fargate-services" {
  source = "../../jb-ccaas-terraform-modules/terraform-aws-ecs-fargate"

  for_each = var.create && var.services != null ? var.services : {}

  create         = var.create
  create_service = var.create_service
  region         = var.region

  # Service
  cluster_arn            = module.cluster.arn
  desired_count          = var.desired_count
  enable_execute_command = var.enable_execute_command
  launch_type            = var.launch_type
  load_balancer          = module.load_balancer
  name                   = var.name
  security_group_ids     = var.security_group_ids
  subnet_ids             = var.subnet_ids
  platform_version       = var.platform_version
  volume_configuration   = each.value.volume_configuration

  create_iam_role                = var.create_iam_role
  create_task_definition         = var.create_task_definition
  create_task_exec_iam_role      = var.create_task_exec_iam_role
  create_task_exec_policy        = var.create_task_exec_policy
  create_tasks_iam_role          = var.create_tasks_iam_role
  enable_autoscaling             = var.enable_autoscaling
  create_security_group          = var.create_security_group
  create_infrastructure_iam_role = var.create_infrastructure_iam_role

  tags = merge(var.tags, each.value.tags)
}