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

type ImageCaptionIndexedAnimationFactory struct {
    AnimationFactoryImageScaleUtil *AnimationFactoryImageScaleUtil
    AnimationFactorySpriteScaleUtil *AnimationFactorySpriteScaleUtil
    CaptionImage IImage
    SpriteMovieImage IImage
    FrameWidth int
    FrameHeight int
    CaptionDx int
    CaptionDy int
    Dx int
    Dy int
    Time int
    SoundInterface ISound
    ScaleProperties *ScaleProperties
    AnimationBehaviorFactory IAnimationBehaviorFactory
}

// Implements AnimationInterfaceFactoryInterface


func (this *ImageCaptionIndexedAnimationFactory) initFields() {
    this.AnimationFactoryImageScaleUtil = AnimationFactoryImageScaleUtil_GetInstance()
    this.AnimationFactorySpriteScaleUtil = AnimationFactorySpriteScaleUtil_GetInstance()
    
    
    
    
    
    
    
    
    
    
    this.ScaleProperties = ScaleProperties_Instance
    
}


func NewImageCaptionIndexedAnimationFactory(captionImage IImage, spriteMovieImage IImage, soundInterface ISound, frameWidth int, frameHeight int, captionDx int, captionDy int, dx int, dy int, time int, animationBehaviorFactory IAnimationBehaviorFactory) *ImageCaptionIndexedAnimationFactory {
    this := &ImageCaptionIndexedAnimationFactory{}
    this.initFields()
        if animationBehaviorFactory == AnimationBehaviorFactory_GetInstance() {
            ForcedLogUtil_Log("Using default AnimationBehaviorFactory with IndexedAnimationFactory", this)
        }
        this.CaptionImage = captionImage
        this.SpriteMovieImage = spriteMovieImage
        this.FrameWidth = frameWidth
        this.FrameHeight = frameHeight
        this.CaptionDx = captionDx
        this.CaptionDy = captionDy
        this.Dx = dx
        this.Dy = dy
        this.Time = time
        this.SoundInterface = soundInterface
        this.AnimationBehaviorFactory = animationBehaviorFactory
    return this
}


// throws Exception
func (this *ImageCaptionIndexedAnimationFactory) GetInstance(instanceId int) IAnimation {
        var scaledImage IImage = this.AnimationFactoryImageScaleUtil.CreateImage(this.CaptionImage, this.CaptionImage.GetWidth(), this.CaptionImage.GetHeight(), this.ScaleProperties.ScaleWidth, this.ScaleProperties.ScaleHeight); _ = scaledImage
        var animationInterface IAnimation = NewImageAnimation(scaledImage, this.AnimationBehaviorFactory.GetOrCreateInstance()); _ = animationInterface
        var sprite *Sprite = this.AnimationFactorySpriteScaleUtil.CreateImage(this.SpriteMovieImage, this.FrameWidth, this.FrameHeight, this.ScaleProperties.ScaleWidth, this.ScaleProperties.ScaleHeight); _ = sprite
        var movieIndexedAnimationInterface IIndexedAnimation = NewSpriteIndexedAnimation(sprite, this.SpriteMovieImage, BasicColorUtil_GetInstance().ZERO_ARRAY, this.AnimationBehaviorFactory.GetOrCreateInstance()); _ = movieIndexedAnimationInterface
        var player *Player = this.SoundInterface.GetPlayerP(); _ = player
        if player == nil {
            panic(errors.New(fmt.Sprint("Sound Was not Initialized")))
        }
        return NewCaptionIndexedAnimation(animationInterface, movieIndexedAnimationInterface, player, this.CaptionDx, this.CaptionDy, this.Dx, this.Dy, this.Time, this.AnimationBehaviorFactory.GetOrCreateInstance())
}


func (this *ImageCaptionIndexedAnimationFactory) SetInitialScale(scaleProperties *ScaleProperties) {
        this.ScaleProperties = scaleProperties
}

