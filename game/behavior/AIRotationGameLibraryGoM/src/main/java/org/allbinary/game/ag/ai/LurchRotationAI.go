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

type LurchRotationAI struct {
    LurchAI
    RotationAnimationInterface IIndexedAnimation
}


func (this *LurchRotationAI) initFields() {
    
}


func NewLurchRotationAI(ownerLayerInterface IAllBinaryLayer, gameInput IGameInput) *LurchRotationAI {
    this := &LurchRotationAI{}
        this.LurchAI = *NewLurchAI(ownerLayerInterface, gameInput)
        this.initFields()
        var rotationAnimationInterfaceCompositeInterface RotationAnimationInterfaceCompositeInterface = any(this.GetOwnerLayerInterface()).(RotationAnimationInterfaceCompositeInterface); _ = rotationAnimationInterfaceCompositeInterface
        this.RotationAnimationInterface = rotationAnimationInterfaceCompositeInterface.GetRotationAnimationInterface()
        this.Update()
    return this
}


func (this *LurchRotationAI) Reverse() {
        this.LurchAI.Reverse()
        this.Update()
}


func (this *LurchRotationAI) Update() {
        if this.CurrentRelativeAngle == 0 {
            var rotationAnimationInterface RotationAnimationInterface = any(this.RotationAnimationInterface).(RotationAnimationInterface); _ = rotationAnimationInterface
            var angleIncrementInfo *AngleIncrementInfo = rotationAnimationInterface.GetAngleInfoP().GetAngleIncrementInfo(); _ = angleIncrementInfo
            var rotationAnimationInterfaceCompositeInterface RotationAnimationInterfaceCompositeInterface = any(this.GetOwnerLayerInterface()).(RotationAnimationInterfaceCompositeInterface); _ = rotationAnimationInterfaceCompositeInterface
            rotationAnimationInterfaceCompositeInterface.SetFrame(int(angleIncrementInfo.RIGHT_FRAME))
        } else if this.CurrentRelativeAngle == 180 {
            var rotationAnimationInterface RotationAnimationInterface = any(this.RotationAnimationInterface).(RotationAnimationInterface); _ = rotationAnimationInterface
            var angleIncrementInfo *AngleIncrementInfo = rotationAnimationInterface.GetAngleInfoP().GetAngleIncrementInfo(); _ = angleIncrementInfo
            var rotationAnimationInterfaceCompositeInterface RotationAnimationInterfaceCompositeInterface = any(this.GetOwnerLayerInterface()).(RotationAnimationInterfaceCompositeInterface); _ = rotationAnimationInterfaceCompositeInterface
            rotationAnimationInterfaceCompositeInterface.SetFrame(int(angleIncrementInfo.LEFT_FRAME))
        }
}

