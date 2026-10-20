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

type PageValidation struct {
    Validation
    LogUtil *LogUtil
    Title *TitleValidation
}

// Implements DomNodeInterface


func (this *PageValidation) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
}




// Java overloads of new PageValidation
func NewPageValidation(args ...any) *PageValidation {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 0 {
            return NewPageValidation_1()
        }
        if len(args) == 1 && Java_IsArg[*Document](args[0], pass) {
            return NewPageValidation_2(Java_Arg[*Document](args[0]))
        }
        if len(args) == 1 && Java_IsArg[*HashMap](args[0], pass) {
            return NewPageValidation_3(Java_Arg[*HashMap](args[0]))
        }
    }
    panic("No Java overload of new PageValidation for the arguments")
}
func NewPageValidation_1() *PageValidation {
    this := &PageValidation{}
    this.Validation = *NewValidation()
    this.initFields()
        this.Title = NewTitleValidation()
    return this
}


func NewPageValidation_2(document *Document) *PageValidation {
    this := &PageValidation{}
    this.Validation = *NewValidation()
    this.initFields()
        var node *Node = DomSearchHelper_GetNode(TitleData_GetInstance().NAME, document.GetElementsByTagName(PageData_GetInstance().NAME)); _ = node
        this.Title = NewTitleValidation(node)
    return this
}


func NewPageValidation_3(hashMap *HashMap) *PageValidation {
    this := &PageValidation{}
    this.Validation = *NewValidation()
    this.initFields()
        this.GetFormData(hashMap)
    return this
}


func (this *PageValidation) GetFormData(hashMap *HashMap) {
        this.Title = NewTitleValidation(hashMap)
}


func (this *PageValidation) IsValid() bool {
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
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                this.LogUtil.PutF("PageValidation", this, this.CommonStrings.IS_VALID)
            }
            if !this.Title.IsValid() {
                valid = false
            }
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                this.LogUtil.PutF("PageValidation: " + fmt.Sprint(valid), this, this.CommonStrings.IS_VALID)
            }
            tryControl1, tryResult1 = 1, valid; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *PageValidation) ValidationInfo() string {
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
            if !this.Title.IsValid() {
                stringBuffer.Append(this.Title.ValidationInfo())
            }
            tryControl1, tryResult1 = 1, fmt.Sprint(stringBuffer); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *PageValidation) ToValidationInfoDoc() *Document {
        return nil
}


func (this *PageValidation) ToValidationInfoNode(document *Document) *Node {
        return nil
}


// throws Exception
func (this *PageValidation) ToXmlNode(document *Document) *Node {
        var node *Node = document.CreateElement(PageData_GetInstance().NAME); _ = node
        node.AppendChild(this.Title.ToXmlNode(document))
        return node
}

