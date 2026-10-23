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

type ModDomHelper struct {
}



func NewModDomHelper() *ModDomHelper {
    this := &ModDomHelper{}
    return this
}


// throws Exception
func ModDomHelper_CreateTextNode(document *Document, name string, text string) *Node {
        var newValueNode *Node = document.CreateElement(name); _ = newValueNode
        var newValueTextNode *Node = document.CreateTextNode(text); _ = newValueTextNode
        newValueNode.AppendChild(newValueTextNode)
        return newValueNode
}


// throws Exception
func ModDomHelper_CreateValueNode(document *Document, text string) *Node {
        return ModDomHelper_CreateTextNode(document, DomData_VALUE, text)
}


// throws Exception
func ModDomHelper_CreateNameNode(document *Document, text string) *Node {
        return ModDomHelper_CreateTextNode(document, DomData_NAME, text)
}


// throws Exception
func ModDomHelper_CreateIndexNode(document *Document, index int) *Node {
        return ModDomHelper_CreateTextNode(document, DomData_INDEX, fmt.Sprint(int(index)))
}


// throws Exception
func ModDomHelper_CreateNameValueNodes_1(document *Document, elementName string, nameText string, valueText string) *Node {
        var newNode *Node = document.CreateElement(elementName); _ = newNode
        newNode.AppendChild(ModDomHelper_CreateNameNode(document, nameText))
        newNode.AppendChild(ModDomHelper_CreateValueNode(document, valueText))
        return newNode
}

// Java overloads of ModDomHelper.createNameValueNodes
func ModDomHelper_CreateNameValueNodes(args ...any) *Node {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 4 && Java_IsArg[*Document](args[0], pass) && Java_IsArg[string](args[1], pass) && Java_IsArg[string](args[2], pass) && Java_IsArg[string](args[3], pass) {
            return ModDomHelper_CreateNameValueNodes_1(Java_Arg[*Document](args[0]), Java_Arg[string](args[1]), Java_Arg[string](args[2]), Java_Arg[string](args[3]))
        }
        if len(args) == 3 && Java_IsArg[*Document](args[0], pass) && Java_IsArg[string](args[1], pass) && Java_IsArg[string](args[2], pass) {
            return ModDomHelper_CreateNameValueNodes_2(Java_Arg[*Document](args[0]), Java_Arg[string](args[1]), Java_Arg[string](args[2]))
        }
        if len(args) == 3 && Java_IsArg[*Document](args[0], pass) && Java_IsArg[string](args[1], pass) && Java_IsArg[*HashMap](args[2], pass) {
            return ModDomHelper_CreateNameValueNodes_3(Java_Arg[*Document](args[0]), Java_Arg[string](args[1]), Java_Arg[*HashMap](args[2]))
        }
        if len(args) == 3 && Java_IsArg[*Document](args[0], pass) && Java_IsArg[*Node](args[1], pass) && Java_IsArg[*HashMap](args[2], pass) {
            return ModDomHelper_CreateNameValueNodes_4(Java_Arg[*Document](args[0]), Java_Arg[*Node](args[1]), Java_Arg[*HashMap](args[2]))
        }
        if len(args) == 4 && Java_IsArg[*Document](args[0], pass) && Java_IsArg[string](args[1], pass) && Java_IsArg[string](args[2], pass) && Java_IsArg[IBasicArrayList](args[3], pass) {
            return ModDomHelper_CreateNameValueNodes_5(Java_Arg[*Document](args[0]), Java_Arg[string](args[1]), Java_Arg[string](args[2]), Java_Arg[IBasicArrayList](args[3]))
        }
    }
    panic("No Java overload of ModDomHelper.createNameValueNodes for the arguments")
}


// throws Exception
func ModDomHelper_CreateNameValueNodes_2(document *Document, elementNameAndNameText string, valueText string) *Node {
        var newNode *Node = ModDomHelper_CreateNameValueNodes(document, elementNameAndNameText, elementNameAndNameText, valueText); _ = newNode
        return newNode
}


// throws Exception
func ModDomHelper_CreateNameValueIndexNodes_1(document *Document, elementNameAndNameText string, index int, valueText string) *Node {
        var newNode *Node = ModDomHelper_CreateNameValueNodes(document, elementNameAndNameText, valueText); _ = newNode
        newNode.AppendChild(ModDomHelper_CreateIndexNode(document, index))
        return newNode
}

// Java overloads of ModDomHelper.createNameValueIndexNodes
func ModDomHelper_CreateNameValueIndexNodes(args ...any) *Node {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 4 && Java_IsArg[*Document](args[0], pass) && Java_IsArg[string](args[1], pass) && Java_IsArg[int](args[2], pass) && Java_IsArg[string](args[3], pass) {
            return ModDomHelper_CreateNameValueIndexNodes_1(Java_Arg[*Document](args[0]), Java_Arg[string](args[1]), Java_Arg[int](args[2]), Java_Arg[string](args[3]))
        }
        if len(args) == 4 && Java_IsArg[*Document](args[0], pass) && Java_IsArg[string](args[1], pass) && Java_IsArg[string](args[2], pass) && Java_IsArg[IBasicArrayList](args[3], pass) {
            return ModDomHelper_CreateNameValueIndexNodes_2(Java_Arg[*Document](args[0]), Java_Arg[string](args[1]), Java_Arg[string](args[2]), Java_Arg[IBasicArrayList](args[3]))
        }
    }
    panic("No Java overload of ModDomHelper.createNameValueIndexNodes for the arguments")
}


