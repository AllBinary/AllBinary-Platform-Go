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

type ImageArrayRotationAnimation struct {
    ImageArrayBaseRotationAnimation
    ExpectedTotalFrames int
}

// The Java class ImageArrayRotationAnimation has subclasses so its variables are this interface
type IImageArrayRotationAnimation interface {
    IImageArrayBaseRotationAnimation
    ImageArrayRotationAnimation_() *ImageArrayRotationAnimation
    Init(imageArray []IImage, angleInfo *AngleInfo, totalAngle int)
    SetImageArray(imageArray []IImage)
}

func (this *ImageArrayRotationAnimation) ImageArrayRotationAnimation_() *ImageArrayRotationAnimation {
    return this
}


func (this *ImageArrayRotationAnimation) initFields() {
    
}


// throws Exception
func ImageArrayRotationAnimation_CreateAnimation(anyType any, animationBehavior IAnimationBehavior) IImageArrayRotationAnimation {
        var imageRotationAnimationInfo *ImageArrayRotationAnimationInfo = any(anyType).(*ImageArrayRotationAnimationInfo); _ = imageRotationAnimationInfo
        return NewImageArrayRotationAnimation(imageRotationAnimationInfo.GetImageArray(), imageRotationAnimationInfo.GetAngleInfoP(), imageRotationAnimationInfo.GetTotalAngle(), animationBehavior)
}


func NewImageArrayRotationAnimation(imageArray []IImage, angleInfo *AngleInfo, totalAngle int, animationBehavior IAnimationBehavior) *ImageArrayRotationAnimation {
    this := &ImageArrayRotationAnimation{}
        this.ImageArrayBaseRotationAnimation = *NewImageArrayBaseRotationAnimation(imageArray, angleInfo, animationBehavior)
        this.initFields()
        this.Init(imageArray, angleInfo, totalAngle)
    return this
}


// throws Exception
func (this *ImageArrayRotationAnimation) Init(imageArray []IImage, angleInfo *AngleInfo, totalAngle int) {
        this.ExpectedTotalFrames = totalAngle / int(this.AngleInfo.GetAngleIncrementInfo().GetAngleIncrement())
        this.AngleInfo.AdjustAngle(0)
        if this.ExpectedTotalFrames != this.GetSize() {
            panic(errors.New(fmt.Sprint("Wrong Number of Frames")))
        }
}


func (this *ImageArrayRotationAnimation) SetImageArray(imageArray []IImage) {
        this.ImageArrayBaseRotationAnimation.SetImageArray(imageArray)
        this.AngleInfo.AdjustAngle(0)
}

