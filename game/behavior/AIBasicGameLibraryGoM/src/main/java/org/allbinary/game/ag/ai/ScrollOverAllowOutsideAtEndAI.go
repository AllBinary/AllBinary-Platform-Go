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

type ScrollOverAllowOutsideAtEndAI struct {
    BasicAI
}



func NewScrollOverAllowOutsideAtEndAI(ownerLayerInterface IAllBinaryLayer, gameInput IGameInput) *ScrollOverAllowOutsideAtEndAI {
    this := &ScrollOverAllowOutsideAtEndAI{}
        this.BasicAI = *NewBasicAI(ownerLayerInterface, gameInput)
    return this
}


// throws Exception
func (this *ScrollOverAllowOutsideAtEndAI) ProcessAI(allBinaryLayerManager IAllBinaryLayerManager) {
        var ownerLayerInterface IAllBinaryLayer = this.GetOwnerLayerInterface(); _ = ownerLayerInterface
        var x int = ownerLayerInterface.GetXP(); _ = x
        var y int = ownerLayerInterface.GetYP(); _ = y
        var width int = ownerLayerInterface.GetWidth(); _ = width
        var height int = ownerLayerInterface.GetHeight(); _ = height
        var displayInfo *DisplayInfoSingleton = DisplayInfoSingleton_GetInstance(); _ = displayInfo
        if x > displayInfo.GetLastWidth() {
            ownerLayerInterface.SetPosition(-width, y, ownerLayerInterface.GetZP())
        }
        if y > displayInfo.GetLastHeight() {
            ownerLayerInterface.SetPosition(x, -height, ownerLayerInterface.GetZP())
        }
        if x < -width {
            ownerLayerInterface.SetPosition(displayInfo.GetLastWidth() + width, y, ownerLayerInterface.GetZP())
        }
        if y < -height {
            ownerLayerInterface.SetPosition(x, displayInfo.GetLastHeight() + height, ownerLayerInterface.GetZP())
        }
}

