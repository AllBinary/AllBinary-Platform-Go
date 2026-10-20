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

type ScrollOverAtHalfAI struct {
    BasicAI
}



func NewScrollOverAtHalfAI(ownerLayerInterface IAllBinaryLayer, gameInput IGameInput) *ScrollOverAtHalfAI {
    this := &ScrollOverAtHalfAI{}
        this.BasicAI = *NewBasicAI(ownerLayerInterface, gameInput)
    return this
}


// throws Exception
func (this *ScrollOverAtHalfAI) ProcessAI(allBinaryLayerManager IAllBinaryLayerManager) {
        var ownerLayerInterface IAllBinaryLayer = this.GetOwnerLayerInterface(); _ = ownerLayerInterface
        var x int = ownerLayerInterface.GetXP(); _ = x
        var y int = ownerLayerInterface.GetYP(); _ = y
        var halfWidth int = ownerLayerInterface.GetHalfWidth(); _ = halfWidth
        var halfHeight int = ownerLayerInterface.GetHalfHeight(); _ = halfHeight
        var displayInfo *DisplayInfoSingleton = DisplayInfoSingleton_GetInstance(); _ = displayInfo
        if x > displayInfo.GetLastWidth() - halfWidth {
            ownerLayerInterface.SetPosition(halfWidth + 1, y, ownerLayerInterface.GetZP())
        }
        if y > displayInfo.GetLastHeight() - halfHeight {
            ownerLayerInterface.SetPosition(x, halfHeight + 1, ownerLayerInterface.GetZP())
        }
        if x < -halfWidth {
            ownerLayerInterface.SetPosition(displayInfo.GetLastWidth() - halfWidth, y, ownerLayerInterface.GetZP())
        }
        if y < -halfHeight {
            ownerLayerInterface.SetPosition(x, displayInfo.GetLastHeight() - halfHeight, ownerLayerInterface.GetZP())
        }
}

