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

type SpriteIndexedAnimation struct {
    IndexedAnimation
    Sprite *Sprite
    Image IImage
    BasicColorArray []*BasicColor
}

// Implements ColorCompositeInterface

// The Java class SpriteIndexedAnimation has subclasses so its variables are this interface
type ISpriteIndexedAnimation interface {
    IIndexedAnimation
    SpriteIndexedAnimation_() *SpriteIndexedAnimation
    GetAnimationSize() int
    SetBasicColorP(basicColor *BasicColor)
    GetBasicColorP() *BasicColor
    PaintFrame(graphics IGraphics, frame int, x int, y int)
    PaintXY(graphics IGraphics, x int, y int)
    Paint(graphics IGraphics)
    NextFrame()
    PreviousFrame()
    GetSize() int
    SetFrame(frame int)
    GetFrame() int
    IsLastFrame() bool
    GetLastFrame() int
    SetSequence(sequence []int)
    GetSequence() []int
    Close()
    Finalize()
}

func (this *SpriteIndexedAnimation) SpriteIndexedAnimation_() *SpriteIndexedAnimation {
    return this
}


func (this *SpriteIndexedAnimation) initFields() {
    
    
    
}


func NewSpriteIndexedAnimation(sprite *Sprite, image IImage, basicColorArray []*BasicColor, animationBehavior IAnimationBehavior) *SpriteIndexedAnimation {
    this := &SpriteIndexedAnimation{}
        this.IndexedAnimation = *NewIndexedAnimation(animationBehavior)
        this.initFields()
        this.Sprite = sprite
        this.Image = image
        this.BasicColorArray = basicColorArray
        if len(this.BasicColorArray) != 0 && this.GetSize() != len(this.BasicColorArray) {
            panic(errors.New(fmt.Sprint(fmt.Sprint(NewStringMaker().Append(CommonLabels_GetInstance().TOTAL_LABEL).Appendint(this.GetSize()).Append("!=").Appendint(len(this.BasicColorArray))))))
        }
    return this
}


// throws Exception
func (this *SpriteIndexedAnimation) GetAnimationSize() int {
        return this.GetSize()
}


func (this *SpriteIndexedAnimation) SetBasicColorP(basicColor *BasicColor) {
        for index := int(0); index < len(this.BasicColorArray); index++ {
            if basicColor == this.BasicColorArray[index] {
                this.SetFrame(index)
                break
            }
        }
}


func (this *SpriteIndexedAnimation) GetBasicColorP() *BasicColor {
        if len(this.BasicColorArray) < this.GetFrame() {
            return this.BasicColorArray[this.GetFrame()]
        } else {
            return BasicColorFactory_GetInstance().WHITE
        }
}


func (this *SpriteIndexedAnimation) PaintFrame(graphics IGraphics, frame int, x int, y int) {
        this.SetFrame(frame)
        this.Sprite.SetPosition(x, y)
        this.Sprite.Paint(graphics)
}


func (this *SpriteIndexedAnimation) PaintXY(graphics IGraphics, x int, y int) {
        this.Sprite.SetPosition(x, y)
        this.Sprite.Paint(graphics)
}


func (this *SpriteIndexedAnimation) Paint(graphics IGraphics) {
        this.Sprite.Paint(graphics)
}


func (this *SpriteIndexedAnimation) NextFrame() {
        this.Sprite.NextFrame()
}


func (this *SpriteIndexedAnimation) PreviousFrame() {
        this.Sprite.PrevFrame()
}


func (this *SpriteIndexedAnimation) GetSize() int {
        return this.Sprite.GetRawFrameCount()
}


func (this *SpriteIndexedAnimation) SetFrame(frame int) {
        this.Sprite.SetFrame(frame)
}


func (this *SpriteIndexedAnimation) GetFrame() int {
        return this.Sprite.GetFrame()
}


func (this *SpriteIndexedAnimation) IsLastFrame() bool {
        if this.Sprite.GetFrame() == this.GetSize() - 1 {
            return true
        } else {
            return false
        }
}


func (this *SpriteIndexedAnimation) GetLastFrame() int {
        return this.GetSize()
}


func (this *SpriteIndexedAnimation) SetSequence(sequence []int) {
}


func (this *SpriteIndexedAnimation) GetSequence() []int {
        return PrimitiveIntUtil_GetArrayInstance()
}


// throws Exception
func (this *SpriteIndexedAnimation) Close() {
}


// throws Throwable
func (this *SpriteIndexedAnimation) Finalize() {
}

