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

type CustomTextAnimationFactory struct {
    LogUtil *LogUtil
    AnimationBehaviorFactory IAnimationBehaviorFactory
    BasicColor *BasicColor
    Text string
    InitScaleHeight int
    Dx int
    Dy int
    Font Font
    ScaleProperties *ScaleProperties
}

// Implements AnimationInterfaceFactoryInterface


func (this *CustomTextAnimationFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    this.BasicColor = BasicColorFactory_GetInstance().BLACK
    
    
    this.Dx = 0
    this.Dy = 0
    
    
}


func CustomTextAnimationFactory_CreateDXY(text string, fontSize int, dx int, dy int, animationBehaviorFactory IAnimationBehaviorFactory) *CustomTextAnimationFactory {
        var customTextAnimationFactory *CustomTextAnimationFactory = NewCustomTextAnimationFactory(text, fontSize, animationBehaviorFactory); _ = customTextAnimationFactory
        customTextAnimationFactory.Dx = dx
        customTextAnimationFactory.Dy = dy
        return customTextAnimationFactory
}


func NewCustomTextAnimationFactory(text string, fontSize int, animationBehaviorFactory IAnimationBehaviorFactory) *CustomTextAnimationFactory {
    this := &CustomTextAnimationFactory{}
    this.initFields()
        this.ScaleProperties = NewScaleProperties()
        this.Text = text
        this.ScaleProperties.ScaleHeight = fontSize - (fontSize / 4)
        this.InitScaleHeight = this.ScaleProperties.ScaleHeight
        this.AnimationBehaviorFactory = animationBehaviorFactory
        this.Font = Font_GetFont(Font_FACE_SYSTEM, Font_STYLE_PLAIN, this.InitScaleHeight)
    return this
}


// throws Exception
func (this *CustomTextAnimationFactory) GetInstance(instanceId int) IAnimation {
        var customTextAnimation ICustomTextAnimation; _ = customTextAnimation
        if this.Dx != 0 || this.Dy != 0 {
            customTextAnimation = NewAdjustCustomTextAnimation(this.Text, this.ScaleProperties.ScaleHeight, this.Dx, this.Dy, this.AnimationBehaviorFactory.GetOrCreateInstance())
        } else {
            customTextAnimation = NewCustomTextAnimation(this.Text, this.ScaleProperties.ScaleHeight, this.AnimationBehaviorFactory.GetOrCreateInstance())
        }
        customTextAnimation.SetBasicColorP(this.BasicColor)
        return customTextAnimation
}


func (this *CustomTextAnimationFactory) SetInitialScale(scaleProperties *ScaleProperties) {
}


func (this *CustomTextAnimationFactory) GetWidth() int {
        return this.Font.StringWidth(this.Text)
}


func (this *CustomTextAnimationFactory) GetHeight() int {
        return this.Font.GetHeight()
}

