from fastapi import FastAPI
import grpc

# python -m grpc_tools.protoc -I../protobufs --python_out=. --pyi_out=. --grpc_python_out=. ../protobufs/auth.proto
from compiled_pb.auth_pb2 import AuthenticationRequest
from compiled_pb.auth_pb2_grpc import AuthStub

app = FastAPI()


@app.get("/token/verify/")
def verify_token(token: str):
    # gRPC client for making RPC calls to the server
    channel = grpc.insecure_channel("localhost:8080")
    client = AuthStub(channel)
    request = AuthenticationRequest(token=token)
    response = client.Authenticate(request)

    data = {
        "user_id": response.user_id,
        "is_token_valid": response.token_valid
    }
    return data
