# grpc-conv-any
grpcレスポンスで取得したany型のバイナリデータを変換する。
→kotlin x springの場合、any型の値が型解決できなくてエラーになったけどgoはしなかった。kotlinもファイル指定すれば見れた。

## setup

```terminal
//gRPCライブラリ
go get -u google.golang.org/grpc

//自動生成用のプラグイン
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

//コード自動生成
protoc --go_out=. --go-grpc_out=require_unimplemented_servers=false:. ./proto/hello.proto
```
