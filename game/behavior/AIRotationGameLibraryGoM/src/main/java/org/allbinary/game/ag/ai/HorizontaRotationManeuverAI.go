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

type HorizontaRotationManeuverAI struct {
    BasicAI
    CircularIndexUtil *CircularIndexUtil
    CurrentSpeedDivisor int
    RotationAnimationInterface IIndexedAnimation
    BaseVelocityInterface IBasicVelocityProperties
    VelocityInterface VelocityInterface
}

// Implements ArtificialIntelligenceTransitionInterface


func (this *HorizontaRotationManeuverAI) initFields() {
    this.CircularIndexUtil = CircularIndexUtil_CreateInstanceAt(10, Integer_MAX_VALUE)
    this.CurrentSpeedDivisor = 5
    
    
    
}


func NewHorizontaRotationManeuverAI(ownerLayerInterface IAllBinaryLayer, gameInput IGameInput) *HorizontaRotationManeuverAI {
    this := &HorizontaRotationManeuverAI{}
        this.BasicAI = *NewBasicAI(ownerLayerInterface, gameInput)
        this.initFields()
        var rotationAnimationInterfaceCompositeInterface RotationAnimationInterfaceCompositeInterface = any(this.GetOwnerLayerInterface()).(RotationAnimationInterfaceCompositeInterface); _ = rotationAnimationInterfaceCompositeInterface
        this.RotationAnimationInterface = rotationAnimationInterfaceCompositeInterface.GetRotationAnimationInterface()
        var velocityInterfaceCompositeInterface VelocityInterfaceCompositeInterface = any(this.GetOwnerLayerInterface()).(VelocityInterfaceCompositeInterface); _ = velocityInterfaceCompositeInterface
        this.BaseVelocityInterface = velocityInterfaceCompositeInterface.GetVelocityProperties()
        this.VelocityInterface = any(this.BaseVelocityInterface).(VelocityInterface)
        var rotationAnimationInterface RotationAnimationInterface = any(this.RotationAnimationInterface).(RotationAnimationInterface); _ = rotationAnimationInterface
        var angleIncrementInfo *AngleIncrementInfo = rotationAnimationInterface.GetAngleInfoP().GetAngleIncrementInfo(); _ = angleIncrementInfo
        rotationAnimationInterfaceCompositeInterface.SetFrame(int(angleIncrementInfo.RIGHT_FRAME))
    return this
}


func (this *HorizontaRotationManeuverAI) Transition() {
        var rotationAnimationInterface RotationAnimationInterface = any(this.RotationAnimationInterface).(RotationAnimationInterface); _ = rotationAnimationInterface
        var angleIncrementInfo *AngleIncrementInfo = rotationAnimationInterface.GetAngleInfoP().GetAngleIncrementInfo(); _ = angleIncrementInfo
        var rotationAnimationInterfaceCompositeInterface RotationAnimationInterfaceCompositeInterface = any(this.GetOwnerLayerInterface()).(RotationAnimationInterfaceCompositeInterface); _ = rotationAnimationInterfaceCompositeInterface
        rotationAnimationInterfaceCompositeInterface.SetFrame(int(angleIncrementInfo.RIGHT_FRAME))
        var allbinaryLayer IAllBinaryLayer = this.GetOwnerLayerInterface(); _ = allbinaryLayer
        allbinaryLayer.SetPosition(allbinaryLayer.GetXP(), 1, allbinaryLayer.GetZP())
}


