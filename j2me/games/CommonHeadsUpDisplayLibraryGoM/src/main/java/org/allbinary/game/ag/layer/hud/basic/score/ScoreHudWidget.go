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

type ScoreHudWidget struct {
    NumberStringHud
}



// throws Exception
func ScoreHudWidget_GetInstance(maxscore int, location int, direction int) *ScoreHudWidget {
        return NewScoreHudWidget(maxscore, location, direction, 2, BasicColorFactory_GetInstance().GREY)
}


func NewScoreHudWidget(maxscore int, location int, direction int, bufferZone int, basicColor *BasicColor) *ScoreHudWidget {
    this := &ScoreHudWidget{}
        this.NumberStringHud = *NewNumberStringHud("Pts ", maxscore, location, direction, bufferZone, basicColor)
        this.UpdateMaxHeight = 14
    return this
}


func (this *ScoreHudWidget) UpdateMeasurement(graphics IGraphics) {
        var font Font = graphics.GetFont(); _ = font
        this.UpdateMaxWidth = font.GetSize() * 5
        this.NumberStringHud.UpdateMeasurement(graphics)
}

