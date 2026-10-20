/*
 *
 * AllBinary Open License Version 1
 * Copyright (c) 2011 AllBinary
 * 
 * By agreeing to this license you and any business entity you represent are
 * legally bound to the AllBinary Open License Version 1 leIIgal agreement.
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

type TitleBodyValidation struct {
    Validation
    LogUtil *LogUtil
    Body string
    TitleValidation *TitleNotRequiredValidation
}

// Implements DomNodeInterface


func (this *TitleBodyValidation) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
}




// Java overloads of new TitleBodyValidation
func NewTitleBodyValidation(args ...any) *TitleBodyValidation {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 0 {
            return NewTitleBodyValidation_1()
        }
        if len(args) == 1 && Java_IsArg[*Document](args[0], pass) {
            return NewTitleBodyValidation_2(Java_Arg[*Document](args[0]))
        }
        if len(args) == 1 && Java_IsArg[*HashMap](args[0], pass) {
            return NewTitleBodyValidation_3(Java_Arg[*HashMap](args[0]))
        }
    }
    panic("No Java overload of new TitleBodyValidation for the arguments")
}
func NewTitleBodyValidation_1() *TitleBodyValidation {
    this := &TitleBodyValidation{}
    this.Validation = *NewValidation()
    this.initFields()
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.CONSTRUCTOR)
        }
        var stringUtil *StringUtil = StringUtil_GetInstance(); _ = stringUtil
        this.Body = stringUtil.EMPTY_STRING
        this.TitleValidation = NewTitleNotRequiredValidation()
    return this
}


func NewTitleBodyValidation_2(document *Document) *TitleBodyValidation {
    this := &TitleBodyValidation{}
    this.Validation = *NewValidation()
    this.initFields()
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            this.LogUtil.PutF(this.CommonStrings.CONSTRUCTOR, this, "Constructor(document)")
        }
        var nodeList *NodeList = document.GetElementsByTagName(BodyData_GetInstance().NAME); _ = nodeList
        for index := int(0); index < nodeList.GetLength(); index++ {
            var aBodyNode *Node = nodeList.Item(index); _ = aBodyNode
            var aBodyValueNode *Node = DomSearchHelper_GetNode(DomData_VALUE, aBodyNode.GetChildNodes()); _ = aBodyValueNode
            this.Body = DomNodeHelper_GetTextNodeValue(aBodyValueNode)
            var titleNode *Node = DomSearchHelper_GetNode(TitleData_GetInstance().NAME, aBodyNode.GetChildNodes()); _ = titleNode
            this.TitleValidation = NewTitleNotRequiredValidation(titleNode)
            if index > 1 {
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                    this.LogUtil.PutF("To Many Body Nodes", this, this.CommonStrings.CONSTRUCTOR)
                }
            }
        }
    return this
}


func NewTitleBodyValidation_3(hashMap *HashMap) *TitleBodyValidation {
    this := &TitleBodyValidation{}
    this.Validation = *NewValidation()
    this.initFields()
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            this.LogUtil.PutF(this.CommonStrings.START, this, "Constructor(hashmap)")
        }
        this.GetFormData(hashMap)
    return this
}


// throws Exception
func (this *TitleBodyValidation) GetFormData(hashMap *HashMap) {
        this.Body = any(hashMap.Get(BodyData_GetInstance().NAME)).(string)
        this.TitleValidation = NewTitleNotRequiredValidation(hashMap)
}


func (this *TitleBodyValidation) IsValid() bool {
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
                this.LogUtil.PutF("TitleBodyValidation", this, this.CommonStrings.IS_VALID)
            }
            if !StringValidationUtil_GetInstance().IsValidNotRequired(this.Body, BodyData_GetInstance().MIN, AbSqlData_MAXBLOB) {
                isValid = false
            }
            if !this.TitleValidation.IsValid() {
                isValid = false
            }
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                this.LogUtil.PutF("TitleBodyValidation: " + fmt.Sprint(isValid), this, this.CommonStrings.IS_VALID)
            }
            tryControl1, tryResult1 = 1, isValid; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *TitleBodyValidation) ValidationInfo() string {
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
            {
                stringBuffer.Append("Error: Data submitted is < ")
                stringBuffer.Appendint(BodyData_GetInstance().MIN)
                stringBuffer.Append(" or > ")
                stringBuffer.Appendint(AbSqlData_MAXBLOB)
            }
            tryControl1, tryResult1 = 1, fmt.Sprint(stringBuffer); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *TitleBodyValidation) ToValidationInfoDoc() *Document {
        return nil
}


func (this *TitleBodyValidation) ToValidationInfoNode(document *Document) *Node {
        return nil
}


// throws Exception
func (this *TitleBodyValidation) ToXmlNode(document *Document) *Node {
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            this.LogUtil.PutF("Title Body: " + this.Body, this, "toXmlNode(document)")
        }
        var node *Node = ModDomHelper_CreateNameValueNodes(document, BodyData_GetInstance().NAME, this.Body); _ = node
        node.AppendChild(this.TitleValidation.ToXmlNode(document))
        return node
}

