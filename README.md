# Terraform `Citrix SDX` Provider

Terraform provider for [Citrix SDX](https://docs.citrix.com/en-us/citrix-hardware-platforms/sdx.html) provides [Infrastructure as Code (IaC)](https://en.wikipedia.org/wiki/Infrastructure_as_code) to manage your ADCs via SDX. Using the terraform provider you can provision VPXs on SDX, start, stop, reboot the VPXs on SDX.

## Requirements

* [Terraform](https://www.terraform.io/downloads.html) 1.x.x
* [Go](https://golang.org/doc/install) 1.11+ (to build the provider plugin)

## Examples

Example terraform scripts can be found in [examples](./examples/) folder

## Installing The Provider

Till the provider is available in terraform registry, you can install the provider via the following steps:

The latest released version of the provider is available on [the Release Page](https://github.com/citrix/terraform-provider-citrixsdx/releases).

Download the zip file as per your operating system and architecture.

Click below against your operating system to know how to install the provider.

<details>
  <summary>Linux</summary>

  1. Extract the zip file and copy the binary to `~/.terraform.d/plugins/registry.terraform.io/citrix/citrixsdx/<VERSION>/linux_amd64` directory. Create the directory if this is not already present.

    1. where `<VERSION>` is the version of the provider you have downloaded. Eg: `0.1.0`

</details>

<details>
  <summary>MacOS</summary>

  1. Extract the zip file and copy the binary to `~/.terraform.d/plugins/registry.terraform.io/citrix/citrixsdx/<VERSION>/darwin_amd64` directory. Create the directory if this is not already present.

    1. where `<VERSION>` is the version of the provider you have downloaded. Eg: `0.1.0`

</details>

<details>
  <summary>Windows</summary>

  1. Extract the zip file and copy the `.exe` file to `%APPDATA%/terraform.d/plugins/registry.terraform.io/citrix/citrixsdx/<version>/<OSARCH>/` directory. Create the directory if this is not already present.

     1. Where, `<version>` is the version of the provider. e.g., `0.1.0`, `<OSARCH>` is the operating system and architecture. e.g., `windows_amd64`(usually this will be the one) or `windows_386`
     2. ![](./media/windows-custom-terraform-provider-plugin-installation/plugin_location.png)
     3. You can check the location of APPDATA by running `echo %APPDATA%` in a command prompt.
     4. ![](./media/windows-custom-terraform-provider-plugin-installation/appdata_location.png)

</details>

## Validate the installation

1. Copy a sample [provider.tf](./examples/provider/provider.tf) file to a new directory.
2. Open command prompt/terminal and run the following command:
3. `terraform init` and `terraform validate`
4. ![](./media/windows-custom-terraform-provider-plugin-installation/terraform-init-validate.png)

## Using the provider

Documentation can be found [here](./PROVIDER_USAGE.md).