// throws Exception
func (this *HorizontaRotationManeuverAI) ProcessAI(allBinaryLayerManager IAllBinaryLayerManager) {
        var rotationAnimationInterface RotationAnimationInterface = (any(this.RotationAnimationInterface).(RotationAnimationInterface)); _ = rotationAnimationInterface
        var angleIncrementInfo *AngleIncrementInfo = rotationAnimationInterface.GetAngleInfoP().GetAngleIncrementInfo(); _ = angleIncrementInfo
        var frame int = this.RotationAnimationInterface.GetFrame(); _ = frame
        var ownerLayerInterface IAllBinaryLayer = this.GetOwnerLayerInterface(); _ = ownerLayerInterface
        if ownerLayerInterface.GetXP() - ownerLayerInterface.GetWidth() <= 0 && frame == int(angleIncrementInfo.LEFT_FRAME) {
            this.Reverse()
            this.Drop()
        }
        var displayInfo *DisplayInfoSingleton = DisplayInfoSingleton_GetInstance(); _ = displayInfo
        if ownerLayerInterface.GetX2() + ownerLayerInterface.GetWidth() > displayInfo.GetLastWidth() && frame == int(angleIncrementInfo.RIGHT_FRAME) {
            this.Reverse()
            this.Accelerate()
            this.Drop()
        }
        var index int = this.CircularIndexUtil.GetIndex(); _ = index
        if index % this.CurrentSpeedDivisor == 0 && index % 2 == 0 {
            this.BasicAI.ProcessKeyAI(Canvas_UP)
            this.VelocityInterface.LimitMaxXYVelocity(this.VelocityInterface.GetMaxForwardVelocity() / this.CurrentSpeedDivisor)
        }
        this.CircularIndexUtil.Next()
        if frame == int(angleIncrementInfo.LEFT_FRAME) {
            this.BasicAI.ProcessKeyAI(Canvas_KEY_NUM0)
        } else if frame == int(angleIncrementInfo.RIGHT_FRAME) {
            this.BasicAI.ProcessKeyAI(Canvas_KEY_POUND)
        }
}


func (this *HorizontaRotationManeuverAI) Reverse() {
        var rotationAnimationInterface RotationAnimationInterface = any(this.RotationAnimationInterface).(RotationAnimationInterface); _ = rotationAnimationInterface
        var angleIncrementInfo *AngleIncrementInfo = rotationAnimationInterface.GetAngleInfoP().GetAngleIncrementInfo(); _ = angleIncrementInfo
        var frame int = this.RotationAnimationInterface.GetFrame(); _ = frame
        if frame == int(angleIncrementInfo.LEFT_FRAME) {
            var rotationAnimationInterfaceCompositeInterface RotationAnimationInterfaceCompositeInterface = any(this.GetOwnerLayerInterface()).(RotationAnimationInterfaceCompositeInterface); _ = rotationAnimationInterfaceCompositeInterface
            rotationAnimationInterfaceCompositeInterface.SetFrame(int(angleIncrementInfo.RIGHT_FRAME))
            this.BaseVelocityInterface.Zero()
        } else if frame == int(angleIncrementInfo.RIGHT_FRAME) {
            var rotationAnimationInterfaceCompositeInterface RotationAnimationInterfaceCompositeInterface = any(this.GetOwnerLayerInterface()).(RotationAnimationInterfaceCompositeInterface); _ = rotationAnimationInterfaceCompositeInterface
            rotationAnimationInterfaceCompositeInterface.SetFrame(int(angleIncrementInfo.LEFT_FRAME))
            this.BaseVelocityInterface.Zero()
        }
}


func (this *HorizontaRotationManeuverAI) Accelerate() {
        if this.CurrentSpeedDivisor > 1 {
            this.CurrentSpeedDivisor--
        }
}


func (this *HorizontaRotationManeuverAI) Drop() {
        var displayInfo *DisplayInfoSingleton = DisplayInfoSingleton_GetInstance(); _ = displayInfo
        var ownerLayerInterface IAllBinaryLayer = this.GetOwnerLayerInterface(); _ = ownerLayerInterface
        var y int = ownerLayerInterface.GetYP(); _ = y
        if ownerLayerInterface.GetY2() + ownerLayerInterface.GetHeight() > displayInfo.GetLastHeight() {
            y = 0
        } else {
            y += ownerLayerInterface.GetHeight() + 1
        }
        ownerLayerInterface.SetPosition(ownerLayerInterface.GetXP(), y, ownerLayerInterface.GetZP())
}


func (this *HorizontaRotationManeuverAI) GetId() int {
        return ArtificialIntelligence_AI_ID
}

