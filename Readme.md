## 生成 proto文件

```
protoc -I . --go_out=plugins=grpc:../pb ./hello.proto
```