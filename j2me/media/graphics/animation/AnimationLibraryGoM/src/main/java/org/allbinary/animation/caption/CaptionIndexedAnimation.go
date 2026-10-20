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

type CaptionIndexedAnimation struct {
    IndexedAnimation
    AnimationInterface IAnimation
    MovieIndexedAnimationInterface IIndexedAnimation
    CaptionDx int
    CaptionDy int
    Dx int
    Dy int
    Time int
    END_EVENT IAllBinaryEventObject
    TimeDelayHelper ITimeDelayHelper
    Player *Player
}


func (this *CaptionIndexedAnimation) initFields() {
    
    
    
    
    
    
    
    
    
    
}


func NewCaptionIndexedAnimation(animationInterface IAnimation, movieIndexedAnimationInterface IIndexedAnimation, player *Player, captionDx int, captionDy int, dx int, dy int, time int, animationBehavior IAnimationBehavior) *CaptionIndexedAnimation {
    this := &CaptionIndexedAnimation{}
        this.IndexedAnimation = *NewIndexedAnimation(animationBehavior)
        this.initFields()
        this.AnimationInterface = animationInterface
        this.MovieIndexedAnimationInterface = movieIndexedAnimationInterface
        this.CaptionDx = captionDx
        this.CaptionDy = captionDy
        this.Dx = dx
        this.Dy = dy
        this.Time = time
        this.END_EVENT = NewAllBinaryEventObject(this)
        var timePerFrame int = this.Time / this.GetSize(); _ = timePerFrame
        this.TimeDelayHelper = NewTimeDelayHelper(timePerFrame)
        this.Player = player
        this.Player.Start()
    return this
}


// throws Exception
func (this *CaptionIndexedAnimation) NextFrame() {
        if this.TimeDelayHelper.IsTime(GameTickTimeDelayHelperFactory_GetInstance().StartTime) {
            this.MovieIndexedAnimationInterface.NextFrame()
        }
        if this.IsLastFrame() {
            AnimationEventHandler_GetInstance().FireEvent(this.END_EVENT)
        }
}


func (this *CaptionIndexedAnimation) IsLastFrame() bool {
        if this.MovieIndexedAnimationInterface.GetFrame() == this.GetSize() - 1 {
            return true
        } else {
            return false
        }
}


func (this *CaptionIndexedAnimation) PreviousFrame() {
        this.MovieIndexedAnimationInterface.PreviousFrame()
}


func (this *CaptionIndexedAnimation) SetFrame(index int) {
        this.MovieIndexedAnimationInterface.SetFrame(index)
}


func (this *CaptionIndexedAnimation) GetFrame() int {
        return this.MovieIndexedAnimationInterface.GetFrame()
}


func (this *CaptionIndexedAnimation) GetSize() int {
        return this.MovieIndexedAnimationInterface.GetSize()
}


func (this *CaptionIndexedAnimation) SetSequence(sequence []int) {
}


func (this *CaptionIndexedAnimation) GetSequence() []int {
        return PrimitiveIntUtil_GetArrayInstance()
}


func (this *CaptionIndexedAnimation) PaintXY(graphics IGraphics, x int, y int) {
        this.AnimationInterface.PaintXY(graphics, x + this.CaptionDx, y + this.CaptionDy)
        this.MovieIndexedAnimationInterface.PaintXY(graphics, x + this.CaptionDx + this.Dx, y + this.CaptionDy + this.Dy)
}


func (this *CaptionIndexedAnimation) PaintThreedXYZ(graphics IGraphics, x int, y int, z int) {
        this.AnimationInterface.PaintThreedXYZ(graphics, x + this.CaptionDx, y + this.CaptionDy, z)
        this.MovieIndexedAnimationInterface.PaintThreedXYZ(graphics, x + this.CaptionDx + this.Dx, y + this.CaptionDy + this.Dy, z)
}

