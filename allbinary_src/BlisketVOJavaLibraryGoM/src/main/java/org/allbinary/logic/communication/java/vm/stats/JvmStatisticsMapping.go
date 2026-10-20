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

type JvmStatisticsMapping struct {
    JvmStatistics
}

// Implements TableMappingInterface

// The Java class JvmStatisticsMapping has subclasses so its variables are this interface
type IJvmStatisticsMapping interface {
    IJvmStatistics
    JvmStatisticsMapping_() *JvmStatisticsMapping
    ToHashMap() *HashMap
    GetKey() any
    ToVector() IBasicArrayList
}

func (this *JvmStatisticsMapping) JvmStatisticsMapping_() *JvmStatisticsMapping {
    return this
}



func NewJvmStatisticsMapping() *JvmStatisticsMapping {
    this := &JvmStatisticsMapping{}
        this.JvmStatistics = *NewJvmStatistics()
    return this
}


func (this *JvmStatisticsMapping) ToHashMap() *HashMap {
        var hashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = hashMap
        hashMap.Put(JvmStatisticsData_FREE_MEMORY_BYTES, this.GetFreeMemoryBytesString())
        hashMap.Put(JvmStatisticsData_FREE_MEMORY_KILOBYTES, this.GetFreeMemoryKilobytesString())
        hashMap.Put(JvmStatisticsData_FREE_MEMORY_MEGABYTES, this.GetFreeMemoryMegabytesString())
        hashMap.Put(JvmStatisticsData_MAX_MEMORY_BYTES, this.GetMaxMemoryBytesString())
        hashMap.Put(JvmStatisticsData_MAX_MEMORY_KILOBYTES, this.GetMaxMemoryKilobytesString())
        hashMap.Put(JvmStatisticsData_MAX_MEMORY_MEGABYTES, this.GetMaxMemoryMegabytesString())
        hashMap.Put(JvmStatisticsData_NUMBER_OF_THREADS, this.GetNumberOfThreadsString())
        hashMap.Put(JvmStatisticsData_TOTAL_MEMORY_BYTES, this.GetTotalMemoryBytesString())
        hashMap.Put(JvmStatisticsData_TOTAL_MEMORY_KILOBYTES, this.GetTotalMemoryKilobytesString())
        hashMap.Put(JvmStatisticsData_TOTAL_MEMORY_MEGABYTES, this.GetTotalMemoryMegabytesString())
        return hashMap
}


// throws Exception
func (this *JvmStatisticsMapping) GetKey() any {
        return JvmStatisticsData_NAME
}


// throws Exception
func (this *JvmStatisticsMapping) ToVector() IBasicArrayList {
        var vector IBasicArrayList = NewBasicArrayListD(); _ = vector
        vector.Add(this.GetFreeMemoryBytesString())
        vector.Add(this.GetFreeMemoryKilobytesString())
        vector.Add(this.GetFreeMemoryMegabytesString())
        vector.Add(this.GetMaxMemoryBytesString())
        vector.Add(this.GetMaxMemoryKilobytesString())
        vector.Add(this.GetMaxMemoryMegabytesString())
        vector.Add(this.GetNumberOfThreadsString())
        vector.Add(this.GetTotalMemoryBytesString())
        vector.Add(this.GetTotalMemoryKilobytesString())
        vector.Add(this.GetTotalMemoryMegabytesString())
        return vector
}

