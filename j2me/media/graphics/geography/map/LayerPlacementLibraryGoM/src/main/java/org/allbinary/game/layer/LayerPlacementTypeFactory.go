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

type LayerPlacementTypeFactory struct {
    MAP *LayerPlacementType
    LEFT *LayerPlacementType
    DOWN *LayerPlacementType
    UP *LayerPlacementType
    RIGHT *LayerPlacementType
}

var LayerPlacementTypeFactory_Instance *LayerPlacementTypeFactory = NewLayerPlacementTypeFactory()

func (this *LayerPlacementTypeFactory) initFields() {
    this.MAP = NewLayerPlacementType(Integer_MAX_VALUE)
    this.LEFT = NewLayerPlacementType(DirectionFactory_GetInstance().LEFT.GetValue())
    this.DOWN = NewLayerPlacementType(DirectionFactory_GetInstance().DOWN.GetValue())
    this.UP = NewLayerPlacementType(DirectionFactory_GetInstance().UP.GetValue())
    this.RIGHT = NewLayerPlacementType(DirectionFactory_GetInstance().RIGHT.GetValue())
}

func NewLayerPlacementTypeFactory() *LayerPlacementTypeFactory {
    this := &LayerPlacementTypeFactory{}
    this.initFields()
    return this
}


func LayerPlacementTypeFactory_GetInstance() *LayerPlacementTypeFactory {
        return LayerPlacementTypeFactory_Instance
}

