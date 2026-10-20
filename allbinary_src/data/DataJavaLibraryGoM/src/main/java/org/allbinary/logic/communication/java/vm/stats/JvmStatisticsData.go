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

type JvmStatisticsData struct {
}

var JvmStatisticsData_JVM_STATISTICS string = "JVM_STATISTICS_"
var JvmStatisticsData_NAME string = fmt.Sprint(JvmStatisticsData_JVM_STATISTICS) + "NAME"
var JvmStatisticsData_TOTAL_MEMORY_BYTES string = fmt.Sprint(JvmStatisticsData_JVM_STATISTICS) + "TOTAL_MEMORY_BYTES"
var JvmStatisticsData_TOTAL_MEMORY_KILOBYTES string = fmt.Sprint(JvmStatisticsData_JVM_STATISTICS) + "TOTAL_MEMORY_KILOBYTES"
var JvmStatisticsData_TOTAL_MEMORY_MEGABYTES string = fmt.Sprint(JvmStatisticsData_JVM_STATISTICS) + "TOTAL_MEMORY_MEGABYTES"
var JvmStatisticsData_FREE_MEMORY_BYTES string = fmt.Sprint(JvmStatisticsData_JVM_STATISTICS) + "FREE_MEMORY_BYTES"
var JvmStatisticsData_FREE_MEMORY_KILOBYTES string = fmt.Sprint(JvmStatisticsData_JVM_STATISTICS) + "FREE_MEMORY_KILOBYTES"
var JvmStatisticsData_FREE_MEMORY_MEGABYTES string = fmt.Sprint(JvmStatisticsData_JVM_STATISTICS) + "FREE_MEMORY_MEGABYTES"
var JvmStatisticsData_MAX_MEMORY_BYTES string = fmt.Sprint(JvmStatisticsData_JVM_STATISTICS) + "MAX_MEMORY_BYTES"
var JvmStatisticsData_MAX_MEMORY_KILOBYTES string = fmt.Sprint(JvmStatisticsData_JVM_STATISTICS) + "MAX_MEMORY_KILOBYTES"
var JvmStatisticsData_MAX_MEMORY_MEGABYTES string = fmt.Sprint(JvmStatisticsData_JVM_STATISTICS) + "MAX_MEMORY_MEGABYTES"
var JvmStatisticsData_NUMBER_OF_THREADS string = fmt.Sprint(JvmStatisticsData_JVM_STATISTICS) + "NUMBER_OF_THREADS"


func NewJvmStatisticsData() *JvmStatisticsData {
    this := &JvmStatisticsData{}
        
    return this
}

