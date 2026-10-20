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
    Logger *Logger
    LogFormatUtil *LogFormatUtil
}

var LogUtil_Instance *LogUtil = NewLogUtil()

func (this *LogUtil) initFields() {
    this.Logger = Logger_GetLogger(reflect.TypeOf((*LogUtil)(nil)).Elem().GetName())
    this.LogFormatUtil = LogFormatUtil_GetInstance()
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


func (this *LogUtil) PutFS(specialMessage string, className string, functionName string) {
        this.PutS(specialMessage, className, functionName, NullUtil_GetInstance().NULL_OBJECT)
}


func (this *LogUtil) Put(specialMessage string, anyType any, functionName string, exception any) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
                }
            }()
            var className string = CommonStrings_GetInstance().EMPTY; _ = className
            var clazz reflect.Type = reflect.TypeOf(anyType); _ = clazz
            if clazz.String() != "" {
                className = clazz.String()
            }
            var message string = this.LogFormatUtil.GetS(className, functionName, specialMessage); _ = message
            if exception != nil {
                this.Logger.Log(Level_SEVERE, message, exception)
            } else {
                this.Logger.Log(Level_INFO, message)
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *LogUtil) PutS(specialMessage string, className string, functionName string, exception any) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
                }
            }()
            var message string = this.LogFormatUtil.GetS(className, functionName, specialMessage); _ = message
            if exception != nil {
                this.Logger.Log(Level_SEVERE, message, exception)
            } else {
                this.Logger.Log(Level_INFO, message)
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}

