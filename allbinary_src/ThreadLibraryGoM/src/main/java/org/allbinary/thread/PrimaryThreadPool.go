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

type PrimaryThreadPool struct {
    ThreadPool
}

var PrimaryThreadPool_THREAD_POOL IThreadPool = NewThreadPool("Primary", 1, ThreadPool_NORMAL_PRIORITY)


func PrimaryThreadPool_GetInstance() IThreadPool {
        return PrimaryThreadPool_THREAD_POOL
}


func NewPrimaryThreadPool(poolName string, numThreads int, priority int) *PrimaryThreadPool {
    this := &PrimaryThreadPool{}
        this.ThreadPool = *NewThreadPool(poolName, numThreads, priority)
    return this
}


func (this *PrimaryThreadPool) RunTask(task Runnable) {
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(StringUtil_GetInstance().ToString(task)).Appendlong(time.Now().UnixMilli())), this, this.ThreadPoolStrings.ADD_TASK)
        this.ThreadPool.RunTask(task)
}


func (this *PrimaryThreadPool) StartTask(task Runnable) {
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(this.ThreadPoolStrings.START_TASK).Append(StringUtil_GetInstance().ToString(task))), this, this.CommonStrings.RUN)
}


func (this *PrimaryThreadPool) CompletedTask(task Runnable) {
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(this.ThreadPoolStrings.COMPLETE_TASK).Append(StringUtil_GetInstance().ToString(task))), this, this.CommonStrings.RUN)
}

