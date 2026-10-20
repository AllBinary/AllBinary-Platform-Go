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

type VectorToImageArrayRotationAnimationFactory struct {
    BaseImageAnimationFactory
    ImageArray []IImage
    AngleIncrement int
}


func (this *VectorToImageArrayRotationAnimationFactory) initFields() {
    this.ImageArray = NullImage_NULL_IMAGE_ARRAY
    
}


func NewVectorToImageArrayRotationAnimationFactory(vectorInfo *VectorInfo, basicColor *BasicColor, animationBehaviorFactory IAnimationBehaviorFactory) *VectorToImageArrayRotationAnimationFactory {
    this := &VectorToImageArrayRotationAnimationFactory{}
        this.BaseImageAnimationFactory = *NewBaseImageAnimationFactory(AnimationFrameToImageUtil_GetInstance().GetInstanceTranslate(vectorInfo.GetWidth(), vectorInfo.GetHeight(), VectorAnimation_CreateVectorAnimation(vectorInfo.GetPoints(), basicColor, animationBehaviorFactory.GetOrCreateInstance())), PrimitiveIntUtil_GetArrayInstance(), 0, 0, 0, 0, animationBehaviorFactory)
        this.initFields()
        this.Init()
    return this
}


// throws Exception
func (this *VectorToImageArrayRotationAnimationFactory) Init() {
        var angleFactory *AngleFactory = AngleFactory_GetInstance(); _ = angleFactory
        this.AngleIncrement = int(angleFactory.TOTAL_ANGLE) / GameConfigurationCentral_GetInstance().GetGameControlFidelity()
        this.ImageArray = ImageToRotationImageArrayUtil_GetInstance().Generate(this.GetImage(), this.GetAngleIncrement(), int(angleFactory.TOTAL_ANGLE))
}


// throws Exception
func (this *VectorToImageArrayRotationAnimationFactory) GetInstance(instanceId int) IAnimation {
        if this.AnimationFactoryInitializationVisitor.Dx != 0 || this.AnimationFactoryInitializationVisitor.Dy != 0 {
            return NewAdjustedImageArrayRotationAnimation(this.ImageArray, AngleInfo_GetInstance(int16(this.GetAngleIncrement())), int(AngleFactory_GetInstance().TOTAL_ANGLE), this.AnimationFactoryInitializationVisitor.Dx, this.AnimationFactoryInitializationVisitor.Dy, this.AnimationBehaviorFactory.GetOrCreateInstance())
        } else {
            return AdjustedImageArrayRotationAnimation_CreateAnimation(this.ImageArray, AngleInfo_GetInstance(int16(this.AngleIncrement)), int(AngleFactory_GetInstance().TOTAL_ANGLE), this.AnimationBehaviorFactory.GetOrCreateInstance())
        }
}


func (this *VectorToImageArrayRotationAnimationFactory) GetAngleIncrement() int {
        return this.AngleIncrement
}


func (this *VectorToImageArrayRotationAnimationFactory) SetInitialScale(scaleProperties *ScaleProperties) {
}

