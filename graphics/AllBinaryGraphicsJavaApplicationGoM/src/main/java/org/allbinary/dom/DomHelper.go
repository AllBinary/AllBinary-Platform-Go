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

type DomHelper struct {
    LogUtil *LogUtil
}

var DomHelper_Instance *DomHelper = NewDomHelper()

func (this *DomHelper) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func DomHelper_GetInstance() *DomHelper {
        return DomHelper_Instance
}


func NewDomHelper() *DomHelper {
    this := &DomHelper{}
    this.initFields()
    return this
}


func (this *DomHelper) CreateDocument_1() *Document {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 *Document) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            var factory *DocumentBuilderFactory = DocumentBuilderFactory_NewInstance(); _ = factory
            var builder *DocumentBuilder = factory.NewDocumentBuilder(); _ = builder
            var document *Document = builder.NewDocument(); _ = document
            tryControl1, tryResult1 = 1, document; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

// Java overloads of DomHelper.createDocument
func (this *DomHelper) CreateDocument(args ...any) *Document {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 0 {
            return this.CreateDocument_1()
        }
        if len(args) == 1 && Java_IsArg[*File](args[0], pass) {
            return this.CreateDocument_2(Java_Arg[*File](args[0]))
        }
    }
    panic("No Java overload of DomHelper.createDocument for the arguments")
}


// throws Exception
func (this *DomHelper) CreateDocument_2(xmlFile *File) *Document {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 *Document) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            }()
            var factory *DocumentBuilderFactory = DocumentBuilderFactory_NewInstance(); _ = factory
            var builder *DocumentBuilder = factory.NewDocumentBuilder(); _ = builder
            var document *Document = builder.Parse(xmlFile); _ = document
            tryControl1, tryResult1 = 1, document; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws Exception
func (this *DomHelper) ToString(document *Document) string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            }()
            var domSource *DOMSource = NewDOMSource(document); _ = domSource
            var byteArrayOutputStream *strings.Builder = &strings.Builder{}; _ = byteArrayOutputStream
            var streamResult *StreamResult = NewStreamResult(byteArrayOutputStream); _ = streamResult
            var copyTransformerFactory *TransformerFactory = TransformerFactory_NewInstance(); _ = copyTransformerFactory
            var copyTransformer *Transformer = copyTransformerFactory.NewTransformer(); _ = copyTransformer
            copyTransformer.Transform(domSource, streamResult)
            tryControl1, tryResult1 = 1, byteArrayOutputStream.String(); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws Exception
func (this *DomHelper) Save(file *File, document *Document) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            }()
            var copyTransformerFactory *TransformerFactory = TransformerFactory_NewInstance(); _ = copyTransformerFactory
            var copyTransformer *Transformer = copyTransformerFactory.NewTransformer(); _ = copyTransformer
            var domSource *DOMSource = NewDOMSource(document); _ = domSource
            file.CreateNewFile()
            var streamResult *StreamResult = NewStreamResult(file); _ = streamResult
            copyTransformer.Transform(domSource, streamResult)
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func (this *DomHelper) GetChildNodeList(nodeName string, nodeList *NodeList) *NodeList {
        var numberOfNodes int = nodeList.GetLength(); _ = numberOfNodes
        for index := int(0); index < numberOfNodes; index++ {
            var node *Node = nodeList.Item(index); _ = node
            if node.GetNodeName().CompareTo(nodeName) == 0 {
                return node.GetChildNodes()
            }
        }
        panic(errors.New(fmt.Sprint(nodeName + " Node Not Found")))
}


// throws Exception
func (this *DomHelper) GetChildrenWithoutTextNodes(nodeName string, nodeList *NodeList) IBasicArrayList {
        var list IBasicArrayList = NewBasicArrayListD(); _ = list
        var childNodeList *NodeList = this.GetChildNodeList(nodeName, nodeList); _ = childNodeList
        var numberOfChildren int = childNodeList.GetLength(); _ = numberOfChildren
        for index := int(0); index < numberOfChildren; index++ {
            var node *Node = childNodeList.Item(index); _ = node
            if node.GetNodeType() != Node_TEXT_NODE {
                list.Add(node)
            }
        }
        return list
}


// throws Exception
func (this *DomHelper) GetWithoutTextNodes(nodeList *NodeList) IBasicArrayList {
        var list IBasicArrayList = NewBasicArrayListD(); _ = list
        var numberOfChildren int = nodeList.GetLength(); _ = numberOfChildren
        for index := int(0); index < numberOfChildren; index++ {
            var node *Node = nodeList.Item(index); _ = node
            if node.GetNodeType() != Node_TEXT_NODE {
                list.Add(node)
            }
        }
        return list
}


// throws Exception
func (this *DomHelper) SearchNodeList_1(nodeName string, nodeList *NodeList) *Node {
        var numberOfNodes int = nodeList.GetLength(); _ = numberOfNodes
        for index := int(0); index < numberOfNodes; index++ {
            var node *Node = nodeList.Item(index); _ = node
            this.LogUtil.PutF("NodeName: " + fmt.Sprint(node.GetNodeName()), this, "searchNodeList")
            if node.GetNodeName().CompareTo(nodeName) == 0 {
                return node
            }
        }
        panic(errors.New(fmt.Sprint(nodeName + " Node Not Found in search")))
}

// Java overloads of DomHelper.searchNodeList
func (this *DomHelper) SearchNodeList(args ...any) *Node {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 2 && Java_IsArg[string](args[0], pass) && Java_IsArg[*NodeList](args[1], pass) {
            return this.SearchNodeList_1(Java_Arg[string](args[0]), Java_Arg[*NodeList](args[1]))
        }
        if len(args) == 2 && Java_IsArg[string](args[0], pass) && Java_IsArg[IBasicArrayList](args[1], pass) {
            return this.SearchNodeList_2(Java_Arg[string](args[0]), Java_Arg[IBasicArrayList](args[1]))
        }
    }
    panic("No Java overload of DomHelper.searchNodeList for the arguments")
}


// throws Exception
func (this *DomHelper) SearchNodeList_2(nodeName string, nodeList IBasicArrayList) *Node {
        var numberOfNodes int = nodeList.Size(); _ = numberOfNodes
        for index := int(0); index < numberOfNodes; index++ {
            var node *Node = any(nodeList.Get(index)).(*Node); _ = node
            this.LogUtil.PutF("NodeName: " + fmt.Sprint(node.GetNodeName()), this, "searchNodeList")
            if node.GetNodeName().CompareTo(nodeName) == 0 {
                return node
            }
        }
        panic(errors.New(fmt.Sprint(nodeName + " Node Not Found in search")))
}

