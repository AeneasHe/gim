import wpath
from scripts.logic.base import *


def sendMessage():
    # 连接 rpc 服务器

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
    messageReq.message_content = content
    messageReq.send_time = 123
    messageReq.is_persist = True

    stub = get_stub()
    metadata = get_metadata()

    response = stub.SendMessage(messageReq, metadata=metadata)

    print(response)


if __name__ == "__main__":
    sendMessage()
