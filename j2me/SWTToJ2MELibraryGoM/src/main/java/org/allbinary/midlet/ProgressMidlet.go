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

//SWTForJ2ME
type ProgressMidlet struct {
    AllBinaryMidlet
    LogUtil *LogUtil
    AbeClientInformation AbeClientInformationInterface
}

// The Java class ProgressMidlet has subclasses so its variables are this interface
type IProgressMidlet interface {
    IAllBinaryMidlet
    ProgressMidlet_() *ProgressMidlet
    PreInit()
    ExitProgress(isProgress bool)
}

func (this *ProgressMidlet) ProgressMidlet_() *ProgressMidlet {
    return this
}


func (this *ProgressMidlet) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
}


func NewProgressMidlet(clientInformationFactory *ClientInformationFactory) *ProgressMidlet {
    this := &ProgressMidlet{}
    this.AllBinaryMidlet = *NewAllBinaryMidlet()
    this.initFields()
        this.PreInit()
        this.AbeClientInformation = clientInformationFactory.GetInstance()
    return this
}


func (this *ProgressMidlet) PreInit() {
}


func (this *ProgressMidlet) ExitProgress(isProgress bool) {
        var processor IProcessor = MidletExitProcessorFactory_GetInstance().GetExitInstance(this); _ = processor
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "exit", e)
            if tryControl2 := func() (tryControl2 int) {
                defer func() {
                    if recovered := recover(); recovered != nil {
                        e2 := Java_ToError(recovered)
                        _ = e2
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "exit", e)
                    }
                }()
                processor.Process()
                return
            }(); tryControl2 == 1 {
                tryControl1 = 1; return
            }
                }
            }()
            var primaryThreadPool IThreadPool = PrimaryThreadPool_GetInstance(); _ = primaryThreadPool
            var runnable Runnable = NewExitRunnable(this, processor, isProgress); _ = runnable
            primaryThreadPool.RunTask(runnable)
            return
        }(); tryControl1 == 1 {
            return
        }
}

