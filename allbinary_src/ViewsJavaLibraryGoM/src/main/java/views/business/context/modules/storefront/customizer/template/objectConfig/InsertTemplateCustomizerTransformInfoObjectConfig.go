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

type InsertTemplateCustomizerTransformInfoObjectConfig struct {
    NoTemplateTransformInfoObjectConfig
    LogUtil *LogUtil
}


func (this *InsertTemplateCustomizerTransformInfoObjectConfig) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func NewInsertTemplateCustomizerTransformInfoObjectConfig(transformInfoInterface TransformInfoInterface) *InsertTemplateCustomizerTransformInfoObjectConfig {
    this := &InsertTemplateCustomizerTransformInfoObjectConfig{}
        this.NoTemplateTransformInfoObjectConfig = *NewNoTemplateTransformInfoObjectConfig(transformInfoInterface)
        this.initFields()
    return this
}


// Java overloaded constructor: Go needs a different name
func NewInsertTemplateCustomizerTransformInfoObjectConfig(transformInfoInterface TransformInfoInterface, document *Document) *InsertTemplateCustomizerTransformInfoObjectConfig {
    this := &InsertTemplateCustomizerTransformInfoObjectConfig{}
        this.NoTemplateTransformInfoObjectConfig = *NewNoTemplateTransformInfoObjectConfig(transformInfoInterface, document)
        this.initFields()
        this.SetDocument(this.Generate(this.ToXmlDoc()))
    return this
}


// Java overloaded constructor: Go needs a different name
func NewInsertTemplateCustomizerTransformInfoObjectConfig(transformInfoInterface TransformInfoInterface, name string, type_ string) *InsertTemplateCustomizerTransformInfoObjectConfig {
    this := &InsertTemplateCustomizerTransformInfoObjectConfig{}
        this.NoTemplateTransformInfoObjectConfig = *NewNoTemplateTransformInfoObjectConfig(transformInfoInterface, name, type_)
        this.initFields()
        this.SetDocument(this.Generate(this.ToXmlDoc()))
    return this
}


// throws Exception
func (this *InsertTemplateCustomizerTransformInfoObjectConfig) Set(aParentViewName string) {
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            this.LogUtil.PutF("Setting: " + aParentViewName, this, "set()")
        }
        var document *Document = this.ToXmlDoc(); _ = document
        var parentComponentsNodeList *NodeList = document.GetElementsByTagName(TransformInfosData_GetInstance().NAME); _ = parentComponentsNodeList
        if parentComponentsNodeList != nil && parentComponentsNodeList.GetLength() > 0 {
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                this.LogUtil.PutF("Replacing", this, "set()")
            }
            var componentNode *Node = parentComponentsNodeList.Item(0).GetChildNodes().Item(0); _ = componentNode
            if componentNode != nil {
                document.RemoveChild(componentNode)
            }
            parentComponentsNodeList.Item(0).AppendChild(NewTransformInfoDomNodeView(aParentViewName).ToXmlNode(document))
        } else {
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                this.LogUtil.PutF("Setting", this, "set()")
            }
            var objectConfigNode *Node = document.GetElementsByTagName(TransformInfoObjectConfigData_GetInstance().NAME).Item(0); _ = objectConfigNode
            var parentComponentsNode *Node = document.CreateElement(TransformInfosData_GetInstance().NAME); _ = parentComponentsNode
            objectConfigNode.AppendChild(parentComponentsNode)
            var transformInfoDomNodeView *TransformInfoDomNodeView = NewTransformInfoDomNodeView(aParentViewName); _ = transformInfoDomNodeView
            parentComponentsNode.AppendChild(transformInfoDomNodeView.ToXmlNode(document))
        }
}

