resource "netscalersdx_device_group" "tf_device_group" {
  name                   = "tf_device_group"
  duration               = 10
  category               = "default"
  criteria_value         = "sample"
  static_device_list_arr = ["10.10.10.10"]
}
