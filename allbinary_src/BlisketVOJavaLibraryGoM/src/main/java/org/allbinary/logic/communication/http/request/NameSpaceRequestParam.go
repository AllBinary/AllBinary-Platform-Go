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

type NameSpaceRequestParam struct {
    LogUtil *LogUtil
    NameSpaceVector IBasicArrayList
    NameSpacePropertiesHashMap *HashMap
    Value string
}


func (this *NameSpaceRequestParam) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
    
}


func NewNameSpaceRequestParam(nameSpace string, value string) *NameSpaceRequestParam {
    this := &NameSpaceRequestParam{}
    this.initFields()
        this.NameSpaceVector = NewBasicArrayListD()
        this.NameSpacePropertiesHashMap = StdUtil_GetInstance().CreateHashMap()
        this.Value = value
        var beginIndex int = strings.Index(nameSpace, string(NameSpaceRequestParamData_NAME)); _ = beginIndex
        if beginIndex < 0 {
            panic(errors.New(fmt.Sprint("Not a NameSpaceRequest")))
        }
        var packages string = nameSpace[beginIndex + len(NameSpaceRequestParamData_NAME) + len(NameSpaceRequestParamData_SEP):]; _ = packages
        var sepTokenizer *Tokenizer = NewTokenizer(NameSpaceRequestParamData_SEP); _ = sepTokenizer
        var nameSpaceWithPropertiesVector IBasicArrayList = sepTokenizer.GetTokensFromString(packages, NewBasicArrayListD()); _ = nameSpaceWithPropertiesVector
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().HTTPREQUEST) {
            this.LogUtil.PutF("\nSep Tokens: " + fmt.Sprint(nameSpaceWithPropertiesVector), this, "NameSpaceRequestParam()")
        }
        var nameSpaceAndPropertiesTokenizer *Tokenizer = NewTokenizer(NameSpaceRequestParamData_PROPERTIES); _ = nameSpaceAndPropertiesTokenizer
        var packageIndex int = 0; _ = packageIndex
        var list IBasicArrayList = NewBasicArrayListD(); _ = list
        var index int = 0; _ = index
        var size int = nameSpaceWithPropertiesVector.Size(); _ = size
        for index < size {
            var nameSpaceWithProperties string = any(nameSpaceWithPropertiesVector.Get(Java_PostIncrement(&index))).(string); _ = nameSpaceWithProperties
            list.Clear()
            var nameSpaceAndPropertiesVector IBasicArrayList = nameSpaceAndPropertiesTokenizer.GetTokensFromString(nameSpaceWithProperties, list); _ = nameSpaceAndPropertiesVector
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().HTTPREQUEST) {
                this.LogUtil.PutF("\nPackageName and PackageName Properties Tokens: " + fmt.Sprint(nameSpaceAndPropertiesVector), this, "NameSpaceRequestParam()")
            }
            if nameSpaceAndPropertiesVector.Size() > 0 {
                var packageName string = any(nameSpaceAndPropertiesVector.Get(0)).(string); _ = packageName
                if packageName == "" {
                    panic(errors.New(fmt.Sprint("Package Name Is Null")))
                }
                this.NameSpaceVector.Add(packageName)
                if nameSpaceAndPropertiesVector.Size() > 1 {
                    var properties string = any(nameSpaceAndPropertiesVector.Get(1)).(string); _ = properties
                    var packagePropertiesHashMap *HashMap = this.GeneratePackagePropertiesHashMap(properties); _ = packagePropertiesHashMap
                    this.NameSpacePropertiesHashMap.Put(fmt.Sprint(packageIndex), packagePropertiesHashMap)
                }
                packageIndex++
            }
        }
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().HTTPREQUEST) {
            this.LogUtil.PutF("NameSpace: " + fmt.Sprint(this.NameSpaceVector), this, "NameSpaceRequestParam()")
        }
    return this
}


func (this *NameSpaceRequestParam) GetPackages() IBasicArrayList {
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().HTTPREQUEST) {
            this.LogUtil.PutF("NameSpace: " + fmt.Sprint(this.NameSpaceVector), this, "getPackages()")
        }
        return this.NameSpaceVector
}


func (this *NameSpaceRequestParam) GetPackageProperties(packageIndex int) *HashMap {
        var packagePropertiesHashMap *HashMap = any(this.NameSpacePropertiesHashMap.Get(fmt.Sprint(packageIndex))).(*HashMap); _ = packagePropertiesHashMap
        if packagePropertiesHashMap == nil {
            packagePropertiesHashMap = StdUtil_GetInstance().CreateHashMap()
        }
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().HTTPREQUEST) {
            this.LogUtil.PutF("Package Properties: " + fmt.Sprint(packagePropertiesHashMap), this, "getPackagesProperties()")
        }
        return packagePropertiesHashMap
}


func (this *NameSpaceRequestParam) GetValue() string {
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().HTTPREQUEST) {
            this.LogUtil.PutF("\nValue: " + this.Value, this, "getValue()")
        }
        return this.Value
}


// throws Exception
func (this *NameSpaceRequestParam) GeneratePackagePropertiesHashMap(properties string) *HashMap {
        var propertiesTokenizer *Tokenizer = NewTokenizer(NameSpaceRequestParamData_PROPERTIESSEPARATOR); _ = propertiesTokenizer
        var packagePropertiesHashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = packagePropertiesHashMap
        var propertyVector IBasicArrayList = propertiesTokenizer.GetTokensFromString(properties, NewBasicArrayListD()); _ = propertyVector
        var list IBasicArrayList = NewBasicArrayListD(); _ = list
        var propertyTokenizer *Tokenizer = NewTokenizer(HtmlNameMathData_GetInstance().EQUALS); _ = propertyTokenizer
        var index int = 0; _ = index
        var size int = propertyVector.Size(); _ = size
        for index < size {
            var property string = any(propertyVector.Get(index)).(string); _ = property
            list.Clear()
            var propertyNameValueVector IBasicArrayList = propertyTokenizer.GetTokensFromString(property, list); _ = propertyNameValueVector
            var index2 int = 0; _ = index2
            for index2 < propertyNameValueVector.Size() {
                var propertyName string = any(propertyNameValueVector.Get(Java_PostIncrement(&index2))).(string); _ = propertyName
                if propertyNameValueVector.Size() > 1 {
                    var propertyValue string = any(propertyNameValueVector.Get(Java_PostIncrement(&index2))).(string); _ = propertyValue
                    if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().HTTPREQUEST) {
                        this.LogUtil.PutF("\nProperty Name: " + propertyName + "\nProperty Value: " + propertyValue, this, "generatePackagePropertiesHashMap()")
                    }
                    packagePropertiesHashMap.Put(propertyName, propertyValue)
                } else {
                    packagePropertiesHashMap.Put(propertyName, StringUtil_GetInstance().EMPTY_STRING)
                }
            }
        }
        return packagePropertiesHashMap
}

