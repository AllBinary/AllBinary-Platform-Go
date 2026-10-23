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

//This is mainly for network communications.
type PathFindingThreadPool struct {
    ThreadPool
}

var PathFindingThreadPool_Instance IThreadPool = NewPathFindingThreadPool("PathFinding", 1, 1)


func PathFindingThreadPool_GetInstance() IThreadPool {
        return PathFindingThreadPool_Instance
}


func NewPathFindingThreadPool(poolName string, numThreads int, priority int) *PathFindingThreadPool {
    this := &PathFindingThreadPool{}
        this.ThreadPool = *NewThreadPool(poolName, numThreads, ThreadPool_NORMAL_PRIORITY)
    return this
}


func (this *PathFindingThreadPool) RunTaskWithPriority(task PriorityRunnable) {
        this.ThreadPool.RunTaskWithPriority(task)
}


func (this *PathFindingThreadPool) RunTask(task Runnable) {
        this.ThreadPool.RunTask(task)
}


func (this *PathFindingThreadPool) StartTask(task Runnable) {
}


func (this *PathFindingThreadPool) CompletedTask(task Runnable) {
}

