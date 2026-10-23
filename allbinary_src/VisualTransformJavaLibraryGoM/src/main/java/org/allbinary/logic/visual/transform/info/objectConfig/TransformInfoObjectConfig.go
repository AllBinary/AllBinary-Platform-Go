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

type TransformInfoObjectConfig struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    OwnerTransformInfoInterface TransformInfoInterface
    Document *Document
    OutputTypeName string
}

// Implements TransformInfoObjectConfigInterface

// The Java class TransformInfoObjectConfig has subclasses so its variables are this interface
type ITransformInfoObjectConfig interface {
    TransformInfoObjectConfig_() *TransformInfoObjectConfig
    CreateDocument()
    GetTransformInfoInterface() TransformInfoInterface
    ToXmlDoc() *Document
    SetDocument(document *Document)
    ContainsView(transformInfoInterface TransformInfoInterface) bool
    GetTemplateAttributes() *NamedNodeMap
    GetName() string
    SetName(name string)
    GetNodeVector(nodeName string) IBasicArrayList
    GetTransformDomNodes(args ...any) IBasicArrayList
    GetTransforms(args ...any) IBasicArrayList
    GetTransformsGroup(group string) IBasicArrayList
    GetGroupTransforms() IBasicArrayList
    GetParentTransforms() IBasicArrayList
    GetRootNode() *Node
    SetOutputTypeName(outputTypeName string)
    GetOutputTypeName() string
    GetInputOutputTypeName() string
    GetInputOutputTypeFile() string
    GetImportUriPath() string
    ToString() string
}

func (this *TransformInfoObjectConfig) TransformInfoObjectConfig_() *TransformInfoObjectConfig {
    return this
}


func (this *TransformInfoObjectConfig) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
    
    
}




// Java overloads of new TransformInfoObjectConfig
func NewTransformInfoObjectConfig(args ...any) *TransformInfoObjectConfig {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[TransformInfoInterface](args[0], pass) {
            return NewTransformInfoObjectConfig_1(Java_Arg[TransformInfoInterface](args[0]))
        }
        if len(args) == 2 && Java_IsArg[TransformInfoInterface](args[0], pass) && Java_IsArg[*Document](args[1], pass) {
            return NewTransformInfoObjectConfig_2(Java_Arg[TransformInfoInterface](args[0]), Java_Arg[*Document](args[1]))
        }
        if len(args) == 3 && Java_IsArg[TransformInfoInterface](args[0], pass) && Java_IsArg[string](args[1], pass) && Java_IsArg[string](args[2], pass) {
            return NewTransformInfoObjectConfig_3(Java_Arg[TransformInfoInterface](args[0]), Java_Arg[string](args[1]), Java_Arg[string](args[2]))
        }
    }
    panic("No Java overload of new TransformInfoObjectConfig for the arguments")
}
func NewTransformInfoObjectConfig_1(transformInfoInterface TransformInfoInterface) *TransformInfoObjectConfig {
    this := &TransformInfoObjectConfig{}
    this.initFields()
        this.OwnerTransformInfoInterface = transformInfoInterface
        this.CreateDocument()
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append("TransformInfo: ")
            if this.OwnerTransformInfoInterface != nil {
                stringBuffer.Append(this.OwnerTransformInfoInterface.GetName())
            } else {
                stringBuffer.Append("No Owner!?#@")
            }
            stringBuffer.Append("\nConstructed with document: ")
            stringBuffer.Append(this.ToString())
            this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "Constructor(TransformInfoInterface)")
        }
    return this
}


func NewTransformInfoObjectConfig_2(transformInfoInterface TransformInfoInterface, document *Document) *TransformInfoObjectConfig {
    this := &TransformInfoObjectConfig{}
    this.initFields()
        this.OwnerTransformInfoInterface = transformInfoInterface
        this.Document = document
        var node *Node = DomSearchHelper_GetNodeNoThrow(OutputTypeData_GetInstance().NAME, this.GetRootNode().GetChildNodes()); _ = node
        if node != nil {
            this.SetOutputTypeName(DomNodeHelper_GetTextNodeValue(node))
        }
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append("TransformInfo: ")
            if this.OwnerTransformInfoInterface != nil {
                stringBuffer.Append(this.OwnerTransformInfoInterface.GetName())
            } else {
                stringBuffer.Append("No Owner!?#@")
            }
            stringBuffer.Append("\nConstructed with document: ")
            stringBuffer.Append(this.ToString())
            this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "Constructor(TransformInfoInterface, Document)")
        }
    return this
}


