# Copyright 2015 gRPC authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
"""The Python implementation of the GRPC helloworld.Greeter client."""

from __future__ import print_function

import grpc

from proto.service_order import service_order_pb2
from proto.service_order import service_order_pb2_grpc


def run():
    channel = grpc.insecure_channel('localhost:58875')
    stub = service_order_pb2_grpc.OrderServiceStub(channel)
    response = stub.GetOrders(service_order_pb2.GetOrdersRequest(page_index=1, page_count=10))
    print(response)


if __name__ == '__main__':
    run()
