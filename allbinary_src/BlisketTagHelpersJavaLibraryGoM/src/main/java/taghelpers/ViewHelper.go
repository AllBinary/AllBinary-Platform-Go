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

type ViewHelper struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    AbeClientInformation AbeClientInformationInterface
    ComponentInterface TransformInterface
    NO_TYPE int
}

// Implements TransformInterface

// The Java class ViewHelper has subclasses so its variables are this interface
type IViewHelper interface {
    ViewHelper_() *ViewHelper
    GetTypeId() int
    View() string
    ToXmlDoc() *Document
    GetViewObject() TransformInterface
    GetTransformInfoInterface() TransformInfoInterface
    GetTransformDocumentInterface() TransformDocumentInterface
    SetTransformDocumentInterface(viewDocumentInterface TransformDocumentInterface)
}

func (this *ViewHelper) ViewHelper_() *ViewHelper {
    return this
}


func (this *ViewHelper) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.AbeClientInformation = ServiceClientInformationInterfaceFactory_GetInstance()
    
    this.NO_TYPE = 0
}


func NewViewHelper(hashMap *HashMap, pageContext *PageContext) *ViewHelper {
    this := &ViewHelper{}
    this.initFields()
        this.ComponentInterface = TransformFactory_GetInstance().GetInstance(this.AbeClientInformation, hashMap, pageContext)
    return this
}


func (this *ViewHelper) GetTypeId() int {
        return this.NO_TYPE
}


// throws Exception
func (this *ViewHelper) View() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Failed to View: "; _ = error_
            if this.ComponentInterface != nil {
                error_ += reflect.TypeOf(this.ComponentInterface).String()
            } else {
                error_ += "Unknown"
            }
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "view()", e)
            }
            panic(e)
                }
            }()
            tryControl1, tryResult1 = 1, this.ComponentInterface.View(); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws Exception
func (this *ViewHelper) ToXmlDoc() *Document {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 *Document) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Failed to View: "; _ = error_
            if this.ComponentInterface != nil {
                error_ += reflect.TypeOf(this.ComponentInterface).String()
            } else {
                error_ += "Unknown"
            }
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "view()", e)
            }
            panic(e)
                }
            }()
            panic(errors.New(fmt.Sprint("Hmm: " + fmt.Sprint(this.ComponentInterface.GetTransformDocumentInterface().GetDoc()))))
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *ViewHelper) GetViewObject() TransformInterface {
        return this.ComponentInterface
}


// throws Exception
func (this *ViewHelper) GetTransformInfoInterface() TransformInfoInterface {
        return this.ComponentInterface.GetTransformInfoInterface()
}


func (this *ViewHelper) GetTransformDocumentInterface() TransformDocumentInterface {
        return nil
}


func (this *ViewHelper) SetTransformDocumentInterface(viewDocumentInterface TransformDocumentInterface) {
        return
}

