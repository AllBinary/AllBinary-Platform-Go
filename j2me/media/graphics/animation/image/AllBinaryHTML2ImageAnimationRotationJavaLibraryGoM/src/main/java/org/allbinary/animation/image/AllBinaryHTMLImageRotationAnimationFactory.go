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

type AllBinaryHTMLImageRotationAnimationFactory struct {
    BaseImageAnimationFactory
    AngleIncrement int16
}


func (this *AllBinaryHTMLImageRotationAnimationFactory) initFields() {
    
}


// throws Exception
func AllBinaryHTMLImageRotationAnimationFactory_CreateFactoryADXY(image IImage, width int, height int, dx int, dy int, angleIncrement int16, animationBehaviorFactory IAnimationBehaviorFactory) *AllBinaryHTMLImageRotationAnimationFactory {
        var htmlImageRotationAnimationFactory *AllBinaryHTMLImageRotationAnimationFactory = NewAllBinaryHTMLImageRotationAnimationFactory(image, width, height, angleIncrement, animationBehaviorFactory); _ = htmlImageRotationAnimationFactory
        htmlImageRotationAnimationFactory.Init(dx, dy)
        return htmlImageRotationAnimationFactory
}


func NewAllBinaryHTMLImageRotationAnimationFactory(image IImage, width int, height int, angleIncrement int16, animationBehaviorFactory IAnimationBehaviorFactory) *AllBinaryHTMLImageRotationAnimationFactory {
    this := &AllBinaryHTMLImageRotationAnimationFactory{}
        this.BaseImageAnimationFactory = *NewBaseImageAnimationFactory(image, PrimitiveIntUtil_GetArrayInstance(), width, height, 0, 0, animationBehaviorFactory)
        this.initFields()
        this.AngleIncrement = angleIncrement
    return this
}


func (this *AllBinaryHTMLImageRotationAnimationFactory) Init(dx int, dy int) {
        this.AnimationFactoryInitializationVisitor.Dx = dx
        this.AnimationFactoryInitializationVisitor.Dy = dy
        this.AnimationFactoryInitializationVisitor.OriginalDx = dx
        this.AnimationFactoryInitializationVisitor.OriginalDy = dy
}


// throws Exception
func (this *AllBinaryHTMLImageRotationAnimationFactory) GetInstance(instanceId int) IAnimation {
        var scaledImage IImage = this.AnimationFactoryImageScaleUtil.CreateImage(this.GetImage(), this.AnimationFactoryInitializationVisitor.Width, this.AnimationFactoryInitializationVisitor.Height, this.ScaleProperties.ScaleWidth, this.ScaleProperties.ScaleHeight); _ = scaledImage
        var copyOfScaledImage IImage = ImageCopyUtil_GetInstance().CreateImage(scaledImage); _ = copyOfScaledImage
        if this.AnimationFactoryInitializationVisitor.Dx != 0 || this.AnimationFactoryInitializationVisitor.Dy != 0 {
            this.AnimationFactoryImageScaleUtil.ProcessAdjust(this)
            return NewAllBinaryAdjustedHTMLImageRotationAnimation(scaledImage, copyOfScaledImage, AngleInfo_GetInstance(this.GetAngleIncrement()), AngleFactory_GetInstance().TOTAL_ANGLE, this.AnimationFactoryInitializationVisitor.Dx, this.AnimationFactoryInitializationVisitor.Dy, this.AnimationBehaviorFactory.GetOrCreateInstance())
        } else {
            return NewAllBinaryHTMLImageRotationAnimation(scaledImage, copyOfScaledImage, AngleInfo_GetInstance(this.AngleIncrement), AngleFactory_GetInstance().TOTAL_ANGLE, this.AnimationBehaviorFactory.GetOrCreateInstance())
        }
}


func (this *AllBinaryHTMLImageRotationAnimationFactory) GetAngleIncrement() int16 {
        return this.AngleIncrement
}

