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

type AdjustedSpriteIndexedAnimation struct {
    SpriteIndexedAnimation
    Dx int
    Dy int
}


func (this *AdjustedSpriteIndexedAnimation) initFields() {
    
    
}


func NewAdjustedSpriteIndexedAnimation(sprite *Sprite, image IImage, basicColorArray []*BasicColor, dx int, dy int, animationBehavior IAnimationBehavior) *AdjustedSpriteIndexedAnimation {
    this := &AdjustedSpriteIndexedAnimation{}
        this.SpriteIndexedAnimation = *NewSpriteIndexedAnimation(sprite, image, basicColorArray, animationBehavior)
        this.initFields()
        this.Dx = dx
        this.Dy = dy
    return this
}


func (this *AdjustedSpriteIndexedAnimation) PaintFrame(graphics IGraphics, frame int, x int, y int) {
        this.SetFrame(frame)
        this.Sprite.SetPosition(x + this.Dx, y + this.Dy)
        this.SpriteIndexedAnimation.Paint(graphics)
}


func (this *AdjustedSpriteIndexedAnimation) PaintXY(g IGraphics, x int, y int) {
        this.Sprite.SetPosition(x + this.Dx, y + this.Dy)
        this.SpriteIndexedAnimation.Paint(g)
}


func (this *AdjustedSpriteIndexedAnimation) Paint(g IGraphics) {
        this.Sprite.SetPosition(this.Dx, this.Dy)
        this.SpriteIndexedAnimation.Paint(g)
}

