cd pkg/proto

#python3 -m grpc_tools.protoc --python_out=../../scripts/pb --grpc_python_out=../../scripts/pb -I. *ext.proto

python3 -m grpc_tools.protoc --python_out=../../scripts/pb --grpc_python_out=../../scripts/pb -I. push.ext.proto


# 生成的文件引用路径可能出错，需要手动修复
# import wpath
# from scripts.pb.

#protoc --python_out=../../scripts/pb/  --grpc_python_out=../../scripts/pb/  *ext.proto