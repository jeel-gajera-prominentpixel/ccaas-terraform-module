module "dynamodb-table" {
  source                      = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-dynamodb?ref=main"
  prefix_company              = "jb"
  lob                         = "itsd"
  prefix_region               = "usw2"
  application                 = "recordings"
  env                         = "sandbox"
  name                        = "jb-dynamodb-table"
  hash_key                    = "id"
  range_key                   = "title"
  table_class                 = "STANDARD"
  deletion_protection_enabled = false
  attributes = [
    {
      name = "id"
      type = "N"
    },
    {
      name = "title"
      type = "S"
    },
  ]
  import_table = {
    input_format           = "DYNAMODB_JSON"
    input_compression_type = "NONE"
    bucket                 = "jb-bucket"
    key_prefix             = "/example"
  }
  stream_enabled                     = true
  stream_view_type                   = "NEW_AND_OLD_IMAGES"
  server_side_encryption_enabled     = true
  server_side_encryption_kms_key_arn = "arn:aws:kms:us-east-1:123456789012:key/abcd1234-a123-456a-a12b-a123b4cd5678"
  global_secondary_indexes = [
    {
      name               = "TitleIndex"
      hash_key           = "title"
      range_key          = "age"
      projection_type    = "INCLUDE"
      non_key_attributes = ["id"]
    }
  ]
  replica_regions = [{
    region_name            = "eu-west-2"
    kms_key_arn            = "arn:aws:kms:us-east-1:123456789012:key/abcd1234-a123-456a-a12b"
    propagate_tags         = true
    point_in_time_recovery = true
  }]
  tags = local.tags
}
