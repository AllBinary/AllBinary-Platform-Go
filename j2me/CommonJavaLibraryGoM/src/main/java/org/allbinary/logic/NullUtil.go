/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2025 AllBinary 
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

//ActualPlatform
type NullUtil struct {
    NULL_OBJECT any
    NULL_CLASS reflect.Type
    NULL_OBJECT_ARRAY []any
    NULL_CHAR_ARRAY []rune
    NULL_BYTE_ARRAY []int8
    NULL_SHORT_ARRAY_ARRAY [][]int16
    NULL_INT_ARRAY []int
    NULL_INT_ARRAY_ARRAY [][]int
    NULL_INT_ARRAY_ARRAY_ARRAY [][][]int
    NULL_FLOAT_ARRAY []float32
}

var NullUtil_Instance *NullUtil = NewNullUtil()

func (this *NullUtil) initFields() {
    this.NULL_OBJECT = new(struct{})
    this.NULL_CLASS = reflect.TypeOf(this.NULL_OBJECT)
    this.NULL_OBJECT_ARRAY = make([]any, 0)
    this.NULL_CHAR_ARRAY = make([]rune, 0)
    this.NULL_BYTE_ARRAY = make([]int8, 0)
    this.NULL_SHORT_ARRAY_ARRAY = func() [][]int16 { array := make([][]int16, 0); for index := range array { array[index] = make([]int16, 0) }; return array }()
    this.NULL_INT_ARRAY = make([]int, 0)
    this.NULL_INT_ARRAY_ARRAY = func() [][]int { array := make([][]int, 0); for index := range array { array[index] = make([]int, 0) }; return array }()
    this.NULL_INT_ARRAY_ARRAY_ARRAY = func() [][][]int { array := make([][][]int, 0); for index := range array { array[index] = make([][]int, 0); for index2 := range array[index] { array[index][index2] = make([]int, 0) } }; return array }()
    this.NULL_FLOAT_ARRAY = make([]float32, 0)
}

func NewNullUtil() *NullUtil {
    this := &NullUtil{}
    this.initFields()
    return this
}


func NullUtil_GetInstance() *NullUtil {
        return NullUtil_Instance
}
