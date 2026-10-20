/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2022 AllBinary 
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

type LazyImageRotationAnimation struct {
    RotationAnimation
    LogUtil *LogUtil
    LayoutIndex int
    InstanceId int
    AnimationInterfaceFactoryInterface IBaseImageAnimationFactory
    NULL_INDEX_ANIMATION IIndexedAnimation
    Animation IIndexedAnimation
    ScaleProperties *ScaleProperties
}

var LazyImageRotationAnimation_SET_REAL_ANIMATION string = "setRealAnimation"

func (this *LazyImageRotationAnimation) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
    
    
    
    this.ScaleProperties = ScaleProperties_Instance
}


func NewLazyImageRotationAnimation(layoutIndex int, instanceId int, scaleProperties *ScaleProperties, animationInterfaceFactoryInterface IBaseImageAnimationFactory, animationBehavior IAnimationBehavior) *LazyImageRotationAnimation {
    this := &LazyImageRotationAnimation{}
        this.RotationAnimation = *NewRotationAnimation(AngleInfo_GetInstance(AngleFactory_GetInstance().QUARTER_TOTAL_ANGLE), CircularIndexUtil_CreateInstance(4), animationBehavior)
        this.initFields()
        this.LayoutIndex = layoutIndex
        this.InstanceId = instanceId
        this.AnimationInterfaceFactoryInterface = animationInterfaceFactoryInterface
        var imageCache IImageCache = ImageCacheFactory_GetInstance(); _ = imageCache
        imageCache.Add(this)
        this.ScaleProperties = scaleProperties
        this.NULL_INDEX_ANIMATION = any(NullRotationAnimationFactory_GetFactoryInstance().GetInstance(0)).(IIndexedAnimation)
        /* Java local class: move it out of the function
type ActualRotationAnimation struct {
    RotationAnimation
    LazyImageRotationAnimation *LazyImageRotationAnimation
    Index int
}


func (this *ActualRotationAnimation) initFields() {
            
            
}
        

        func NewLazyImageRotationAnimation(lazyImageRotationAnimation *LazyImageRotationAnimation) *LazyImageRotationAnimation {
            this := &LazyImageRotationAnimation{}
                this.RotationAnimation = *NewRotationAnimation(AngleInfo_GetInstance(AngleFactory_GetInstance().QUARTER_TOTAL_ANGLE), CircularIndexUtil_CreateInstance(4), this.AnimationBehavior)
                this.initFields()
                this.LazyImageRotationAnimation = lazyImageRotationAnimation
            return this
        }
        

        func (this *LazyImageRotationAnimation) SetFrame(index int) {
                this.Index = index
        }
        

        func (this *LazyImageRotationAnimation) GetFrame() int {
                return this.Index
        }
        

        func (this *LazyImageRotationAnimation) PaintXY(graphics IGraphics, x int, y int) {
                if tryControl1 := func() (tryControl1 int) {
                    defer func() {
                        if recovered := recover(); recovered != nil {
                            e := Java_ToError(recovered)
                            _ = e
                    var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
                    logUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.PROCESS, e)
                        }
                    }()
                    ImageCacheFactory_GetInstance().InsertFirst(this)
                    this.LazyImageRotationAnimation.Animation = this.LazyImageRotationAnimation.NULL_INDEX_ANIMATION
                    return
                }(); tryControl1 == 1 {
                    return
                }
        }
        

        func (this *LazyImageRotationAnimation) PaintThreedXYZ(graphics IGraphics, x int, y int, z int) {
                if tryControl1 := func() (tryControl1 int) {
                    defer func() {
                        if recovered := recover(); recovered != nil {
                            e := Java_ToError(recovered)
                            _ = e
                    var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
                    logUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.PROCESS, e)
                        }
                    }()
                    ImageCacheFactory_GetInstance().InsertFirst(this)
                    this.LazyImageRotationAnimation.Animation = this.LazyImageRotationAnimation.NULL_INDEX_ANIMATION
                    return
                }(); tryControl1 == 1 {
                    return
                }
        }
*/
        this.Animation = NewActualRotationAnimation(this)
    return this
}


func (this *LazyImageRotationAnimation) SetRealAnimation() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, LazyImageRotationAnimation_SET_REAL_ANIMATION, e)
                }
            }()
            var animation IIndexedAnimation = this.Animation; _ = animation
            this.AnimationInterfaceFactoryInterface.SetInitialScale(this.ScaleProperties)
            this.Animation = any(this.AnimationInterfaceFactoryInterface.GetInstance(this.InstanceId)).(IIndexedAnimation)
            this.Animation.SetState(animation)
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *LazyImageRotationAnimation) SetScale(scaleX float32, scaleY float32) {
        this.Animation.SetScale(scaleX, scaleY)
}


func (this *LazyImageRotationAnimation) GetAnimationBehavior() IAnimationBehavior {
        return this.Animation.GetAnimationBehavior()
}


// throws Exception
func (this *LazyImageRotationAnimation) Set(gl GL) {
        this.Animation.Set(gl)
}


func (this *LazyImageRotationAnimation) SetAlpha(alpha int) {
        this.Animation.SetAlpha(alpha)
}


