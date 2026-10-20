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

type TransformInfoDomNode struct {
    LogUtil *LogUtil
    TransformInfoInterface TransformInfoInterface
    MappedName string
}


func (this *TransformInfoDomNode) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
}


func NewTransformInfoDomNode(node *Node) *TransformInfoDomNode {
    this := &TransformInfoDomNode{}
    this.initFields()
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            this.LogUtil.PutF("Constructing TransformInfo From TransformInfoDomNode: " + fmt.Sprint(node.GetNodeName()), this, "TransformInfoDomNode(Node node)")
        }
        var transformInfoData *TransformInfoData = TransformInfoData_GetInstance(); _ = transformInfoData
        var attributes *NamedNodeMap = node.GetAttributes(); _ = attributes
        var attrNode *Attr = any(attributes.GetNamedItem(transformInfoData.NAME)).(*Attr); _ = attrNode
        var name string = attrNode.GetValue(); _ = name
        var nodeList *NodeList = node.GetChildNodes(); _ = nodeList
        if nodeList.GetLength() > 1 {
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                stringBuffer.Append("Constructing Complete TransformInfo Node For: ")
                stringBuffer.Append(name)
                stringBuffer.Append(" with ")
                stringBuffer.Appendint(int(nodeList.GetLength()))
                stringBuffer.Append(" children")
                this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "TransformInfoObjectConfigComponent(Node node)")
            }
            var objectFileNameNode *Node = DomSearchHelper_GetNode(transformInfoData.OBJECTFILENAME, nodeList); _ = objectFileNameNode
            var objectFileName string = DomNodeHelper_GetTextNodeValue(objectFileNameNode); _ = objectFileName
            var objectConfigFileNameNode *Node = DomSearchHelper_GetNode(transformInfoData.OBJECTCONFIGFILENAME, nodeList); _ = objectConfigFileNameNode
            var objectConfigFileName string = DomNodeHelper_GetTextNodeValue(objectConfigFileNameNode); _ = objectConfigFileName
            var templateFileNameNode *Node = DomSearchHelper_GetNode(transformInfoData.TEMPLATEFILENAME, nodeList); _ = templateFileNameNode
            var templateFileName string = DomNodeHelper_GetTextNodeValue(templateFileNameNode); _ = templateFileName
            var dataFileNameNode *Node = DomSearchHelper_GetNode(transformInfoData.DATAFILENAME, nodeList); _ = dataFileNameNode
            var dataFileName string = DomNodeHelper_GetTextNodeValue(dataFileNameNode); _ = dataFileName
            this.TransformInfoInterface = any(NewTransformInfo(name, objectFileName, objectConfigFileName, templateFileName, dataFileName)).(TransformInfoInterface)
        } else {
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                stringBuffer.Append("Constructing Partial TransformInfo Node For: ")
                stringBuffer.Append(name)
                stringBuffer.Append(" with ")
                stringBuffer.Appendint(int(nodeList.GetLength()))
                stringBuffer.Append(" children")
                this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "TransformInfoObjectConfigComponent(Node node)")
            }
            this.TransformInfoInterface = any(NewTransformInfo(name)).(TransformInfoInterface)
        }
        var templateAttrNode *Attr = any(attributes.GetNamedItem(transformInfoData.MAPPED)).(*Attr); _ = templateAttrNode
        if templateAttrNode != nil {
            this.MappedName = templateAttrNode.GetValue()
        } else {
            this.MappedName = this.GetTransformInfoInterface().GetName()
        }
    return this
}


// Java overloaded constructor: Go needs a different name
func NewTransformInfoDomNode(transformInfoInterface TransformInfoInterface) *TransformInfoDomNode {
    this := &TransformInfoDomNode{}
    this.initFields()
        this.TransformInfoInterface = transformInfoInterface
        this.MappedName = transformInfoInterface.GetName()
    return this
}


// Java overloaded constructor: Go needs a different name
func NewTransformInfoDomNode(transformInfoInterface TransformInfoInterface, mappedName string) *TransformInfoDomNode {
    this := &TransformInfoDomNode{}
    this.initFields()
        this.TransformInfoInterface = transformInfoInterface
        this.MappedName = mappedName
    return this
}


// Java overloaded constructor: Go needs a different name
func NewTransformInfoDomNode(name string, mappedName string) *TransformInfoDomNode {
    this := &TransformInfoDomNode{}
    this.initFields()
        this.TransformInfoInterface = NewTransformInfo(name)
        this.MappedName = mappedName
    return this
}


func (this *TransformInfoDomNode) GetTransformInfoInterface() TransformInfoInterface {
        return this.TransformInfoInterface
}


func (this *TransformInfoDomNode) GetMappedName() string {
        return this.MappedName
}


func (this *TransformInfoDomNode) GetReplaceKey() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append("<transform:component name=\"")
        stringBuffer.Append("body")
        stringBuffer.Append("\"/>")
        var key string = fmt.Sprint(stringBuffer); _ = key
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            this.LogUtil.PutF("Component Key: " + key, this, "getReplaceKey()")
        }
        return key
}

