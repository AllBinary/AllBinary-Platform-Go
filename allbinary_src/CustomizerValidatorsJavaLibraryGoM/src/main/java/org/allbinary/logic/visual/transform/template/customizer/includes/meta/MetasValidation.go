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
 * II
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

type MetasValidation struct {
    Validation
    LogUtil *LogUtil
    StoreFrontInterface StoreFrontInterface
    MetaValidationVector IBasicArrayList
}

// Implements DomNodeInterface


func (this *MetasValidation) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
}




// Java overloads of new MetasValidation
func NewMetasValidation(args ...any) *MetasValidation {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[string](args[0], pass) {
            return NewMetasValidation_1(Java_Arg[string](args[0]))
        }
        if len(args) == 1 && Java_IsArg[*Document](args[0], pass) {
            return NewMetasValidation_2(Java_Arg[*Document](args[0]))
        }
    }
    panic("No Java overload of new MetasValidation for the arguments")
}
func NewMetasValidation_1(storeName string) *MetasValidation {
    this := &MetasValidation{}
    this.Validation = *NewValidation()
    this.initFields()
        this.StoreFrontInterface = StoreFrontFactory_GetInstance(storeName)
        this.MetaValidationVector = NewBasicArrayListD()
        var contentValue string = this.StoreFrontInterface.GetName() + " E-Commerce Site"; _ = contentValue
        var keywordsMetaValidation *MetaValidation = NewMetaValidation(HtmlMetaAttributeDataFactory_GetInstance().NAME, "Keywords", HtmlMetaAttributeValuesData_GetInstance().KEYWORDS, contentValue); _ = keywordsMetaValidation
        var abstractMetaValidation *MetaValidation = NewMetaValidation(HtmlMetaAttributeDataFactory_GetInstance().NAME, "Abstract", HtmlMetaAttributeValuesData_GetInstance().ABSTRACT, contentValue); _ = abstractMetaValidation
        var descriptionMetaValidation *MetaValidation = NewMetaValidation(HtmlMetaAttributeDataFactory_GetInstance().NAME, "Description", HtmlMetaAttributeValuesData_GetInstance().DESCRIPTION, contentValue); _ = descriptionMetaValidation
        this.MetaValidationVector.Add(keywordsMetaValidation)
        this.MetaValidationVector.Add(abstractMetaValidation)
        this.MetaValidationVector.Add(descriptionMetaValidation)
    return this
}


func NewMetasValidation_2(document *Document) *MetasValidation {
    this := &MetasValidation{}
    this.Validation = *NewValidation()
    this.initFields()
        panic(errors.New(fmt.Sprint(CommonStrings_GetInstance().NOT_IMPLEMENTED)))
    return this
}


func (this *MetasValidation) IsValid() bool {
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


func (this *MetasValidation) ValidationInfo() string {
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


func (this *MetasValidation) ToValidationInfoDoc() *Document {
        return nil
}


func (this *MetasValidation) ToValidationInfoNode(document *Document) *Node {
        return nil
}


// throws Exception
func (this *MetasValidation) ToXmlNode(document *Document) *Node {
        var node *Node = document.CreateElement(HtmlMetasData_GetInstance().NAME); _ = node
        var metaArray []*MetaValidation = any(this.MetaValidationVector.ToArrayType(make([]*MetaValidation, this.MetaValidationVector.Size()))).([]*MetaValidation); _ = metaArray
        var size int = len(metaArray); _ = size
        for i := int(0); i < size; i++ {
            var metaValidation *MetaValidation = metaArray[i]; _ = metaValidation
            node.AppendChild(metaValidation.ToXmlNode(document))
        }
        return node
}

