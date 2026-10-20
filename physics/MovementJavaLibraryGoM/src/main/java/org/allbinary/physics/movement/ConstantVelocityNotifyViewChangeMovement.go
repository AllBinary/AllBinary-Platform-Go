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

type ConstantVelocityNotifyViewChangeMovement struct {
    Movement
    BasicVelocityProperties IBasicVelocityProperties
    SpeedBasicDecimal *BasicDecimal
    AxisMathVectorUtil *AxisMathVectorUtil
}


func (this *ConstantVelocityNotifyViewChangeMovement) initFields() {
    
    
    this.AxisMathVectorUtil = AxisMathVectorUtil_GetInstance()
}


func NewConstantVelocityNotifyViewChangeMovement() *ConstantVelocityNotifyViewChangeMovement {
    this := &ConstantVelocityNotifyViewChangeMovement{}
    this.Movement = *NewMovement()
    this.initFields()
        this.SpeedBasicDecimal = BasicDecimal_ZERO_BIGDECIMAL
        this.BasicVelocityProperties = NewBasicVelocityProperties()
    return this
}


func (this *ConstantVelocityNotifyViewChangeMovement) Init(speedBasicDecimal *BasicDecimal, angle int, otherAngle int) {
        this.SpeedBasicDecimal = speedBasicDecimal
        var angleFactory *AngleFactory = AngleFactory_GetInstance(); _ = angleFactory
        this.BasicVelocityProperties.SetVelocityWithBigDecimal(speedBasicDecimal, angleFactory.GetAt(angle), angleFactory.GetAt(otherAngle))
}


func (this *ConstantVelocityNotifyViewChangeMovement) MoveOutsideRadius(layer IAllBinaryLayer, radius int64, angle int, otherAngle int) {
        var scaleFactorValue int = this.SpeedBasicDecimal.GetScaledFactorValue(); _ = scaleFactorValue
        var xVector int = int((this.AxisMathVectorUtil.CalculateX(radius, angle) / int64(scaleFactorValue))); _ = xVector
        var yVector int = int((this.AxisMathVectorUtil.CalculateY(radius, angle) / int64(scaleFactorValue))); _ = yVector
        var zVector int = int((this.AxisMathVectorUtil.CalculateZ(radius, otherAngle) / int64(scaleFactorValue))); _ = zVector
        layer.MoveDXYZ(xVector, yVector, zVector)
}


// throws Exception
func (this *ConstantVelocityNotifyViewChangeMovement) Process(layer IAllBinaryGameLayer) {
        var x int = this.BasicVelocityProperties.GetVelocityXBasicDecimalP().GetScaled(); _ = x
        var y int = this.BasicVelocityProperties.GetVelocityYBasicDecimalP().GetScaled(); _ = y
        var z int = this.BasicVelocityProperties.GetVelocityZBasicDecimalP().GetScaled(); _ = z
        layer.MoveDXYZ(x, y, z)
        if x != 0 || y != 0 || z != 0 {
            layer.OnViewPositionChangeEvent()
        }
}


func (this *ConstantVelocityNotifyViewChangeMovement) Stop() {
        this.BasicVelocityProperties.Zero()
}

