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

type BasicVelocityProperties struct {
    LogUtil *LogUtil
    VelocityXBasicDecimal *BasicDecimal
    VelocityYBasicDecimal *BasicDecimal
    VelocityZBasicDecimal *BasicDecimal
    AxisMathVectorUtil *AxisMathVectorUtil
    DirectionUtil *DirectionUtil
}

// Implements BasicVelocityInterface

// The Java class BasicVelocityProperties has subclasses so its variables are this interface
type IBasicVelocityProperties interface {
    BasicVelocityProperties_() *BasicVelocityProperties
    Zero()
    GetVelocityXBasicDecimalP() *BasicDecimal
    GetVelocityYBasicDecimalP() *BasicDecimal
    GetVelocityZBasicDecimalP() *BasicDecimal
    SetVelocityWithBigDecimalAndDirection(magnitudeBasicDecimal *BasicDecimal, direction *Direction, otherDirection *Direction)
    SetVelocityWithDirection(magnitude int64, direction *Direction, otherDirection *Direction)
    AddVelocityWithBigDecimalAndDirection(magnitudeBasicDecimal *BasicDecimal, direction *Direction, otherDirection *Direction)
    AddVelocityWithDirection(magnitude int64, direction *Direction, otherDirection *Direction)
    SetVelocityWithBigDecimal(magnitudeBasicDecimal *BasicDecimal, angle IAngle, otherAngle IAngle)
    AddVelocityWithBigDecimal(magnitudeBasicDecimal *BasicDecimal, angle IAngle, otherAngle IAngle)
    SetVelocity(magnitude int64, angle IAngle, otherAngle IAngle)
    AddVelocity(magnitude int64, angle IAngle, otherAngle IAngle)
    SetVelocityi(magnitude int64, angle int, otherAngle int)
    AddVelocityi(magnitude int64, angle int, otherAngle int)
    ToString() string
}

func (this *BasicVelocityProperties) BasicVelocityProperties_() *BasicVelocityProperties {
    return this
}


func (this *BasicVelocityProperties) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
    
    this.AxisMathVectorUtil = AxisMathVectorUtil_GetInstance()
    this.DirectionUtil = DirectionUtil_GetInstance()
}


func NewBasicVelocityProperties() *BasicVelocityProperties {
    this := &BasicVelocityProperties{}
    this.initFields()
        this.VelocityXBasicDecimal = NewBasicDecimal(0)
        this.VelocityYBasicDecimal = NewBasicDecimal(0)
        this.VelocityZBasicDecimal = NewBasicDecimal(0)
    return this
}


func (this *BasicVelocityProperties) Zero() {
        this.VelocityXBasicDecimal.Setint(0)
        this.VelocityYBasicDecimal.Setint(0)
        this.VelocityZBasicDecimal.Setint(0)
}


func (this *BasicVelocityProperties) GetVelocityXBasicDecimalP() *BasicDecimal {
        return this.VelocityXBasicDecimal
}


func (this *BasicVelocityProperties) GetVelocityYBasicDecimalP() *BasicDecimal {
        return this.VelocityYBasicDecimal
}


func (this *BasicVelocityProperties) GetVelocityZBasicDecimalP() *BasicDecimal {
        return this.VelocityZBasicDecimal
}


func (this *BasicVelocityProperties) SetVelocityWithBigDecimalAndDirection(magnitudeBasicDecimal *BasicDecimal, direction *Direction, otherDirection *Direction) {
        var angle IAngle = this.DirectionUtil.GetAngle(direction); _ = angle
        var otherAngle IAngle = this.DirectionUtil.GetAngle(otherDirection); _ = otherAngle
        this.SetVelocityWithBigDecimal(magnitudeBasicDecimal, angle, otherAngle)
}


func (this *BasicVelocityProperties) SetVelocityWithDirection(magnitude int64, direction *Direction, otherDirection *Direction) {
        var angle IAngle = this.DirectionUtil.GetAngle(direction); _ = angle
        var otherAngle IAngle = this.DirectionUtil.GetAngle(otherDirection); _ = otherAngle
        this.SetVelocity(magnitude, angle, otherAngle)
}


func (this *BasicVelocityProperties) AddVelocityWithBigDecimalAndDirection(magnitudeBasicDecimal *BasicDecimal, direction *Direction, otherDirection *Direction) {
        var angle IAngle = this.DirectionUtil.GetAngle(direction); _ = angle
        var otherAngle IAngle = this.DirectionUtil.GetAngle(otherDirection); _ = otherAngle
        this.AddVelocityWithBigDecimal(magnitudeBasicDecimal, angle, otherAngle)
}


