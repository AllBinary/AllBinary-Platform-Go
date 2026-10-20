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

type ExplosionResources struct {
    EXPLOSION_90_RESOURCE string
    EXPLOSION_60_RESOURCE string
    HALF_EXPLOSION_RESOURCE string
    THIRD_EXPLOSION_RESOURCE string
}

var ExplosionResources_Instance *ExplosionResources = NewExplosionResources()

func (this *ExplosionResources) initFields() {
    this.EXPLOSION_90_RESOURCE = "/explosion_sprite_90_by_90.png"
    this.EXPLOSION_60_RESOURCE = "/explosion_sprite_60_by_60.png"
    this.HALF_EXPLOSION_RESOURCE = "/explosion_sprite_30_by_30.png"
    this.THIRD_EXPLOSION_RESOURCE = "/explosion_sprite_15_by_15.png"
}


func ExplosionResources_GetInstance() *ExplosionResources {
        return ExplosionResources_Instance
}


func NewExplosionResources() *ExplosionResources {
    this := &ExplosionResources{}
    this.initFields()
    return this
}

