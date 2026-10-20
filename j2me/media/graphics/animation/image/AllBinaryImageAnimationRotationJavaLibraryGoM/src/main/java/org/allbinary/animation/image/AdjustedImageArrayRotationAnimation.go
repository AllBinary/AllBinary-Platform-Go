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

//TWB - Adjustments should be done in the resource creation and not at the animation level
type AdjustedImageArrayRotationAnimation struct {
    ImageArrayRotationAnimation
    Dx int
    Dy int
}

// The Java class AdjustedImageArrayRotationAnimation has subclasses so its variables are this interface
type IAdjustedImageArrayRotationAnimation interface {
    IImageArrayRotationAnimation
    AdjustedImageArrayRotationAnimation_() *AdjustedImageArrayRotationAnimation
    Init(dx int, dy int)
    PaintXY(graphics IGraphics, x int, y int)
    SetDx(dx int)
    GetDx() int
    SetDy(dy int)
    GetDy() int
}

func (this *AdjustedImageArrayRotationAnimation) AdjustedImageArrayRotationAnimation_() *AdjustedImageArrayRotationAnimation {
    return this
}


func (this *AdjustedImageArrayRotationAnimation) initFields() {
    
    
}


// throws Exception
func AdjustedImageArrayRotationAnimation_CreateAdjustedImageArrayRotationAnimationInfo(anyType any, animationBehavior IAnimationBehavior) IAdjustedImageArrayRotationAnimation {
        var imageRotationAnimationInfo *ImageArrayRotationAnimationInfo = any(anyType).(*ImageArrayRotationAnimationInfo); _ = imageRotationAnimationInfo
        return NewAdjustedImageArrayRotationAnimation(imageRotationAnimationInfo.GetImageArray(), imageRotationAnimationInfo.GetAngleInfoP(), imageRotationAnimationInfo.GetTotalAngle(), imageRotationAnimationInfo.GetDx(), imageRotationAnimationInfo.GetDy(), animationBehavior)
}


// throws Exception
func AdjustedImageArrayRotationAnimation_CreateAnimation(imageArray []IImage, angleInfo *AngleInfo, totalAngle int, animationBehavior IAnimationBehavior) IAdjustedImageArrayRotationAnimation {
        return NewAdjustedImageArrayRotationAnimation(imageArray, angleInfo, totalAngle, -(imageArray[0].GetWidth() >> 2), -(imageArray[0].GetHeight() >> 2), animationBehavior)
}


func NewAdjustedImageArrayRotationAnimation(imageArray []IImage, angleInfo *AngleInfo, totalAngle int, dx int, dy int, animationBehavior IAnimationBehavior) *AdjustedImageArrayRotationAnimation {
    this := &AdjustedImageArrayRotationAnimation{}
        this.ImageArrayRotationAnimation = *NewImageArrayRotationAnimation(imageArray, angleInfo, totalAngle, animationBehavior)
        this.initFields()
        this.Init(dx, dy)
    return this
}


// throws Exception
func (this *AdjustedImageArrayRotationAnimation) Init(dx int, dy int) {
        this.SetDx(dx)
        this.SetDy(dy)
}


func (this *AdjustedImageArrayRotationAnimation) PaintXY(graphics IGraphics, x int, y int) {
        this.ImageArrayRotationAnimation.PaintXY(graphics, x + this.Dx, y + this.Dy)
}


func (this *AdjustedImageArrayRotationAnimation) SetDx(dx int) {
        this.Dx = dx
}


func (this *AdjustedImageArrayRotationAnimation) GetDx() int {
        return this.Dx
}


func (this *AdjustedImageArrayRotationAnimation) SetDy(dy int) {
        this.Dy = dy
}


func (this *AdjustedImageArrayRotationAnimation) GetDy() int {
        return this.Dy
}

