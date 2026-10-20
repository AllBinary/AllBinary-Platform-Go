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

type RelativeLayerRelationship struct {
    RelativeRelationship
    Layer IAllBinaryLayer
}

// The Java class RelativeLayerRelationship has subclasses so its variables are this interface
type IRelativeLayerRelationship interface {
    IRelativeRelationship
    RelativeLayerRelationship_() *RelativeLayerRelationship
    GetX() int
    GetY() int
    GetZ() int
}

func (this *RelativeLayerRelationship) RelativeLayerRelationship_() *RelativeLayerRelationship {
    return this
}


func (this *RelativeLayerRelationship) initFields() {
    this.Layer = AllBinaryLayer_NULL_ALLBINARY_LAYER
}


func NewRelativeLayerRelationship(layer IAllBinaryLayer, point IGPoint, typesAllowedList IBasicArrayList) *RelativeLayerRelationship {
    this := &RelativeLayerRelationship{}
        this.RelativeRelationship = *NewRelativeRelationship(point, typesAllowedList)
        this.initFields()
        this.Layer = layer
    return this
}


func (this *RelativeLayerRelationship) GetX() int {
        return this.Layer.GetXP() + this.RelativeRelationship.GetX()
}


func (this *RelativeLayerRelationship) GetY() int {
        return this.Layer.GetYP() + this.RelativeRelationship.GetY()
}


func (this *RelativeLayerRelationship) GetZ() int {
        return this.Layer.GetZP() + this.RelativeRelationship.GetZ()
}

