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

type JvmStatistics struct {
    Runtime *Runtime
}

// The Java class JvmStatistics has subclasses so its variables are this interface
type IJvmStatistics interface {
    JvmStatistics_() *JvmStatistics
    GetFreeMemoryBytes() int64
    GetFreeMemoryKilobytes() int64
    GetFreeMemoryMegabytes() int64
    GetMaxMemoryBytes() int64
    GetMaxMemoryKilobytes() int64
    GetMaxMemoryMegabytes() int64
    GetTotalMemoryBytes() int64
    GetTotalMemoryKilobytes() int64
    GetTotalMemoryMegabytes() int64
    GetFreeMemoryBytesLong() int64
    GetFreeMemoryKilobytesLong() int64
    GetFreeMemoryMegabytesLong() int64
    GetMaxMemoryBytesLong() int64
    GetMaxMemoryKilobytesLong() int64
    GetMaxMemoryMegabytesLong() int64
    GetTotalMemoryBytesLong() int64
    GetTotalMemoryKilobytesLong() int64
    GetTotalMemoryMegabytesLong() int64
    GetFreeMemoryBytesString() string
    GetFreeMemoryKilobytesString() string
    GetFreeMemoryMegabytesString() string
    GetMaxMemoryBytesString() string
    GetMaxMemoryKilobytesString() string
    GetMaxMemoryMegabytesString() string
    GetTotalMemoryBytesString() string
    GetTotalMemoryKilobytesString() string
    GetTotalMemoryMegabytesString() string
    GetNumberOfThreads() int64
    GetNumberOfThreadsLong() int64
    GetNumberOfThreadsString() string
    CollectGarbage()
}

func (this *JvmStatistics) JvmStatistics_() *JvmStatistics {
    return this
}


func (this *JvmStatistics) initFields() {
    
}


func NewJvmStatistics() *JvmStatistics {
    this := &JvmStatistics{}
    this.initFields()
        this.Runtime = Runtime_GetRuntime()
    return this
}


func (this *JvmStatistics) GetFreeMemoryBytes() int64 {
        return this.Runtime.FreeMemory()
}


func (this *JvmStatistics) GetFreeMemoryKilobytes() int64 {
        return this.GetFreeMemoryBytes() / 1024
}


func (this *JvmStatistics) GetFreeMemoryMegabytes() int64 {
        return this.GetFreeMemoryKilobytes() / 1024
}


func (this *JvmStatistics) GetMaxMemoryBytes() int64 {
        return this.Runtime.MaxMemory()
}


func (this *JvmStatistics) GetMaxMemoryKilobytes() int64 {
        return this.GetMaxMemoryBytes() / 1024
}


func (this *JvmStatistics) GetMaxMemoryMegabytes() int64 {
        return this.GetMaxMemoryKilobytes() / 1024
}


func (this *JvmStatistics) GetTotalMemoryBytes() int64 {
        return this.Runtime.TotalMemory()
}


func (this *JvmStatistics) GetTotalMemoryKilobytes() int64 {
        return this.GetTotalMemoryBytes() / 1024
}


func (this *JvmStatistics) GetTotalMemoryMegabytes() int64 {
        return this.GetTotalMemoryKilobytes() / 1024
}


func (this *JvmStatistics) GetFreeMemoryBytesLong() int64 {
        return int64(this.GetFreeMemoryBytes())
}


func (this *JvmStatistics) GetFreeMemoryKilobytesLong() int64 {
        return int64(this.GetFreeMemoryKilobytes())
}


func (this *JvmStatistics) GetFreeMemoryMegabytesLong() int64 {
        return int64(this.GetFreeMemoryMegabytes())
}


func (this *JvmStatistics) GetMaxMemoryBytesLong() int64 {
        return int64(this.GetMaxMemoryBytes())
}


func (this *JvmStatistics) GetMaxMemoryKilobytesLong() int64 {
        return int64(this.GetMaxMemoryKilobytes())
}


func (this *JvmStatistics) GetMaxMemoryMegabytesLong() int64 {
        return int64(this.GetMaxMemoryMegabytes())
}


func (this *JvmStatistics) GetTotalMemoryBytesLong() int64 {
        return int64(this.GetTotalMemoryBytes())
}


func (this *JvmStatistics) GetTotalMemoryKilobytesLong() int64 {
        return int64(this.GetTotalMemoryKilobytes())
}


func (this *JvmStatistics) GetTotalMemoryMegabytesLong() int64 {
        return int64(this.GetTotalMemoryMegabytes())
}


func (this *JvmStatistics) GetFreeMemoryBytesString() string {
        return fmt.Sprint(this.GetFreeMemoryBytesLong())
}


func (this *JvmStatistics) GetFreeMemoryKilobytesString() string {
        return fmt.Sprint(this.GetFreeMemoryKilobytesLong())
}


func (this *JvmStatistics) GetFreeMemoryMegabytesString() string {
        return fmt.Sprint(this.GetFreeMemoryMegabytesLong())
}


func (this *JvmStatistics) GetMaxMemoryBytesString() string {
        return fmt.Sprint(this.GetMaxMemoryBytesLong())
}


func (this *JvmStatistics) GetMaxMemoryKilobytesString() string {
        return fmt.Sprint(this.GetMaxMemoryKilobytesLong())
}


func (this *JvmStatistics) GetMaxMemoryMegabytesString() string {
        return fmt.Sprint(this.GetMaxMemoryMegabytesLong())
}


func (this *JvmStatistics) GetTotalMemoryBytesString() string {
        return fmt.Sprint(this.GetTotalMemoryBytesLong())
}


func (this *JvmStatistics) GetTotalMemoryKilobytesString() string {
        return fmt.Sprint(this.GetTotalMemoryKilobytesLong())
}


func (this *JvmStatistics) GetTotalMemoryMegabytesString() string {
        return fmt.Sprint(this.GetTotalMemoryMegabytesLong())
}


func (this *JvmStatistics) GetNumberOfThreads() int64 {
        return Thread_ActiveCount()
}


func (this *JvmStatistics) GetNumberOfThreadsLong() int64 {
        return int64(this.GetNumberOfThreads())
}


func (this *JvmStatistics) GetNumberOfThreadsString() string {
        return fmt.Sprint(this.GetNumberOfThreadsLong())
}


func (this *JvmStatistics) CollectGarbage() {
        this.Runtime.Gc()
}

