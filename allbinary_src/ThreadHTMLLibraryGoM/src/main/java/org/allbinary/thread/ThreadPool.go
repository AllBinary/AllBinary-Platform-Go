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

type ThreadPool struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    NULL_RUNNABLE *NullRunnable
    ThreadPoolStrings *ThreadPoolStrings
    ThreadObjectUtil *ThreadObjectUtil
    IsAlive bool
    TaskQueue IBasicArrayList
    NumThreads int
    RunningTask bool
    CurrentPriorityRunnable PriorityRunnable
}

// The Java class ThreadPool has subclasses so its variables are this interface
type IThreadPool interface {
    ThreadPool_() *ThreadPool
    RunAPriorityTask()
    RunATask()
    Init()
    RunTaskWithPriority(task PriorityRunnable)
    RunTask(task Runnable)
    GetTask() Runnable
    Clear()
    Close()
    Join()
    IsBusy() bool
    ThreadStarted()
    ThreadStopped()
    StartTask(task Runnable)
    CompletedTask(task Runnable)
}

func (this *ThreadPool) ThreadPool_() *ThreadPool {
    return this
}

var ThreadPool_NORMAL_PRIORITY int = 5

func (this *ThreadPool) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.NULL_RUNNABLE = NullRunnable_GetInstance()
    this.ThreadPoolStrings = ThreadPoolStrings_GetInstance()
    this.ThreadObjectUtil = ThreadObjectUtil_GetInstance()
    
    this.TaskQueue = NewBasicArrayListD()
    
    
    this.CurrentPriorityRunnable = this.ThreadObjectUtil.NULL_PRIORITY_RUNNABLE
}


func NewThreadPool(poolName string, numThreads int, priority int) *ThreadPool {
    this := &ThreadPool{}
    this.initFields()
    return this
}


// throws Exception
func (this *ThreadPool) RunAPriorityTask() {
        if !this.CurrentPriorityRunnable.IsDone() {
            this.CurrentPriorityRunnable.Run()
        } else {
            var runnable Runnable = this.GetTask(); _ = runnable
            if runnable == this.NULL_RUNNABLE {
                return
            }
            this.CurrentPriorityRunnable = any(runnable).(PriorityRunnable)
            if !(this.CurrentPriorityRunnable == this.ThreadObjectUtil.NULL_PRIORITY_RUNNABLE) {
                this.CurrentPriorityRunnable.Reset()
                this.CurrentPriorityRunnable.Run()
            }
        }
}


// throws Exception
func (this *ThreadPool) RunATask() {
        var runnable Runnable = this.GetTask(); _ = runnable
        if runnable != nil {
            runnable.Run()
        }
}


func (this *ThreadPool) Init() {
        if !this.IsAlive {
            this.IsAlive = true
            this.TaskQueue = NewBasicArrayListD()
        }
}


// synchronized
func (this *ThreadPool) RunTaskWithPriority(task PriorityRunnable) {
        if !this.IsAlive {
            this.Init()
        }
        if task != nil {
            var size int = this.TaskQueue.Size(); _ = size
            var runnable PriorityRunnable; _ = runnable
            var lowerPriorityRunnable PriorityRunnable = this.ThreadObjectUtil.NULL_PRIORITY_RUNNABLE; _ = lowerPriorityRunnable
            for index := int(0); index < size; index++ {
                runnable = any(this.TaskQueue.Get(index)).(PriorityRunnable)
                if runnable.GetPriority() > task.GetPriority() {
                    lowerPriorityRunnable = runnable
                    break
                }
            }
            if lowerPriorityRunnable == this.ThreadObjectUtil.NULL_PRIORITY_RUNNABLE {
                this.TaskQueue.Add(task)
            } else {
                var index int = this.TaskQueue.IndexOf(lowerPriorityRunnable); _ = index
                this.TaskQueue.AddAt(index, task)
            }
        }
}


// synchronized
func (this *ThreadPool) RunTask(task Runnable) {
        if !this.IsAlive {
            this.Init()
        }
        if task != nil {
            this.TaskQueue.Add(task)
        }
}


// throws InterruptedException
// synchronized
func (this *ThreadPool) GetTask() Runnable {
        if this.TaskQueue.IsEmpty() {
            return this.NULL_RUNNABLE
        }
        return any(this.TaskQueue.RemoveAt(0)).(Runnable)
}


// synchronized
func (this *ThreadPool) Clear() {
        this.TaskQueue.Clear()
}


// synchronized
func (this *ThreadPool) Close() {
        if this.IsAlive {
            this.IsAlive = false
            this.TaskQueue.Clear()
            this.CurrentPriorityRunnable = this.ThreadObjectUtil.NULL_PRIORITY_RUNNABLE
        }
}


func (this *ThreadPool) Join() {
        this.IsAlive = false
        this.TaskQueue.Clear()
        this.CurrentPriorityRunnable = this.ThreadObjectUtil.NULL_PRIORITY_RUNNABLE
}


func (this *ThreadPool) IsBusy() bool {
        if !this.IsAlive {
            return false
        }
        if this.TaskQueue.Size() > 0 {
            return true
        }
        if this.RunningTask {
            return true
        }
        return false
}


func (this *ThreadPool) ThreadStarted() {
}


func (this *ThreadPool) ThreadStopped() {
        if this.NumThreads == 1 {
            this.IsAlive = false
            this.TaskQueue.Clear()
            this.CurrentPriorityRunnable = this.ThreadObjectUtil.NULL_PRIORITY_RUNNABLE
        }
}


func (this *ThreadPool) StartTask(task Runnable) {
}


func (this *ThreadPool) CompletedTask(task Runnable) {
}

