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

type Log struct {
    LogFormatUtil *LogFormatUtil
    SpecialMessage string
    AnyType any
    FunctionName string
    Exception any
}


func (this *Log) initFields() {
    this.LogFormatUtil = LogFormatUtil_GetInstance()
    
    
    
    
}


func NewLog(specialMessage string, anyType any, functionName string, exception any) *Log {
    this := &Log{}
    this.initFields()
        this.SpecialMessage = specialMessage
        this.AnyType = anyType
        this.FunctionName = functionName
        this.Exception = exception
    return this
}


func (this *Log) GetSpecialMessage() string {
        return this.SpecialMessage
}


func (this *Log) GetObject() any {
        return this.AnyType
}


func (this *Log) GetFunctionName() string {
        return this.FunctionName
}


func (this *Log) GetThrowable() any {
        return this.Exception
}


func (this *Log) ToString() string {
        var className string = CommonStrings_GetInstance().EMPTY; _ = className
        var clazz reflect.Type = reflect.TypeOf(this.AnyType); _ = clazz
        if clazz.String() != "" {
            className = clazz.String()
        }
        return this.LogFormatUtil.Get(className, this.FunctionName, this.SpecialMessage, this.Exception)
}

func (this *Log) String() string {
    return this.ToString()
}

