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

type BasicConstantVelocityMovement struct {
    Movement
    VelocityProperties IBasicVelocityProperties
    SpeedBasicDecimal *BasicDecimal
    AxisMathVectorUtil *AxisMathVectorUtil
}

// Implements VelocityInterfaceCompositeInterface

// The Java class BasicConstantVelocityMovement has subclasses so its variables are this interface
type IBasicConstantVelocityMovement interface {
    IMovement
    BasicConstantVelocityMovement_() *BasicConstantVelocityMovement
    Init(speedBasicDecimal *BasicDecimal, angle int, otherAngle int)
    MoveOutsideRadius(layer IAllBinaryLayer, radius int64, angle int, otherAngle int)
    Process(layer IAllBinaryGameLayer)
    ToString() string
    Stop()
    GetVelocityProperties() IBasicVelocityProperties
    SetVelocityProperties(velocityProperties IBasicVelocityProperties)
    SetSpeedBasicDecimal(speedBasicDecimal *BasicDecimal)
    GetSpeedBasicDecimal() *BasicDecimal
}

func (this *BasicConstantVelocityMovement) BasicConstantVelocityMovement_() *BasicConstantVelocityMovement {
    return this
}


func (this *BasicConstantVelocityMovement) initFields() {
    
    this.SpeedBasicDecimal = BasicDecimal_ZERO_BIGDECIMAL
    this.AxisMathVectorUtil = AxisMathVectorUtil_GetInstance()
}


func NewBasicConstantVelocityMovement(basicDecimal *BasicDecimal, velocityProperties IBasicVelocityProperties) *BasicConstantVelocityMovement {
    this := &BasicConstantVelocityMovement{}
    this.Movement = *NewMovement()
    this.initFields()
        this.SetSpeedBasicDecimal(basicDecimal)
        this.VelocityProperties = velocityProperties
    return this
}


func (this *BasicConstantVelocityMovement) Init(speedBasicDecimal *BasicDecimal, angle int, otherAngle int) {
        this.SpeedBasicDecimal = speedBasicDecimal
        var angleFactory *AngleFactory = AngleFactory_GetInstance(); _ = angleFactory
        this.VelocityProperties.SetVelocityWithBigDecimal(speedBasicDecimal, angleFactory.GetAt(angle), angleFactory.GetAt(otherAngle))
}


func (this *BasicConstantVelocityMovement) MoveOutsideRadius(layer IAllBinaryLayer, radius int64, angle int, otherAngle int) {
        var scaleFactorValue int = this.SpeedBasicDecimal.GetScaledFactorValue(); _ = scaleFactorValue
        var xVector int = int((this.AxisMathVectorUtil.CalculateX(radius, angle) / int64(scaleFactorValue))); _ = xVector
        var yVector int = int((this.AxisMathVectorUtil.CalculateY(radius, angle) / int64(scaleFactorValue))); _ = yVector
        var zVector int = int((this.AxisMathVectorUtil.CalculateZ(radius, otherAngle) / int64(scaleFactorValue))); _ = zVector
        layer.MoveDXYZ(xVector, yVector, zVector)
}


// throws Exception
func (this *BasicConstantVelocityMovement) Process(layer IAllBinaryGameLayer) {
        layer.MoveDXYZ(this.VelocityProperties.GetVelocityXBasicDecimalP().GetScaled(), this.VelocityProperties.GetVelocityYBasicDecimalP().GetScaled(), this.VelocityProperties.GetVelocityZBasicDecimalP().GetScaled())
}


func (this *BasicConstantVelocityMovement) ToString() string {
        return fmt.Sprint(this.VelocityProperties)
}

func (this *BasicConstantVelocityMovement) String() string {
    return this.ToString()
}


func (this *BasicConstantVelocityMovement) Stop() {
        this.VelocityProperties.Zero()
}


func (this *BasicConstantVelocityMovement) GetVelocityProperties() IBasicVelocityProperties {
        return this.VelocityProperties
}


func (this *BasicConstantVelocityMovement) SetVelocityProperties(velocityProperties IBasicVelocityProperties) {
        this.VelocityProperties = velocityProperties
}


func (this *BasicConstantVelocityMovement) SetSpeedBasicDecimal(speedBasicDecimal *BasicDecimal) {
        this.SpeedBasicDecimal = speedBasicDecimal
}


func (this *BasicConstantVelocityMovement) GetSpeedBasicDecimal() *BasicDecimal {
        return this.SpeedBasicDecimal
}

