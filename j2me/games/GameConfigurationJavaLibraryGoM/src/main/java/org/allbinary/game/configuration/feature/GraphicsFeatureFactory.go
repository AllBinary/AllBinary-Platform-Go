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

type GraphicsFeatureFactory struct {
    NONE *GraphicsFeature
    TRANSPARENT_IMAGE_CREATION *GraphicsFeature
    IMAGE_GRAPHICS *GraphicsFeature
    IMAGE_ROTATION_ON_THE_FLY *GraphicsFeature
    IMAGE_TO_ARRAY_GRAPHICS *GraphicsFeature
    SPRITE_QUARTER_ROTATION_GRAPHICS *GraphicsFeature
    SPRITE_FULL_GRAPHICS *GraphicsFeature
    VECTOR_GRAPHICS *GraphicsFeature
    VECTOR_TO_IMAGE_GRAPHICS *GraphicsFeature
}

var GraphicsFeatureFactory_Instance *GraphicsFeatureFactory = NewGraphicsFeatureFactory()

func (this *GraphicsFeatureFactory) initFields() {
    this.NONE = NewGraphicsFeature(StringUtil_GetInstance().NULL_STRING)
    this.TRANSPARENT_IMAGE_CREATION = NewGraphicsFeature("Transparent Image Creation")
    this.IMAGE_GRAPHICS = NewGraphicsFeature("Image Graphics")
    this.IMAGE_ROTATION_ON_THE_FLY = NewGraphicsFeature("Image Rotated On The Fly")
    this.IMAGE_TO_ARRAY_GRAPHICS = NewGraphicsFeature("Image To Array Graphics")
    this.SPRITE_QUARTER_ROTATION_GRAPHICS = NewGraphicsFeature("Sprite Quarter Rotation Graphics")
    this.SPRITE_FULL_GRAPHICS = NewGraphicsFeature("Sprite Full Rotation Graphics")
    this.VECTOR_GRAPHICS = NewGraphicsFeature("BasicArrayList Graphics")
    this.VECTOR_TO_IMAGE_GRAPHICS = NewGraphicsFeature("BasicArrayList To Image Graphics")
}


func GraphicsFeatureFactory_GetInstance() *GraphicsFeatureFactory {
        return GraphicsFeatureFactory_Instance
}


func NewGraphicsFeatureFactory() *GraphicsFeatureFactory {
    this := &GraphicsFeatureFactory{}
    this.initFields()
    return this
}

