from google.protobuf import message
import wpath
import grpc
import scripts.pb.conn.ext_pb2 as conn
import scripts.pb.logic.ext_pb2 as pb
import scripts.pb.logic.ext_pb2_grpc as pb_grpc
import time


def run():
    # 连接 rpc 服务器
    channel = grpc.insecure_channel("127.0.0.1:50001")
    stub = pb_grpc.LogicExtStub(channel)

    metadata = [
        ("user_id", "2"),
        ("device_id", "2"),
        ("token", "0"),
        ("request_id", str(int(time.time() * 1000))),
    ]

    text = conn.Text()
    text.text = b"hello gim"
    content = text.SerializeToString()
    # 反序列化用ParseFromString

    messageReq = pb.SendMessageReq()
    messageReq.receiver_type = conn.ReceiverType.RT_USER
    messageReq.receiver_id = 1
    # 如果要设置通知用户，如下方法之一设置repeated格式的数据
    # messageReq.to_user_ids.extend([1, 2])
    # messageReq.to_user_ids[:] = [1, 2]
    messageReq.message_type = conn.MessageType.MT_TEXT
    messageReq.message_content = text  # .encode("utf-8")
    messageReq.send_time = 123
    messageReq.is_persist = True

    response = stub.SendMessage(messageReq, metadata=metadata)

    print(response)


if __name__ == "__main__":
    run()