func (this *BasicVelocityProperties) AddVelocityWithDirection(magnitude int64, direction *Direction, otherDirection *Direction) {
        var angle IAngle = this.DirectionUtil.GetAngle(direction); _ = angle
        var otherAngle IAngle = this.DirectionUtil.GetAngle(otherDirection); _ = otherAngle
        this.AddVelocity(magnitude, angle, otherAngle)
}


func (this *BasicVelocityProperties) SetVelocityWithBigDecimal(magnitudeBasicDecimal *BasicDecimal, angle IAngle, otherAngle IAngle) {
        var magnitude int64 = magnitudeBasicDecimal.GetUnscaled(); _ = magnitude
        this.SetVelocity(magnitude, angle, otherAngle)
}


func (this *BasicVelocityProperties) AddVelocityWithBigDecimal(magnitudeBasicDecimal *BasicDecimal, angle IAngle, otherAngle IAngle) {
        var magnitude int64 = magnitudeBasicDecimal.GetUnscaled(); _ = magnitude
        this.AddVelocity(magnitude, angle, otherAngle)
}


func (this *BasicVelocityProperties) SetVelocity(magnitude int64, angle IAngle, otherAngle IAngle) {
        this.SetVelocityi(magnitude, int(angle.GetValue()), int(otherAngle.GetValue()))
}


func (this *BasicVelocityProperties) AddVelocity(magnitude int64, angle IAngle, otherAngle IAngle) {
        this.AddVelocityi(magnitude, int(angle.GetValue()), int(otherAngle.GetValue()))
}


func (this *BasicVelocityProperties) SetVelocityi(magnitude int64, angle int, otherAngle int) {
        var xVector int64 = (this.AxisMathVectorUtil.CalculateX(magnitude, angle) / int64(this.VelocityXBasicDecimal.GetScaledFactorValue())); _ = xVector
        var yVector int64 = (this.AxisMathVectorUtil.CalculateY(magnitude, angle) / int64(this.VelocityYBasicDecimal.GetScaledFactorValue())); _ = yVector
        var zVector int64 = (this.AxisMathVectorUtil.CalculateZ(magnitude, otherAngle) / int64(this.VelocityZBasicDecimal.GetScaledFactorValue())); _ = zVector
        this.VelocityXBasicDecimal.Setlong(xVector)
        this.VelocityYBasicDecimal.Setlong(yVector)
        this.VelocityZBasicDecimal.Setlong(zVector)
}


func (this *BasicVelocityProperties) AddVelocityi(magnitude int64, angle int, otherAngle int) {
        var xVector int64 = (this.AxisMathVectorUtil.CalculateX(magnitude, angle) / int64(this.VelocityXBasicDecimal.GetScaledFactorValue())); _ = xVector
        var yVector int64 = (this.AxisMathVectorUtil.CalculateY(magnitude, angle) / int64(this.VelocityYBasicDecimal.GetScaledFactorValue())); _ = yVector
        var zVector int64 = (this.AxisMathVectorUtil.CalculateZ(magnitude, otherAngle) / int64(this.VelocityZBasicDecimal.GetScaledFactorValue())); _ = zVector
        this.VelocityXBasicDecimal.Addlong(xVector)
        this.VelocityYBasicDecimal.Addlong(yVector)
        this.VelocityZBasicDecimal.Addlong(zVector)
}


func (this *BasicVelocityProperties) ToString() string {
        var commonSeps *CommonSeps = CommonSeps_GetInstance(); _ = commonSeps
        var positionStrings *PositionStrings = PositionStrings_GetInstance(); _ = positionStrings
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        var stringUtil *StringUtil = StringUtil_GetInstance(); _ = stringUtil
        stringBuffer.Append("Velocity ")
        stringBuffer.Append(positionStrings.X_LABEL)
        stringBuffer.Append(stringUtil.ToString(this.VelocityXBasicDecimal))
        stringBuffer.Append(commonSeps.SPACE)
        stringBuffer.Append(positionStrings.Y_LABEL)
        stringBuffer.Append(stringUtil.ToString(this.VelocityYBasicDecimal))
        stringBuffer.Append(commonSeps.SPACE)
        stringBuffer.Append(positionStrings.Z_LABEL)
        stringBuffer.Append(stringUtil.ToString(this.VelocityZBasicDecimal))
        return fmt.Sprint(stringBuffer)
}

func (this *BasicVelocityProperties) String() string {
    return this.ToString()
}

