syntax = "proto3";

package helloworld; //包名

// MathService 微服务
service MathService {
  rpc Add (AddRequest) returns (AddReply) {}
}

// AddRequest 请求数据格式
message AddRequest {
  int32 a = 1;
  int32 b = 2;
}

// AddReply 响应数据格式
message AddReply {
  int32 c = 1;
}

