# 安装 python 下的 protoc 编译器
pip install grpcio-tools

# 编译 proto 文件
python -m grpc_tools.protoc --python_out=. --grpc_python_out=. -I. helloworld.proto

python -m grpc_tools.protoc: python 下的 protoc 编译器通过 python 模块(module) 实现, 所以说这一步非常省心
--python_out=. : 编译生成处理 protobuf 相关的代码的路径, 这里生成到当前目录
--grpc_python_out=. : 编译生成处理 grpc 相关的代码的路径, 这里生成到当前目录
-I. helloworld.proto : proto 文件的路径, 这里的 proto 文件在当前目录


# 如果报错
got an unexpected keyword argument 'serialized_options'
很有可能是，终端上的 protoc 版本 与python库内的protobuf版本不一样。
pip install -U protobuf
