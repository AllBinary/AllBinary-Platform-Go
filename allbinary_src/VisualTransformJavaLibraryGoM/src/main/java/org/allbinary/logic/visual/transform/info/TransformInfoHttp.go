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

type TransformInfoHttp struct {
    TransformInfo
    LogUtil *LogUtil
    AbContext *AbContext
}

// Implements TransformInfoHttpInterface


func (this *TransformInfoHttp) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
}


func NewTransformInfoHttp(propertiesHashMap *HashMap, pageContext *PageContext, crud bool) *TransformInfoHttp {
    this := &TransformInfoHttp{}
        this.TransformInfo = *NewTransformInfo()
        this.initFields()
        this.AbContext = NewAbContext(propertiesHashMap, pageContext)
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            this.LogUtil.PutF("Properties HashMap: " + fmt.Sprint(propertiesHashMap), this, "Constructor(HashMap, PageContext, boolean)")
        }
        this.Override(propertiesHashMap)
    return this
}


// Java overloaded constructor: Go needs a different name
func NewTransformInfoHttp(request *HttpServletRequest, propertiesHashMap *HashMap, pageContext *PageContext) *TransformInfoHttp {
    this := &TransformInfoHttp{}
        this.TransformInfo = *NewTransformInfo()
        this.initFields()
        this.AbContext = NewAbContext(propertiesHashMap, pageContext)
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            this.LogUtil.PutF("Properties HashMap: " + fmt.Sprint(propertiesHashMap), this, "Constructor(HttpServletRequest, HashMap, PageContext)")
        }
        this.Set(NewRequestParams(any(request).(*HttpServletRequest)).ToHashMap())
        this.Override(propertiesHashMap)
    return this
}


// Java overloaded constructor: Go needs a different name
func NewTransformInfoHttp(databaseHashMap *HashMap, propertiesHashMap *HashMap, pageContext *PageContext) *TransformInfoHttp {
    this := &TransformInfoHttp{}
        this.TransformInfo = *NewTransformInfo()
        this.initFields()
        this.AbContext = NewAbContext(propertiesHashMap, pageContext)
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            this.LogUtil.PutF("Properties HashMap: " + fmt.Sprint(propertiesHashMap), this, "Constructor(HashMap, HashMap, PageContext)")
        }
        this.Set(databaseHashMap)
    return this
}


// Java overloaded constructor: Go needs a different name
func NewTransformInfoHttp(propertiesHashMap *HashMap, pageContext *PageContext) *TransformInfoHttp {
    this := &TransformInfoHttp{}
        this.TransformInfo = *NewTransformInfo()
        this.initFields()
        this.AbContext = NewAbContext(propertiesHashMap, pageContext)
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            this.LogUtil.PutF("Properties HashMap: " + fmt.Sprint(propertiesHashMap), this, "Constructor(HashMap, PageContext)")
        }
        this.Set(propertiesHashMap)
    return this
}


// Java overloaded constructor: Go needs a different name
func NewTransformInfoHttp(parentViewOfThisTransformInfoInterface *TransformInfoHttp) *TransformInfoHttp {
    this := &TransformInfoHttp{}
        this.TransformInfo = *NewTransformInfo()
        this.initFields()
        this.AbContext = NewAbContext(parentViewOfThisTransformInfoInterface.GetPropertiesHashMap(), parentViewOfThisTransformInfoInterface.GetPageContext())
    return this
}


// throws Exception
func (this *TransformInfoHttp) Set(hashMap *HashMap) {
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            this.LogUtil.PutF("Properties HashMap: " + fmt.Sprint(hashMap), this, "set()")
        }
        var transformInfoData *TransformInfoData = TransformInfoData_GetInstance(); _ = transformInfoData
        this.SetName(any(hashMap.Get(transformInfoData.NAME)).(string))
        this.SetStoreName(any(hashMap.Get(StoreFrontData_GetInstance().NAME)).(string))
        this.SetObjectFile(any(hashMap.Get(transformInfoData.OBJECTFILENAME)).(string))
        this.SetObject(any(hashMap.Get(transformInfoData.OBJECT)).(string))
        var stringValidationUtil *StringValidationUtil = StringValidationUtil_GetInstance(); _ = stringValidationUtil
        var objectConfigString string = any(hashMap.Get(transformInfoData.OBJECTCONFIG)).(string); _ = objectConfigString
        if stringValidationUtil.IsValidRequired(objectConfigString, 10, AbSqlData_MAXBLOB) {
            this.SetObjectConfig(objectConfigString)
        } else {
            this.SetObjectConfig(fmt.Sprint(NewTransformInfoObjectConfig(this)))
        }
        if !stringValidationUtil.IsEmpty(any(hashMap.Get(transformInfoData.OBJECTCONFIGFILENAME)).(string)) {
            this.SetObjectConfigFile(any(hashMap.Get(transformInfoData.OBJECTCONFIGFILENAME)).(string))
        }
        this.SetTemplateFile(any(hashMap.Get(transformInfoData.TEMPLATEFILENAME)).(string))
        this.SetTemplate(any(hashMap.Get(transformInfoData.TEMPLATE)).(string))
        this.SetDataFile(any(hashMap.Get(transformInfoData.DATAFILENAME)).(string))
        this.SetData(any(hashMap.Get(transformInfoData.DATA)).(string))
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPERFACTORY) {
            this.LogUtil.PutF(this.Log(), this, "set()")
        }
}


func (this *TransformInfoHttp) GetWeblisketSession() WeblisketSessionInterface {
        return this.AbContext.GetWeblisketSession()
}


func (this *TransformInfoHttp) GetPageContext() *PageContext {
        return this.AbContext.GetPageContext()
}


func (this *TransformInfoHttp) GetPropertiesHashMap() *HashMap {
        return this.AbContext.GetPropertiesHashMap()
}


// throws Exception
func (this *TransformInfoHttp) GetTemplateFilePath() *AbPath {
        return NewAbPath(TransformInfo_URLGLOBALS.GetMainPath() + TransformInfo_FREEBLISKET_PATH_GLOBALS.GetInstance().XSLPATH, this.GetTemplateFile())
}


// throws Exception
func (this *TransformInfoHttp) GetObjectConfigFilePath() *AbPath {
        return NewAbPath(TransformInfo_URLGLOBALS.GetMainPath() + TransformInfo_FREEBLISKET_PATH_GLOBALS.GetInstance().XSLPATH, this.GetObjectConfigFile())
}


// throws Exception
func (this *TransformInfoHttp) GetDataFilePath() *AbPath {
        return NewAbPath(TransformInfo_URLGLOBALS.GetMainPath() + TransformInfo_FREEBLISKET_PATH_GLOBALS.GetInstance().XSLPATH, this.GetDataFile())
}