func (this *LazyImageRotationAnimation) SetDx(dx int) {
        this.Animation.SetDx(dx)
}


func (this *LazyImageRotationAnimation) SetDy(dy int) {
        this.Animation.SetDy(dy)
}


func (this *LazyImageRotationAnimation) SetMaxScale(maxScaleX float32, maxScaleY float32) {
        this.Animation.SetMaxScale(maxScaleX, maxScaleY)
}


func (this *LazyImageRotationAnimation) NextRotation() {
        var rotationAnimation IRotationAnimation = (any(this.Animation).(IRotationAnimation)); _ = rotationAnimation
        rotationAnimation.NextRotation()
}


func (this *LazyImageRotationAnimation) PreviousRotation() {
        var rotationAnimation IRotationAnimation = (any(this.Animation).(IRotationAnimation)); _ = rotationAnimation
        rotationAnimation.PreviousRotation()
}


func (this *LazyImageRotationAnimation) NextRotationX() {
        var rotationAnimation IRotationAnimation = (any(this.Animation).(IRotationAnimation)); _ = rotationAnimation
        rotationAnimation.NextRotationX()
}


func (this *LazyImageRotationAnimation) PreviousRotationX() {
        var rotationAnimation IRotationAnimation = (any(this.Animation).(IRotationAnimation)); _ = rotationAnimation
        rotationAnimation.PreviousRotationX()
}


func (this *LazyImageRotationAnimation) NextRotationZ() {
        var rotationAnimation IRotationAnimation = (any(this.Animation).(IRotationAnimation)); _ = rotationAnimation
        rotationAnimation.NextRotationZ()
}


func (this *LazyImageRotationAnimation) PreviousRotationZ() {
        var rotationAnimation IRotationAnimation = (any(this.Animation).(IRotationAnimation)); _ = rotationAnimation
        rotationAnimation.PreviousRotationZ()
}


func (this *LazyImageRotationAnimation) ChangeBasicColor(basicColor *BasicColor) {
        this.Animation.ChangeBasicColor(basicColor)
}


func (this *LazyImageRotationAnimation) GetBasicColorP() *BasicColor {
        return this.Animation.GetBasicColorP()
}


func (this *LazyImageRotationAnimation) GetChangeBasicColor() *BasicColor {
        return this.Animation.GetChangeBasicColor()
}


func (this *LazyImageRotationAnimation) GetChangeColor() int {
        return this.Animation.GetChangeColor()
}


func (this *LazyImageRotationAnimation) GetColor() int {
        return this.Animation.GetColor()
}


func (this *LazyImageRotationAnimation) GetDx() int {
        return this.Animation.GetDx()
}


func (this *LazyImageRotationAnimation) GetDy() int {
        return this.Animation.GetDy()
}


func (this *LazyImageRotationAnimation) IsThreed() bool {
        return this.Animation.IsThreed()
}


// throws Exception
func (this *LazyImageRotationAnimation) NextFrame() {
        this.Animation.NextFrame()
}


func (this *LazyImageRotationAnimation) Reset() {
        this.Animation.Reset()
}


func (this *LazyImageRotationAnimation) SetFrame(index int) {
        this.Animation.SetFrame(index)
}


func (this *LazyImageRotationAnimation) GetFrame() int {
        return this.Animation.GetFrame()
}


// throws Exception
func (this *LazyImageRotationAnimation) GetAnimationSize() int {
        return this.Animation.GetAnimationSize()
}


func (this *LazyImageRotationAnimation) GetSize() int {
        return this.Animation.GetSize()
}


func (this *LazyImageRotationAnimation) PreviousFrame() {
        this.Animation.PreviousFrame()
}


func (this *LazyImageRotationAnimation) IsLastFrame() bool {
        return this.Animation.IsLastFrame()
}


func (this *LazyImageRotationAnimation) SetSequence(sequence []int) {
        this.Animation.SetSequence(sequence)
}


func (this *LazyImageRotationAnimation) GetSequence() []int {
        return this.Animation.GetSequence()
}


func (this *LazyImageRotationAnimation) GetWidth() int {
        return this.Animation.GetWidth()
}


func (this *LazyImageRotationAnimation) PaintXY(graphics IGraphics, x int, y int) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.PROCESS, e)
                }
            }()
            this.Animation.PaintXY(graphics, x, y)
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *LazyImageRotationAnimation) PaintThreedXYZ(graphics IGraphics, x int, y int, z int) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.PROCESS, e)
                }
            }()
            this.Animation.PaintThreedXYZ(graphics, x, y, z)
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *LazyImageRotationAnimation) ToString() string {
        var commonSeps *CommonSeps = CommonSeps_GetInstance(); _ = commonSeps
        var image IImage = this.AnimationInterfaceFactoryInterface.GetImage(); _ = image
        return fmt.Sprint(NewStringMaker().Append(this.RotationAnimation.ToString()).Append(commonSeps.SPACE).Append(image.GetName()).Append(commonSeps.SPACE).Appendint(int(image.GetWidth())).Append(commonSeps.SPACE).Appendint(int(image.GetHeight())))
}

func (this *LazyImageRotationAnimation) String() string {
    return this.ToString()
}

