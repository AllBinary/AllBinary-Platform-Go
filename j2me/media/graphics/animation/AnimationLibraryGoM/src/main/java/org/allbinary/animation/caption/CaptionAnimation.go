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

type CaptionAnimation struct {
    Animation
    AnimationInterface IAnimation
    OverlayAnimationInterface IAnimation
    CaptionDx int
    CaptionDy int
    Dx int
    Dy int
}


func (this *CaptionAnimation) initFields() {
    
    
    
    
    
    
}


func NewCaptionAnimation(animationInterface IAnimation, overlayAnimationInterface IAnimation, captionDx int, captionDy int, dx int, dy int) *CaptionAnimation {
    this := &CaptionAnimation{}
    this.Animation = *NewAnimation()
    this.initFields()
        this.AnimationInterface = animationInterface
        this.OverlayAnimationInterface = overlayAnimationInterface
        this.CaptionDx = captionDx
        this.CaptionDy = captionDy
        this.Dx = dx
        this.Dy = dy
    return this
}


// throws Exception
func (this *CaptionAnimation) NextFrame() {
}


func (this *CaptionAnimation) PaintXY(graphics IGraphics, x int, y int) {
        this.AnimationInterface.PaintXY(graphics, x + this.CaptionDx, y + this.CaptionDy)
        this.OverlayAnimationInterface.PaintXY(graphics, x + this.CaptionDx + this.Dx, y + this.CaptionDy + this.Dy)
}


func (this *CaptionAnimation) PaintThreedXYZ(graphics IGraphics, x int, y int, z int) {
        this.AnimationInterface.PaintThreedXYZ(graphics, x + this.CaptionDx, y + this.CaptionDy, z)
        this.OverlayAnimationInterface.PaintThreedXYZ(graphics, x + this.CaptionDx + this.Dx, y + this.CaptionDy + this.Dy, z)
}

