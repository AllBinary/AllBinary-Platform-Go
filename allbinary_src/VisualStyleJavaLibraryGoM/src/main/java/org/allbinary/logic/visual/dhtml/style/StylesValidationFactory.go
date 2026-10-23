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

type StylesValidationFactory struct {
    LogUtil *LogUtil
}

var StylesValidationFactory_Instance *StylesValidationFactory = NewStylesValidationFactory()

func (this *StylesValidationFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func StylesValidationFactory_GetInstance_1() *StylesValidationFactory {
        return StylesValidationFactory_Instance
}

// Java overloads of StylesValidationFactory.getInstance
func StylesValidationFactory_GetInstance(args ...any) any {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 0 {
            return StylesValidationFactory_GetInstance_1()
        }
        if len(args) == 1 && Java_IsArg[*HashMap](args[0], pass) {
            return StylesValidationFactory_GetInstance_2(Java_Arg[*HashMap](args[0]))
        }
    }
    panic("No Java overload of StylesValidationFactory.getInstance for the arguments")
}


func NewStylesValidationFactory() *StylesValidationFactory {
    this := &StylesValidationFactory{}
    this.initFields()
    return this
}


// throws Exception
func (this *StylesValidationFactory) GetInstance(document *Document) IBasicArrayList {
        var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
        var nodeList *NodeList = document.GetElementsByTagName(StylesData_GetInstance().NAME); _ = nodeList
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            this.LogUtil.PutF("Style Present: " + XmlDocumentHelper_ToString(document), this, commonStrings.GET_INSTANCE)
        }
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            this.LogUtil.PutF("Styles Present: " + fmt.Sprint(nodeList.GetLength()), this, commonStrings.GET_INSTANCE)
        }
        for index := int(0); index < nodeList.GetLength(); index++ {
            var stylesNode *Node = nodeList.Item(index); _ = stylesNode
            var styleNodeList IBasicArrayList = DomSearchHelper_GetAllNodes(StyleData_GetInstance().NAME, stylesNode.GetChildNodes()); _ = styleNodeList
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                this.LogUtil.PutF("Number Of Style Nodes Present: " + fmt.Sprint(styleNodeList.Size()), this, commonStrings.GET_INSTANCE)
            }
            for styleNodesIndex := int(0); styleNodesIndex < styleNodeList.Size(); styleNodesIndex++ {
                var styleNode *Node = any(styleNodeList.Get(styleNodesIndex)).(*Node); _ = styleNode
                var cssElementStyleNodeList IBasicArrayList = DomSearchHelper_GetAllNodes(CssElementData_GetInstance().NAME, styleNode.GetChildNodes()); _ = cssElementStyleNodeList
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                    this.LogUtil.PutF("Number Of Element Nodes Present: " + fmt.Sprint(cssElementStyleNodeList.Size()), this, commonStrings.GET_INSTANCE)
                }
                return CssElementsValidationFactory_GetInstance(cssElementStyleNodeList)
            }
        }
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            this.LogUtil.PutF("No Style Present", this, commonStrings.GET_INSTANCE)
        }
        return NewBasicArrayListD()
}


// throws Exception
func StylesValidationFactory_GetInstance_2(hashMap *HashMap) IBasicArrayList {
        var stylesDocument *Document = any(hashMap.Get(NameSpaceRequestParamData_DOCUMENT)).(*Document); _ = stylesDocument
        return StylesValidationFactory_GetInstance().(*StylesValidationFactory).GetInstance(stylesDocument)
}

