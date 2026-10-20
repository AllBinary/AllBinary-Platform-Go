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

type PickedUpLayerTypeFactory struct {
    NONE *PickedUpLayerType
    BEAM *PickedUpLayerType
    MINE *PickedUpLayerType
    PROJECTILE *PickedUpLayerType
    ENHANCEMENT *PickedUpLayerType
    MORPH *PickedUpLayerType
    PART *PickedUpLayerType
    HEAL *PickedUpLayerType
    LIFE *PickedUpLayerType
    POINTS *PickedUpLayerType
}

var PickedUpLayerTypeFactory_Instance *PickedUpLayerTypeFactory = NewPickedUpLayerTypeFactory()

func (this *PickedUpLayerTypeFactory) initFields() {
    this.NONE = NewPickedUpLayerType(StringUtil_GetInstance().NULL_STRING)
    this.BEAM = NewPickedUpLayerType("BEAM")
    this.MINE = NewPickedUpLayerType("MINE")
    this.PROJECTILE = NewPickedUpLayerType("PROJECTILE")
    this.ENHANCEMENT = NewPickedUpLayerType("Enhancement")
    this.MORPH = NewPickedUpLayerType("Morph")
    this.PART = NewPickedUpLayerType("Part")
    this.HEAL = NewPickedUpLayerType("Heal")
    this.LIFE = NewPickedUpLayerType("Life")
    this.POINTS = NewPickedUpLayerType("Points")
}

func NewPickedUpLayerTypeFactory() *PickedUpLayerTypeFactory {
    this := &PickedUpLayerTypeFactory{}
    this.initFields()
    return this
}


func PickedUpLayerTypeFactory_GetInstance() *PickedUpLayerTypeFactory {
        return PickedUpLayerTypeFactory_Instance
}

