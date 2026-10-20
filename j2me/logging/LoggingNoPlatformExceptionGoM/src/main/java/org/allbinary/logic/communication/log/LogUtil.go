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

//NoPlatform
type LogUtil struct {
}

var LogUtil_Instance *LogUtil = NewLogUtil()


func LogUtil_GetInstance() *LogUtil {
        return LogUtil_Instance
}


func NewLogUtil() *LogUtil {
    this := &LogUtil{}
    return this
}


func (this *LogUtil) PutL(log *Log) {
        panic(errors.New(fmt.Sprint()))
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
        panic(errors.New(fmt.Sprint()))
}

