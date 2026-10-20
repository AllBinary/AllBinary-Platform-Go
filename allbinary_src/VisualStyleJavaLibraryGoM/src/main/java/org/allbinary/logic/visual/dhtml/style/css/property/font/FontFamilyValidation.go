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

type FontFamilyValidation struct {
    FontFamilyView
    LogUtil *LogUtil
}

// Implements ValidationInterface


func (this *FontFamilyValidation) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func NewFontFamilyValidation() *FontFamilyValidation {
    this := &FontFamilyValidation{}
        this.FontFamilyView = *NewFontFamilyView()
        this.initFields()
    return this
}


// Java overloaded constructor: Go needs a different name
func NewFontFamilyValidation(node *Node) *FontFamilyValidation {
    this := &FontFamilyValidation{}
        this.FontFamilyView = *NewFontFamilyView(node)
        this.initFields()
    return this
}


// throws Exception
func (this *FontFamilyValidation) IsValid() bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEWERROR) {
                this.LogUtil.Put("Failed to validate data", this, this.CommonStrings.IS_VALID, e)
            }
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            var valid bool = true; _ = valid
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                this.LogUtil.PutF("FontFamilyValidation", this, this.CommonStrings.IS_VALID)
            }
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                this.LogUtil.PutF("CssValidation: " + fmt.Sprint(valid), this, this.CommonStrings.IS_VALID)
            }
            tryControl1, tryResult1 = 1, valid; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws Exception
func (this *FontFamilyValidation) ValidationInfo() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEWERROR) {
                this.LogUtil.Put("Failed to generate validation info error", this, "validationInfo()", e)
            }
            tryControl1, tryResult1 = 1, "Error Validating Data"; return
                }
            }()
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append("Css Font Family Property Validation Error")
            tryControl1, tryResult1 = 1, fmt.Sprint(stringBuffer); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws Exception
func (this *FontFamilyValidation) ToValidationInfoDoc() *Document {
        return nil
}


// throws Exception
func (this *FontFamilyValidation) ToValidationInfoNode(document *Document) *Node {
        return nil
}

