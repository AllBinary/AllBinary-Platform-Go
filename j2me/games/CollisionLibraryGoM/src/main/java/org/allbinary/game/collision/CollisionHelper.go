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

type CollisionHelper struct {
    SourceLayerInterface IAllBinaryLayer
}

// The Java class CollisionHelper has subclasses so its variables are this interface
type ICollisionHelper interface {
    CollisionHelper_() *CollisionHelper
    IsCollidable(layerInterface ICollidableCompositeLayer) bool
    GetOwnerLayerInterface() IAllBinaryLayer
    SetOwnerLayerInterface(sourceLayerInterface IAllBinaryLayer)
}

func (this *CollisionHelper) CollisionHelper_() *CollisionHelper {
    return this
}


func (this *CollisionHelper) initFields() {
    
}


func NewCollisionHelper(sourceLayerInterface IAllBinaryLayer) *CollisionHelper {
    this := &CollisionHelper{}
    this.initFields()
        this.SourceLayerInterface = sourceLayerInterface
    return this
}


func (this *CollisionHelper) IsCollidable(layerInterface ICollidableCompositeLayer) bool {
        return true
}


func (this *CollisionHelper) GetOwnerLayerInterface() IAllBinaryLayer {
        return this.SourceLayerInterface
}


func (this *CollisionHelper) SetOwnerLayerInterface(sourceLayerInterface IAllBinaryLayer) {
        this.SourceLayerInterface = sourceLayerInterface
}

