package patch

var WinBat = `
@ECHO OFF
@ECHO ON
go get github.com/pauln/go-datefmt
go get github.com/golang/protobuf/protoc-gen-go
go get github.com/golang/protobuf/proto
go get github.com/jteeuwen/go-bindata/...
go generate github.com/tuneinc/truss/gengokit/template
go install -ldflags "-X 'main.Version=7dc4d5d85c' -X 'main.VersionDate=Mon May 28 22:12:59 UTC 2018'" github.com/tuneinc/truss/cmd/truss
@ECHO OFF
    `
