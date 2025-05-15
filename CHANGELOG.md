# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.7.1]

### Changed

- **Go-version**: Upgraded go version to 1.24

### Fixed

- **netscalersdx_ns_device_profile**: Handled read func to not set some attributes from NetScaler SDX [#86]

## [0.7.0]

### Fixed

- **netscalersdx_ns_device_profile**: Handled read func to not set some attributes from NetScaler SDX [#86]

[#86]: https://github.com/netscaler/terraform-provider-netscalersdx/issues/86

## [0.6.0]

### Added

- **Import Support** `netscalersdx_provision_vpx`
- **New Resource** `current_hostname` [#76]
- **New Resource** `mps_ssl_certkey` [#56]

### Changed

- Updated dependent libraries due to security issues.
- Updated the documentation using `tfplugindocs`
- **netscalersdx_provision_vpx**: Updated read logic to handle network_interface attribute in more generic manner

[#76]: https://github.com/netscaler/terraform-provider-netscalersdx/issues/76
[#56]: https://github.com/netscaler/terraform-provider-netscalersdx/issues/56

## [0.5.0]

### Added

- **New Resource** `system_settings` [#55]
- **New Resource** `ssl_settings` [#57]
- **New Resource** `cipher_config` [#58]
- **New Resource** `mps` [#59]
- **New Resource** `ntp_sync` [#61]

### Changed

- **netscalersdx_provision_vpx**: Added Validate func to `licence` attribute to validate the input from user [#12]. Removed deprecated attributes and updated the attribute types

[#12]: https://github.com/netscaler/terraform-provider-netscalersdx/issues/12
[#55]: https://github.com/netscaler/terraform-provider-netscalersdx/issues/55
[#57]: https://github.com/netscaler/terraform-provider-netscalersdx/issues/57
[#58]: https://github.com/netscaler/terraform-provider-netscalersdx/issues/58
[#59]: https://github.com/netscaler/terraform-provider-netscalersdx/issues/59
[#61]: https://github.com/netscaler/terraform-provider-netscalersdx/issues/61

## [0.4.0]

### Added

- **New Resource** `aclrule`
- **New Resource** `aaa_server`

### Fixed

- **netscalersdx_provision_vpx**: handled double encoding for list of string values [#40]

## [0.3.0]

### Added

- **New Resource** `snmp_trap`
- **New Resource** `snmp_manager`
- **New Resource** `snmp_view`
- **New Resource** `mpsuser`
- **New Resource** `mpsgroup`
- **New Resource** `mps_feature`
- **New Resource** `current_timezone`
- **New Resource** `snmp_alarm_config`

## [0.2.0]

### Added

- **New Resource** `blx_device_profile`
- **New Resource** `cipher_group`
- **New Resource** `device_group`
- **New Resource** `device_profile`
- **New Resource** `ldap_server`
- **New Resource** `ns_device_profile` ([#5])
- **New Resource** `ns_save_config`
- **New Resource** `ntp_param`
- **New Resource** `ntp_server`
- **New Resource** `radius_server`
- **New Resource** `smtp_server`
- **New Resource** `snmp_user`
- **New Resource** `static_route`
- **New Resource** `syslog_params`
- **New Resource** `syslog_server`
- **New Resource** `tacacs_server`
- Migrated the provider to terraform-plugin-framework.
- Updated the resource name from `citrixsdx` to `netscalersdx`.


## [0.1.0] - 2022-08-08

### Added

* **New Resource** `provision_vpx`
* **New Resource** `vpx_state`

[Unreleased]: https://github.com/netscaler/terraform-provider-netscalersdx/compare/v0.7.1...HEAD
[0.7.1]: https://github.com/netscaler/terraform-provider-netscalersdx/compare/v0.7.0...v0.7.1
[0.7.0]: https://github.com/netscaler/terraform-provider-netscalersdx/compare/v0.6.0...v0.7.0
[0.6.0]: https://github.com/netscaler/terraform-provider-netscalersdx/compare/v0.5.0...v0.6.0
[0.5.0]: https://github.com/netscaler/terraform-provider-netscalersdx/compare/v0.4.0...v0.5.0
[0.4.0]: https://github.com/netscaler/terraform-provider-netscalersdx/compare/v0.3.0...v0.4.0
[0.3.0]: https://github.com/netscaler/terraform-provider-netscalersdx/compare/v0.2.0...v0.3.0
[0.2.0]: https://github.com/netscaler/terraform-provider-netscalersdx/compare/v0.1.0...0.2.0
[0.1.0]: https://github.com/netscaler/terraform-provider-netscalersdx/releases/tag/v0.1.0
[#40]: https://github.com/netscaler/terraform-provider-netscalersdx/issues/40
[#5]: https://github.com/netscaler/terraform-provider-netscalersdx/issues/5