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

type TransformInfoPropertiesFactory struct {
    LogUtil *LogUtil
}

var TransformInfoPropertiesFactory_Instance *TransformInfoPropertiesFactory = NewTransformInfoPropertiesFactory()

func (this *TransformInfoPropertiesFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func TransformInfoPropertiesFactory_GetInstance() *TransformInfoPropertiesFactory {
        return TransformInfoPropertiesFactory_Instance
}


func NewTransformInfoPropertiesFactory() *TransformInfoPropertiesFactory {
    this := &TransformInfoPropertiesFactory{}
    this.initFields()
    return this
}


// throws Exception
func (this *TransformInfoPropertiesFactory) GetInstance(node *Node) *TransformInfoProperties {
        var transformInfoChildNodeList *NodeList = node.GetChildNodes(); _ = transformInfoChildNodeList
        var transformInfoData *TransformInfoData = TransformInfoData_GetInstance(); _ = transformInfoData
        var attributes *NamedNodeMap = node.GetAttributes(); _ = attributes
        var attrNode *Attr = any(attributes.GetNamedItem(transformInfoData.NAME)).(*Attr); _ = attrNode
        var name string = attrNode.GetValue(); _ = name
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            this.LogUtil.PutF(fmt.Sprint(CommonLabels_GetInstance().NEXT) + " View Name: " + name, this, "toTransformInfoPropertiesHashMap()")
        }
        var labelNode *Node = DomSearchHelper_GetNodeNoThrow(transformInfoData.LABEL, transformInfoChildNodeList); _ = labelNode
        var label string = DomNodeHelper_GetTextNodesValue(labelNode); _ = label
        if label == "" {
            label = name
        }
        var descriptionNode *Node = DomSearchHelper_GetNode(transformInfoData.DESCRIPTION, transformInfoChildNodeList); _ = descriptionNode
        var description string = DomNodeHelper_GetTextNodesValue(descriptionNode); _ = description
        var objectFileNode *Node = DomSearchHelper_GetNode(transformInfoData.OBJECTFILENAME, transformInfoChildNodeList); _ = objectFileNode
        var objectFileName string = DomNodeHelper_GetTextNodeValue(objectFileNode); _ = objectFileName
        var objectConfigFileNode *Node = DomSearchHelper_GetNode(transformInfoData.OBJECTCONFIGFILENAME, transformInfoChildNodeList); _ = objectConfigFileNode
        var objectConfigFileName string = DomNodeHelper_GetTextNodeValue(objectConfigFileNode); _ = objectConfigFileName
        var templateFileNode *Node = DomSearchHelper_GetNode(transformInfoData.TEMPLATEFILENAME, transformInfoChildNodeList); _ = templateFileNode
        var templateFileName string = DomNodeHelper_GetTextNodeValue(templateFileNode); _ = templateFileName
        var transformInfoProperties *TransformInfoProperties = NewTransformInfoProperties(name, label, description, objectFileName, objectConfigFileName, templateFileName); _ = transformInfoProperties
        return transformInfoProperties
}

