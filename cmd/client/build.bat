cd ../../../../../
set GOPATH=%cd%
set GOARCH=amd64
set GOOS=windows
cd src/demo/grpc_test/cmd/client
go build -v -ldflags="-s -w"