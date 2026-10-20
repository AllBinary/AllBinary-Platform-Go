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

type AdjustedImageArrayAnimation struct {
    ImageArrayAnimation
    Dx int
    Dy int
}


func (this *AdjustedImageArrayAnimation) initFields() {
    
    
}


func NewAdjustedImageArrayAnimation(imageArray []IImage, dx int, dy int, animationBehavior IAnimationBehavior) *AdjustedImageArrayAnimation {
    this := &AdjustedImageArrayAnimation{}
        this.ImageArrayAnimation = *NewImageArrayAnimation(imageArray, animationBehavior)
        this.initFields()
        this.Init(dx, dy)
    return this
}


// throws Exception
func (this *AdjustedImageArrayAnimation) Init(dx int, dy int) {
        this.SetDx(dx)
        this.SetDy(dy)
}


func (this *AdjustedImageArrayAnimation) PaintXY(graphics IGraphics, x int, y int) {
        this.ImageArrayAnimation.PaintXY(graphics, x + this.Dx, y + this.Dy)
}


func (this *AdjustedImageArrayAnimation) SetDx(dx int) {
        this.Dx = dx
}


func (this *AdjustedImageArrayAnimation) GetDx() int {
        return this.Dx
}


func (this *AdjustedImageArrayAnimation) SetDy(dy int) {
        this.Dy = dy
}


func (this *AdjustedImageArrayAnimation) GetDy() int {
        return this.Dy
}

