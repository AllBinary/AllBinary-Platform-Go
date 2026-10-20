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

package ag

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
    . "org/allbinary"
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
var _ = NullUtil_GetInstance

type ReverseVelocityBoundsVisitor struct {
    LogUtil *LogUtil
    VelocityInterface VelocityInterface
    LayerInterface LayerInterface
}

// Implements BoundsVisitorInterface


func (this *ReverseVelocityBoundsVisitor) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
}


func NewReverseVelocityBoundsVisitor(layerInterface LayerInterface) *ReverseVelocityBoundsVisitor {
    this := &ReverseVelocityBoundsVisitor{}
    this.initFields()
        this.LayerInterface = layerInterface
        var velocityInterfaceCompositeInterface VelocityInterfaceCompositeInterface = any(layerInterface).(VelocityInterfaceCompositeInterface); _ = velocityInterfaceCompositeInterface
        this.VelocityInterface = any(velocityInterfaceCompositeInterface.GetVelocityProperties()).(VelocityInterface)
    return this
}


func (this *ReverseVelocityBoundsVisitor) GetX() int {
        return this.LayerInterface.GetXP()
}


func (this *ReverseVelocityBoundsVisitor) GetY() int {
        return this.LayerInterface.GetYP()
}


func (this *ReverseVelocityBoundsVisitor) MinX() {
        if this.VelocityInterface.GetVelocityXBasicDecimalP().GetUnscaled() < 0 {
            this.VelocityInterface.GetVelocityXBasicDecimalP().Multiplyint(-1)
        }
}


func (this *ReverseVelocityBoundsVisitor) MaxX() {
        if this.VelocityInterface.GetVelocityXBasicDecimalP().GetUnscaled() > 0 {
            this.VelocityInterface.GetVelocityXBasicDecimalP().Multiplyint(-1)
        }
}


func (this *ReverseVelocityBoundsVisitor) MinY() {
        if this.VelocityInterface.GetVelocityYBasicDecimalP().GetUnscaled() < 0 {
            this.VelocityInterface.GetVelocityYBasicDecimalP().Multiplyint(-1)
        }
}


func (this *ReverseVelocityBoundsVisitor) MaxY() {
        if this.VelocityInterface.GetVelocityYBasicDecimalP().GetUnscaled() > 0 {
            this.VelocityInterface.GetVelocityYBasicDecimalP().Multiplyint(-1)
        }
}

