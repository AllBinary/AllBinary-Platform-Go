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

//Android
type ResourceUtil struct {
    LogUtil *LogUtil
    Context any
    Resources any
    HashMap *ABHashtable[any, any]
}

var ResourceUtil_Instance *ResourceUtil = NewResourceUtil()

func (this *ResourceUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.Context = NullUtil_GetInstance().NULL_OBJECT
    this.Resources = NullUtil_GetInstance().NULL_OBJECT
    this.HashMap = StdUtil_GetInstance().CreateHashtable()
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


func (this *ResourceUtil) GetContext() *Context {
        return any(this.Context).(*Context)
}


func (this *ResourceUtil) SetContextFromActivity(activity *Activity) {
        this.Context = activity
}


func (this *ResourceUtil) SetContext(aContext *Context) {
        this.Context = aContext
}


func (this *ResourceUtil) GetResourceId(resource string) int {
        var value int = any(this.HashMap.Get(resource)).(int); _ = value
        if DebugFactory_GetInstance() != NoDebug_GetInstance() {
            PreLogUtil_Put(fmt.Sprint(NewStringMaker().Append(resource).Append(CommonSeps_GetInstance().COLON).Append(fmt.Sprint(value))), this, "getResourceId")
        }
        return value
}


func (this *ResourceUtil) AddResource(resource string, value int) {
        if DebugFactory_GetInstance() != NoDebug_GetInstance() {
            PreLogUtil_Put(fmt.Sprint(NewStringMaker().Append(resource).Append(CommonSeps_GetInstance().COLON).Append(fmt.Sprint(value))), this, "addResource")
            if this.ContainsDuplicate(resource, value) {
                ForcedLogUtil_Log(fmt.Sprint(NewStringMaker().Append("Found Duplicate Resource: ").Append(resource)), this)
            }
        }
        this.HashMap.Put(resource, value)
}


func (this *ResourceUtil) ContainsDuplicate(resource string, value int) bool {
        var objectArray []any = HashtableUtil_GetInstance().GetKeysAsArray(this.HashMap); _ = objectArray
        for index := int(len(objectArray)); Java_PreDecrement(&index) >= 0;  {
            var integer int = any(this.HashMap.Get(any(objectArray[index]).(any))).(int); _ = integer
            if resource != objectArray[index] {
                if value != integer {
                    if int(value) == int(integer) {
                        return true
                    }
                }
            }
        }
        return false
}


func (this *ResourceUtil) SetResources(resources *Resources) {
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append("Resource Loader: ").Append(reflect.TypeOf(resources).String())), this, "setResources")
        this.Resources = resources
}


// throws Exception
func (this *ResourceUtil) GetResourceAsStream(resource string) IInputStream {
        var integer int = any(this.HashMap.Get(resource)).(int); _ = integer
        var id int = int(integer); _ = id
        var resources *Resources = (any(this.Resources).(*Resources)); _ = resources
        var inputStream IInputStream = resources.OpenRawResource(id); _ = inputStream
        if inputStream != nil {
            return inputStream
        }
        return inputStream
}

