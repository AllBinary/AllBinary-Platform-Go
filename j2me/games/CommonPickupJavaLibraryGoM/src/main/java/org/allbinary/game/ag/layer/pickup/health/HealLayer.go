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

type HealLayer struct {
    AllBinaryLayer
}

// Implements HealthVisitorInterface



func NewHealLayer() *HealLayer {
    this := &HealLayer{}
        this.AllBinaryLayer = *NewAllBinaryLayer(StringUtil_GetInstance().EMPTY_STRING, NewRectangle(PointFactory_GetInstance().ZERO_ZERO, 0, 0), ViewPosition_GetInstanceD())
    return this
}


func (this *HealLayer) Paint(graphics IGraphics) {
}


func (this *HealLayer) Visit(healthInterface HealthInterface) {
        if healthInterface.IsAlive() {
            healthInterface.HealMax()
        } else {
            this.LogUtil.PutF("Heal attempt on dead: Does not currently occur", this, "visit")
        }
}

