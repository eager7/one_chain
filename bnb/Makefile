#/bin/bash
# This is how we want to name the binary output
# These are the values we want to pass for Version and BuildTime
GITTAG=1.0.0
BUILD_TIME=`date +%Y%m%d%H%M%S`
# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS=-ldflags "-X main.Version=${GITTAG} -X main.Build_Time=${BUILD_TIME} -s -w"

mod:
	export GOPROXY="https://athens.azurefd.net" && GO111MODULE=on go build ${LDFLAGS} -o build/bnb main.go

depends:
	GO111MODULE=on go mod download

tidy:
	export GOPROXY="https://athens.azurefd.net" && GO111MODULE=on go mod tidy

replace:
	GO111MODULE=on go mod edit -replace=golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2=github.com/golang/crypto@v0.0.0-20190222235706-ffb98f73852f
	GO111MODULE=on go mod edit -replace=golang.org/x/net@v0.0.0-20190328230028-74de082e2cca=github.com/golang/net@v0.0.0-20180711010518-15845e8f865b
	GO111MODULE=on go mod edit -replace=golang.org/x/text@v0.3.1-0.20180807135948-17ff2d5776d2=github.com/golang/text@v0.3.0
	GO111MODULE=on go mod edit -replace=golang.org/x/sys@v0.0.0-20181205085412-a5c9d58dba9a=github.com/golang/sys@v0.0.0-20190215142949-d0b11bdaac8a

clean:
	-rm -rf build

check:
	GO111MODULE=on golangci-lint run ./...
	