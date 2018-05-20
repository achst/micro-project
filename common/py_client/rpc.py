import grpc
import random
from base import get_one_rpc_node
from proto.py.service_order.service_order_pb2_grpc import OrderServiceStub


class OrderServiceClient(OrderServiceStub):

    def __init__(self):
        super(OrderServiceClient, self).__init__(grpc.insecure_channel(get_one_rpc_node("go.micro.srv.order")))

