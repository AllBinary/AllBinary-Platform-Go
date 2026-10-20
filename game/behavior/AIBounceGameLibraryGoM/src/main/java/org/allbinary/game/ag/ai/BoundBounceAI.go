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

type BoundBounceAI struct {
    BasicAI
    CurrentRelativeAngle int
    LayerBounds ILayerBounds
    BoundsVisitorInterface BoundsVisitorInterface
}


func (this *BoundBounceAI) initFields() {
    this.CurrentRelativeAngle = 0
    
    
}


func NewBoundBounceAI(ownerLayerInterface IAllBinaryLayer, gameInput IGameInput, layerBounds ILayerBounds, boundsVisitorInterface BoundsVisitorInterface) *BoundBounceAI {
    this := &BoundBounceAI{}
        this.BasicAI = *NewBasicAI(ownerLayerInterface, gameInput)
        this.initFields()
        this.LayerBounds = layerBounds
        this.BoundsVisitorInterface = boundsVisitorInterface
    return this
}


func (this *BoundBounceAI) Paint(graphics IGraphics) {
        var rectange *Rectangle = this.LayerBounds.GetRectangleP(); _ = rectange
        graphics.DrawRect(rectange.GetPoint().GetX(), rectange.GetPoint().GetY(), rectange.GetWidth(), rectange.GetHeight())
}


// throws Exception
func (this *BoundBounceAI) ProcessAI(allBinaryLayerManager IAllBinaryLayerManager) {
        this.LayerBounds.Visit(this.BoundsVisitorInterface)
}

