package main

import (
    "github.com/sbinet/go-python"
    "fmt"
)

func init() {
    err := python.Initialize()
    if err != nil {
        panic(err.Error())
    }
}

var PyStr = python.PyString_FromString
var GoStr = python.PyString_AS_STRING

func main() {
    fmt.Printf("python version:" + python.Py_GetProgramName() + " \n")

    hello := ImportModule("/home/user/runqing", "hello")
    fmt.Printf("[MODULE] repr(hello) = %s\n", GoStr(hello.Repr()))

    var n_name string = "node1"
    var n_cpu string = "123"
    var n_mem string = "888"
    var n_pnum string = "110"
    var p_cpu string = "50"
    var p_mem string = "300"

    // init a list
    Nodes := python.PyList_New(0)

    // replace the condition with the nodes list, can use range
    // process nodeinfo to store them in a list
    for i := 1; i < 3; i++ {
        Nodes = ProcessNode(n_name, n_cpu, n_mem, n_pnum, hello, Nodes)
    }
    
    fmt.Printf("[CALL] transformDataFormat('NodeList') = %s\n", Nodes)

    // processed by the ranker and get the selected node's name
    NodeName := ProcessByRanker(p_cpu, p_mem, hello, Nodes)
    fmt.Printf("[CALL] transformDataFormat('SelectedNodeName') = %s\n", GoStr(NodeName))
}

func ProcessNode(n_name string, n_cpu string, n_mem string, n_pnum string, hello *python.PyObject, nodes *python.PyObject) (*python.PyObject) {
    // set parameters to a tuple
    NArgs := python.PyTuple_New(5)
    python.PyTuple_SetItem(NArgs, 0, PyStr(n_name))
    python.PyTuple_SetItem(NArgs, 1, PyStr(n_cpu))
    python.PyTuple_SetItem(NArgs, 2, PyStr(n_mem))
    python.PyTuple_SetItem(NArgs, 3, PyStr(n_pnum))
    python.PyTuple_SetItem(NArgs, 4, nodes)

    // get func name and call it
    transformDataFormat := hello.GetAttrString("transformDataFormat")
    nodes = transformDataFormat.Call(NArgs, python.Py_None)
    return nodes
}

func ProcessByRanker(p_cpu string, p_mem string, hello *python.PyObject, nodes *python.PyObject) (*python.PyObject) {
    // set parameters to a tuple
    RArgs := python.PyTuple_New(3)
    python.PyTuple_SetItem(RArgs, 0, PyStr(p_cpu))
    python.PyTuple_SetItem(RArgs, 1, PyStr(p_mem))
    python.PyTuple_SetItem(RArgs, 2, nodes)

    // get func name and call it
    RankNodesWithModel := hello.GetAttrString("RankNodesWithModel")
    NodeName := RankNodesWithModel.Call(RArgs, python.Py_None)
    return NodeName
}

// ImportModule will import python module from given directory
func ImportModule(dir, name string) *python.PyObject {
    sysModule := python.PyImport_ImportModule("sys") // import sys
    path := sysModule.GetAttrString("path")                    // path = sys.path
    python.PyList_Insert(path, 0, PyStr(dir))                     // path.insert(0, dir)
    return python.PyImport_ImportModule(name)            // return __import__(name)
}
