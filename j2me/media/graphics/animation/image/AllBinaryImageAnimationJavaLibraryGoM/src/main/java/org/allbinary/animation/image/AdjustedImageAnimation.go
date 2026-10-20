/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2022 AllBinary 
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

type AdjustedImageAnimation struct {
    ImageAnimation
    Dx int
    Dy int
}


func (this *AdjustedImageAnimation) initFields() {
    
    
}


func NewAdjustedImageAnimation(image IImage, dx int, dy int, animationBehavior IAnimationBehavior) *AdjustedImageAnimation {
    this := &AdjustedImageAnimation{}
        this.ImageAnimation = *NewImageAnimation(image, animationBehavior)
        this.initFields()
        this.Init(dx, dy)
    return this
}


// throws Exception
func (this *AdjustedImageAnimation) Init(dx int, dy int) {
        this.SetDx(dx)
        this.SetDy(dy)
}


func (this *AdjustedImageAnimation) PaintXY(graphics IGraphics, x int, y int) {
        this.ImageAnimation.PaintXY(graphics, x + this.Dx, y + this.Dy)
}


func (this *AdjustedImageAnimation) SetDx(dx int) {
        this.Dx = dx
}


func (this *AdjustedImageAnimation) GetDx() int {
        return this.Dx
}


func (this *AdjustedImageAnimation) SetDy(dy int) {
        this.Dy = dy
}


func (this *AdjustedImageAnimation) GetDy() int {
        return this.Dy
}

