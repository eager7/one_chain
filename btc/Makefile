#/bin/bash
# This is how we want to name the binary output
# These are the values we want to pass for Version and BuildTime
GITTAG=1.0.0
BUILD_TIME=`date +%Y%m%d%H%M%S`
# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS=-ldflags "-X main.Version=${GITTAG} -X main.Build_Time=${BUILD_TIME} -s -w"

eth_cli:
#	GO111MODULE=on go mod edit -replace golang.org/x/crypto@v0.0.0-20180904163835-0709b304e793=github.com/golang/crypto@v0.0.0-20190222235706-ffb98f73852f
	export GOPROXY="https://athens.azurefd.net" && GO111MODULE=on go build ${LDFLAGS} -o build/go_eth_client main.go

depends:
	GO111MODULE=on go mod download

tidy:
	export GOPROXY=https://goproxy.io && GO111MODULE=on go mod tidy

clean:
	-rm -rf build

check:
	GO111MODULE=on golangci-lint run ./...
	