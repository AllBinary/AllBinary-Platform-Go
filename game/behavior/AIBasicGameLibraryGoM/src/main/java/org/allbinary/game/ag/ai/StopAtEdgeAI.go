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

type StopAtEdgeAI struct {
    BasicAI
}



func NewStopAtEdgeAI(ownerLayerInterface IAllBinaryLayer, gameInput IGameInput) *StopAtEdgeAI {
    this := &StopAtEdgeAI{}
        this.BasicAI = *NewBasicAI(ownerLayerInterface, gameInput)
    return this
}


// throws Exception
func (this *StopAtEdgeAI) ProcessAI(allBinaryLayerManager IAllBinaryLayerManager) {
        var ownerLayerInterface IAllBinaryLayer = this.GetOwnerLayerInterface(); _ = ownerLayerInterface
        var velocityInterfaceCompositeInterface VelocityInterfaceCompositeInterface = (any(ownerLayerInterface).(VelocityInterfaceCompositeInterface)); _ = velocityInterfaceCompositeInterface
        var velocityProperties IBasicVelocityProperties = velocityInterfaceCompositeInterface.GetVelocityProperties(); _ = velocityProperties
        var x int = ownerLayerInterface.GetXP(); _ = x
        var y int = ownerLayerInterface.GetYP(); _ = y
        var x2 int = ownerLayerInterface.GetX2(); _ = x2
        var y2 int = ownerLayerInterface.GetY2(); _ = y2
        var width int = ownerLayerInterface.GetWidth(); _ = width
        var height int = ownerLayerInterface.GetHeight(); _ = height
        var displayInfo *DisplayInfoSingleton = DisplayInfoSingleton_GetInstance(); _ = displayInfo
        if x2 + velocityProperties.GetVelocityXBasicDecimalP().GetScaled() > displayInfo.GetLastWidth() {
            ownerLayerInterface.SetPosition(displayInfo.GetLastWidth() - width, y, ownerLayerInterface.GetZP())
            velocityProperties.GetVelocityYBasicDecimalP().Setint(0)
        }
        if y2 + velocityProperties.GetVelocityYBasicDecimalP().GetScaled() > displayInfo.GetLastHeight() {
            ownerLayerInterface.SetPosition(x, displayInfo.GetLastHeight() - height, ownerLayerInterface.GetZP())
            velocityProperties.GetVelocityXBasicDecimalP().Setint(0)
        }
        if x + velocityProperties.GetVelocityXBasicDecimalP().GetScaled() < 0 {
            ownerLayerInterface.SetPosition(0, y, ownerLayerInterface.GetZP())
            velocityProperties.GetVelocityYBasicDecimalP().Setint(0)
        }
        if y + velocityProperties.GetVelocityYBasicDecimalP().GetScaled() < 0 {
            ownerLayerInterface.SetPosition(x, 0, ownerLayerInterface.GetZP())
            velocityProperties.GetVelocityXBasicDecimalP().Setint(0)
        }
}

