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

type J2SEMath struct {
}

var J2SEMath_Instance *J2SEMath = NewJ2SEMath()

func NewJ2SEMath() *J2SEMath {
    this := &J2SEMath{}
    return this
}


func J2SEMath_GetInstance() *J2SEMath {
        return J2SEMath_Instance
}


func (this *J2SEMath) Round(value float32) int {
        return int64(math.Floor(float64(value) + 0.5))
}


func (this *J2SEMath) Abs(value float32) float32 {
        return func() float32 { if (value < 0.0) { return -value }; return value }()
}


func (this *J2SEMath) Absdouble(value float64) float64 {
        return func() float64 { if (value < 0.0) { return -value }; return value }()
}

