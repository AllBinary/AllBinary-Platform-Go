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

type GuiLog struct {
    LogFormatUtil *LogFormatUtil
    CommonStrings *CommonStrings
}

var GuiLog_Instance *GuiLog = NewGuiLog()

func (this *GuiLog) initFields() {
    this.LogFormatUtil = LogFormatUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
}


func GuiLog_GetInstance() *GuiLog {
        return GuiLog_Instance
}


func NewGuiLog() *GuiLog {
    this := &GuiLog{}
    this.initFields()
    return this
}


// synchronized
func (this *GuiLog) ShowDialog(msg string) {
        var error_ *JDialog = NewJDialog(); _ = error_
        var size int = len(msg) * 9; _ = size
        var x int = size; _ = x
        if x > 550 {
            x = 550
        }
        error_.GetContentPane().SetLayout(NewGridLayout(1, 1))
        if size > 255 {
            var y int = 550; _ = y
            error_.SetSize(x, y)
            error_.GetContentPane().Add(NewJScrollPane(NewJTextArea(msg)))
        } else if size > 0 {
            var y int = 50; _ = y
            error_.SetSize(x, y)
            error_.GetContentPane().Add(NewJScrollPane(NewJLabel(msg)))
        }
        error_.Show()
}


// synchronized
func (this *GuiLog) Put_1(specialMessage string, anyType any, functionName string) string {
        return this.Put(specialMessage, anyType, functionName, nil)
}

// Java overloads of GuiLog.put
func (this *GuiLog) Put(args ...any) string {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 3 && Java_IsArg[string](args[0], pass) && Java_IsArg[any](args[1], pass) && Java_IsArg[string](args[2], pass) {
            return this.Put_1(Java_Arg[string](args[0]), Java_Arg[any](args[1]), Java_Arg[string](args[2]))
        }
        if len(args) == 4 && Java_IsArg[string](args[0], pass) && Java_IsArg[any](args[1], pass) && Java_IsArg[string](args[2], pass) && Java_IsArg[error](args[3], pass) {
            return this.Put_2(Java_Arg[string](args[0]), Java_Arg[any](args[1]), Java_Arg[string](args[2]), Java_Arg[error](args[3]))
        }
        if len(args) == 3 && Java_IsArg[string](args[0], pass) && Java_IsArg[string](args[1], pass) && Java_IsArg[string](args[2], pass) {
            return this.Put_3(Java_Arg[string](args[0]), Java_Arg[string](args[1]), Java_Arg[string](args[2]))
        }
        if len(args) == 4 && Java_IsArg[string](args[0], pass) && Java_IsArg[string](args[1], pass) && Java_IsArg[string](args[2], pass) && Java_IsArg[error](args[3], pass) {
            return this.Put_4(Java_Arg[string](args[0]), Java_Arg[string](args[1]), Java_Arg[string](args[2]), Java_Arg[error](args[3]))
        }
    }
    panic("No Java overload of GuiLog.put for the arguments")
}


// synchronized
func (this *GuiLog) Put_2(specialMessage string, anyType any, functionName string, exception error) string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            PreLogUtil_PutOE(this.CommonStrings.EXCEPTION, this, "put", e)
            tryControl1, tryResult1 = 1, "Logging Error"; return
                }
            }()
            var data string = this.LogFormatUtil.Get(specialMessage, reflect.TypeOf(anyType).String(), functionName, exception); _ = data
            this.ShowDialog(data)
            fmt.Println(data)
            tryControl1, tryResult1 = 1, data; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// synchronized
func (this *GuiLog) Put_3(specialMessage string, className string, functionName string) string {
        return this.Put(specialMessage, className, functionName, nil)
}


// synchronized
func (this *GuiLog) Put_4(specialMessage string, className string, functionName string, exception error) string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            PreLogUtil_PutOE(this.CommonStrings.EXCEPTION, this, "put", e)
            tryControl1, tryResult1 = 1, "Logging Error"; return
                }
            }()
            var data string = this.LogFormatUtil.Get(specialMessage, className, functionName, exception); _ = data
            this.ShowDialog(data)
            fmt.Println(data)
            tryControl1, tryResult1 = 1, data; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

