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

    Nodes := python.PyList_New(0)

    // replace the condition with the nodes list, can use range
    for i := 1; i < 3; i++ {
        Args := python.PyTuple_New(5)
        python.PyTuple_SetItem(Args, 0, PyStr(n_name))
        python.PyTuple_SetItem(Args, 1, PyStr(n_cpu))
        python.PyTuple_SetItem(Args, 2, PyStr(n_mem))
        python.PyTuple_SetItem(Args, 3, PyStr(n_pnum))
        python.PyTuple_SetItem(Args, 4, Nodes)

        transformDataFormat := hello.GetAttrString("transformDataFormat")
        Nodes = transformDataFormat.Call(Args, python.Py_None)
    }
    fmt.Printf("[CALL] transformDataFormat('NodeList') = %s\n", Nodes)

    RArgs := python.PyTuple_New(3)
    python.PyTuple_SetItem(RArgs, 0, PyStr(p_cpu))
    python.PyTuple_SetItem(RArgs, 1, PyStr(p_mem))
    python.PyTuple_SetItem(RArgs, 2, Nodes)

    RankNodesWithModel := hello.GetAttrString("RankNodesWithModel")
    NodeName := RankNodesWithModel.Call(RArgs, python.Py_None)
    fmt.Printf("[CALL] transformDataFormat('SelectedNodeName') = %s\n", GoStr(NodeName))
}

// ImportModule will import python module from given directory
func ImportModule(dir, name string) *python.PyObject {
    sysModule := python.PyImport_ImportModule("sys") // import sys
    path := sysModule.GetAttrString("path")                    // path = sys.path
    python.PyList_Insert(path, 0, PyStr(dir))                     // path.insert(0, dir)
    return python.PyImport_ImportModule(name)            // return __import__(name)
}
