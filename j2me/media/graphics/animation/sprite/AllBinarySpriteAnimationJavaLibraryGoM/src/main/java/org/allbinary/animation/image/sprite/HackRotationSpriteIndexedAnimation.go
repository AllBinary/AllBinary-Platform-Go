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

type HackRotationSpriteIndexedAnimation struct {
    RotationAnimation
    Sprite *Sprite
    Image IImage
}

// The Java class HackRotationSpriteIndexedAnimation has subclasses so its variables are this interface
type IHackRotationSpriteIndexedAnimation interface {
    IRotationAnimation
    HackRotationSpriteIndexedAnimation_() *HackRotationSpriteIndexedAnimation
    GetAnimationSize() int
    PaintFrame(g IGraphics, frame int, x int, y int)
    PaintXY(g IGraphics, x int, y int)
    Paint(g IGraphics)
    NextRotation()
    PreviousRotation()
    GetSize() int
    SetFrame(frame int)
    GetFrame() int
    SetSequence(sequence []int)
    GetSequence() []int
}

func (this *HackRotationSpriteIndexedAnimation) HackRotationSpriteIndexedAnimation_() *HackRotationSpriteIndexedAnimation {
    return this
}


func (this *HackRotationSpriteIndexedAnimation) initFields() {
    
    
}


func NewHackRotationSpriteIndexedAnimation(sprite *Sprite, image IImage, angleInfo *AngleInfo, animationBehavior IAnimationBehavior) *HackRotationSpriteIndexedAnimation {
    this := &HackRotationSpriteIndexedAnimation{}
        this.RotationAnimation = *NewRotationAnimation(angleInfo, CircularIndexUtil_CreateInstance(360 / int(angleInfo.GetAngleIncrementInfo().GetAngleIncrement())), animationBehavior)
        this.initFields()
        this.Sprite = sprite
        this.Image = image
    return this
}


// throws Exception
func (this *HackRotationSpriteIndexedAnimation) GetAnimationSize() int {
        return this.GetSize()
}


func (this *HackRotationSpriteIndexedAnimation) PaintFrame(g IGraphics, frame int, x int, y int) {
        this.SetFrame(frame)
        this.PaintXY(g, x, y)
}


func (this *HackRotationSpriteIndexedAnimation) PaintXY(g IGraphics, x int, y int) {
        this.Sprite.SetPosition(x, y)
        this.Paint(g)
}


func (this *HackRotationSpriteIndexedAnimation) Paint(g IGraphics) {
        this.Sprite.Paint(g)
}


func (this *HackRotationSpriteIndexedAnimation) NextRotation() {
        this.Sprite.NextFrame()
}


func (this *HackRotationSpriteIndexedAnimation) PreviousRotation() {
        this.Sprite.PrevFrame()
}


func (this *HackRotationSpriteIndexedAnimation) GetSize() int {
        return this.Sprite.GetRawFrameCount()
}


func (this *HackRotationSpriteIndexedAnimation) SetFrame(frame int) {
        this.Sprite.SetFrame(frame)
}


func (this *HackRotationSpriteIndexedAnimation) GetFrame() int {
        return this.Sprite.GetFrame()
}


func (this *HackRotationSpriteIndexedAnimation) SetSequence(sequence []int) {
}


func (this *HackRotationSpriteIndexedAnimation) GetSequence() []int {
        return PrimitiveIntUtil_GetArrayInstance()
}