func NewTransformInfoObjectConfig_3(transformInfoInterface TransformInfoInterface, name string, type_ string) *TransformInfoObjectConfig {
    this := &TransformInfoObjectConfig{}
    this.initFields()
        this.OwnerTransformInfoInterface = transformInfoInterface
        this.CreateDocument()
        var configNameAttr *Attr = this.Document.CreateAttribute(TransformInfoObjectConfigData_GetInstance().NAME); _ = configNameAttr
        configNameAttr.SetValue(name)
        var objectConfigNode *Node = this.Document.GetElementsByTagName(TransformInfoObjectConfigData_GetInstance().NAME).Item(0); _ = objectConfigNode
        objectConfigNode.AppendChild(configNameAttr)
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append("TransformInfo: ")
            if this.OwnerTransformInfoInterface != nil {
                stringBuffer.Append(this.OwnerTransformInfoInterface.GetName())
            } else {
                stringBuffer.Append("No Owner!?#@")
            }
            stringBuffer.Append("\nConstructed with document: ")
            stringBuffer.Append(this.ToString())
            this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "Constructor(TransformInfoInterface, name, type)")
        }
    return this
}


func (this *TransformInfoObjectConfig) CreateDocument() {
        this.Document = DomDocumentHelper_Create()
        var objectConfigNode *Node = this.Document.CreateElement(TransformInfoObjectConfigData_GetInstance().NAME); _ = objectConfigNode
        this.Document.AppendChild(objectConfigNode)
}


func (this *TransformInfoObjectConfig) GetTransformInfoInterface() TransformInfoInterface {
        return this.OwnerTransformInfoInterface
}


// throws Exception
func (this *TransformInfoObjectConfig) ToXmlDoc() *Document {
        return this.Document
}


func (this *TransformInfoObjectConfig) SetDocument(document *Document) {
        this.Document = document
}


func (this *TransformInfoObjectConfig) ContainsView(transformInfoInterface TransformInfoInterface) bool {
        var objectConfigNode *Node = this.Document.GetElementsByTagName(TransformInfoObjectConfigData_GetInstance().NAME).Item(0); _ = objectConfigNode
        var viewNodeList *NodeList = objectConfigNode.GetChildNodes(); _ = viewNodeList
        var numberOfViews int = viewNodeList.GetLength(); _ = numberOfViews
        for index := int(0); index < numberOfViews; index++ {
            var viewNode *Node = viewNodeList.Item(index); _ = viewNode
            var viewAttributes *NamedNodeMap = viewNode.GetAttributes(); _ = viewAttributes
            var attrNode *Attr = any(viewAttributes.GetNamedItem(TransformInfoData_GetInstance().NAME)).(*Attr); _ = attrNode
            if strings.Compare(transformInfoInterface.GetName(), attrNode.GetValue()) == 0 {
                return true
            }
        }
        return false
}


// throws Exception
func (this *TransformInfoObjectConfig) GetTemplateAttributes() *NamedNodeMap {
        var componentNode *Node = this.Document.GetElementsByTagName(TransformInfoObjectConfigData_GetInstance().NAME).Item(0); _ = componentNode
        return componentNode.GetAttributes()
}


// throws Exception
func (this *TransformInfoObjectConfig) GetName() string {
        var attrNode *Attr = any(this.GetTemplateAttributes().GetNamedItem(TransformInfoObjectConfigData_GetInstance().NAME)).(*Attr); _ = attrNode
        return attrNode.GetValue()
}


// throws Exception
func (this *TransformInfoObjectConfig) SetName(name string) {
        var attrNode *Attr = any(this.GetTemplateAttributes().GetNamedItem(TransformInfoObjectConfigData_GetInstance().NAME)).(*Attr); _ = attrNode
        attrNode.SetValue(name)
}


