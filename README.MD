# go-zero k8s test

```shell

goctl api go --api user.api --dir .

goctl rpc protoc user.proto --go_out=./ --go-grpc_out=./  --zrpc_out=./

go mod tidy

cd user/api
goctl docker -go user.go
mv Dockerfile ../../Dockerfile-user-api

cd user/rpc
goctl docker -go user.go
mv Dockerfile ../../Dockerfile-user-rpc

```