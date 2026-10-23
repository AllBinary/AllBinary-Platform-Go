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

//ActualPlatform
type LogFactory struct {
}

var LogFactory_LOG *Log = NewLog(StringUtil_GetInstance().EMPTY_STRING, NullUtil_GetInstance().NULL_OBJECT, StringUtil_GetInstance().EMPTY_STRING, NullUtil_GetInstance().NULL_OBJECT)

func NewLogFactory() *LogFactory {
    this := &LogFactory{}
    return this
}


func LogFactory_GetInstanceD() *Log {
        var stringUtil *StringUtil = StringUtil_GetInstance(); _ = stringUtil
        var nullUtil *NullUtil = NullUtil_GetInstance(); _ = nullUtil
        return NewLog(stringUtil.EMPTY_STRING, nullUtil.NULL_OBJECT, stringUtil.EMPTY_STRING, nullUtil.NULL_OBJECT)
}


func LogFactory_GetInstance(specialMessage string, anyType any, functionName string, exception any) *Log {
        return LogFactory_LOG
}


func LogFactory_GetInstanceF(specialMessage string, anyType any, functionName string) *Log {
        return LogFactory_LOG
}

