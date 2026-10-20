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

type BasicComponentView struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    AbeClientInformation AbeClientInformationInterface
    TransformDocumentInterface TransformDocumentInterface
    TransformInfoInterface TransformInfoInterface
    NO_TYPE int
}

// Implements TransformInterface


func (this *BasicComponentView) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.AbeClientInformation = ServiceClientInformationInterfaceFactory_GetInstance()
    
    
    this.NO_TYPE = 0
}


func NewBasicComponentView(transformInfoInterface TransformInfoInterface) *BasicComponentView {
    this := &BasicComponentView{}
    this.initFields()
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            this.LogUtil.PutF("View Name: " + transformInfoInterface.GetName(), this, this.CommonStrings.CONSTRUCTOR)
        }
        this.SetTransformInfoInterface(transformInfoInterface)
        this.SetTransformDocumentInterface(TransformDocumentFactory_GetInstance())
    return this
}


func (this *BasicComponentView) GetTypeId() int {
        return this.NO_TYPE
}


func (this *BasicComponentView) GetTransformDocumentInterface() TransformDocumentInterface {
        return this.TransformDocumentInterface
}


func (this *BasicComponentView) SetTransformDocumentInterface(transformDocumentInterface TransformDocumentInterface) {
        this.TransformDocumentInterface = transformDocumentInterface
}


func (this *BasicComponentView) GetTransformInfoInterface() TransformInfoInterface {
        return this.TransformInfoInterface
}


func (this *BasicComponentView) SetTransformInfoInterface(transformInfoInterface TransformInfoInterface) {
        this.TransformInfoInterface = transformInfoInterface
}


func (this *BasicComponentView) ToXmlDoc() *Document {
        return this.TransformDocumentInterface.GetDoc()
}


// throws Exception
func (this *BasicComponentView) GetDoc() *Document {
        var document *Document = this.GetTransformInfoInterface().GetDataDocument(); _ = document
        var dataNode *Node = this.GetTransformDocumentInterface().GetDoc().ImportNode(document.GetFirstChild(), true); _ = dataNode
        if dataNode != nil {
            this.GetTransformDocumentInterface().GetBaseNode().AppendChild(dataNode)
        }
        return this.GetTransformDocumentInterface().GetDoc()
}


// throws Exception
func (this *BasicComponentView) View() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPERERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "view()", e)
            }
            panic(e)
                }
            }()
            var success string = XmlDocumentHelper_ToString(this.GetDoc()); _ = success
            var result string = NewBasicTransformer(this.AbeClientInformation, this.GetTransformInfoInterface()).Translate(success); _ = result
            tryControl1, tryResult1 = 1, result; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

