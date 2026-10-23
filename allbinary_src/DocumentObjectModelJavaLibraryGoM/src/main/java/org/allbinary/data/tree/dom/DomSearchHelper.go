/*
 *
 * AllBinary Open License Version 1
 * Copyright (c) 2011 AllBinary
 * 
 * By agreeing to this license you and any business entity you represent are
 * legally bound to the AllBinary Open License Version 1 legal agreement.
 * 
 * You may obtain the AllBinary Open License Version 1 legal agreement from
 * AllBinary or the root directory of AllBinary's AllBinary Platform repository.
 * 
 * Created By: Travis Berthelot
 *  
 */
// Generated Code Do Not Modify

package allbinary

import (
    "errors"
    "fmt"
    "math"
    "math/rand"
    "os"
    "reflect"
    "runtime"
    "strconv"
    "strings"
    "time"
    "unicode"
)

var _ = errors.New
var _ = fmt.Sprint
var _ = math.Abs
var _ = rand.Float64
var _ = os.Exit
var _ = reflect.TypeOf
var _ = runtime.GC
var _ = strconv.Itoa
var _ = strings.Contains
var _ = time.Now
var _ = unicode.IsDigit

type DomSearchHelper struct {
}



func NewDomSearchHelper() *DomSearchHelper {
    this := &DomSearchHelper{}
    return this
}


func DomSearchHelper_GetNodeNoThrow(nodeName string, nodeList *NodeList) *Node {
        var numberOfNodes int = nodeList.GetLength(); _ = numberOfNodes
        for index := int(0); index < numberOfNodes; index++ {
            var node *Node = nodeList.Item(index); _ = node
            if node.GetNodeName().CompareTo(nodeName) == 0 {
                return node
            }
        }
        return nil
}


// throws Exception
func DomSearchHelper_GetNode(nodeName string, nodeList *NodeList) *Node {
        var node *Node = DomSearchHelper_GetNodeNoThrow(nodeName, nodeList); _ = node
        if node != nil {
            return node
        } else {
            panic(errors.New(fmt.Sprint(nodeName + " Node Not Found in search")))
        }
}


// throws Exception
func DomSearchHelper_GetChildNodeList(nodeName string, nodeList *NodeList) *NodeList {
        var node *Node = DomSearchHelper_GetNode(nodeName, nodeList); _ = node
        return node.GetChildNodes()
}


// throws Exception
func DomSearchHelper_GetAllNodes(nodeName string, nodeList *NodeList) IBasicArrayList {
        var vector IBasicArrayList = DomSearchHelper_GetAllNodesNoThrow(nodeName, nodeList); _ = vector
        if vector == nil {
            panic(errors.New(fmt.Sprint(nodeName + " Node Not Found in search")))
        } else {
            return vector
        }
}


func DomSearchHelper_GetAllNodesNoThrow(nodeName string, nodeList *NodeList) IBasicArrayList {
        var nodeVector IBasicArrayList = NewBasicArrayListD(); _ = nodeVector
        var numberOfNodes int = nodeList.GetLength(); _ = numberOfNodes
        for index := int(0); index < numberOfNodes; index++ {
            var node *Node = nodeList.Item(index); _ = node
            if node.GetNodeName().CompareTo(nodeName) == 0 {
                nodeVector.Add(node)
            }
        }
        return nodeVector
}