// throws Exception
func ModDomHelper_CreateNameValueNodes_3(document *Document, rootNode string, hashMap *HashMap) *Node {
        var node *Node = document.CreateElement(rootNode); _ = node
        return ModDomHelper_CreateNameValueNodes(document, node, hashMap)
}


// throws Exception
func ModDomHelper_CreateNameValueNodes_4(document *Document, node *Node, hashMap *HashMap) *Node {
        var stringUtil *StringUtil = StringUtil_GetInstance(); _ = stringUtil
        var keySet Set = hashMap.KeySet(); _ = keySet
        var keyArray []any = keySet.ToArray(); _ = keyArray
        var size int = len(keyArray); _ = size
        for i := int(0); i < size; i++ {
            var objectKey any = keyArray[i]; _ = objectKey
            var name string = any(objectKey).(string); _ = name
            var anyType any = any(hashMap.Get(objectKey)).(any); _ = anyType
            if anyType == nil {
                var value string = stringUtil.EMPTY_STRING; _ = value
                node.AppendChild(ModDomHelper_CreateNameValueNodes(document, name, value))
            } else if func() bool { _, ok := any(anyType).(string); return ok }() {
                var value string = stringUtil.GetNonNull(any(anyType).(string)); _ = value
                node.AppendChild(ModDomHelper_CreateNameValueNodes(document, name, value))
            } else if func() bool { _, ok := any(anyType).(*HashMap); return ok }() {
                node.AppendChild(ModDomHelper_CreateNameValueNodes(document, name, hashMap))
            } else {
                panic(errors.New(fmt.Sprint("HashMap value is the wrong instance and is: " + fmt.Sprint(ClassUtil_ViewAll(anyType, CommonSeps_GetInstance().NEW_LINE)))))
            }
        }
        return node
}


// throws Exception
func ModDomHelper_CreateNameValueNodes_5(document *Document, rootNodeName string, duplicateNodeName string, valueVector IBasicArrayList) *Node {
        var stringUtil *StringUtil = StringUtil_GetInstance(); _ = stringUtil
        var node *Node = document.CreateElement(rootNodeName); _ = node
        var size int = valueVector.Size(); _ = size
        for i := int(0); i < size; i++ {
            var value string = stringUtil.GetNonNull(any(valueVector.Get(i)).(string)); _ = value
            node.AppendChild(ModDomHelper_CreateNameValueNodes(document, duplicateNodeName, value))
        }
        return node
}


// throws Exception
func ModDomHelper_CreateNameValueIndexNodes_2(document *Document, rootNodeName string, duplicateNodeName string, valueVector IBasicArrayList) *Node {
        var stringUtil *StringUtil = StringUtil_GetInstance(); _ = stringUtil
        var node *Node = document.CreateElement(rootNodeName); _ = node
        var size int = valueVector.Size(); _ = size
        for index := int(0); index < size; index++ {
            var value string = stringUtil.GetNonNull(any(valueVector.Get(index)).(string)); _ = value
            node.AppendChild(ModDomHelper_CreateNameValueIndexNodes(document, duplicateNodeName, index, value))
        }
        return node
}


// throws Exception
func ModDomHelper_CreateNodeWithValueNodes_1(document *Document, rootNode string, hashMap *HashMap) *Node {
        var node *Node = document.CreateElement(rootNode); _ = node
        return ModDomHelper_CreateNodeWithValueNodes(document, node, hashMap)
}

// Java overloads of ModDomHelper.createNodeWithValueNodes
func ModDomHelper_CreateNodeWithValueNodes(args ...any) *Node {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 3 && Java_IsArg[*Document](args[0], pass) && Java_IsArg[string](args[1], pass) && Java_IsArg[*HashMap](args[2], pass) {
            return ModDomHelper_CreateNodeWithValueNodes_1(Java_Arg[*Document](args[0]), Java_Arg[string](args[1]), Java_Arg[*HashMap](args[2]))
        }
        if len(args) == 3 && Java_IsArg[*Document](args[0], pass) && Java_IsArg[*Node](args[1], pass) && Java_IsArg[*HashMap](args[2], pass) {
            return ModDomHelper_CreateNodeWithValueNodes_2(Java_Arg[*Document](args[0]), Java_Arg[*Node](args[1]), Java_Arg[*HashMap](args[2]))
        }
    }
    panic("No Java overload of ModDomHelper.createNodeWithValueNodes for the arguments")
}


// throws Exception
func ModDomHelper_CreateNodeWithValueNodes_2(document *Document, node *Node, hashMap *HashMap) *Node {
        var stringUtil *StringUtil = StringUtil_GetInstance(); _ = stringUtil
        var keySet Set = hashMap.KeySet(); _ = keySet
        var keyArray []any = keySet.ToArray(); _ = keyArray
        var size int = len(keyArray); _ = size
        for i := int(0); i < size; i++ {
            var name string = any(keyArray[i]).(string); _ = name
            var anyType any = any(hashMap.Get(name)).(any); _ = anyType
            var value string = stringUtil.GetNonNull(any(anyType).(string)); _ = value
            node.AppendChild(ModDomHelper_CreateTextNode(document, name, value))
        }
        return node
}

