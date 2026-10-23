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

//Sets Template Type
type InsertCustomizerValidationView struct {
    HttpStoreComponentView
    LogUtil *LogUtil
    ViewName string
}

// Implements ValidationComponentInterface


func (this *InsertCustomizerValidationView) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
}


func NewInsertCustomizerValidationView(transformInfoInterface TransformInfoInterface) *InsertCustomizerValidationView {
    this := &InsertCustomizerValidationView{}
        this.HttpStoreComponentView = *NewHttpStoreComponentView(transformInfoInterface)
        this.initFields()
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            this.LogUtil.PutF("View Name: " + transformInfoInterface.GetName(), this, this.CommonStrings.CONSTRUCTOR)
        }
        var requestHashMap *HashMap = NewRequestParams(this.GetPageContext()).ToHashMap(); _ = requestHashMap
        this.ViewName = any(requestHashMap.Get(TransformInfoData_GetInstance().NAME)).(string)
    return this
}


func (this *InsertCustomizerValidationView) IsValid() bool {
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
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                this.LogUtil.PutF("Started Validation", this, this.CommonStrings.IS_VALID)
            }
            var isValid bool = true; _ = isValid
            if StringValidationUtil_GetInstance().IsEmpty(this.ViewName) {
                isValid = false
            } else {
                var objectConfig INoTemplateTransformInfoObjectConfig = NewNoTemplateTransformInfoObjectConfig(this.GetTransformInfoInterface(), this.GetTransformInfoInterface().GetObjectConfigInterface().ToXmlDoc()); _ = objectConfig
                var componentVector IBasicArrayList = objectConfig.GetGroupTransforms(); _ = componentVector
                var size int = componentVector.Size(); _ = size
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                    var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                    stringBuffer.Append("Pointing ")
                    stringBuffer.Appendint(componentVector.Size())
                    stringBuffer.Append(" Components")
                    this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, this.CommonStrings.IS_VALID)
                }
                for index := int(0); index < size; index++ {
                    var transformInfo ITransformInfo = any(componentVector.Get(index)).(ITransformInfo); _ = transformInfo
                    var transformInfoName string = transformInfo.GetName(); _ = transformInfoName
                    if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                        this.LogUtil.PutF("Setting Template Component with ViewName: " + transformInfoName, this, this.CommonStrings.IS_VALID)
                    }
                    var componentInterface TransformInterface = TransformFactory_GetInstance().GetInstance(this.AbeClientInformation, transformInfoName, this.GetTransformInfoInterface()); _ = componentInterface
                    var templateViewObjectConfig *InsertTemplateCustomizerTransformInfoObjectConfig = NewInsertTemplateCustomizerTransformInfoObjectConfig(componentInterface.GetTransformInfoInterface(), componentInterface.GetTransformInfoInterface().GetObjectConfigInterface().ToXmlDoc()); _ = templateViewObjectConfig
                    var storeName string = this.GetTransformInfoInterface().GetStoreName(); _ = storeName
                    var endIndex int = strings.Index(this.ViewName, string(storeName)) + len(storeName); _ = endIndex
                    var storePrepend string = this.ViewName[0:endIndex]; _ = storePrepend
                    var viewNamePostfix string = this.ViewName[endIndex:len(this.ViewName)]; _ = viewNamePostfix
                    var newViewNameStringBuffer *StringMaker = NewStringMaker(); _ = newViewNameStringBuffer
                    newViewNameStringBuffer.Append(storePrepend)
                    var commonSeps *CommonSeps = CommonSeps_GetInstance(); _ = commonSeps
                    var transformInfosData *TransformInfosData = TransformInfosData_GetInstance(); _ = transformInfosData
                    if strings.Index(transformInfoName, string(transformInfosData.SMALL_PREVIEW)) >= 0 {
                        newViewNameStringBuffer.Append(commonSeps.SPACE)
                        newViewNameStringBuffer.Append(transformInfosData.SMALL_PREVIEW)
                    } else if strings.Index(transformInfoName, string(transformInfosData.PREVIEW)) >= 0 {
                        newViewNameStringBuffer.Append(commonSeps.SPACE)
                        newViewNameStringBuffer.Append(transformInfosData.PREVIEW)
                    }
                    newViewNameStringBuffer.Append(commonSeps.SPACE)
                    newViewNameStringBuffer.Append(viewNamePostfix)
                    if strings.Index(transformInfoName, string(GeneratorTransformInfoData_NAME)) >= 0 {
                        newViewNameStringBuffer.Append(commonSeps.SPACE)
                        newViewNameStringBuffer.Append(GeneratorTransformInfoData_NAME)
                    }
                    var newViewName string = fmt.Sprint(newViewNameStringBuffer); _ = newViewName
                    if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                        stringBuffer.Append("Template Component: ")
                        stringBuffer.Append(transformInfoName)
                        stringBuffer.Append("\n Now Pointing To View Name: ")
                        stringBuffer.Append(newViewName)
                        this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, this.CommonStrings.IS_VALID)
                    }
                    templateViewObjectConfig.Set(newViewName)
                    if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                        stringBuffer.Append("Template Component: ")
                        stringBuffer.Append(transformInfoName)
                        stringBuffer.Append("\n with ObjectConfig: ")
                        stringBuffer.Append(fmt.Sprint(templateViewObjectConfig))
                        this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, this.CommonStrings.IS_VALID)
                    }
                    componentInterface.GetTransformInfoInterface().SetObjectConfigInterface(templateViewObjectConfig)
                    var updatedTransformInfoHashMap *HashMap = componentInterface.GetTransformInfoInterface().ToHashMap(); _ = updatedTransformInfoHashMap
                    TransformInfoEntityBuilder_GetInstance().Update(updatedTransformInfoHashMap)
                }
            }
            tryControl1, tryResult1 = 1, isValid; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *InsertCustomizerValidationView) ValidationInfo() string {
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
            stringBuffer.Append("Error: Template Name Is Empty.")
            tryControl1, tryResult1 = 1, fmt.Sprint(stringBuffer); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *InsertCustomizerValidationView) ToValidationInfoDoc() *Document {
        return nil
}


func (this *InsertCustomizerValidationView) ToValidationInfoNode(document *Document) *Node {
        return nil
}


// throws Exception
func (this *InsertCustomizerValidationView) View() string {
        return this.Views.ValidationOnlyTempUtil.GetInstance().View(this)
}

