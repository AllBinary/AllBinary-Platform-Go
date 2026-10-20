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

type AllBinaryScreenRelationalImageAnimationInterfaceFactory struct {
    BaseImageAnimationFactory
    LastImage IImage
}


func (this *AllBinaryScreenRelationalImageAnimationInterfaceFactory) initFields() {
    this.LastImage = NullImage_NULL_IMAGE
}


func NewAllBinaryScreenRelationalImageAnimationInterfaceFactory(image IImage, animationBehaviorFactory IAnimationBehaviorFactory) *AllBinaryScreenRelationalImageAnimationInterfaceFactory {
    this := &AllBinaryScreenRelationalImageAnimationInterfaceFactory{}
        this.BaseImageAnimationFactory = *NewBaseImageAnimationFactory(image, PrimitiveIntUtil_GetArrayInstance(), 0, 0, 0, 0, animationBehaviorFactory)
        this.initFields()
        if this.LastImage != NullImage_NULL_IMAGE {
            var androidImageInterface AndroidImageInterface = any(this.LastImage).(AndroidImageInterface); _ = androidImageInterface
            androidImageInterface.GetBitmap().Recycle()
        }
        var scale float32 = ScreenRelationalUtil_GetInstance().GetScaleImage(image); _ = scale
        this.LastImage = ImageScaleUtil_GetInstance().CreateImage(GameFeatureImageCacheFactory_GetInstance(), this.GetImage(), scale, scale, false)
    return this
}


// throws Exception
func (this *AllBinaryScreenRelationalImageAnimationInterfaceFactory) GetInstance(instanceId int) IAnimation {
        return NewImageAnimation(this.LastImage, this.AnimationBehaviorFactory.GetOrCreateInstance())
}

