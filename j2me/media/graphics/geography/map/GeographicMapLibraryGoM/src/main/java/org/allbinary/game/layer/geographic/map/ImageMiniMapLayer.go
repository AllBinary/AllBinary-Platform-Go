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

type ImageMiniMapLayer struct {
    MiniMapLayer
    Image IImage
    Anchor int
}


func (this *ImageMiniMapLayer) initFields() {
    this.Image = NullImage_NULL_IMAGE
    this.Anchor = Anchor_TOP_LEFT
}


func NewImageMiniMapLayer(geographicMapInterface IBasicGeographicMap, viewPosition IViewPositionBase) *ImageMiniMapLayer {
    this := &ImageMiniMapLayer{}
        this.MiniMapLayer = *NewMiniMapLayer(geographicMapInterface, viewPosition)
        this.initFields()
    return this
}


// throws Exception
func (this *ImageMiniMapLayer) Init() {
        var aWidth int = this.AllBinaryTiledLayer.GetWidth(); _ = aWidth
        var aHeight int = this.AllBinaryTiledLayer.GetHeight(); _ = aHeight
        this.Image = PaintableToImageUtil_GetImage(this.AllBinaryTiledLayer, aWidth, aHeight)
}


func (this *ImageMiniMapLayer) Paint(graphics IGraphics) {
        graphics.DrawImage(this.Image, this.X, this.Y, this.Anchor)
        this.PaintDots(graphics)
}

