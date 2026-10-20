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

type GameLayerUtil struct {
}



func NewGameLayerUtil() *GameLayerUtil {
    this := &GameLayerUtil{}
    return this
}


func GameLayerUtil_IsOnScreen(layer IAllBinaryLayer) bool {
        var displayInfo *DisplayInfoSingleton = DisplayInfoSingleton_GetInstance(); _ = displayInfo
        var viewPosition IViewPositionBase = layer.GetViewPosition(); _ = viewPosition
        return GameLayerUtil_IsInside(viewPosition.GetX(), viewPosition.GetY(), viewPosition.GetX2(), viewPosition.GetY2(), 0, 0, displayInfo.GetLastWidth(), displayInfo.GetLastHeight())
}


func GameLayerUtil_IsInside(rectX1 int, rectY1 int, rectX2 int, rectY2 int, rect2X1 int, rect2Y1 int, rect2X2 int, rect2Y2 int) bool {
        if rectX1 > rect2X1 && rectX2 < rect2X2 && rectY1 > rect2Y1 && rectY2 < rect2Y2 {
            return true
        } else {
            return false
        }
}

