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

type DomNodeHelper struct {
}



func NewDomNodeHelper() *DomNodeHelper {
    this := &DomNodeHelper{}
    return this
}


func DomNodeHelper_GetFirstChildElement(parentNode *Node) *Node {
        var node *Node = parentNode.GetFirstChild(); _ = node
        if node != nil && node.GetNodeType() != Node_ELEMENT_NODE {
            var nodeList *NodeList = parentNode.GetChildNodes(); _ = nodeList
            for index := int(0); index < nodeList.GetLength(); index++ {
                node = nodeList.Item(index)
                if node.GetNodeType() == 1 {
                    break
                }
            }
        }
        return node
}


func DomNodeHelper_GetTextNodeValue_1(node *Node) string {
        var nodeTextNode *Node = node.GetFirstChild(); _ = nodeTextNode
        if nodeTextNode != nil {
            return nodeTextNode.GetNodeValue()
        } else {
            return StringUtil_GetInstance().EMPTY_STRING
        }
}

// Java overloads of DomNodeHelper.getTextNodeValue
func DomNodeHelper_GetTextNodeValue(args ...any) string {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[*Node](args[0], pass) {
            return DomNodeHelper_GetTextNodeValue_1(Java_Arg[*Node](args[0]))
        }
        if len(args) == 2 && Java_IsArg[string](args[0], pass) && Java_IsArg[*NodeList](args[1], pass) {
            return DomNodeHelper_GetTextNodeValue_2(Java_Arg[string](args[0]), Java_Arg[*NodeList](args[1]))
        }
    }
    panic("No Java overload of DomNodeHelper.getTextNodeValue for the arguments")
}


// throws Exception
func DomNodeHelper_GetTextNodeValue_2(nodeName string, nodeList *NodeList) string {
        var node *Node = DomSearchHelper_GetNode(nodeName, nodeList); _ = node
        return DomNodeHelper_GetTextNodeValue(node)
}


func DomNodeHelper_GetTextNodesValue(node *Node) string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        var nodeList *NodeList = node.GetChildNodes(); _ = nodeList
        for index := int(0); index < nodeList.GetLength(); index++ {
            var nodeTextNode *Node = nodeList.Item(index); _ = nodeTextNode
            if nodeTextNode.GetNodeType() == Node_TEXT_NODE {
                if nodeTextNode != nil {
                    stringBuffer.Append(nodeTextNode.GetNodeValue())
                }
            }
        }
        return fmt.Sprint(stringBuffer)
}

