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

type HttpComponentView struct {
    TransformInfoHttpComposite
    LogUtil *LogUtil
    AbeClientInformation AbeClientInformationInterface
    DomNodeInterfaceVector *BasicArrayList
    TransformDocumentInterface TransformDocumentInterface
    NO_TYPE int
}

// Implements TransformInterface


func (this *HttpComponentView) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.AbeClientInformation = ServiceClientInformationInterfaceFactory_GetInstance()
    
    
    this.NO_TYPE = 0
}


func NewHttpComponentView(transformInfoInterface TransformInfoInterface) *HttpComponentView {
    this := &HttpComponentView{}
        this.TransformInfoHttpComposite = *NewTransformInfoHttpComposite(transformInfoInterface)
        this.initFields()
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            this.LogUtil.PutF("View Name: " + transformInfoInterface.GetName(), this, this.CommonStrings.CONSTRUCTOR)
        }
        this.DomNodeInterfaceVector = NewBasicArrayListD()
        this.SetTransformDocumentInterface(TransformHttpRequestDocumentFactory_GetInstance(this.GetPageContext(), this.GetWeblisketSession()))
    return this
}


func (this *HttpComponentView) GetTypeId() int {
        return this.NO_TYPE
}


func (this *HttpComponentView) GetTransformDocumentInterface() TransformDocumentInterface {
        return this.TransformDocumentInterface
}


func (this *HttpComponentView) SetTransformDocumentInterface(transformDocumentInterface TransformDocumentInterface) {
        this.TransformDocumentInterface = transformDocumentInterface
}


func (this *HttpComponentView) AddDomNodeInterface(domNodeInterface DomNodeInterface) {
        this.DomNodeInterfaceVector.Add(domNodeInterface)
}


// throws Exception
func (this *HttpComponentView) ToXmlDoc() *Document {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 *Document) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEWERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "toXmlDoc()", e)
            }
            panic(e)
                }
            }()
            var size int = this.DomNodeInterfaceVector.Size(); _ = size
            for index := int(0); index < size; index++ {
                var domNodeInterface DomNodeInterface = any(this.DomNodeInterfaceVector.Get(index)).(DomNodeInterface); _ = domNodeInterface
                this.TransformDocumentInterface.GetBaseNode().AppendChild(domNodeInterface.ToXmlNode(this.TransformDocumentInterface.GetDoc()))
            }
            tryControl1, tryResult1 = 1, this.GetTransformDocumentInterface().GetDoc(); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws Exception
func (this *HttpComponentView) GetDoc() *Document {
        var document *Document = this.GetTransformInfoInterface().GetDataDocument(); _ = document
        var node *Node = DomNodeHelper_GetFirstChildElement(document); _ = node
        if node != nil {
            var dataNode *Node = this.GetTransformDocumentInterface().GetDoc().ImportNode(node, true); _ = dataNode
            if dataNode != nil {
                this.GetTransformDocumentInterface().GetBaseNode().AppendChild(dataNode)
            }
        }
        return this.GetTransformDocumentInterface().GetDoc()
}


// throws Exception
func (this *HttpComponentView) View() string {
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
            this.ToXmlDoc()
            var success string = XmlDocumentHelper_ToString(this.GetDoc()); _ = success
            var result string = NewBasicTransformer(this.AbeClientInformation, this.GetTransformInfoInterface()).Translate(success); _ = result
            tryControl1, tryResult1 = 1, result; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

