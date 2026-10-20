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

type HudResources struct {
    RESOURCE_LIFE string
    RESOURCE_HEALTH string
}

var HudResources_Instance *HudResources = NewHudResources()

func (this *HudResources) initFields() {
    this.RESOURCE_LIFE = "/hud_life_10_by_10.png"
    this.RESOURCE_HEALTH = "/hud_health_10_by_10.png"
}


func HudResources_GetInstance() *HudResources {
        return HudResources_Instance
}


func NewHudResources() *HudResources {
    this := &HudResources{}
    this.initFields()
    return this
}

