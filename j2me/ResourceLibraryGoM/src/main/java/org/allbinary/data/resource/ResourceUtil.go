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

//J2SE - J2ME
type ResourceUtil struct {
    LogUtil *LogUtil
}

var ResourceUtil_ClassLoader any = NullUtil_GetInstance().NULL_OBJECT
var ResourceUtil_Instance *ResourceUtil = NewResourceUtil()

func (this *ResourceUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func ResourceUtil_GetInstance() *ResourceUtil {
        return ResourceUtil_Instance
}


func NewResourceUtil() *ResourceUtil {
    this := &ResourceUtil{}
    this.initFields()
    return this
}


func (this *ResourceUtil) SetLoadingPaths(path string, ext string) {
        this.LogUtil.PutF(CommonStrings_GetInstance().NOT_IMPLEMENTED, this, "setLoadingPaths")
}


func (this *ResourceUtil) SetClassLoader(classLoader *ClassLoader) {
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append("Resource Loader: ").Append(reflect.TypeOf(classLoader).String())), this, "setClassLoader")
        ResourceUtil_ClassLoader = classLoader
}


// throws Exception
func (this *ResourceUtil) GetResourceAsStream(resource string) IInputStream {
        var inputStream IInputStream = this.GetResourceAsStreamAtStart(resource, 2); _ = inputStream
        if inputStream == nil {
            inputStream = this.GetResourceAsStreamAtStart(resource, 1)
            if inputStream == nil {
                panic(errors.New(fmt.Sprint(fmt.Sprint(NewStringMaker().Append("Unable to obtain: ").Append(resource)))))
            }
        }
        return inputStream
}


// throws Exception
func (this *ResourceUtil) GetResourceAsStreamAtStart(resource string, startIndex int) IInputStream {
        var commonSeps *CommonSeps = CommonSeps_GetInstance(); _ = commonSeps
        var index int = strings.Index(resource, string(commonSeps.COLON)); _ = index
        var resourcePath string = resource[index + startIndex:]; _ = resourcePath
        var inputStream IInputStream = reflect.TypeOf(resource).GetResourceAsStream(resourcePath); _ = inputStream
        if inputStream != nil {
            return inputStream
        }
        var classLoader *ClassLoader = any(ResourceUtil_ClassLoader).(*ClassLoader); _ = classLoader
        inputStream = classLoader.GetResourceAsStream(resourcePath)
        if inputStream != nil {
            return inputStream
        }
        inputStream = Thread_CurrentThread().GetContextClassLoader().GetResourceAsStream(resourcePath)
        if inputStream != nil {
            return inputStream
        }
        return inputStream
}


func (this *ResourceUtil) AddResource(resource string, value int) {
}

