TEST?=$$(go list ./acctest/... | grep -v 'vendor')
HOSTNAME=registry.terraform.io
NAMESPACE=netscaler
NAME=netscalersdx
BINARY=terraform-provider-${NAME}
VERSION=0.2.0
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
	# Usage: make testacc VPX_IP=10.10.10.10 VPX_NETMASK=255.255.255.0 VPX_GATEWAY=10.10.10.1 VPX_IMAGE=NSVPX-XEN-13.1-17.42_nc_64.xva VPX_PROFILE=vpx_provile1 NETWORK=10.10.10.11 GATEWAY=10.10.10.12
	- rm internal/acctest/netscalersdx.acctest.log 
	TF_ACC=1 TF_ACC_LOG_PATH=./netscalersdx.acctest.log TF_LOG=TRACE VPX_IP=$(VPX_IP) VPX_NETMASK=$(VPX_NETMASK) VPX_GATEWAY=$(VPX_GATEWAY) VPX_IMAGE=$(VPX_IMAGE) VPX_PROFILE=$(VPX_PROFILE) go test terraform-provider-netscalersdx/internal/acctest -v

# start-debug: debug-build
# 	~/go/bin/dlv exec --accept-multiclient --continue --headless ./${BINARY} -- -debug