// throws Exception
func (this *TransformInfoObjectConfig) GetNodeVector(nodeName string) IBasicArrayList {
        var componentsNodeList *NodeList = this.Document.GetElementsByTagName(nodeName); _ = componentsNodeList
        if componentsNodeList != nil && componentsNodeList.GetLength() > 0 {
            var viewNodeVector IBasicArrayList = DomSearchHelper_GetAllNodes(TransformInfoData_GetInstance().NAME, componentsNodeList.Item(0).GetChildNodes()); _ = viewNodeVector
            var numberOfViews int = viewNodeVector.Size(); _ = numberOfViews
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                stringBuffer.Append("Number Of ")
                stringBuffer.Append(nodeName)
                stringBuffer.Append(" Nodes: ")
                stringBuffer.Appendint(numberOfViews)
                this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "getNodeVector(nodename)")
            }
            return viewNodeVector
        }
        return NewBasicArrayListD()
}


// throws Exception
func (this *TransformInfoObjectConfig) GetTransformDomNodes_1(nodeName string) IBasicArrayList {
        var viewVector IBasicArrayList = NewBasicArrayListD(); _ = viewVector
        var viewNodeVector IBasicArrayList = this.GetNodeVector(nodeName); _ = viewNodeVector
        var size int = viewNodeVector.Size(); _ = size
        for index := int(0); index < size; index++ {
            var viewNode *Node = any(viewNodeVector.Get(index)).(*Node); _ = viewNode
            viewVector.Add(NewTransformInfoDomNode(viewNode))
        }
        return viewVector
}

// Java overloads of TransformInfoObjectConfig.getTransformDomNodes
func (this *TransformInfoObjectConfig) GetTransformDomNodes(args ...any) IBasicArrayList {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[string](args[0], pass) {
            return this.GetTransformDomNodes_1(Java_Arg[string](args[0]))
        }
        if len(args) == 0 {
            return this.GetTransformDomNodes_2()
        }
    }
    panic("No Java overload of TransformInfoObjectConfig.getTransformDomNodes for the arguments")
}


// throws Exception
func (this *TransformInfoObjectConfig) GetTransforms_1(nodeName string) IBasicArrayList {
        var viewVector IBasicArrayList = NewBasicArrayListD(); _ = viewVector
        var viewNodeVector IBasicArrayList = this.GetNodeVector(nodeName); _ = viewNodeVector
        var size int = viewNodeVector.Size(); _ = size
        for index := int(0); index < size; index++ {
            var viewNode *Node = any(viewNodeVector.Get(index)).(*Node); _ = viewNode
            viewVector.Add(NewTransformInfoDomNode(viewNode).GetTransformInfoInterface())
        }
        return viewVector
}

// Java overloads of TransformInfoObjectConfig.getTransforms
func (this *TransformInfoObjectConfig) GetTransforms(args ...any) IBasicArrayList {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[string](args[0], pass) {
            return this.GetTransforms_1(Java_Arg[string](args[0]))
        }
        if len(args) == 0 {
            return this.GetTransforms_2()
        }
    }
    panic("No Java overload of TransformInfoObjectConfig.getTransforms for the arguments")
}


