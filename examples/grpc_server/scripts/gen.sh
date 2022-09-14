echo "生成 server rpc 代码"

GIM_OUT=../rpc
protoc \
--go_out=${GIM_OUT} \
--go-grpc_out=${GIM_OUT} \
--go-grpc_opt=require_unimplemented_servers=false \
server.proto

