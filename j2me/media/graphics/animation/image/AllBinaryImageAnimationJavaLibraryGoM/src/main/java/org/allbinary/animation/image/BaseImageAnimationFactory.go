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

type BaseImageAnimationFactory struct {
    AnimationFactoryImageScaleUtil *AnimationFactoryImageScaleUtil
    Image IImage
    AnimationBehaviorFactory IAnimationBehaviorFactory
    SequenceArray []int
    AnimationFactoryInitializationVisitor *AnimationFactoryInitializationVisitor
    ScaleProperties *ScaleProperties
}

// Implements AnimationInterfaceFactoryInterface

// The Java class BaseImageAnimationFactory has subclasses so its variables are this interface
type IBaseImageAnimationFactory interface {
    BaseImageAnimationFactory_() *BaseImageAnimationFactory
    GetInstance(instanceId int) IAnimation
    GetImage() IImage
    GetSequenceArray() []int
    SetInitialScale(scaleProperties *ScaleProperties)
    GetAnimationFactoryInitializationVisitorP() *AnimationFactoryInitializationVisitor
    GetScalePropertiesP() *ScaleProperties
    ToString() string
}

func (this *BaseImageAnimationFactory) BaseImageAnimationFactory_() *BaseImageAnimationFactory {
    return this
}


func (this *BaseImageAnimationFactory) initFields() {
    this.AnimationFactoryImageScaleUtil = AnimationFactoryImageScaleUtil_GetInstance()
    
    
    
    
    this.ScaleProperties = ScaleProperties_Instance
}


// throws Exception
func BaseImageAnimationFactory_CreateFactoryBase(image IImage, sequenceArray []int, width int, height int, animationBehaviorFactory IAnimationBehaviorFactory) IBaseImageAnimationFactory {
        return NewBaseImageAnimationFactory(image, sequenceArray, width, height, 0, 0, animationBehaviorFactory)
}


func NewBaseImageAnimationFactory(image IImage, sequenceArray []int, width int, height int, dx int, dy int, animationBehaviorFactory IAnimationBehaviorFactory) *BaseImageAnimationFactory {
    this := &BaseImageAnimationFactory{}
    this.initFields()
        this.AnimationFactoryInitializationVisitor = NewAnimationFactoryInitializationVisitor()
        this.Image = image
        this.AnimationFactoryInitializationVisitor.Width = width
        this.AnimationFactoryInitializationVisitor.Height = height
        this.AnimationBehaviorFactory = animationBehaviorFactory
        this.SequenceArray = sequenceArray
        this.AnimationFactoryInitializationVisitor.Dx = dx
        this.AnimationFactoryInitializationVisitor.Dy = dy
    return this
}


// throws Exception
func (this *BaseImageAnimationFactory) GetInstance(instanceId int) IAnimation {
        return NullAnimationFactory_GetFactoryInstance().GetInstance(instanceId)
}


func (this *BaseImageAnimationFactory) GetImage() IImage {
        return this.Image
}


func (this *BaseImageAnimationFactory) GetSequenceArray() []int {
        return this.SequenceArray
}


func (this *BaseImageAnimationFactory) SetInitialScale(scaleProperties *ScaleProperties) {
        this.ScaleProperties = scaleProperties
        if this.ScaleProperties.ShouldScale {
            this.ScaleProperties.ScaleWidth = int((float32(this.AnimationFactoryInitializationVisitor.Width) * this.ScaleProperties.ScaleX))
            this.ScaleProperties.ScaleHeight = int((float32(this.AnimationFactoryInitializationVisitor.Height) * this.ScaleProperties.ScaleY))
        } else {
        }
}


func (this *BaseImageAnimationFactory) GetAnimationFactoryInitializationVisitorP() *AnimationFactoryInitializationVisitor {
        return this.AnimationFactoryInitializationVisitor
}


func (this *BaseImageAnimationFactory) GetScalePropertiesP() *ScaleProperties {
        return this.ScaleProperties
}


func (this *BaseImageAnimationFactory) ToString() string {
        var commonSeps *CommonSeps = CommonSeps_GetInstance(); _ = commonSeps
        var commonLabels *CommonLabels = CommonLabels_GetInstance(); _ = commonLabels
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(this..ToString())
        stringBuffer.Append(commonLabels.HEIGHT_LABEL)
        stringBuffer.Appendint(int(this.Image.GetHeight()))
        stringBuffer.Append(commonSeps.SPACE)
        stringBuffer.Append(fmt.Sprint(this.AnimationFactoryInitializationVisitor))
        return fmt.Sprint(stringBuffer)
}

func (this *BaseImageAnimationFactory) String() string {
    return this.ToString()
}

