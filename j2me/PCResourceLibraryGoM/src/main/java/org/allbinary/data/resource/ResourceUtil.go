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

//J2SE - PC
type ResourceUtil struct {
    LogUtil *LogUtil
    Path string
    Ext string
}

var ResourceUtil_Instance *ResourceUtil = NewResourceUtil()

func (this *ResourceUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.Path = StringUtil_GetInstance().EMPTY_STRING
    this.Ext = StringUtil_GetInstance().EMPTY_STRING
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
        this.Path = path
        this.Ext = ext
}


func (this *ResourceUtil) SetClassLoader(classLoader *ClassLoader) {
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
        var inputStream IInputStream = NewFileInputStream(fmt.Sprint(NewStringMaker().Append(this.Path).Append(resource).Append(this.Ext))); _ = inputStream
        if inputStream != nil {
            var byteArray []int8 = make([]int8, inputStream.Available()); _ = byteArray
            StreamUtil_GetInstance().GetByteArray(inputStream, &strings.Builder{}, byteArray)
            return NewByteArrayInputStream(byteArray)
        }
        return inputStream
}


func (this *ResourceUtil) AddResource(resource string, value int) {
}

