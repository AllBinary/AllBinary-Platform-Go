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

type CustomGravityConstantVelocityMovement struct {
    BasicConstantVelocityMovement
    CustomGravity int
    GravityUtil *GravityUtil
}


func (this *CustomGravityConstantVelocityMovement) initFields() {
    
    this.GravityUtil = GravityUtil_GetInstance()
}


func NewCustomGravityConstantVelocityMovement(speedBasicDecimal *BasicDecimal, customGravity int) *CustomGravityConstantVelocityMovement {
    this := &CustomGravityConstantVelocityMovement{}
        this.BasicConstantVelocityMovement = *NewBasicConstantVelocityMovement(speedBasicDecimal, NewVelocityProperties(int(speedBasicDecimal.GetUnscaled()), int(speedBasicDecimal.GetUnscaled())))
        this.initFields()
        this.CustomGravity = customGravity
    return this
}


// throws Exception
func (this *CustomGravityConstantVelocityMovement) Process(layer IAllBinaryGameLayer) {
        this.BasicConstantVelocityMovement.Process(layer)
        this.GravityUtil.Process(this.GetVelocityProperties(), this.CustomGravity)
        var velocityProperties *VelocityProperties = any(this.GetVelocityProperties()).(*VelocityProperties); _ = velocityProperties
        velocityProperties.LimitMaxYForwardVelocity()
}

