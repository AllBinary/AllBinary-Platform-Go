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

type Layer struct {
    LogUtil *LogUtil
    X int
    Y int
    Z int
    Width int
    Height int
    Visible bool
}

// The Java class Layer has subclasses so its variables are this interface
type ILayer interface {
    Layer_() *Layer
    SetPosition(x int, y int, z int)
    MoveDXY(dx int, dy int)
    MoveDXYZ(dx int, dy int, dz int)
    GetXP() int
    GetYP() int
    GetZP() int
    GetWidth() int
    GetHeight() int
    GetDepth() int64
    SetVisible(visible bool)
    IsVisible() bool
    Paint(graphics IGraphics)
    SetLayerWidth(width int)
    SetLayerHeight(height int)
    ToStringAppend(stringBuffer *StringMaker)
}

func (this *Layer) Layer_() *Layer {
    return this
}

var Layer_ID int = SmallIntegerSingletonFactory_GetInstance().GetAt(0)

func (this *Layer) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
    this.Z = 3
    
    
    this.Visible = true
}


func NewLayer(width int, height int) *Layer {
    this := &Layer{}
    this.initFields()
        this.SetLayerWidth(width)
        this.SetLayerHeight(height)
    return this
}


func (this *Layer) SetPosition(x int, y int, z int) {
        this.X = x
        this.Y = y
        this.Z = z
}


func (this *Layer) MoveDXY(dx int, dy int) {
        this.X += dx
        this.Y += dy
}


func (this *Layer) MoveDXYZ(dx int, dy int, dz int) {
        this.X += dx
        this.Y += dy
        this.Z += dz
}


func (this *Layer) GetXP() int {
        return this.X
}


func (this *Layer) GetYP() int {
        return this.Y
}


func (this *Layer) GetZP() int {
        return this.Z
}


func (this *Layer) GetWidth() int {
        return this.Width
}


func (this *Layer) GetHeight() int {
        return this.Height
}


func (this *Layer) GetDepth() int64 {
        return 0
}


func (this *Layer) SetVisible(visible bool) {
        this.Visible = visible
}


func (this *Layer) IsVisible() bool {
        return this.Visible
}


func (this *Layer) Paint(graphics IGraphics) {
        var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
        var canvasStrings *CanvasStrings = CanvasStrings_GetInstance(); _ = canvasStrings
        this.LogUtil.PutF(commonStrings.NOT_IMPLEMENTED, this, canvasStrings.PAINT)
}


func (this *Layer) SetLayerWidth(width int) {
        if width < 0 {
            panic(errors.New(fmt.Sprint()))
        }
        this.Width = width
}


func (this *Layer) SetLayerHeight(height int) {
        if height < 0 {
            panic(errors.New(fmt.Sprint()))
        }
        this.Height = height
}


func (this *Layer) ToStringAppend(stringBuffer *StringMaker) {
}

