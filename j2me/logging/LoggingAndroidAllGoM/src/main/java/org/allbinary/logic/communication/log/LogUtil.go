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
type LogUtil struct {
    CommonSeps *CommonSeps
    LogFormatUtil *LogFormatUtil
    LABEL string
}

var LogUtil_Instance *LogUtil = NewLogUtil()

func (this *LogUtil) initFields() {
    this.CommonSeps = CommonSeps_GetInstance()
    this.LogFormatUtil = LogFormatUtil_GetInstance()
    this.LABEL = "org.allbinary"
}


func LogUtil_GetInstance() *LogUtil {
        return LogUtil_Instance
}


func NewLogUtil() *LogUtil {
    this := &LogUtil{}
    this.initFields()
    return this
}


func (this *LogUtil) PutL(log *Log) {
        var specialMessage string = log.GetSpecialMessage(); _ = specialMessage
        var anyType any = log.GetObject(); _ = anyType
        var functionName string = log.GetFunctionName(); _ = functionName
        var exception any = log.GetThrowable(); _ = exception
        this.Put(specialMessage, anyType, functionName, exception)
}


func (this *LogUtil) PutF(specialMessage string, anyType any, functionName string) {
        this.Put(specialMessage, anyType, functionName, NullUtil_GetInstance().NULL_OBJECT)
}


func (this *LogUtil) PutFS(specialMessage string, string_ string, functionName string) {
        this.PutS(specialMessage, string_, functionName, NullUtil_GetInstance().NULL_OBJECT)
}


func (this *LogUtil) Put(specialMessage string, anyType any, functionName string, exception any) {
        var className string = this.LABEL; _ = className
        className = fmt.Sprint(NewStringMaker().Append(reflect.TypeOf(anyType).String()).Append(this.CommonSeps.COLON).Append(strconv.FormatInt(int64(TsUtil_GetInstance().HashCode(anyType)), 16)))
        var message string = this.LogFormatUtil.Get(className, functionName, specialMessage, exception); _ = message
        android_Util.Log.I(this.LABEL, message)
}


func (this *LogUtil) PutS(specialMessage string, string_ string, functionName string, exception any) {
        var className string = this.LABEL; _ = className
        className = fmt.Sprint(NewStringMaker().Append(string_).Append(this.CommonSeps.COLON).Append(strconv.FormatInt(int64(TsUtil_GetInstance().HashCode(string_)), 16)))
        var message string = this.LogFormatUtil.Get(className, functionName, specialMessage, exception); _ = message
        android_Util.Log.I(this.LABEL, message)
}

