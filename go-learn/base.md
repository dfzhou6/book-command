# go 基础

## go包管理模式
1. gopath
2. go module

### gopath
1. 只会去GOPATH的src目录下查找项目目录，所以项目目录只能放到GOPATH的src下。
2. GO111MODULE必须设置为off，可使用 `go env -w GO111MODULE=off` 设置。

### go module
1. go moudle模式下，可以将项目放置到任何地方，解决了GOPATH的问题。
2. 必须指定 `GO111MODULE=on`，并使用 `go mod init` 或者自己新建 `go.mod` 文件，文件内容如下：
    ```
    module zdftest

    go 1.15
    ```