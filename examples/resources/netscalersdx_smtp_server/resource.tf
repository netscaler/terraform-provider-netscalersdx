resource "netscalersdx_smtp_server" "tf_smtp_server" {
  server_name    = "tf_smtp_server"
  sender_mail_id = "abc.com"
  password       = "secret"
  port           = 587
  username       = "user"
  is_ssl         = "true"
  is_auth        = "true"
}
