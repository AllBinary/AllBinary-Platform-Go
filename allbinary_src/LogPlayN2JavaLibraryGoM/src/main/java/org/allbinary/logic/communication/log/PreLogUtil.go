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
type PreLogUtil struct {
}

var PreLogUtil_LOG_SUCCESS string = "org.allbinary: "


func NewPreLogUtil() *PreLogUtil {
    this := &PreLogUtil{}
    return this
}


func PreLogUtil_Put(specialMessage string, anyType any, functionName string) {
        var className string = PreLogUtil_GetClassName(anyType); _ = className
        if className == "" {
            className = CommonStrings_GetInstance().EMPTY
        }
        className = fmt.Sprint(NewStringMaker().Append(className).Append(CommonSeps_GetInstance().FORWARD_SLASH).Append(StringUtil_GetInstance().ToString(anyType)))
        var message string = LogFormatUtil_GetInstance().GetS(className, functionName, specialMessage); _ = message
        PlayN_GetInstance().Log().Debug(fmt.Sprint(PreLogUtil_LOG_SUCCESS) + message)
}


func PreLogUtil_PutOE(specialMessage string, anyType any, functionName string, exception any) {
        var className string = PreLogUtil_GetClassName(anyType); _ = className
        if className == "" {
            className = CommonStrings_GetInstance().EMPTY
        }
        className = fmt.Sprint(NewStringMaker().Append(className).Append(CommonSeps_GetInstance().FORWARD_SLASH).Append(StringUtil_GetInstance().ToString(anyType)))
        var message string = LogFormatUtil_GetInstance().Get(className, functionName, specialMessage, exception); _ = message
        PlayN_GetInstance().Log().Error(fmt.Sprint(PreLogUtil_LOG_SUCCESS) + message, any(exception).(error))
}


func PreLogUtil_PutS(specialMessage string, className string, functionName string) {
        var message string = LogFormatUtil_GetInstance().GetS(className, functionName, specialMessage); _ = message
        PlayN_GetInstance().Log().Debug(fmt.Sprint(PreLogUtil_LOG_SUCCESS) + message)
}


func PreLogUtil_PutSE(specialMessage string, className string, functionName string, exception any) {
        var message string = LogFormatUtil_GetInstance().Get(className, functionName, specialMessage, exception); _ = message
        PlayN_GetInstance().Log().Error(fmt.Sprint(PreLogUtil_LOG_SUCCESS) + message, any(exception).(error))
}


func PreLogUtil_GetClassName(anyType any) string {
    panic("abstract method getClassName")
}

