## About

truss-patch is a patch tool to customize the [github.com/tuneinc/truss](https://github.com/tuneinc/truss)


## Usage

```bash
# install the dependencies
go get -u github.com/golang/protobuf/protoc-gen-go
go get -u github.com/golang/protobuf/proto
go get -u github.com/jteeuwen/go-bindata/...

# install the truss tool
go get -u github.com/tuneinc/truss

# install truss-patch
go get -u -insecure cshgitlab.cn-csh.celestica.com/micro-service/truss-patch

# do one-time truss patch
truss-patch

# use the truss tool to generate your code
truss ...
```