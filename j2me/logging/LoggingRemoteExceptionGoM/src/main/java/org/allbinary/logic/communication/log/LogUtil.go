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
    MESSAGE string
    IsFirstException bool
    TimeDelayHelper ITimeDelayHelper
    AbeClientInformation AbeClientInformationInterface
    StringBuffer *StringMaker
}

var LogUtil_Instance *LogUtil = NewLogUtil()

func (this *LogUtil) initFields() {
    this.MESSAGE = "message"
    this.IsFirstException = true
    this.TimeDelayHelper = NewTimeDelayHelper(200000)
    
    this.StringBuffer = NewStringMaker()
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
        if log == nil {
            return
        }
        var specialMessage string = log.GetSpecialMessage(); _ = specialMessage
        var anyType any = log.GetObject(); _ = anyType
        var functionName string = log.GetFunctionName(); _ = functionName
        var exception any = log.GetThrowable(); _ = exception
        this.Put(specialMessage, anyType, functionName, exception)
}


func (this *LogUtil) PutF_1(specialMessage string, anyType any, functionName string) {
}

// Java overloads of LogUtil.putF
func (this *LogUtil) PutF(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 3 && Java_IsArg[string](args[0], pass) && Java_IsArg[any](args[1], pass) && Java_IsArg[string](args[2], pass) {
            this.PutF_1(Java_Arg[string](args[0]), Java_Arg[any](args[1]), Java_Arg[string](args[2]))
            return
        }
        if len(args) == 3 && Java_IsArg[string](args[0], pass) && Java_IsArg[string](args[1], pass) && Java_IsArg[string](args[2], pass) {
            this.PutF_2(Java_Arg[string](args[0]), Java_Arg[string](args[1]), Java_Arg[string](args[2]))
            return
        }
    }
    panic("No Java overload of LogUtil.putF for the arguments")
}


func (this *LogUtil) PutF_2(specialMessage string, className string, functionName string) {
}


func (this *LogUtil) Put(specialMessage string, anyType any, functionName string, exception any) {
        if exception != nil {
            if strings.Compare(reflect.TypeOf(exception).String(), reflect.TypeOf((*XmlRpcException)(nil)).Elem().GetName()) == 0 {
                return
            }
            if strings.Compare(reflect.TypeOf(exception).String(), reflect.TypeOf((*IOException)(nil)).Elem().GetName()) == 0 {
                return
            }
            if this.IsFirstException || this.TimeDelayHelper.IsTimeTNT() {
                var className string = CommonStrings_GetInstance().EMPTY; _ = className
                this.IsFirstException = false
                if anyType != nil && reflect.TypeOf(anyType).String() != "" {
                    className = reflect.TypeOf(anyType).String()
                }
                var message string = LogFormatUtil_GetInstance().Get(className, functionName, specialMessage, exception); _ = message
                if tryControl1 := func() (tryControl1 int) {
                    defer func() {
                        if recovered := recover(); recovered != nil {
                            e := Java_ToError(recovered)
                            _ = e
                    fmt.Println("Exception")
                    fmt.Println(e)
                        }
                    }()
                    fmt.Println("Eeeek")
                    if this.AbeClientInformation == nil {
                        panic(errors.New(fmt.Sprint()))
                    }
                    var hashtable *ABHashtable[any, any] = this.AbeClientInformation.ToHashtable(); _ = hashtable
                    this.StringBuffer.Delete(0, this.StringBuffer.Length())
                    this.StringBuffer.Append(message)
                    this.StringBuffer.Append(CommonSeps_GetInstance().SPACE)
                    this.StringBuffer.Append(SpecialMessageUtil_GetInstance().Get())
                    hashtable.Put(this.MESSAGE, fmt.Sprint(this.StringBuffer))
                    NewXmlRpcRemoteLogClient(this.AbeClientInformation).Get(hashtable)
                    return
                }(); tryControl1 == 1 {
                    return
                }
            }
        }
}

