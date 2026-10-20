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

type ExitRunnable struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    Midlet IAllBinaryMidlet
    Processor IProcessor
    IsProgress bool
}

// Implements Runnable


func (this *ExitRunnable) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
    
    
}


func NewExitRunnable(midlet IAllBinaryMidlet, processor IProcessor, isProgress bool) *ExitRunnable {
    this := &ExitRunnable{}
    this.initFields()
        this.Midlet = midlet
        this.IsProgress = isProgress
        this.Processor = processor
    return this
}


func (this *ExitRunnable) Run() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.RUN, e)
            if tryControl2 := func() (tryControl2 int) {
                defer func() {
                    if recovered := recover(); recovered != nil {
                        e2 := Java_ToError(recovered)
                        _ = e2
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.RUN, e)
                    }
                }()
                this.Processor.Process()
                return
            }(); tryControl2 == 1 {
                tryControl1 = 1; return
            }
                }
            }()
            this.LogUtil.PutF(this.CommonStrings.START_RUNNABLE, this, this.CommonStrings.RUN)
            this.Midlet.DestroyAppInRunnable(false, this.IsProgress)
            this.Midlet.SetDisplay(NullCanvas_NULL_CANVAS)
            this.Midlet.NotifyDestroyed()
            this.LogUtil.PutF(this.CommonStrings.END_RUNNABLE, this, this.CommonStrings.RUN)
            this.Processor.Process()
            return
        }(); tryControl1 == 1 {
            return
        }
}

