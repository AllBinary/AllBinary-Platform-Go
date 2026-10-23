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
    PoolName string
    Priority int
    IsAlive bool
    TaskQueue IBasicArrayList
    ThreadID int
    NumThreads int
    RunningTask bool
}

// The Java class ThreadPool has subclasses so its variables are this interface
type IThreadPool interface {
    ThreadPool_() *ThreadPool
    Init()
    RunAPriorityTask()
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
    CreateName() string
}

func (this *ThreadPool) ThreadPool_() *ThreadPool {
    return this
}

var ThreadPool_ROOT_NAME string = "-PooledThread-"
var ThreadPool_NORMAL_PRIORITY int = Thread_NORM_PRIORITY

func (this *ThreadPool) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.NULL_RUNNABLE = NullRunnable_GetInstance()
    this.ThreadPoolStrings = ThreadPoolStrings_GetInstance()
    this.ThreadObjectUtil = ThreadObjectUtil_GetInstance()
    
    
    this.IsAlive = false
    this.TaskQueue = BasicArrayListUtil_GetInstance().GetImmutableInstance()
    
    
    
}


func NewThreadPool(poolName string, numThreads int, priority int) *ThreadPool {
    this := &ThreadPool{}
    this.initFields()
        this.PoolName = poolName
        this.Priority = priority
        this.NumThreads = numThreads
    return this
}


func (this *ThreadPool) Init() {
        if !this.IsAlive {
            this.IsAlive = true
            this.TaskQueue = NewBasicArrayListD()
            var pooledThread *PooledThread; _ = pooledThread
            for i := int(0); i < this.NumThreads; i++ {
                pooledThread = NewPooledThread(this)
                pooledThread.SetPriority(this.Priority)
                pooledThread.Start()
            }
        }
}


// throws Exception
func (this *ThreadPool) RunAPriorityTask() {
        panic(errors.New(fmt.Sprint()))
}


func (this *ThreadPool) RunTaskWithPriority(task PriorityRunnable) {
        { // synchronized (this)
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
                Java_Notify(this)
            }
        }
}


func (this *ThreadPool) RunTask(task Runnable) {
        { // synchronized (this)
            if !this.IsAlive {
                this.Init()
            }
            if task != nil {
                this.TaskQueue.Add(task)
                Java_Notify(this)
            }
        }
}


// throws InterruptedException
func (this *ThreadPool) GetTask() Runnable {
        { // synchronized (this)
            for this.TaskQueue.Size() == 0 {
                if !this.IsAlive {
                    return this.NULL_RUNNABLE
                }
                Java_Wait(this)
            }
            var runnable Runnable = any(this.TaskQueue.RemoveAt(0)).(Runnable); _ = runnable
            return runnable
        }
}


func (this *ThreadPool) Clear() {
        { // synchronized (this)
            if this.IsAlive {
                this.TaskQueue.Clear()
            }
        }
}


func (this *ThreadPool) Close() {
        { // synchronized (this)
            if this.IsAlive {
                this.IsAlive = false
                this.TaskQueue.Clear()
            }
        }
}


func (this *ThreadPool) Join() {
        { // synchronized (this)
            this.IsAlive = false
            this.TaskQueue.Clear()
            Java_NotifyAll(this)
        }
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
        }
}


func (this *ThreadPool) StartTask(task Runnable) {
}


func (this *ThreadPool) CompletedTask(task Runnable) {
}


func (this *ThreadPool) CreateName() string {
        return fmt.Sprint(NewStringMaker().Append(this.PoolName).Append(ThreadPool_ROOT_NAME).Appendint(Java_PostIncrement(&this.ThreadID)))
}

