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

type HealthBarThreedAnimation struct {
    HealthBarAnimation
    Rectangle *Rectangle
    BasicColorFactory *BasicColorFactory
    Object3dProcessor IObject3dProcessor
}


func (this *HealthBarThreedAnimation) initFields() {
    
    this.BasicColorFactory = BasicColorFactory_GetInstance()
    this.Object3dProcessor = Object3dProcessorUtil_GetInstance().GetInstanceObject3dProcessor()
}


func NewHealthBarThreedAnimation(layerInterface IAllBinaryLayer, location int) *HealthBarThreedAnimation {
    this := &HealthBarThreedAnimation{}
        this.HealthBarAnimation = *NewHealthBarAnimation(layerInterface, location)
        this.initFields()
        this.BasicColor = this.BasicColorFactory.GREEN
        this.ColorP = int(this.BasicColor)
        this.Rectangle = NewRectangle(0, 0, 1, 1, this.BasicColor)
        this.Rectangle.SetDoubleSidedEnabled(true)
        this.Rectangle.SetNormalsEnabled(false)
        this.Rectangle.SetTexturesEnabled(false)
    return this
}


func (this *HealthBarThreedAnimation) UpdateColor() {
        var colorBufferList *ColorBufferList = this.Rectangle.GetVertices().GetColor4BufferList(); _ = colorBufferList
        var size int = colorBufferList.Size(); _ = size
        for index := int(size); Java_PreDecrement(&index) >= 0;  {
            colorBufferList.Set(index, this.BasicColor)
        }
}


func (this *HealthBarThreedAnimation) UpdateSize() {
        var number3dBufferList *Number3dBufferList = this.Rectangle.GetVertices().GetPointsOnFacesNumber3dBufferList(); _ = number3dBufferList
        var width int = this.X2; _ = width
        var height int = this.Thickness * 2; _ = height
        var segsH int = 1; _ = segsH
        var segsW int = 1; _ = segsW
        var row int; _ = row
        var col int; _ = col
        var w float32 = float32(width / segsW); _ = w
        var h float32 = float32(height / segsH); _ = h
        var width5 float32 = float32(width) / 2; _ = width5
        var height5 float32 = float32(height) / 2; _ = height5
        var index int = 0; _ = index
        for row = 0; row <= segsH; row++ {
            for col = 0; col <= segsW; col++ {
                number3dBufferList.Set(Java_PostIncrement(&index), float32(col) * w - width5, float32(row) * h - height5, 0)
            }
        }
}


func (this *HealthBarThreedAnimation) OnHealthChange(newX2 int) {
        this.HealthBarAnimation.OnHealthChange(newX2)
        this.UpdateColor()
        this.UpdateSize()
}


func (this *HealthBarThreedAnimation) PaintThreedXYZ(graphics IGraphics, x int, y int, z int) {
        var openGLESGraphics IOpenGLESGraphics = any(graphics).(IOpenGLESGraphics); _ = openGLESGraphics
        var gl GL10 = openGLESGraphics.GetGl10(); _ = gl
        var viewPosition IViewPositionBase = this.AllbinaryLayer.GetViewPosition(); _ = viewPosition
        var positionNumber3d *Number3d = this.Rectangle.GetPosition(); _ = positionNumber3d
        positionNumber3d.X = -(float32(viewPosition.GetX()))
        positionNumber3d.Z = -(float32(viewPosition.GetY()))
        positionNumber3d.Y = float32(viewPosition.GetZ() + this.AllbinaryLayer.GetHeight())
        this.Object3dProcessor.DrawObject(gl, this.Rectangle)
}

