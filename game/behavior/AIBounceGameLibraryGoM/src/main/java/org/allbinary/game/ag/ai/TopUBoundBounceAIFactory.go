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

type TopUBoundBounceAIFactory struct {
}

// Implements ArtificialIntelligenceInterfaceFactoryInterface


func NewTopUBoundBounceAIFactory() *TopUBoundBounceAIFactory {
    this := &TopUBoundBounceAIFactory{}
    return this
}


// throws Exception
func (this *TopUBoundBounceAIFactory) GetInstance(hashtable *ABHashtable[any, any], ownerLayerInterface IAllBinaryLayer, gameInput IGameInput) ArtificialIntelligenceInterface {
        var velocityInterfaceCompositeInterface VelocityInterfaceCompositeInterface = any(ownerLayerInterface).(VelocityInterfaceCompositeInterface); _ = velocityInterfaceCompositeInterface
        var velocityInterface VelocityInterface = any(velocityInterfaceCompositeInterface.GetVelocityProperties()).(VelocityInterface); _ = velocityInterface
        var maxDistancePerTick int = (velocityInterface.GetMaxForwardVelocity() >> velocityInterface.GetVelocityXBasicDecimalP().GetScaledFactor()); _ = maxDistancePerTick
        var halfWidth int = ownerLayerInterface.GetWidth() * 2 + maxDistancePerTick + 1; _ = halfWidth
        var halfHeight int = ownerLayerInterface.GetHeight() * 2 + maxDistancePerTick + 1; _ = halfHeight
        var point IGPoint = PointFactory_GetInstance().CreateXY(halfWidth, halfHeight + 60); _ = point
        var displayInfo *DisplayInfoSingleton = DisplayInfoSingleton_GetInstance(); _ = displayInfo
        var rectangle *Rectangle = NewRectangle(point, (displayInfo.GetLastWidth() - halfWidth - point.GetX()), (displayInfo.GetLastHeight() - halfHeight - point.GetY())); _ = rectangle
        return NewBoundBounceAI(ownerLayerInterface, gameInput, NewTopULayerBounds(rectangle), NewReverseVelocityBoundsVisitor(ownerLayerInterface))
}

