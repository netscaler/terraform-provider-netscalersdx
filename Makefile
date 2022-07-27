TEST?=$$(go list ./citrixsdx/... | grep -v 'vendor')
HOSTNAME=registry.terraform.io
NAMESPACE=citrix
NAME=citrixsdx
BINARY=terraform-provider-${NAME}
VERSION=9.9.9
# OS_ARCH=darwin_amd64
# Generalise OS_ARCH to support other platforms
OS_ARCH=$(shell go env GOOS)_$(shell go env GOARCH)

default: install

docgen:
	tfplugindocs generate --examples-dir examples/.

build: fmt
	go build -o ${BINARY}

debug-build: fmt
	go build -gcflags="all=-N -l" -o ${BINARY}
	cp -f ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

release:
	goreleaser release --rm-dist

fmt:
	go fmt ./...

tffmt:
	terraform fmt -list=true -recursive examples


install: build
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

testacc:
	# Usage: make testacc VPX_IP=10.222.74.177 VPX_NETMASK=255.255.255.0 VPX_GATEWAY=10.222.74.129 VPX_IMAGE=NSVPX-XEN-13.1-17.42_nc_64.xva VPX_PROFILE=nsroot_Notnsroot250
	- rm -i citrixsdx/citrixsdx.acctest.log
	TF_ACC=1 TF_ACC_LOG_PATH=./citrixsdx.acctest.log TF_LOG=TRACE VPX_IP=$(VPX_IP) VPX_NETMASK=$(VPX_NETMASK) VPX_GATEWAY=$(VPX_GATEWAY) VPX_IMAGE=$(VPX_IMAGE) VPX_PROFILE=$(VPX_PROFILE) go test terraform-provider-citrixsdx/citrixsdx -v

start-debug: debug-build
	~/go/bin/dlv exec --accept-multiclient --continue --headless ./${BINARY} -- -debug
