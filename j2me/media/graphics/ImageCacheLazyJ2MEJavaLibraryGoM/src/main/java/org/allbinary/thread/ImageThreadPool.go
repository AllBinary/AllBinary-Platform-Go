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
type ImageThreadPool struct {
    ThreadPool
    LogUtil *LogUtil
}

var ImageThreadPool_Instance IThreadPool = NewImageThreadPool("ImageLoader", 1, ThreadPool_NORMAL_PRIORITY)

func (this *ImageThreadPool) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func ImageThreadPool_GetInstance() IThreadPool {
        return ImageThreadPool_Instance
}


func NewImageThreadPool(poolName string, numThreads int, priority int) *ImageThreadPool {
    this := &ImageThreadPool{}
        this.ThreadPool = *NewThreadPool(poolName, numThreads, priority)
        this.initFields()
    return this
}


// synchronized
func (this *ImageThreadPool) RunTask(task Runnable) {
        this.ThreadPool.RunTask(task)
}


func (this *ImageThreadPool) StartTask(task Runnable) {
}


func (this *ImageThreadPool) CompletedTask(task Runnable) {
}

