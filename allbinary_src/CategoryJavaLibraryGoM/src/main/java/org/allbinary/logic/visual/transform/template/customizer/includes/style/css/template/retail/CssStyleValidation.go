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

type CssStyleValidation struct {
    Validation
    LogUtil *LogUtil
    CssStyleElementVector IBasicArrayList
}

// Implements DomNodeInterface

// The Java class CssStyleValidation has subclasses so its variables are this interface
type ICssStyleValidation interface {
    IValidation
    CssStyleValidation_() *CssStyleValidation
    IsValid() bool
    ValidationInfo() string
    ToValidationInfoDoc() *Document
    ToValidationInfoNode(document *Document) *Node
    ToXmlNode(document *Document) *Node
}

func (this *CssStyleValidation) CssStyleValidation_() *CssStyleValidation {
    return this
}


func (this *CssStyleValidation) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
}




// Java overloads of new CssStyleValidation
func NewCssStyleValidation(args ...any) *CssStyleValidation {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 0 {
            return NewCssStyleValidation_1()
        }
        if len(args) == 1 && Java_IsArg[*Document](args[0], pass) {
            return NewCssStyleValidation_2(Java_Arg[*Document](args[0]))
        }
        if len(args) == 1 && Java_IsArg[*HashMap](args[0], pass) {
            return NewCssStyleValidation_3(Java_Arg[*HashMap](args[0]))
        }
    }
    panic("No Java overload of new CssStyleValidation for the arguments")
}
func NewCssStyleValidation_1() *CssStyleValidation {
    this := &CssStyleValidation{}
    this.Validation = *NewValidation()
    this.initFields()
        this.CssStyleElementVector = NewBasicArrayListD()
    return this
}


func NewCssStyleValidation_2(document *Document) *CssStyleValidation {
    this := &CssStyleValidation{}
    this.Validation = *NewValidation()
    this.initFields()
        this.CssStyleElementVector = StylesValidationFactory_GetInstance().(*StylesValidationFactory).GetInstance(document)
    return this
}


func NewCssStyleValidation_3(hashMap *HashMap) *CssStyleValidation {
    this := &CssStyleValidation{}
    this.Validation = *NewValidation()
    this.initFields()
        this.CssStyleElementVector = StylesValidationFactory_GetInstance(hashMap)
    return this
}


func (this *CssStyleValidation) IsValid() bool {
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
            var size int = this.CssStyleElementVector.Size(); _ = size
            for i := int(0); i < size; i++ {
                var styleValidationInterface ValidationInterface = any(this.CssStyleElementVector.Get(i)).(ValidationInterface); _ = styleValidationInterface
                if !styleValidationInterface.IsValid() {
                    isValid = false
                }
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


func (this *CssStyleValidation) ValidationInfo() string {
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


func (this *CssStyleValidation) ToValidationInfoDoc() *Document {
        return nil
}


func (this *CssStyleValidation) ToValidationInfoNode(document *Document) *Node {
        return nil
}


// throws Exception
func (this *CssStyleValidation) ToXmlNode(document *Document) *Node {
        var node *Node = document.CreateElement(StylesData_GetInstance().NAME); _ = node
        var styleNode *Node = document.CreateElement(StyleData_GetInstance().NAME); _ = styleNode
        node.AppendChild(styleNode)
        var size int = this.CssStyleElementVector.Size(); _ = size
        for i := int(0); i < size; i++ {
            var styleDomNodeInterface DomNodeInterface = any(this.CssStyleElementVector.Get(i)).(DomNodeInterface); _ = styleDomNodeInterface
            styleNode.AppendChild(styleDomNodeInterface.ToXmlNode(document))
        }
        return node
}

