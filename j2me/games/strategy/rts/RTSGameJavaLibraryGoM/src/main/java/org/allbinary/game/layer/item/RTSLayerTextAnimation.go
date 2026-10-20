/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2003 AllBinary 
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

type RTSLayerTextAnimation struct {
    Animation
    Image IImage
    Text string
    MyFontProcessor IMyFontProcessor
    FontHeight int
}

// Implements UpdateMyFontInterface


func (this *RTSLayerTextAnimation) initFields() {
    
    
    this.MyFontProcessor = NewUpdateMyFontProcessor(this)
    this.FontHeight = 0
}


func NewRTSLayerTextAnimation(text string, image IImage) *RTSLayerTextAnimation {
    this := &RTSLayerTextAnimation{}
    this.Animation = *NewAnimation()
    this.initFields()
        this.Text = text
        this.Image = image
    return this
}


func (this *RTSLayerTextAnimation) UpdateMeasurement(graphics IGraphics) {
        var font Font = graphics.GetFont(); _ = font
        this.FontHeight = font.GetHeight()
        this.MyFontProcessor = MyFontProcessor_GetInstance()
}


func (this *RTSLayerTextAnimation) PaintXY(graphics IGraphics, x int, y int) {
        this.MyFontProcessor.Process(graphics)
        this.Animation.PaintXY(graphics, x, y)
        var adjustedCostY int = this.Image.GetHeight() - this.FontHeight; _ = adjustedCostY
        graphics.DrawString(this.Text, x, y + adjustedCostY, 0)
}

