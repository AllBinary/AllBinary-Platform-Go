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

type BasicVelocityInterface interface {
    GetVelocityXBasicDecimalP() *BasicDecimal
    GetVelocityYBasicDecimalP() *BasicDecimal
    Zero()
    SetVelocityWithBigDecimalAndDirection(magnitudeBasicDecimal *BasicDecimal, direction *Direction, otherDirection *Direction)
    AddVelocityWithBigDecimalAndDirection(magnitudeBasicDecimal *BasicDecimal, direction *Direction, otherDirection *Direction)
    SetVelocityWithDirection(magnitude int64, direction *Direction, otherDirection *Direction)
    AddVelocityWithDirection(magnitude int64, direction *Direction, otherDirection *Direction)
    SetVelocityWithBigDecimal(magnitudeBasicDecimal *BasicDecimal, angle IAngle, otherAngle IAngle)
    AddVelocityWithBigDecimal(magnitudeBasicDecimal *BasicDecimal, angle IAngle, otherAngle IAngle)
    SetVelocity(magnitude int64, angle IAngle, otherAngle IAngle)
    AddVelocity(magnitude int64, angle IAngle, otherAngle IAngle)
    AddVelocityi(magnitude int64, angle int, otherAngle int)
}


