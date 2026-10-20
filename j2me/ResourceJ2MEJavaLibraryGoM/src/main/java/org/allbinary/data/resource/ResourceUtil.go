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

//J2ME
type ResourceUtil struct {
    LogUtil *LogUtil
}

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


// throws Exception
func (this *ResourceUtil) GetResourceAsStream(resource string) IInputStream {
        var inputStream IInputStream = reflect.TypeOf(resource).GetResourceAsStream(resource); _ = inputStream
        if inputStream == nil {
            inputStream = this.GetResourceAsStreamAtStart(resource, 2)
            if inputStream == nil {
                inputStream = this.GetResourceAsStreamAtStart(resource, 1)
                if inputStream == nil {
                    var stringMaker *StringMaker = NewStringMaker(); _ = stringMaker
                    var index int = strings.LastIndex(resource, string('/')); _ = index
                    var resourcePath string = resource[index + 1:]; _ = resourcePath
                    inputStream = reflect.TypeOf(resource).GetResourceAsStream(resourcePath)
                    if inputStream == nil {
                        resourcePath = resource[index:]
                        inputStream = reflect.TypeOf(resource).GetResourceAsStream(resourcePath)
                        if inputStream == nil {
                            var RES string = "res"; _ = RES
                            resourcePath = fmt.Sprint(stringMaker.Append(RES).Append(resource[index:]))
                            inputStream = reflect.TypeOf(resource).GetResourceAsStream(resourcePath)
                            if inputStream == nil {
                                stringMaker.Delete(0, stringMaker.Length())
                                resourcePath = fmt.Sprint(stringMaker.Append("/").Append(RES).Append(resource[index:]))
                                inputStream = reflect.TypeOf(resource).GetResourceAsStream(resourcePath)
                                var COLON string = CommonSeps_GetInstance().COLON; _ = COLON
                                if inputStream == nil {
                                    stringMaker.Delete(0, stringMaker.Length())
                                    resourcePath = fmt.Sprint(stringMaker.Append(RES).Append(COLON).Append(resource[index:]))
                                    inputStream = reflect.TypeOf(resource).GetResourceAsStream(resourcePath)
                                    if inputStream == nil {
                                        stringMaker.Delete(0, stringMaker.Length())
                                        resourcePath = fmt.Sprint(stringMaker.Append(RES).Append(COLON).Append(resource[index + 1:]))
                                        inputStream = reflect.TypeOf(resource).GetResourceAsStream(resourcePath)
                                        if inputStream == nil {
                                            var RESOURCE_STRING string = "resource"; _ = RESOURCE_STRING
                                            stringMaker.Delete(0, stringMaker.Length())
                                            resourcePath = fmt.Sprint(stringMaker.Append(RESOURCE_STRING).Append(COLON).Append(resource[index:]))
                                            inputStream = reflect.TypeOf(resource).GetResourceAsStream(resourcePath)
                                            if inputStream == nil {
                                                stringMaker.Delete(0, stringMaker.Length())
                                                resourcePath = fmt.Sprint(stringMaker.Append(RESOURCE_STRING).Append(COLON).Append(resource[index + 1:]))
                                                inputStream = reflect.TypeOf(resource).GetResourceAsStream(resourcePath)
                                                if inputStream == nil {
                                                    stringMaker.Delete(0, stringMaker.Length())
                                                    panic(errors.New(fmt.Sprint(fmt.Sprint(stringMaker.Append("Unable to obtain: ").Append(resource)))))
                                                }
                                            }
                                        }
                                    }
                                }
                            }
                        }
                    }
                }
            }
        }
        return inputStream
}


// throws Exception
func (this *ResourceUtil) GetResourceAsStreamAtStart(resource string, startIndex int) IInputStream {
        var stringMaker *StringMaker = NewStringMaker(); _ = stringMaker
        var index int = strings.Index(resource, string(CommonSeps_GetInstance().COLON)); _ = index
        var resourcePath string = resource[index + startIndex:]; _ = resourcePath
        var inputStream IInputStream = reflect.TypeOf(resource).GetResourceAsStream(resourcePath); _ = inputStream
        if inputStream != nil {
            return inputStream
        }
        return inputStream
}


func (this *ResourceUtil) AddResource(resource string, value int) {
}

