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

type Replace struct {
    LogUtil *LogUtil
    J2seUtil *J2SEUtil
    REPLACERS_ string
    ALL string
    FOUND_KEY string
    HashMap *HashMap
}


func (this *Replace) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.J2seUtil = J2SEUtil_GetInstance()
    this.REPLACERS_ = "Replacers: "
    this.ALL = "all"
    this.FOUND_KEY = "found key total: "
    
}




// Java overloads of new Replace
func NewReplace(args ...any) *Replace {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 2 && Java_IsArg[string](args[0], pass) && Java_IsArg[string](args[1], pass) {
            return NewReplace_1(Java_Arg[string](args[0]), Java_Arg[string](args[1]))
        }
        if len(args) == 2 && Java_IsArg[[]string](args[0], pass) && Java_IsArg[[]string](args[1], pass) {
            return NewReplace_2(Java_Arg[[]string](args[0]), Java_Arg[[]string](args[1]))
        }
        if len(args) == 1 && Java_IsArg[*HashMap](args[0], pass) {
            return NewReplace_3(Java_Arg[*HashMap](args[0]))
        }
    }
    panic("No Java overload of new Replace for the arguments")
}
func NewReplace_1(key string, value string) *Replace {
    this := &Replace{}
    this.initFields()
        this.HashMap = StdUtil_GetInstance().CreateHashMap()
        this.HashMap.Put([]string{key}, []string{value})
    return this
}


func NewReplace_2(keys []string, values []string) *Replace {
    this := &Replace{}
    this.initFields()
        this.HashMap = StdUtil_GetInstance().CreateHashMap()
        this.HashMap.Put(keys, values)
    return this
}


func NewReplace_3(hashMap *HashMap) *Replace {
    this := &Replace{}
    this.initFields()
        this.HashMap = hashMap
    return this
}


func (this *Replace) All(replace string) string {
        return this.UpTo(replace, Integer_MIN_VALUE)
}


func (this *Replace) UpTo(replace string, total int) string {
        var newStringBuffer *StringMaker = NewStringMaker(); _ = newStringBuffer
        var totalNumberOfReplaces int = 0; _ = totalNumberOfReplaces
        var keys []string; _ = keys
        var values []string; _ = values
        var keyArray []any = this.J2seUtil.GetHashMapAsArray(this.HashMap); _ = keyArray
        var size int = len(keyArray); _ = size
        for index2 := int(0); index2 < size; index2++ {
            keys = any(keyArray[index2]).([]string)
            values = any(this.HashMap.Get(keys)).([]string)
            var foundTotal int64 = 0; _ = foundTotal
            var index int = 0; _ = index
            var nextKey string; _ = nextKey
            var end int; _ = end
            var nextSecondKey string; _ = nextSecondKey
            var begin2 int; _ = begin2
            for index < len(replace) {
                var begin int = replace.IndexOf(keys[0], index); _ = begin
                if begin != -1 {
                    foundTotal++
                    if foundTotal % 100 == 0 {
                        fmt.Println(this.FOUND_KEY + fmt.Sprint(foundTotal))
                    }
                    nextKey = keys[0]
                    end = begin + len(nextKey)
                    newStringBuffer.Delete(0, newStringBuffer.Length())
                    newStringBuffer.Append(replace[0:begin])
                    newStringBuffer.Append(values[0])
                    index = newStringBuffer.Length()
                    newStringBuffer.Append(replace[end:len(replace)])
                    replace = fmt.Sprint(newStringBuffer)
                    if len(keys) > 1 {
                        nextSecondKey = keys[1]
                        begin2 = replace.IndexOf(nextSecondKey, index)
                        if begin2 != -1 {
                            var end2 int = begin2 + len(nextSecondKey); _ = end2
                            newStringBuffer.Delete(0, newStringBuffer.Length())
                            newStringBuffer.Append(replace[0:begin2])
                            newStringBuffer.Append(values[1])
                            index = newStringBuffer.Length()
                            newStringBuffer.Append(replace[end2:len(replace)])
                            replace = fmt.Sprint(newStringBuffer)
                        }
                    }
                    totalNumberOfReplaces++
                    if total != Integer_MIN_VALUE && totalNumberOfReplaces >= total {
                        break
                    }
                } else {
                    break
                }
            }
        }
        if totalNumberOfReplaces > 0 {
            if LogConfigTypes_LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().REPLACE) {
                this.LogUtil.PutF("Total Number Of Replaces: " + fmt.Sprint(totalNumberOfReplaces), this, this.ALL)
            }
        }
        return replace
}


func (this *Replace) Line(replace string) string {
        var newStringBuffer *StringMaker = NewStringMaker(); _ = newStringBuffer
        var totalNumberOfReplaces int = 0; _ = totalNumberOfReplaces
        var key string; _ = key
        var value string; _ = value
        var keyArray []any = this.J2seUtil.GetHashMapAsArray(this.HashMap); _ = keyArray
        var size int = len(keyArray); _ = size
        for index2 := int(0); index2 < size; index2++ {
            key = any(keyArray[index2]).(string)
            value = any(this.HashMap.Get(key)).(string)
            var index int = 0; _ = index
            for index < len(replace) {
                newStringBuffer.Delete(0, newStringBuffer.Length())
                var begin int = replace.IndexOf(key, index); _ = begin
                if begin != -1 {
                    fmt.Println(this.FOUND_KEY)
                    var end int = replace.IndexOf('\n', begin + len(key)) + 1; _ = end
                    if end >= 0 {
                        newStringBuffer.Append(replace[0:begin])
                        index = newStringBuffer.Length()
                        newStringBuffer.Append(replace[end:len(replace)])
                        replace = fmt.Sprint(newStringBuffer)
                        totalNumberOfReplaces++
                    }
                } else {
                    break
                }
            }
        }
        if LogConfigTypes_LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().REPLACE) {
            this.LogUtil.PutF("Total Number Of Replaces: " + fmt.Sprint(totalNumberOfReplaces), this, this.ALL)
        }
        return replace
}


func Replace_Main(args []string) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
                }
            }()
            var stringUtil *StringUtil = StringUtil_GetInstance(); _ = stringUtil
            var appUrlGlobals *AppUrlGlobals = NewAppUrlGlobals(); _ = appUrlGlobals
            appUrlGlobals.SetWebappPath("G:/mnt/bc/mydev/work/allbinary_src/ToolsJavaLibrary/tools")
            URLGLOBALS.Init(appUrlGlobals)
            var testString string = "\"super gun\" big tank \"goodgunriflegun"; _ = testString
            var replace *Replace = NewReplace("gun", "abcdefghijklmnopqrstuxwxyzabcdefghijklmnopqrstuxwxyz"); _ = replace
            fmt.Println("Old String: " + testString)
            testString = replace.All(testString)
            fmt.Println("New String: " + testString)
            var replaceCat *Replace = NewReplace("abcdefghijklmnopqrstuxwxyzabcdefghijklmnopqrstuxwxyz", stringUtil.EMPTY_STRING); _ = replaceCat
            testString = replaceCat.All(testString)
            fmt.Println("New String: " + testString)
            return
        }(); tryControl1 == 1 {
            return
        }
}

