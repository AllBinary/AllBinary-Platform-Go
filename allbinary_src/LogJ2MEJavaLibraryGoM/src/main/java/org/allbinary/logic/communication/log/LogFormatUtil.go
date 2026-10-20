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
type LogFormatUtil struct {
    TimeStampUtil *TimeStampUtil
    CommonSeps *CommonSeps
    NONE string
    LOG_ERROR string
    EMPTY string
    STACK_TRACE string
    TIME string
    CLASS_NAME string
    FUNCTION_CALL string
    SPECIAL_MESSAGE string
    StringUtil *StringUtil
    ExceptionUtil *ExceptionUtil
    NullUtil *NullUtil
}

var LogFormatUtil_Instance *LogFormatUtil = NewLogFormatUtil()

func (this *LogFormatUtil) initFields() {
    this.TimeStampUtil = TimeStampUtil_GetInstance()
    this.CommonSeps = CommonSeps_GetInstance()
    this.NONE = "None"
    this.LOG_ERROR = "\nLog-Error: "
    this.EMPTY = "Empty"
    this.STACK_TRACE = "\nStack Trace: "
    this.TIME = "Time: "
    this.CLASS_NAME = this.CommonSeps.SPACE
    this.FUNCTION_CALL = CommonLabels_GetInstance().COLON_SEP
    this.SPECIAL_MESSAGE = "> "
    this.StringUtil = StringUtil_GetInstance()
    this.ExceptionUtil = ExceptionUtil_GetInstance()
    this.NullUtil = NullUtil_GetInstance()
}


func LogFormatUtil_GetInstance() *LogFormatUtil {
        return LogFormatUtil_Instance
}


func NewLogFormatUtil() *LogFormatUtil {
    this := &LogFormatUtil{}
    this.initFields()
    return this
}


func (this *LogFormatUtil) Get(className string, functionName string, specialMessage string, exception any) string {
        var stringBuffer *StringMaker = this.GetF(className, functionName); _ = stringBuffer
        stringBuffer.Append(this.GetO(exception))
        stringBuffer.Append(this.SPECIAL_MESSAGE)
        stringBuffer.Append(specialMessage)
        stringBuffer.Append(this.CommonSeps.NEW_LINE)
        return fmt.Sprint(stringBuffer)
}


func (this *LogFormatUtil) GetS(className string, functionName string, specialMessage string) string {
        var stringBuffer *StringMaker = this.GetF(className, functionName); _ = stringBuffer
        stringBuffer.Append(this.SPECIAL_MESSAGE)
        stringBuffer.Append(specialMessage)
        stringBuffer.Append(this.CommonSeps.NEW_LINE)
        return fmt.Sprint(stringBuffer)
}


func (this *LogFormatUtil) GetF(className string, functionName string) *StringMaker {
        if functionName == "" {
            functionName = this.NONE
        }
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(this.TIME)
        stringBuffer.Append(this.TimeStampUtil.GetAsString())
        stringBuffer.Append(this.CLASS_NAME)
        stringBuffer.Append(className)
        stringBuffer.Append(this.FUNCTION_CALL)
        stringBuffer.Append(functionName)
        return stringBuffer
}


func (this *LogFormatUtil) GetO(exception any) string {
        if exception != this.NullUtil.NULL_OBJECT {
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append(this.LOG_ERROR)
            var exceptionAsString string = fmt.Sprint(exception); _ = exceptionAsString
            if exceptionAsString != "" {
                stringBuffer.Append(exceptionAsString)
            } else {
                stringBuffer.Append(this.EMPTY)
            }
            stringBuffer.Append(this.STACK_TRACE)
            if exception != this.NullUtil.NULL_OBJECT {
                stringBuffer.Append(this.ExceptionUtil.GetStackTrace(any(exception).(error)))
            } else {
                stringBuffer.Append(this.StringUtil.NULL_STRING)
            }
            return fmt.Sprint(stringBuffer)
        } else {
            return StringUtil_GetInstance().EMPTY_STRING
        }
}
