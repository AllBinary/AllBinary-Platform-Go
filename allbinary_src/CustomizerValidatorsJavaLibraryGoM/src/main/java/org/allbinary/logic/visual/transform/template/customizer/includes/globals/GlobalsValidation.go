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

type GlobalsValidation struct {
    Validation
    LogUtil *LogUtil
    StoreFrontInterface StoreFrontInterface
}

// Implements DomNodeInterface


func (this *GlobalsValidation) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
}




// Java overloads of new GlobalsValidation
func NewGlobalsValidation(args ...any) *GlobalsValidation {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[string](args[0], pass) {
            return NewGlobalsValidation_1(Java_Arg[string](args[0]))
        }
        if len(args) == 1 && Java_IsArg[*Document](args[0], pass) {
            return NewGlobalsValidation_2(Java_Arg[*Document](args[0]))
        }
    }
    panic("No Java overload of new GlobalsValidation for the arguments")
}
func NewGlobalsValidation_1(storeName string) *GlobalsValidation {
    this := &GlobalsValidation{}
    this.Validation = *NewValidation()
    this.initFields()
        this.StoreFrontInterface = StoreFrontFactory_GetInstance(storeName)
    return this
}


func NewGlobalsValidation_2(document *Document) *GlobalsValidation {
    this := &GlobalsValidation{}
    this.Validation = *NewValidation()
    this.initFields()
        panic(errors.New(fmt.Sprint(CommonStrings_GetInstance().NOT_IMPLEMENTED)))
    return this
}


func (this *GlobalsValidation) IsValid() bool {
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
            var valid bool = true; _ = valid
            tryControl1, tryResult1 = 1, valid; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *GlobalsValidation) ValidationInfo() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
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


func (this *GlobalsValidation) ToValidationInfoDoc() *Document {
        return nil
}


func (this *GlobalsValidation) ToValidationInfoNode(document *Document) *Node {
        return nil
}


// throws Exception
func (this *GlobalsValidation) ToXmlNode(document *Document) *Node {
        var node *Node = document.CreateElement(GlobalsData_NAME); _ = node
        node.AppendChild(NewStoreFrontView(this.StoreFrontInterface).ToXmlNode(document))
        return node
}

