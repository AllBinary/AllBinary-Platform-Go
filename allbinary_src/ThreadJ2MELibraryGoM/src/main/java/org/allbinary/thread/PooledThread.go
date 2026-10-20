/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2025 AllBinary 
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

type PooledThread struct {
    Thread
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    ThreadObjectUtil *ThreadObjectUtil
    ThreadPool IThreadPool
    INTERRUPT_EXCEPTION string
}


func (this *PooledThread) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.ThreadObjectUtil = ThreadObjectUtil_GetInstance()
    
    this.INTERRUPT_EXCEPTION = "Exit InterruptedException"
}


func NewPooledThread(threadPool IThreadPool) *PooledThread {
    this := &PooledThread{}
        this.Thread = *NewThread(threadPool.CreateName())
        this.initFields()
        this.LogUtil.PutF(this.CommonStrings.CONSTRUCTOR, this, this.CommonStrings.CONSTRUCTOR)
        this.ThreadPool = threadPool
    this.Thread.Target = this
    return this
}


func (this *PooledThread) Run() {
        this.ThreadPool.ThreadStarted()
        for true {
            var task2 Runnable = this.ThreadObjectUtil.NULL_PRIORITY_RUNNABLE; _ = task2
            if tryControl1 := func() (tryControl1 int) {
                defer func() {
                    if recovered := recover(); recovered != nil {
                        ex := Java_ToError(recovered)
                        _ = ex
                var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
                logUtil.PutF(this.INTERRUPT_EXCEPTION, this, this.CommonStrings.RUN)
                tryControl1 = 2; return
                    }
                }()
                task2 = this.ThreadPool.GetTask()
                this.ThreadPool.ThreadPool_().RunningTask = true
                this.ThreadPool.StartTask(task2)
                return
            }(); tryControl1 == 1 {
                return
            } else if tryControl1 == 2 {
                break
            }
            if task2 == this.ThreadObjectUtil.NULL_PRIORITY_RUNNABLE {
                break
            }
            if tryControl1 := func() (tryControl1 int) {
                defer func() {
                    if recovered := recover(); recovered != nil {
                        e := Java_ToError(recovered)
                        _ = e
                var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
                logUtil.Put(fmt.Sprint(NewStringMaker().Append(this.CommonStrings.EXCEPTION_LABEL).Append(StringUtil_GetInstance().ToString(task2))), this, this.CommonStrings.RUN, e)
                    }
                }()
                task2.Run()
                this.ThreadPool.CompletedTask(task2)
                this.ThreadPool.ThreadPool_().RunningTask = false
                return
            }(); tryControl1 == 1 {
                return
            }
        }
        this.ThreadPool.ThreadStopped()
}
