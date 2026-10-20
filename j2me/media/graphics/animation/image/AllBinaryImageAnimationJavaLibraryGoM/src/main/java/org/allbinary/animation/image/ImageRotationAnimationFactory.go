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

type ImageRotationAnimationFactory struct {
    BaseImageAnimationFactory
    AngleIncrement int16
}


func (this *ImageRotationAnimationFactory) initFields() {
    
}


func NewImageRotationAnimationFactory(image IImage, width int, height int, dx int, dy int, angleIncrement int16, animationBehaviorFactory IAnimationBehaviorFactory) *ImageRotationAnimationFactory {
    this := &ImageRotationAnimationFactory{}
        this.BaseImageAnimationFactory = *NewBaseImageAnimationFactory(image, PrimitiveIntUtil_GetArrayInstance(), width, height, 0, 0, animationBehaviorFactory)
        this.initFields()
        this.AngleIncrement = angleIncrement
        this.AnimationFactoryInitializationVisitor.Dx = dx
        this.AnimationFactoryInitializationVisitor.Dy = dy
    return this
}


// throws Exception
func (this *ImageRotationAnimationFactory) GetInstance(instanceId int) IAnimation {
        if this.AnimationFactoryInitializationVisitor.Dx != 0 || this.AnimationFactoryInitializationVisitor.Dy != 0 {
            this.AnimationFactoryImageScaleUtil.ProcessAdjust(this)
            return NewAdjustedImageBaseRotationAnimation(this.GetImage(), AngleInfo_GetInstance(this.AngleIncrement), AngleFactory_GetInstance().TOTAL_ANGLE, this.AnimationFactoryInitializationVisitor.Dx, this.AnimationFactoryInitializationVisitor.Dy, this.AnimationBehaviorFactory.GetOrCreateInstance())
        } else {
            return NewImageBaseRotationAnimation(this.GetImage(), AngleInfo_GetInstance(this.AngleIncrement), AngleFactory_GetInstance().TOTAL_ANGLE, this.AnimationBehaviorFactory.GetOrCreateInstance())
        }
}


func (this *ImageRotationAnimationFactory) GetAngleIncrementP() int16 {
        return this.AngleIncrement
}

