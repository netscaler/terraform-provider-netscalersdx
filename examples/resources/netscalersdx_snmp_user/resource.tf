resource "netscalersdx_snmp_user" "tf_snmp_user" {
  name             = "tf_snmp_user"
  security_level   = 2  # Please refer docs for these values      
  auth_protocol    = 1
  auth_password    = "Verysecret@123"
  privacy_protocol = 1
  privacy_password = "Verysecret@123"
}
