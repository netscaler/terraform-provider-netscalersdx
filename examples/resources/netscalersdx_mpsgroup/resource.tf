resource "netscalersdx_mpsgroup" "tf_mpsgroup" {
  name                     = "tf_mpsgroup"
  permission               = "readonly"
  select_individual_entity = "false"
  assign_all_apps          = "true"
  description              = "test-terraform"
  role                     = "nonadmin"
}
