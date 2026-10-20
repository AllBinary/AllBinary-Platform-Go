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

type GenericCssStyleValidation struct {
    Validation
    LogUtil *LogUtil
}

// Implements DomNodeInterface


func (this *GenericCssStyleValidation) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}




// Java overloads of new GenericCssStyleValidation
func NewGenericCssStyleValidation(args ...any) *GenericCssStyleValidation {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 0 {
            return NewGenericCssStyleValidation_1()
        }
        if len(args) == 1 && Java_IsArg[*Document](args[0], pass) {
            return NewGenericCssStyleValidation_2(Java_Arg[*Document](args[0]))
        }
        if len(args) == 1 && Java_IsArg[*HashMap](args[0], pass) {
            return NewGenericCssStyleValidation_3(Java_Arg[*HashMap](args[0]))
        }
    }
    panic("No Java overload of new GenericCssStyleValidation for the arguments")
}
func NewGenericCssStyleValidation_1() *GenericCssStyleValidation {
    this := &GenericCssStyleValidation{}
    this.Validation = *NewValidation()
    this.initFields()
    return this
}


func NewGenericCssStyleValidation_2(document *Document) *GenericCssStyleValidation {
    this := &GenericCssStyleValidation{}
    this.Validation = *NewValidation()
    this.initFields()
        var nodeList *NodeList = document.GetElementsByTagName(StylesData_GetInstance().NAME); _ = nodeList
    return this
}


func NewGenericCssStyleValidation_3(hashMap *HashMap) *GenericCssStyleValidation {
    this := &GenericCssStyleValidation{}
    this.Validation = *NewValidation()
    this.initFields()
        this.GetFormData(hashMap)
    return this
}


// throws Exception
func (this *GenericCssStyleValidation) GetFormData(hashMap *HashMap) {
}


func (this *GenericCssStyleValidation) IsValid() bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put("Failed to validate form", this, this.CommonStrings.IS_VALID, e)
            }
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            var isValid bool = true; _ = isValid
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.IS_VALID)
            }
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                this.LogUtil.PutF("End: " + fmt.Sprint(isValid), this, this.CommonStrings.IS_VALID)
            }
            tryControl1, tryResult1 = 1, isValid; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *GenericCssStyleValidation) ValidationInfo() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEWERROR) {
                this.LogUtil.Put("Failed to generate validation error info", this, "validationInfo()", e)
            }
            tryControl1, tryResult1 = 1, "Error Validating Form"; return
                }
            }()
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            tryControl1, tryResult1 = 1, fmt.Sprint(stringBuffer); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *GenericCssStyleValidation) ToValidationInfoDoc() *Document {
        return nil
}


func (this *GenericCssStyleValidation) ToValidationInfoNode(document *Document) *Node {
        return nil
}


// throws Exception
func (this *GenericCssStyleValidation) ToXmlNode(document *Document) *Node {
        var node *Node = ModDomHelper_CreateNameValueNodes(document, BodyData_GetInstance().NAME, StringUtil_GetInstance().EMPTY_STRING); _ = node
        return node
}

