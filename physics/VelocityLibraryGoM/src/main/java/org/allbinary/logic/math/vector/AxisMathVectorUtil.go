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

type AxisMathVectorUtil struct {
    NoDecimalTrigTable *NoDecimalTrigTable
}

var AxisMathVectorUtil_Instance *AxisMathVectorUtil = NewAxisMathVectorUtil()

func (this *AxisMathVectorUtil) initFields() {
    this.NoDecimalTrigTable = NoDecimalTrigTable_GetInstance()
}

func NewAxisMathVectorUtil() *AxisMathVectorUtil {
    this := &AxisMathVectorUtil{}
    this.initFields()
    return this
}


func AxisMathVectorUtil_GetInstance() *AxisMathVectorUtil {
        return AxisMathVectorUtil_Instance
}


func (this *AxisMathVectorUtil) CalculateX(magnitude int64, angle int) int64 {
        return magnitude * this.NoDecimalTrigTable.Cos(angle) / 10
}


func (this *AxisMathVectorUtil) CalculateY(magnitude int64, angle int) int64 {
        return magnitude * this.NoDecimalTrigTable.Sin(angle) / 10
}


func (this *AxisMathVectorUtil) CalculateZ(magnitude int64, otherAngle int) int64 {
        return magnitude * this.NoDecimalTrigTable.Sin(otherAngle) / 10
}

