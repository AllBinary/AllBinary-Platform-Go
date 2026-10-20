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

type EditThemeValidationView struct {
    ThemeCustomizerView
    LogUtil *LogUtil
}

// Implements ValidationComponentInterface


func (this *EditThemeValidationView) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func NewEditThemeValidationView(transformInfoInterface TransformInfoInterface) *EditThemeValidationView {
    this := &EditThemeValidationView{}
        this.ThemeCustomizerView = *NewThemeCustomizerView(transformInfoInterface)
        this.initFields()
    return this
}


func (this *EditThemeValidationView) IsValid() bool {
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
            var categoryFactoryInterface CategoryFactoryInterface = NewStoreThemeCategoryFactory(this.GetTransformInfoInterface()); _ = categoryFactoryInterface
            var categoryLoaderInterface CategoryLoaderInterface = CategoryLoaderFactory_GetInstance(categoryFactoryInterface); _ = categoryLoaderInterface
            var rootStoreThemeCategoryInterface StoreThemeCategoryInterface = any(categoryFactoryInterface.GetRootInstance()).(StoreThemeCategoryInterface); _ = rootStoreThemeCategoryInterface
            var storeThemeCategoryInterface StoreThemeCategoryInterface = any(categoryLoaderInterface.GetAll(any(rootStoreThemeCategoryInterface).(CategoryInterface))).(StoreThemeCategoryInterface); _ = storeThemeCategoryInterface
            this.ValidationInterface = NewStoreThemeCategoryPathValidationView(any(storeThemeCategoryInterface).(StoreThemeCategoryInterface))
            isValid = this.ValidationInterface.IsValid()
            var domNodeInterface DomNodeInterface = any(this.ValidationInterface).(DomNodeInterface); _ = domNodeInterface
            var document *Document = DomDocumentHelper_Create(); _ = document
            document.AppendChild(domNodeInterface.ToXmlNode(document))
            var documentString string = XmlDocumentHelper_ToString(document); _ = documentString
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEWERROR) {
                this.LogUtil.PutF("Setting Data: " + documentString, this, this.CommonStrings.CONSTRUCTOR)
            }
            this.GetTransformInfoInterface().SetData(documentString)
            tryControl1, tryResult1 = 1, isValid; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *EditThemeValidationView) ValidationInfo() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEWERROR) {
                this.LogUtil.Put("Failed to generate validation error info", this, "validationInfo()", e)
            }
            tryControl1, tryResult1 = 1, "Error Getting Validation Info"; return
                }
            }()
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append(this.ValidationInterface.ValidationInfo())
            tryControl1, tryResult1 = 1, fmt.Sprint(stringBuffer); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *EditThemeValidationView) ToValidationInfoDoc() *Document {
        return nil
}


func (this *EditThemeValidationView) ToValidationInfoNode(document *Document) *Node {
        return nil
}

