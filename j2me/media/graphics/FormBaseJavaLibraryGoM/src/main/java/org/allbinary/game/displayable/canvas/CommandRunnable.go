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

type CommandRunnable struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    CommandFormInputProcessor ICommandFormInputProcessor
    Command ICommand
}

// Implements Runnable


func (this *CommandRunnable) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
    
}


func NewCommandRunnable(commandFormInputProcessor ICommandFormInputProcessor, command ICommand) *CommandRunnable {
    this := &CommandRunnable{}
    this.initFields()
        this.CommandFormInputProcessor = commandFormInputProcessor
        this.Command = command
    return this
}


func (this *CommandRunnable) Run() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.RUN, e)
                }
            }()
            this.LogUtil.PutF(this.CommonStrings.START_RUNNABLE, this, this.CommonStrings.RUN)
            var canvas IMyCanvas = this.CommandFormInputProcessor.GetCanvas(); _ = canvas
            var commandListener CommandListener = canvas.GetCustomCommandListener(); _ = commandListener
            commandListener.CommandAction(this.Command, canvas)
            this.LogUtil.PutF(this.CommonStrings.END_RUNNABLE, this, this.CommonStrings.RUN)
            return
        }(); tryControl1 == 1 {
            return
        }
}

