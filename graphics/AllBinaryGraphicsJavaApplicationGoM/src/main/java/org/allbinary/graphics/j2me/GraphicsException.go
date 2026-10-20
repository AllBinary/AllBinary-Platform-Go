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

type GraphicsException struct {
    Exception
    LogUtil *LogUtil
    GuiLog *GuiLog
}


func (this *GraphicsException) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.GuiLog = GuiLog_GetInstance()
}




// Java overloads of new GraphicsException
func NewGraphicsException(args ...any) *GraphicsException {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 3 && Java_IsArg[string](args[0], pass) && Java_IsArg[any](args[1], pass) && Java_IsArg[string](args[2], pass) {
            return NewGraphicsException_1(Java_Arg[string](args[0]), Java_Arg[any](args[1]), Java_Arg[string](args[2]))
        }
        if len(args) == 3 && Java_IsArg[string](args[0], pass) && Java_IsArg[string](args[1], pass) && Java_IsArg[string](args[2], pass) {
            return NewGraphicsException_2(Java_Arg[string](args[0]), Java_Arg[string](args[1]), Java_Arg[string](args[2]))
        }
    }
    panic("No Java overload of new GraphicsException for the arguments")
}
func NewGraphicsException_1(msg string, obj any, method string) *GraphicsException {
    this := &GraphicsException{}
        this.Exception = *NewException(msg)
        this.initFields()
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
                }
            }()
            this.GuiLog.ShowDialog(msg)
            this.LogUtil.Put(msg, obj, method, this)
            return
        }(); tryControl1 == 1 {
            return
        }
    return this
}


func NewGraphicsException_2(msg string, className string, method string) *GraphicsException {
    this := &GraphicsException{}
        this.Exception = *NewException(msg)
        this.initFields()
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
                }
            }()
            this.GuiLog.ShowDialog(msg)
            this.LogUtil.Put(msg, className, method, this)
            return
        }(); tryControl1 == 1 {
            return
        }
    return this
}

