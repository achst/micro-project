# encoding:utf-8

from flask import Flask
from werkzeug.serving import run_simple

from config import Config
from common import py_client
from common.py_client.rpc import OrderServiceClient

app = Flask(__name__)


@app.route('/api/order/list', methods=['GET'])
def get_order_list():
    from proto.py.service_order import service_order_pb2

    my_client = OrderServiceClient()
    response = my_client.GetOrders(service_order_pb2.GetOrdersRequest(page_index=1, page_count=10))
    return str(response)


def init():
    py_client.init_rpc(Config)


def main():
    run_simple('localhost', 4000, app)


if __name__ == '__main__':
    init()
    main()
