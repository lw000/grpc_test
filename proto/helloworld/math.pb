syntax = "proto3";

package helloworld; //包名

// MathService 微服务
service MathService {
  rpc Add (AddRequest) returns (AddReply) {}
  rpc Sub (SubRequest) returns (SubReply) {}
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

// SubRequest 请求数据格式
message SubRequest {
  int32 a = 1;
  int32 b = 2;
}

// SubReply 响应数据格式
message SubReply {
  int32 c = 1;
}

