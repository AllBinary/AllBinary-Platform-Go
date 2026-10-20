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

package ag

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
    . "org/allbinary"
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
var _ = NullUtil_GetInstance

type HealthVectorData struct {
}

var HealthVectorData_Points [][][]int = [][][]int{[][]int{[]int{2, 1}, []int{3, 2}, []int{3, 2}, []int{4, 1}, []int{4, 1}, []int{5, 1}, []int{5, 1}, []int{6, 2}, []int{6, 2}, []int{6, 3}, []int{6, 3}, []int{3, 6}, []int{3, 6}, []int{0, 3}, []int{0, 3}, []int{0, 2}, []int{0, 2}, []int{1, 1}, []int{1, 1}, []int{2, 1}, []int{1000, 1000}}}

func NewHealthVectorData() *HealthVectorData {
    this := &HealthVectorData{}
    return this
}

