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

type LeftToRightImageAnimationFactory struct {
    BaseImageAnimationFactory
}



func NewLeftToRightImageAnimationFactory(image IImage, sequenceArray []int, dx int, dy int, animationBehaviorFactory IAnimationBehaviorFactory) *LeftToRightImageAnimationFactory {
    this := &LeftToRightImageAnimationFactory{}
        this.BaseImageAnimationFactory = *NewBaseImageAnimationFactory(image, sequenceArray, image.GetWidth(), image.GetHeight(), dx, dy, animationBehaviorFactory)
    return this
}


// throws Exception
func (this *LeftToRightImageAnimationFactory) GetInstance(instanceId int) IAnimation {
        var scaledImage IImage = this.AnimationFactoryImageScaleUtil.CreateImage(this.GetImage(), this.AnimationFactoryInitializationVisitor.Width, this.AnimationFactoryInitializationVisitor.Height, this.ScaleProperties.ScaleWidth, this.ScaleProperties.ScaleHeight); _ = scaledImage
        if this.AnimationFactoryInitializationVisitor.Dx != 0 || this.AnimationFactoryInitializationVisitor.Dy != 0 {
            this.AnimationFactoryImageScaleUtil.ProcessAdjust(this)
            return NewAdjustedLeftToRightImageAnimation(scaledImage, this.GetSequenceArray(), this.AnimationFactoryInitializationVisitor.Dx, this.AnimationFactoryInitializationVisitor.Dy, this.AnimationBehaviorFactory.GetOrCreateInstance())
        } else {
            return NewLeftToRightImageAnimation(scaledImage, this.GetSequenceArray(), this.AnimationBehaviorFactory.GetOrCreateInstance())
        }
}

