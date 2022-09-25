echo "生成 rpc 代码"

CLIENT_OUT=../client/pb
protoc \
--go_out=${CLIENT_OUT} \
--go-grpc_out=${CLIENT_OUT} \
--go-grpc_opt=require_unimplemented_servers=false \
server.proto

SERVER_OUT=../server/pb
protoc \
--go_out=${SERVER_OUT} \
--go-grpc_out=${SERVER_OUT} \
--go-grpc_opt=require_unimplemented_servers=false \
server.proto

