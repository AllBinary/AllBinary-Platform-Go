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

type GenericStoreTransformInfoObjectConfig struct {
    TransformInfoObjectConfig
    LogUtil *LogUtil
    AbeClientInformation AbeClientInformationInterface
}


func (this *GenericStoreTransformInfoObjectConfig) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
}


func NewGenericStoreTransformInfoObjectConfig(abeClientInformation AbeClientInformationInterface, transformInfoInterface TransformInfoInterface) *GenericStoreTransformInfoObjectConfig {
    this := &GenericStoreTransformInfoObjectConfig{}
        this.TransformInfoObjectConfig = *NewTransformInfoObjectConfig(transformInfoInterface)
        this.initFields()
        this.AbeClientInformation = abeClientInformation
    return this
}


// Java overloaded constructor: Go needs a different name
func NewGenericStoreTransformInfoObjectConfig(abeClientInformation AbeClientInformationInterface, transformInfoInterface TransformInfoInterface, document *Document) *GenericStoreTransformInfoObjectConfig {
    this := &GenericStoreTransformInfoObjectConfig{}
        this.TransformInfoObjectConfig = *NewTransformInfoObjectConfig(transformInfoInterface, document)
        this.initFields()
        this.AbeClientInformation = abeClientInformation
        this.SetDocument(this.Generate(this.ToXmlDoc()))
    return this
}


// Java overloaded constructor: Go needs a different name
func NewGenericStoreTransformInfoObjectConfig(abeClientInformation AbeClientInformationInterface, transformInfoInterface TransformInfoInterface, name string, type_ string) *GenericStoreTransformInfoObjectConfig {
    this := &GenericStoreTransformInfoObjectConfig{}
        this.TransformInfoObjectConfig = *NewTransformInfoObjectConfig(transformInfoInterface, name, type_)
        this.initFields()
        this.AbeClientInformation = abeClientInformation
        this.SetDocument(this.Generate(this.ToXmlDoc()))
    return this
}


// throws Exception
func (this *GenericStoreTransformInfoObjectConfig) Generate(objectConfigDocument *Document) *Document {
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append("TransformInfo: ")
            if this.GetTransformInfoInterface() != nil {
                stringBuffer.Append(this.GetTransformInfoInterface().GetName())
            } else {
                stringBuffer.Append("No Owner!?#@")
            }
            this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "generate()")
        }
        var transformInfoHttpStoreInterface *TransformInfoHttp = any(this.GetTransformInfoInterface()).(*TransformInfoHttp); _ = transformInfoHttpStoreInterface
        var objectConfigDocumentString string = XmlDocumentHelper_ToString(objectConfigDocument); _ = objectConfigDocumentString
        var replaceHashMap *HashMap = this.CreateReplaceHashMap(transformInfoHttpStoreInterface, objectConfigDocumentString); _ = replaceHashMap
        return this.Generate(objectConfigDocumentString, replaceHashMap)
}


// throws Exception
func (this *GenericStoreTransformInfoObjectConfig) CreateReplaceHashMap(transformInfoHttpStoreInterface *TransformInfoHttp, objectConfigDocumentString string) *HashMap {
        var storeName string = transformInfoHttpStoreInterface.GetStoreName(); _ = storeName
        var hashMap *HashMap = this.CreateHashMap(transformInfoHttpStoreInterface, objectConfigDocumentString); _ = hashMap
        var transformInfoObjectConfigData *TransformInfoObjectConfigData = TransformInfoObjectConfigData_GetInstance(); _ = transformInfoObjectConfigData
        hashMap.Put(transformInfoObjectConfigData.VARKEY + StoreFrontData_GetInstance().NAME, storeName)
        var pageName string = TransformTemplateCustomizerUtil_GetInstance().GetPageNameHack(this.GetTransformInfoInterface().GetName(), storeName); _ = pageName
        hashMap.Put(transformInfoObjectConfigData.VARKEY + TransformInfoData_GetInstance().PARTIAL, pageName)
        return hashMap
}


// throws Exception
func (this *GenericStoreTransformInfoObjectConfig) CreateHashMap(transformInfoHttpStoreInterface *TransformInfoHttp, objectConfigDocumentString string) *HashMap {
        var hashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = hashMap
        var storeName string = transformInfoHttpStoreInterface.GetStoreName(); _ = storeName
        var propertiesHashMap *HashMap = transformInfoHttpStoreInterface.GetPropertiesHashMap(); _ = propertiesHashMap
        var templateNameOverride string = StringUtil_GetInstance().GetNonNull(any(propertiesHashMap.Get(TransformInfoData_GetInstance().PARTIAL)).(string)); _ = templateNameOverride
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            this.LogUtil.PutF("TemplateNameOverride: " + templateNameOverride, this, "generate()")
        }
        var templateNameStringBuffer *StringMaker = NewStringMaker(); _ = templateNameStringBuffer
        templateNameStringBuffer.Append(storeName)
        templateNameStringBuffer.Append(templateNameOverride)
        templateNameStringBuffer.Append(CommonSeps_GetInstance().SPACE)
        templateNameStringBuffer.Append(RootTransformInfoData_NAME)
        var templateNameKey string = TransformInfoObjectConfigData_GetInstance().VARKEY + TransformInfoData_GetInstance().OWNER; _ = templateNameKey
        if strings.Index(objectConfigDocumentString, string(templateNameKey)) != -1 {
            var templateInterface TransformTemplateInterface = TransformTemplateFactory_GetInstance().GetInstance(this.AbeClientInformation, fmt.Sprint(templateNameStringBuffer), transformInfoHttpStoreInterface.GetPropertiesHashMap(), transformInfoHttpStoreInterface.GetPageContext()); _ = templateInterface
            var selectedTemplate string = templateInterface.GetName(); _ = selectedTemplate
            hashMap.Put(templateNameKey, selectedTemplate)
        }
        return hashMap
}


// throws Exception
// Java overloaded method: Go needs a different name
func (this *GenericStoreTransformInfoObjectConfig) Generate(objectConfigDocumentString string, hashMap *HashMap) *Document {
        var replace *Replace = NewReplace(hashMap); _ = replace
        var newObjectConfigDocument *Document = DomDocumentHelper_Create(replace.All(objectConfigDocumentString)); _ = newObjectConfigDocument
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            this.LogUtil.PutF("Final ObjectConfig: " + XmlDocumentHelper_ToString(newObjectConfigDocument), this, "generate()")
        }
        return newObjectConfigDocument
}

