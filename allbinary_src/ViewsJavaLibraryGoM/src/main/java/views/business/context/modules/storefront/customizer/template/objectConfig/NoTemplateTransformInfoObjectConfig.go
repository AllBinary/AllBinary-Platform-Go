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

type NoTemplateTransformInfoObjectConfig struct {
    TransformInfoObjectConfig
    LogUtil *LogUtil
}


func (this *NoTemplateTransformInfoObjectConfig) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func NewNoTemplateTransformInfoObjectConfig(transformInfoInterface TransformInfoInterface) *NoTemplateTransformInfoObjectConfig {
    this := &NoTemplateTransformInfoObjectConfig{}
        this.TransformInfoObjectConfig = *NewTransformInfoObjectConfig(transformInfoInterface)
        this.initFields()
    return this
}


// Java overloaded constructor: Go needs a different name
func NewNoTemplateTransformInfoObjectConfig(transformInfoInterface TransformInfoInterface, document *Document) *NoTemplateTransformInfoObjectConfig {
    this := &NoTemplateTransformInfoObjectConfig{}
        this.TransformInfoObjectConfig = *NewTransformInfoObjectConfig(transformInfoInterface, document)
        this.initFields()
        this.SetDocument(this.Generate(this.ToXmlDoc()))
    return this
}


// Java overloaded constructor: Go needs a different name
func NewNoTemplateTransformInfoObjectConfig(transformInfoInterface TransformInfoInterface, name string, type_ string) *NoTemplateTransformInfoObjectConfig {
    this := &NoTemplateTransformInfoObjectConfig{}
        this.TransformInfoObjectConfig = *NewTransformInfoObjectConfig(transformInfoInterface, name, type_)
        this.initFields()
        this.SetDocument(this.Generate(this.ToXmlDoc()))
    return this
}


// throws Exception
func (this *NoTemplateTransformInfoObjectConfig) Generate(objectConfigDocument *Document) *Document {
        var docString string = XmlDocumentHelper_ToString(objectConfigDocument); _ = docString
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            this.LogUtil.PutF("Initial ObjectConfig: " + docString, this, "generate()")
        }
        var storeName string = this.GetTransformInfoInterface().GetStoreName(); _ = storeName
        var viewName string = this.GetTransformInfoInterface().GetName(); _ = viewName
        var hashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = hashMap
        var VARKEY string = TransformInfoObjectConfigData_GetInstance().VARKEY; _ = VARKEY
        hashMap.Put(VARKEY + fmt.Sprint(StoreFrontData_GetInstance().NAME), storeName)
        hashMap.Put(VARKEY + fmt.Sprint(TransformInfoData_GetInstance().OWNER), viewName)
        var pageName string = TransformTemplateCustomizerUtil_GetInstance().GetPageNameHack(this.GetTransformInfoInterface().GetName(), this.GetTransformInfoInterface().GetStoreName()); _ = pageName
        hashMap.Put(VARKEY + fmt.Sprint(TransformInfoData_GetInstance().PARTIAL), pageName)
        var replace *Replace = NewReplace(hashMap); _ = replace
        var newObjectConfigDocument *Document = DomDocumentHelper_Create(replace.All(docString)); _ = newObjectConfigDocument
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            this.LogUtil.PutF("Final ObjectConfig: " + XmlDocumentHelper_ToString(newObjectConfigDocument), this, "generate()")
        }
        return newObjectConfigDocument
}

