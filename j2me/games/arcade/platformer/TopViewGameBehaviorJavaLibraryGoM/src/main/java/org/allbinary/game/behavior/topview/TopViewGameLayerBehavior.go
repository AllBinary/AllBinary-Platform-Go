/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2022 AllBinary 
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

type TopViewGameLayerBehavior struct {
    GameLayerBehavior
    MaxGravityActionIndex int
    IsJumpAction bool
    IsJumpOver bool
    IsFallingWithoutJumpAttempt bool
    GravityActionIndex int
}

// The Java class TopViewGameLayerBehavior has subclasses so its variables are this interface
type ITopViewGameLayerBehavior interface {
    IGameLayerBehavior
    TopViewGameLayerBehavior_() *TopViewGameLayerBehavior
    Gravity()
    Land(velocityProperties *VelocityProperties)
    LandReset()
    Up(velocityProperties *VelocityProperties, acceleration *BasicAccelerationProperties, jumpBehavior *InitialJumpBehavior, accelerationMultiplier int)
    InputFrames(velocityProperties *VelocityProperties)
}

func (this *TopViewGameLayerBehavior) TopViewGameLayerBehavior_() *TopViewGameLayerBehavior {
    return this
}


func (this *TopViewGameLayerBehavior) initFields() {
    
    this.IsJumpAction = true
    this.IsJumpOver = false
    this.IsFallingWithoutJumpAttempt = false
    this.GravityActionIndex = 0
}


func NewTopViewGameLayerBehavior(maxGravityActionIndex int) *TopViewGameLayerBehavior {
    this := &TopViewGameLayerBehavior{}
    this.GameLayerBehavior = *NewGameLayerBehavior()
    this.initFields()
        this.MaxGravityActionIndex = maxGravityActionIndex
    return this
}


func (this *TopViewGameLayerBehavior) Gravity() {
        if this.GravityActionIndex == 0 {
            this.GravityActionIndex++
            this.IsFallingWithoutJumpAttempt = true
        }
}


func (this *TopViewGameLayerBehavior) Land(velocityProperties *VelocityProperties) {
        velocityProperties.GetVelocityYBasicDecimalP().Setint(0)
        this.LandReset()
}


func (this *TopViewGameLayerBehavior) LandReset() {
        this.GravityActionIndex = 0
        this.IsFallingWithoutJumpAttempt = false
        this.IsJumpAction = true
        this.IsJumpOver = false
}


func (this *TopViewGameLayerBehavior) Up(velocityProperties *VelocityProperties, acceleration *BasicAccelerationProperties, jumpBehavior *InitialJumpBehavior, accelerationMultiplier int) {
        if !this.IsJumpOver {
            if this.GravityActionIndex < this.MaxGravityActionIndex {
                var acceleration2 int = -acceleration.GetForward() * accelerationMultiplier; _ = acceleration2
                velocityProperties.GetVelocityYBasicDecimalP().Addint(acceleration2)
                velocityProperties.LimitXYToForwardAndReverseMaxVelocity()
                this.GravityActionIndex++
            } else {
            }
        } else {
        }
        if this.IsJumpAction {
            jumpBehavior.Process()
            this.IsJumpAction = false
        }
}


func (this *TopViewGameLayerBehavior) InputFrames(velocityProperties *VelocityProperties) {
        if this.GravityActionIndex > 0 && velocityProperties.GetVelocityYBasicDecimalP().GetUnscaled() > 0 {
            this.IsJumpOver = true
        }
}

