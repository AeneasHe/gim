import wpath
from scripts.logic.base import *


def registerDevice():
    req = pb.RegisterDeviceReq()
    req.type = 1  #
    req.brand = "iphone"  # 厂商
    req.model = "iphone 12"  # 机型
    req.system_version = "1.0.0"  # 系统版本
    req.sdk_version = "1.0.0"  # sdk版本号

    stub = get_stub()
    meta_data = get_metadata()

    res = stub.RegisterDevice(req)
    print(res)


if __name__ == "__main__":
    registerDevice()
