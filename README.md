# sf-express-sdk-go

顺丰软件开发工具包(go)

## 相关

- [顺丰丰桥](https://qiao.sf-express.com)

## API Reference

- [sf-express-sdk-go - GoDoc](https://godoc.org/github.com/taadis/sf-express-sdk-go)

## 安装

### GOPATH

可以使用 `go get` 把 SDK 下载到你的 GOPATH 工作目录

```
go get github.com/taadis/sf-express-sdk-go
```

### Go Modules

也可以使用模块管理依赖, 这也是未来推荐的方式

```
// 获取最新版本
go get github.com/taadis/sf-express-sdk-go@latest

// 当然你也可以指定具体的版本
// go get github.com/taadis/sf-express-sdk-go@v0.0.1
```

## 测试

为了方便测试以及保护顾客编码和校检码, 这里可以通过附加参数 `-agrs` 来指定你的顾客编码和校检码

```
go test -v -args -clientCode 你的顾客编码 -checkWord 你的校检码
```

> 更多的示例可以参考各 `_test` 文件, 或者提个 issue 吧...