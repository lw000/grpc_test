cd ../../../../../
set GOPATH=%cd%
set GOARCH=amd64
set GOOS=linux
cd src/demo/grpc_test/cmd/server
go build -v -ldflags="-s -w"