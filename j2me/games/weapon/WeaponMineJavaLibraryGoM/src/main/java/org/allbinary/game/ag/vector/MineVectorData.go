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

type MineVectorData struct {
}

var MineVectorData_Points [][][]int = [][][]int{[][]int{[]int{10, 5}, []int{2, 6}, []int{1000, 1000}, []int{2, 2}, []int{10, 9}, []int{1000, 1000}, []int{10, 2}, []int{2, 9}, []int{1000, 1000}, []int{2, 5}, []int{10, 6}, []int{1000, 1000}, []int{7, 9}, []int{5, 2}, []int{1000, 1000}, []int{7, 2}, []int{5, 9}, []int{1000, 1000}}}


func NewMineVectorData() *MineVectorData {
    this := &MineVectorData{}
    return this
}

