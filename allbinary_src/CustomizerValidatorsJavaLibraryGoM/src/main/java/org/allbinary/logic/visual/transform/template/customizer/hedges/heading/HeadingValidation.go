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

type HeadingValidation struct {
    Validation
    LogUtil *LogUtil
    Title *TitleValidation
    Logo *LogoValidation
}

// Implements DomNodeInterface


func (this *HeadingValidation) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
}




// Java overloads of new HeadingValidation
func NewHeadingValidation(args ...any) *HeadingValidation {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 0 {
            return NewHeadingValidation_1()
        }
        if len(args) == 1 && Java_IsArg[*Document](args[0], pass) {
            return NewHeadingValidation_2(Java_Arg[*Document](args[0]))
        }
        if len(args) == 1 && Java_IsArg[*HashMap](args[0], pass) {
            return NewHeadingValidation_3(Java_Arg[*HashMap](args[0]))
        }
    }
    panic("No Java overload of new HeadingValidation for the arguments")
}
func NewHeadingValidation_1() *HeadingValidation {
    this := &HeadingValidation{}
    this.Validation = *NewValidation()
    this.initFields()
        this.Title = NewTitleValidation()
        this.Logo = NewLogoValidation()
    return this
}


func NewHeadingValidation_2(document *Document) *HeadingValidation {
    this := &HeadingValidation{}
    this.Validation = *NewValidation()
    this.initFields()
        var headingNodeList *NodeList = document.GetElementsByTagName(HeadingData_NAME); _ = headingNodeList
        var titleNode *Node = DomSearchHelper_GetNode(TitleData_GetInstance().NAME, headingNodeList.Item(0).GetChildNodes()); _ = titleNode
        var logoNode *Node = DomSearchHelper_GetNode(LogoData_GetInstance().NAME, headingNodeList.Item(0).GetChildNodes()); _ = logoNode
        this.Title = NewTitleValidation(titleNode)
        this.Logo = NewLogoValidation(logoNode)
    return this
}


func NewHeadingValidation_3(hashMap *HashMap) *HeadingValidation {
    this := &HeadingValidation{}
    this.Validation = *NewValidation()
    this.initFields()
        this.GetFormData(hashMap)
    return this
}


// throws Exception
func (this *HeadingValidation) GetFormData(hashMap *HashMap) {
        this.Title = NewTitleValidation(hashMap)
        this.Logo = NewLogoValidation(hashMap)
}


func (this *HeadingValidation) IsValid() bool {
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
            var titleValid bool = true; _ = titleValid
            var logoValid bool = true; _ = logoValid
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.IS_VALID)
            }
            if !this.Title.IsValid() {
                titleValid = false
            }
            if !this.Logo.IsValid() {
                logoValid = false
            } else {
                this.Logo.ProcessLogoFile()
            }
            if !titleValid && !logoValid {
                valid = false
            }
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                this.LogUtil.PutF("HeadingValidation: " + fmt.Sprint(valid), this, this.CommonStrings.IS_VALID)
            }
            tryControl1, tryResult1 = 1, valid; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *HeadingValidation) ValidationInfo() string {
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
            var valid bool = true; _ = valid
            var titleValid bool = true; _ = titleValid
            var logoValid bool = true; _ = logoValid
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            if !this.Title.IsValid() {
                titleValid = false
            }
            if !this.Logo.IsValid() {
                logoValid = false
            }
            if !titleValid && !logoValid {
                valid = false
            }
            if !valid {
                stringBuffer.Append("Include a title and/or logo for the header.")
            }
            tryControl1, tryResult1 = 1, fmt.Sprint(stringBuffer); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *HeadingValidation) ToValidationInfoDoc() *Document {
        return nil
}


func (this *HeadingValidation) ToValidationInfoNode(document *Document) *Node {
        return nil
}


// throws Exception
func (this *HeadingValidation) ToXmlNode(document *Document) *Node {
        var node *Node = document.CreateElement(HeadingData_NAME); _ = node
        node.AppendChild(this.Logo.ToXmlNode(document))
        node.AppendChild(this.Title.ToXmlNode(document))
        return node
}

