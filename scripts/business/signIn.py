from google.protobuf import message
import wpath
import grpc
import scripts.pb.conn.ext_pb2 as conn
import scripts.pb.business.ext_pb2 as pb
import scripts.pb.business.ext_pb2_grpc as pb_grpc
import time


def run():
    # 连接 rpc 服务器
    channel = grpc.insecure_channel("127.0.0.1:50301")
    stub = pb_grpc.BusinessExtStub(channel)

    metadata = [
        ("user_id", "3"),
        ("device_id", "3"),
        ("token", "0"),
        ("request_id", str(int(time.time() * 1000))),
    ]

    signInReq = conn.SignInReq()
    signInReq.phone = "18800002222"

    response = stub.SignIn(signInReq, metadata=metadata)

    print(response)


if __name__ == "__main__":
    run()
