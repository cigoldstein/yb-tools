TOP_BUILDDIR = ../.
include ../global.mk

.PHONEY: all protobuf yb-connect

all: test yb-connect

yb-connect:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -ldflags=" -X 'github.com/yugabyte/yb-tools/yb-connect/cmd.Version=$$(git rev-parse HEAD)'" -o ${TOP_BUILDDIR}/bin/yb-connect main.go

