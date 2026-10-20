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

type BasicLayerPlacement struct {
    LayerType *LayerPlacementType
    Width int
    Height int
}


func (this *BasicLayerPlacement) initFields() {
    
    
    
}


func NewBasicLayerPlacement(layerType *LayerPlacementType, width int, height int) *BasicLayerPlacement {
    this := &BasicLayerPlacement{}
    this.initFields()
        this.LayerType = layerType
        this.Width = width
        this.Height = height
    return this
}


func (this *BasicLayerPlacement) GetWidth() int {
        return this.Width
}


func (this *BasicLayerPlacement) GetHeight() int {
        return this.Height
}


func (this *BasicLayerPlacement) GetLayerType() *LayerPlacementType {
        return this.LayerType
}


func (this *BasicLayerPlacement) SetLayerType(layerType *LayerPlacementType) {
        this.LayerType = layerType
}

