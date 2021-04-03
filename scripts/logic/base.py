from google.protobuf import message
import wpath
import grpc
import scripts.pb.conn.ext_pb2 as conn
import scripts.pb.logic.ext_pb2 as pb
import scripts.pb.logic.ext_pb2_grpc as pb_grpc
import time


def get_metadata():
    metadata = [
        ("user_id", "2"),
        ("device_id", "2"),
        ("token", "0"),
        ("request_id", str(int(time.time() * 1000))),
    ]
    return metadata


def get_stub():
    channel = grpc.insecure_channel("127.0.0.1:50001")
    stub = pb_grpc.LogicExtStub(channel)
    return stub
