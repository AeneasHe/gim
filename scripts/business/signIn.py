import wpath
from scripts.business.base import *


def run():
    signInReq = pb.SignInReq()
    signInReq.phone_number = "18800003333"
    signInReq.code = "0"
    signInReq.device_id = 3

    stub = get_stub()
    metadata = get_metadata()

    response = stub.SignIn(signInReq, metadata=metadata)

    print(response)


if __name__ == "__main__":
    run()
