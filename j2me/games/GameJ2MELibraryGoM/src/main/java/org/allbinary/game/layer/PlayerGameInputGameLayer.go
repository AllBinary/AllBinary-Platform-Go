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

type PlayerGameInputGameLayer struct {
    AllBinaryGameLayer
    PlayerGameInput IPlayerGameInput
}

// Implements PlayerGameInputCompositeInterface


func (this *PlayerGameInputGameLayer) initFields() {
    
}


func NewPlayerGameInputGameLayer(playerInputId int) *PlayerGameInputGameLayer {
    this := &PlayerGameInputGameLayer{}
        this.AllBinaryGameLayer = *NewAllBinaryGameLayer(StringUtil_GetInstance().EMPTY_STRING, NewRectangle(PointFactory_GetInstance().ZERO_ZERO, 0, 0), NewStaticViewPosition(0, 0, 0))
        this.initFields()
        this.PlayerGameInput = NewPlayerGameInput(this.GetGameKeyEventList(), NewBasicArrayListD(), playerInputId)
        var canvasStrings *CanvasStrings = CanvasStrings_GetInstance(); _ = canvasStrings
        this.LogUtil.PutF("Danger Danger Danger: Should Not Be Called Except For Testing Input", this, canvasStrings.PAINT)
    return this
}


func (this *PlayerGameInputGameLayer) GetPlayerGameInput() IPlayerGameInput {
        return this.PlayerGameInput
}


func (this *PlayerGameInputGameLayer) InitInputProcessors() {
}


// throws Exception
func (this *PlayerGameInputGameLayer) ProcessInput(myManager IAllBinaryLayerManager) {
}


func (this *PlayerGameInputGameLayer) Paint(graphics IGraphics) {
}

