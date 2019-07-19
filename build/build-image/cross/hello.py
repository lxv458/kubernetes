import yaml
import sys
import os
from sklearn.datasets import dump_svmlight_file
import xgboost as xgb
from xgboost import DMatrix
from sklearn.datasets import load_svmlight_file
import pickle
import time

a = 10

def b(xixi):
	return xixi + "runqing"

def transformDataFormat(node_name, node_cpu, node_mem, node_pnum, Nodes = []):
	node = {}
	node['name'] = node_name
	node['cpu'] = float(node_cpu)
	node['mem'] = float(node_mem)
	node['pnum'] = float(node_pnum)
	Nodes.append(node)
	return Nodes

def RankNodesWithModel(pod_cpu, pod_mem, nodes):
	# return nodes[0]['name']
	output = []
	y = []
	for node in nodes:
		entry = [float(pod_cpu), float(pod_mem), node['cpu'], node['mem'], node['pnum']]
		output.append(entry)
		# y is useless
		y.append(1)
	dump_svmlight_file(output, y, 'real_test')

	# load data
	x_test, y = load_svmlight_file("real_test")

	group_test = [len(nodes)]
	test_dmatrix = DMatrix(x_test)
	test_dmatrix.set_group(group_test)
	
	# load model from file
	loaded_model = pickle.load(open("xuelengmillion.pickle.dat", "rb"))
	
	# predict
	start_time = time.time()
	print "start to predict with model"
	pred = loaded_model.predict(test_dmatrix)
	end_time = time.time()
	predict_time = end_time - start_time
	print "prediction time:", predict_time
	print pred

	pred_list = pred.tolist()
	selected_node_name = nodes[pred_list.index(min(pred))]['name']
	
	return selected_node_name