// throws Exception
func (this *TransformInfoObjectConfig) GetTransformsGroup(group string) IBasicArrayList {
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            this.LogUtil.PutF("Started: " + group, this, "getTransformsGroup()")
        }
        var viewVector IBasicArrayList = NewBasicArrayListD(); _ = viewVector
        var GROUP string = TransformInfosData_GetInstance().GROUP; _ = GROUP
        var componentsNodeList *NodeList = this.Document.GetElementsByTagName(GROUP); _ = componentsNodeList
        if componentsNodeList != nil && componentsNodeList.GetLength() > 0 {
            var componentsNode *Node = componentsNodeList.Item(0); _ = componentsNode
            var length int = componentsNodeList.GetLength(); _ = length
            for index := int(0); index < length; index++ {
                var node *Node = componentsNodeList.Item(index); _ = node
                var attributes *NamedNodeMap = node.GetAttributes(); _ = attributes
                var attrNode *Attr = any(attributes.GetNamedItem(GROUP)).(*Attr); _ = attrNode
                var value string = attrNode.GetValue(); _ = value
                if strings.Compare(value, group) == 0 {
                    componentsNode = node
                    break
                }
            }
            var viewNodeVector IBasicArrayList = DomSearchHelper_GetAllNodes(TransformInfoData_GetInstance().NAME, componentsNode.GetChildNodes()); _ = viewNodeVector
            var numberOfViews int = viewNodeVector.Size(); _ = numberOfViews
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                this.LogUtil.PutF("Number Of Nodes: " + fmt.Sprint(numberOfViews), this, "getTransformsGroup()")
            }
            var size int = viewNodeVector.Size(); _ = size
            for index := int(0); index < size; index++ {
                var viewNode *Node = any(viewNodeVector.Get(index)).(*Node); _ = viewNode
                viewVector.Add(NewTransformInfoDomNode(viewNode))
            }
        } else {
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                if componentsNodeList == nil {
                    this.LogUtil.PutF("Number Of Nodes: NULL", this, "getTransformsGroup()")
                } else {
                    this.LogUtil.PutF("Number Of Nodes: 0", this, "getTransformsGroup()")
                }
            }
        }
        return viewVector
}


// throws Exception
func (this *TransformInfoObjectConfig) GetTransformDomNodes_2() IBasicArrayList {
        return this.GetTransformDomNodes(TransformInfosData_GetInstance().GROUP)
}


// throws Exception
func (this *TransformInfoObjectConfig) GetTransforms_2() IBasicArrayList {
        return this.GetTransforms(TransformInfosData_GetInstance().NAME)
}


// throws Exception
func (this *TransformInfoObjectConfig) GetGroupTransforms() IBasicArrayList {
        return this.GetTransforms(TransformInfosData_GetInstance().GROUP)
}


// throws Exception
func (this *TransformInfoObjectConfig) GetParentTransforms() IBasicArrayList {
        return this.GetTransforms(TransformInfoData_GetInstance().PARENT)
}


func (this *TransformInfoObjectConfig) GetRootNode() *Node {
        return this.Document.GetElementsByTagName(TransformInfoObjectConfigData_GetInstance().NAME).Item(0)
}


func (this *TransformInfoObjectConfig) SetOutputTypeName(outputTypeName string) {
        this.OutputTypeName = outputTypeName
}


// throws Exception
func (this *TransformInfoObjectConfig) GetOutputTypeName() string {
        return this.OutputTypeName
}


// throws Exception
func (this *TransformInfoObjectConfig) GetInputOutputTypeName() string {
        var node *Node = DomSearchHelper_GetNode(InputOutputTypeData_GetInstance().NAME, this.GetRootNode().GetChildNodes()); _ = node
        return DomNodeHelper_GetTextNodeValue(node)
}


// throws Exception
func (this *TransformInfoObjectConfig) GetInputOutputTypeFile() string {
        var node *Node = DomSearchHelper_GetNode(InputOutputTypeData_GetInstance().FILE, this.GetRootNode().GetChildNodes()); _ = node
        return DomNodeHelper_GetTextNodeValue(node)
}


// throws Exception
func (this *TransformInfoObjectConfig) GetImportUriPath() string {
        var node *Node = DomSearchHelper_GetNode(XslData_GetInstance().ROOT_IMPORT_URI, this.GetRootNode().GetChildNodes()); _ = node
        return DomNodeHelper_GetTextNodeValue(node)
}


func (this *TransformInfoObjectConfig) ToString() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEWERROR) {
                this.LogUtil.PutF("DOM Document error", this, this.CommonStrings.TOSTRING)
            }
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            if this.Document != nil {
                tryControl1, tryResult1 = 1, XmlDocumentHelper_ToString(this.Document); return
            } else {
                tryControl1, tryResult1 = 1, nil; return
            }
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

func (this *TransformInfoObjectConfig) String() string {
    return this.ToString()
}

