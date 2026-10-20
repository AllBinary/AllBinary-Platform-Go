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

type NoZBasicConstantVelocityMovement struct {
    Movement
    VelocityProperties IBasicVelocityProperties
    SpeedBasicDecimal *BasicDecimal
    AxisMathVectorUtil *AxisMathVectorUtil
}

// Implements VelocityInterfaceCompositeInterface


func (this *NoZBasicConstantVelocityMovement) initFields() {
    
    this.SpeedBasicDecimal = BasicDecimal_ZERO_BIGDECIMAL
    this.AxisMathVectorUtil = AxisMathVectorUtil_GetInstance()
}


func NewNoZBasicConstantVelocityMovement(basicDecimal *BasicDecimal, velocityProperties IBasicVelocityProperties) *NoZBasicConstantVelocityMovement {
    this := &NoZBasicConstantVelocityMovement{}
    this.Movement = *NewMovement()
    this.initFields()
        this.SetSpeedBasicDecimal(basicDecimal)
        this.VelocityProperties = velocityProperties
    return this
}


func (this *NoZBasicConstantVelocityMovement) Init(speedBasicDecimal *BasicDecimal, angle int, otherAngle int) {
        this.SpeedBasicDecimal = speedBasicDecimal
        var angleFactory *AngleFactory = AngleFactory_GetInstance(); _ = angleFactory
        this.VelocityProperties.SetVelocityWithBigDecimal(speedBasicDecimal, angleFactory.GetAt(angle), angleFactory.GetAt(otherAngle))
}


func (this *NoZBasicConstantVelocityMovement) MoveOutsideRadius(layer IAllBinaryLayer, radius int64, angle int, otherAngle int) {
        var scaleFactorValue int = this.SpeedBasicDecimal.GetScaledFactorValue(); _ = scaleFactorValue
        var xVector int = int((this.AxisMathVectorUtil.CalculateX(radius, angle) / int64(scaleFactorValue))); _ = xVector
        var yVector int = int((this.AxisMathVectorUtil.CalculateY(radius, angle) / int64(scaleFactorValue))); _ = yVector
        layer.MoveDXYZ(xVector, yVector, 0)
}


// throws Exception
func (this *NoZBasicConstantVelocityMovement) Process(layer IAllBinaryGameLayer) {
        layer.MoveDXYZ(this.VelocityProperties.GetVelocityXBasicDecimalP().GetScaled(), this.VelocityProperties.GetVelocityYBasicDecimalP().GetScaled(), 0)
}


func (this *NoZBasicConstantVelocityMovement) ToString() string {
        return fmt.Sprint(this.VelocityProperties)
}

func (this *NoZBasicConstantVelocityMovement) String() string {
    return this.ToString()
}


func (this *NoZBasicConstantVelocityMovement) Stop() {
        this.VelocityProperties.Zero()
}


func (this *NoZBasicConstantVelocityMovement) GetVelocityProperties() IBasicVelocityProperties {
        return this.VelocityProperties
}


func (this *NoZBasicConstantVelocityMovement) SetVelocityProperties(velocityProperties IBasicVelocityProperties) {
        this.VelocityProperties = velocityProperties
}


func (this *NoZBasicConstantVelocityMovement) SetSpeedBasicDecimal(speedBasicDecimal *BasicDecimal) {
        this.SpeedBasicDecimal = speedBasicDecimal
}


func (this *NoZBasicConstantVelocityMovement) GetSpeedBasicDecimal() *BasicDecimal {
        return this.SpeedBasicDecimal
}

