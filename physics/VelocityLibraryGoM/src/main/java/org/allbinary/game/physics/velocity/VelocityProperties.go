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

type VelocityProperties struct {
    BasicVelocityProperties
    MaxForwardVelocity int
    MaxReverseVelocity int
}

// Implements VelocityInterface

var VelocityProperties_NULL_VELOCITY_PROPERTIES *VelocityProperties = NewVelocityProperties(0, 0)

func (this *VelocityProperties) initFields() {
    
    
}


func NewVelocityProperties(maxForwardVelocity int, maxReverseVelocity int) *VelocityProperties {
    this := &VelocityProperties{}
    this.BasicVelocityProperties = *NewBasicVelocityProperties()
    this.initFields()
        this.SetMaxForwardVelocity(maxForwardVelocity)
        this.SetMaxReverseVelocity(maxReverseVelocity)
    return this
}


func (this *VelocityProperties) GetMaxForwardVelocity() int {
        return this.MaxForwardVelocity
}


func (this *VelocityProperties) SetMaxForwardVelocity(maxForwardVelocity int) {
        this.MaxForwardVelocity = maxForwardVelocity
}


func (this *VelocityProperties) GetMaxReverseVelocity() int {
        return this.MaxReverseVelocity
}


func (this *VelocityProperties) SetMaxReverseVelocity(maxReverseVelocity int) {
        this.MaxReverseVelocity = maxReverseVelocity
}


func (this *VelocityProperties) LimitMaxXYForwardVelocity() {
        this.LimitMaxXYVelocity(this.GetMaxForwardVelocity())
}


func (this *VelocityProperties) LimitMaxXYReverseVelocity() {
        this.LimitMaxXYVelocity(this.GetMaxReverseVelocity())
}


func (this *VelocityProperties) LimitMaxYForwardVelocity() {
        this.LimitMaxYVelocity(this.GetMaxForwardVelocity())
}


func (this *VelocityProperties) LimitMaxYReverseVelocity() {
        this.LimitMaxYVelocity(this.GetMaxReverseVelocity())
}


func (this *VelocityProperties) LimitMaxXForwardVelocity() {
        this.LimitMaxXVelocity(this.GetMaxForwardVelocity())
}


func (this *VelocityProperties) LimitMaxXReverseVelocity() {
        this.LimitMaxXVelocity(this.GetMaxReverseVelocity())
}


func (this *VelocityProperties) IsOverXYMaxForwardVelocity() bool {
        return this.IsOverXYMaxVelocity(this.GetMaxForwardVelocity())
}


func (this *VelocityProperties) IsOverXYMaxReverseVelocity() bool {
        return this.IsOverXYMaxVelocity(this.GetMaxReverseVelocity())
}


func (this *VelocityProperties) LimitXYToForwardAndReverseMaxVelocity() {
        this.LimitMaxXYForwardVelocity()
        this.LimitMaxXYReverseVelocity()
}


func (this *VelocityProperties) LimitMaxXYVelocity(maxVelocity int) {
        this.LimitMaxXVelocity(maxVelocity)
        this.LimitMaxYVelocity(maxVelocity)
}


func (this *VelocityProperties) LimitMaxPositiveYVelocity(maxVelocity int) {
        if this.VelocityYBasicDecimal.GetUnscaled() > int64(maxVelocity) {
            this.VelocityYBasicDecimal.Setint(maxVelocity)
        }
}


func (this *VelocityProperties) LimitMaxNegativeYVelocity(maxVelocity int) {
        if this.VelocityYBasicDecimal.GetUnscaled() < int64(-maxVelocity) {
            this.VelocityYBasicDecimal.Setint(-maxVelocity)
        }
}


func (this *VelocityProperties) LimitMaxYVelocity(maxVelocity int) {
        this.LimitMaxPositiveYVelocity(maxVelocity)
        this.LimitMaxNegativeYVelocity(maxVelocity)
}


func (this *VelocityProperties) LimitMaxXVelocity(maxVelocity int) {
        if this.VelocityXBasicDecimal.GetUnscaled() > int64(maxVelocity) {
            this.VelocityXBasicDecimal.Setint(maxVelocity)
        }
        if this.VelocityXBasicDecimal.GetUnscaled() < int64(-maxVelocity) {
            this.VelocityXBasicDecimal.Setint(-maxVelocity)
        }
}


func (this *VelocityProperties) IsOverXYMaxVelocity(maxVelocity int) bool {
        if this.VelocityXBasicDecimal.GetUnscaled() > int64(maxVelocity) {
            return true
        }
        if this.VelocityXBasicDecimal.GetUnscaled() < int64(-maxVelocity) {
            return true
        }
        if this.VelocityYBasicDecimal.GetUnscaled() > int64(maxVelocity) {
            return true
        }
        if this.VelocityYBasicDecimal.GetUnscaled() < int64(-maxVelocity) {
            return true
        }
        return false
}


func (this *VelocityProperties) SetVelocityi(magnitude int64, angle int, otherAngle int) {
        this.BasicVelocityProperties.SetVelocityi(magnitude, angle, otherAngle)
        this.LimitXYToForwardAndReverseMaxVelocity()
}


func (this *VelocityProperties) AddVelocityi(magnitude int64, angle int, otherAngle int) {
        this.BasicVelocityProperties.AddVelocityi(magnitude, angle, otherAngle)
        this.LimitXYToForwardAndReverseMaxVelocity()
}

