import sys
a = 10

def b(xixi):
	return xixi + "runqing"

def transformDataFormat(node_name, node_cpu, node_mem, node_pnum, Nodes = []):
	node = {}
	node['name'] = node_name
	node['cpu'] = node_cpu
	node['mem'] = node_mem
	node['pnum'] = node_pnum
	Nodes.append(node)
	return Nodes

def RankNodesWithModel(pod_cpu, pod_mem, nodes):
	return nodes[0]['name']
