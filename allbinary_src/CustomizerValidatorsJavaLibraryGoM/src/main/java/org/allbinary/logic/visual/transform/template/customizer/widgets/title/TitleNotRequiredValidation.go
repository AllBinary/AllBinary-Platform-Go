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

type TitleNotRequiredValidation struct {
    TitleView
    LogUtil *LogUtil
}

// Implements ValidationInterface


func (this *TitleNotRequiredValidation) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}




// Java overloads of new TitleNotRequiredValidation
func NewTitleNotRequiredValidation(args ...any) *TitleNotRequiredValidation {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 0 {
            return NewTitleNotRequiredValidation_1()
        }
        if len(args) == 1 && Java_IsArg[*Node](args[0], pass) {
            return NewTitleNotRequiredValidation_2(Java_Arg[*Node](args[0]))
        }
        if len(args) == 1 && Java_IsArg[*HashMap](args[0], pass) {
            return NewTitleNotRequiredValidation_3(Java_Arg[*HashMap](args[0]))
        }
    }
    panic("No Java overload of new TitleNotRequiredValidation for the arguments")
}
func NewTitleNotRequiredValidation_1() *TitleNotRequiredValidation {
    this := &TitleNotRequiredValidation{}
        this.TitleView = *NewTitleView()
        this.initFields()
    return this
}


func NewTitleNotRequiredValidation_2(node *Node) *TitleNotRequiredValidation {
    this := &TitleNotRequiredValidation{}
        this.TitleView = *NewTitleView(node)
        this.initFields()
    return this
}


func NewTitleNotRequiredValidation_3(hashMap *HashMap) *TitleNotRequiredValidation {
    this := &TitleNotRequiredValidation{}
        this.TitleView = *NewTitleView(hashMap)
        this.initFields()
    return this
}


func (this *TitleNotRequiredValidation) IsValid() bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEWERROR) {
                this.LogUtil.Put("Failed to validate form", this, this.CommonStrings.IS_VALID, e)
            }
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.IS_VALID)
            }
            var valid bool = true; _ = valid
            if !StringValidationUtil_GetInstance().IsValidNotRequired(this.GetTitle(), 0, TitleData_GetInstance().MAXLEN) {
                valid = false
            }
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                stringBuffer.Append("Title: ")
                stringBuffer.Append(this.GetTitle())
                stringBuffer.Append("End: ")
                stringBuffer.Appendboolean(valid)
                this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, this.CommonStrings.IS_VALID)
            }
            tryControl1, tryResult1 = 1, valid; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *TitleNotRequiredValidation) ValidationInfo() string {
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
            stringBuffer.Append("Title is not valid.")
            tryControl1, tryResult1 = 1, fmt.Sprint(stringBuffer); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *TitleNotRequiredValidation) ToValidationInfoDoc() *Document {
        return nil
}


func (this *TitleNotRequiredValidation) ToValidationInfoNode(document *Document) *Node {
        return nil
}

