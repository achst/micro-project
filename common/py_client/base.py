import requests
import random


class RpcConfig(object):

    registry_uri = None


def get_rpc_nodes(micro_name):
    registry_uri = RpcConfig.registry_uri
    data = {
        "service": micro_name,
    }
    response = requests.get(registry_uri, data)
    if response.status_code != 200:
        raise Exception('get_rpc_nodes status error')
    data = response.json()
    if not data or not isinstance(data, list) or len(data) <= 0:
        raise Exception('get_rpc_nodes not find any service node from registry')
    if 'nodes' not in data[0]:
        raise Exception('get_rpc_nodes get valid data')
    return data[0]['nodes']


def get_one_rpc_node(micro_name):
    nodes = get_rpc_nodes(micro_name)
    one_node = random.choice(nodes)
    node_address = one_node['address'] + ":" + str(one_node['port'])
    return node_address
