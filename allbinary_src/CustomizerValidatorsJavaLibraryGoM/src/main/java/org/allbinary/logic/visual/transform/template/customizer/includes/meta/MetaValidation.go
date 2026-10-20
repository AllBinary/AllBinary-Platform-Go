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
 * III
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

type MetaValidation struct {
    Validation
    LogUtil *LogUtil
    Label string
    Name string
    HttpEquiv string
    Content string
}

// Implements DomNodeInterface


func (this *MetaValidation) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
    
    
}




// Java overloads of new MetaValidation
func NewMetaValidation(args ...any) *MetaValidation {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 4 && Java_IsArg[*HtmlMetaAttributeData](args[0], pass) && Java_IsArg[string](args[1], pass) && Java_IsArg[string](args[2], pass) && Java_IsArg[string](args[3], pass) {
            return NewMetaValidation_1(Java_Arg[*HtmlMetaAttributeData](args[0]), Java_Arg[string](args[1]), Java_Arg[string](args[2]), Java_Arg[string](args[3]))
        }
        if len(args) == 1 && Java_IsArg[*Document](args[0], pass) {
            return NewMetaValidation_2(Java_Arg[*Document](args[0]))
        }
    }
    panic("No Java overload of new MetaValidation for the arguments")
}
func NewMetaValidation_1(htmlMetaAttributeData *HtmlMetaAttributeData, label string, attributeValue string, contentValue string) *MetaValidation {
    this := &MetaValidation{}
    this.Validation = *NewValidation()
    this.initFields()
        this.Label = label
        if Java_Equals(htmlMetaAttributeData, HtmlMetaAttributeDataFactory_GetInstance().HTTP_EQUIV) {
            this.HttpEquiv = attributeValue
        } else if Java_Equals(htmlMetaAttributeData, HtmlMetaAttributeDataFactory_GetInstance().NAME) {
            this.Name = attributeValue
        }
        this.Content = contentValue
    return this
}


func NewMetaValidation_2(document *Document) *MetaValidation {
    this := &MetaValidation{}
    this.Validation = *NewValidation()
    this.initFields()
        panic(errors.New(fmt.Sprint(CommonStrings_GetInstance().NOT_IMPLEMENTED)))
    return this
}


func (this *MetaValidation) IsValid() bool {
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


func (this *MetaValidation) ValidationInfo() string {
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


func (this *MetaValidation) ToValidationInfoDoc() *Document {
        return nil
}


func (this *MetaValidation) ToValidationInfoNode(document *Document) *Node {
        return nil
}


func (this *MetaValidation) ToHashMap() *HashMap {
        var hashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = hashMap
        var htmlMetaAttributeDataFactory *HtmlMetaAttributeDataFactory = HtmlMetaAttributeDataFactory_GetInstance(); _ = htmlMetaAttributeDataFactory
        hashMap.Put(HtmlMetaData_GetInstance().LABEL, this.Label)
        hashMap.Put(fmt.Sprint(htmlMetaAttributeDataFactory.NAME), this.Name)
        hashMap.Put(fmt.Sprint(htmlMetaAttributeDataFactory.HTTP_EQUIV), this.HttpEquiv)
        hashMap.Put(fmt.Sprint(htmlMetaAttributeDataFactory.CONTENT), this.Content)
        return hashMap
}


// throws Exception
func (this *MetaValidation) ToXmlNode(document *Document) *Node {
        var node *Node = ModDomHelper_CreateNameValueNodes(document, HtmlMetaData_GetInstance().NAME, this.ToHashMap()); _ = node
        return node
}

