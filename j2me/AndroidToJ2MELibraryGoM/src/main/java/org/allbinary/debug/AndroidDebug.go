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

type AndroidDebug struct {
    StartTime int64
    Running bool
    BufferSize int
}

// Implements DebugInterface


func (this *AndroidDebug) initFields() {
    this.StartTime = Long_MAX_VALUE
    this.Running = false
    this.BufferSize = 32 * 1024 * 1024
}

func NewAndroidDebug() *AndroidDebug {
    this := &AndroidDebug{}
    this.initFields()
    return this
}


func (this *AndroidDebug) Start() {
        this.SetStartTime(time.Now().UnixMilli())
        Debug_StartMethodTracing("trace", this.BufferSize)
        this.SetRunning(true)
}


func (this *AndroidDebug) Stop() {
        this.StartTime = Long_MAX_VALUE
        Debug_StopMethodTracing()
        this.SetRunning(false)
}


func (this *AndroidDebug) GetMaxTime() int64 {
        return 120000
}


func (this *AndroidDebug) SetStartTime(startTime int64) {
        this.StartTime = startTime
}


func (this *AndroidDebug) GetStartTime() int64 {
        return this.StartTime
}


func (this *AndroidDebug) SetRunning(running bool) {
        this.Running = running
}


func (this *AndroidDebug) IsRunning() bool {
        return this.Running
}

