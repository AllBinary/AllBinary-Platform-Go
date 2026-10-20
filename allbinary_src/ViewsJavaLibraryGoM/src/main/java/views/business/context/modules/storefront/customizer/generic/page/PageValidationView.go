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

type PageValidationView struct {
    PageView
    LogUtil *LogUtil
}

// Implements ValidationComponentInterface


func (this *PageValidationView) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func NewPageValidationView(transformInfoInterface TransformInfoInterface) *PageValidationView {
    this := &PageValidationView{}
        this.PageView = *NewPageView(transformInfoInterface)
        this.initFields()
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            this.LogUtil.PutF("View Name: " + transformInfoInterface.GetName(), this, "PageViewValidation()")
        }
    return this
}


func (this *PageValidationView) IsValid() bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEWERROR) {
                this.LogUtil.Put("Failed to validate", this, this.CommonStrings.IS_VALID, e)
            }
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            var isValid bool = true; _ = isValid
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.IS_VALID)
            }
            var transformInfoEntityInterface *TransformInfoEntity = TransformInfoEntityBuilder_GetInstance(); _ = transformInfoEntityInterface
            var pageObjectConfigInterface TransformInfoObjectConfigInterface = this.GetTransformInfoInterface().GetObjectConfigInterface(); _ = pageObjectConfigInterface
            var allViewsToBeModifiedVector *BasicArrayList = pageObjectConfigInterface.GetGroupTransforms(); _ = allViewsToBeModifiedVector
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                this.LogUtil.PutF("Views To Be Modified: " + fmt.Sprint(allViewsToBeModifiedVector.Size()), this, "get(transformInfoInterface)")
            }
            var size int = allViewsToBeModifiedVector.Size(); _ = size
            for index := int(0); index < size; index++ {
                var transformInfoInterface *TransformInfo = any(allViewsToBeModifiedVector.Get(index)).(*TransformInfo); _ = transformInfoInterface
                var viewNameOfViewToBeModified string = transformInfoInterface.GetName(); _ = viewNameOfViewToBeModified
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                    var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                    stringBuffer.Append(this.GetTransformInfoInterface().GetName())
                    stringBuffer.Append(" is modifying view: ")
                    stringBuffer.Append(viewNameOfViewToBeModified)
                    this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "insert()")
                }
                var httpTransformInfoInterface TransformInfoHttpInterface = any(this.GetTransformInfoInterface()).(TransformInfoHttpInterface); _ = httpTransformInfoInterface
                var specifiedTransformInfoInterface TransformInfoInterface = transformInfoEntityInterface.Get(viewNameOfViewToBeModified, httpTransformInfoInterface.GetPropertiesHashMap(), httpTransformInfoInterface.GetPageContext()); _ = specifiedTransformInfoInterface
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                    var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                    stringBuffer.Append(this.GetTransformInfoInterface().GetName())
                    stringBuffer.Append(" is adding data to view: ")
                    stringBuffer.Append(viewNameOfViewToBeModified)
                    this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "insert()")
                }
                var hashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = hashMap
                var title string = specifiedTransformInfoInterface.GetName()[this.GetTransformInfoInterface().GetStoreName().Length():]; _ = title
                if strings.Compare(title, "index") == 0 {
                    hashMap.Put(TitleData_GetInstance().TEXT, fmt.Sprint(this.GetTransformInfoInterface().GetStoreName()) + " - Home Page")
                } else {
                    hashMap.Put(TitleData_GetInstance().TEXT, fmt.Sprint(this.GetTransformInfoInterface().GetStoreName()) + " -" + title)
                }
                var pageValidation *PageValidation = NewPageValidation(hashMap); _ = pageValidation
                if pageValidation.IsValid() == false {
                    isValid = false
                }
                if isValid == true {
                    var document *Document = DomDocumentHelper_Create(); _ = document
                    document.AppendChild(pageValidation.ToXmlNode(document))
                    var documentString string = XmlDocumentHelper_ToString(document); _ = documentString
                    if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                        stringBuffer.Append(viewNameOfViewToBeModified)
                        stringBuffer.Append(" is changing data in ")
                        stringBuffer.Append(fmt.Sprint(specifiedTransformInfoInterface.GetDataFilePath()))
                        stringBuffer.Append(" to the following data:\n")
                        stringBuffer.Append(documentString)
                        this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, this.CommonStrings.IS_VALID)
                    }
                    CustomizerUtil_GetInstance().Write(specifiedTransformInfoInterface, documentString)
                }
            }
            tryControl1, tryResult1 = 1, isValid; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *PageValidationView) ValidationInfo() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEWERROR) {
                this.LogUtil.Put("Failed to generate validation error info", this, "validationInfo()", e)
            }
            tryControl1, tryResult1 = 1, "An auto generated page name was invalid."; return
                }
            }()
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            var transformInfoEntityInterface *TransformInfoEntity = TransformInfoEntityBuilder_GetInstance(); _ = transformInfoEntityInterface
            var objectConfig TransformInfoObjectConfigInterface = this.GetTransformInfoInterface().GetObjectConfigInterface(); _ = objectConfig
            var httpTransformInfoInterface TransformInfoHttpInterface = any(this.GetTransformInfoInterface()).(TransformInfoHttpInterface); _ = httpTransformInfoInterface
            var allViewsToBeModifiedVector *BasicArrayList = objectConfig.GetGroupTransforms(); _ = allViewsToBeModifiedVector
            var size int = allViewsToBeModifiedVector.Size(); _ = size
            for index := int(0); index < size; index++ {
                var transformInfoInterface *TransformInfo = any(allViewsToBeModifiedVector.Get(index)).(*TransformInfo); _ = transformInfoInterface
                var viewNameOfViewToBeModified string = transformInfoInterface.GetName(); _ = viewNameOfViewToBeModified
                var specifiedTransformInfoInterface TransformInfoInterface = transformInfoEntityInterface.Get(viewNameOfViewToBeModified, httpTransformInfoInterface.GetPropertiesHashMap(), httpTransformInfoInterface.GetPageContext()); _ = specifiedTransformInfoInterface
                var hashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = hashMap
                var title string = specifiedTransformInfoInterface.GetName()[this.GetTransformInfoInterface().GetStoreName().Length():]; _ = title
                if strings.Compare(title, "index") == 0 {
                    hashMap.Put(TitleData_GetInstance().TEXT, fmt.Sprint(this.GetTransformInfoInterface().GetStoreName()) + " - Home Page")
                } else {
                    hashMap.Put(TitleData_GetInstance().TEXT, fmt.Sprint(this.GetTransformInfoInterface().GetStoreName()) + " -" + title)
                }
                var pageValidation *PageValidation = NewPageValidation(hashMap); _ = pageValidation
                if pageValidation.IsValid() == false {
                    stringBuffer.Append("TransformInfo Name for PageValidation:" + specifiedTransformInfoInterface.GetName())
                    stringBuffer.Append("PageValidation:" + fmt.Sprint(hashMap))
                    stringBuffer.Append("PageValidation Info:" + fmt.Sprint(pageValidation.ValidationInfo()))
                }
            }
            tryControl1, tryResult1 = 1, fmt.Sprint(stringBuffer); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *PageValidationView) ToValidationInfoDoc() *Document {
        return nil
}


func (this *PageValidationView) ToValidationInfoNode(document *Document) *Node {
        return nil
}

