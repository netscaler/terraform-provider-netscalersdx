

data "netscalersdx_mps_feature" "tf_mps_feature" {
  feature_name = "Device_SSL_Cert"
}

output "name" {
  value = data.netscalersdx_mps_feature.tf_mps_feature.id
}
