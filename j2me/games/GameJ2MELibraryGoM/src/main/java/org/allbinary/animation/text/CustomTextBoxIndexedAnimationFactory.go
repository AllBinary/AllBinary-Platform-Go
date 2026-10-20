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

type CustomTextBoxIndexedAnimationFactory struct {
    LogUtil *LogUtil
    ScaleProperties *ScaleProperties
    MaxLength int
}

// Implements AnimationInterfaceFactoryInterface


func (this *CustomTextBoxIndexedAnimationFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
}


func NewCustomTextBoxIndexedAnimationFactory(fontSize int, maxLength int) *CustomTextBoxIndexedAnimationFactory {
    this := &CustomTextBoxIndexedAnimationFactory{}
    this.initFields()
        this.ScaleProperties = NewScaleProperties()
        this.ScaleProperties.ScaleHeight = fontSize
        this.MaxLength = maxLength
    return this
}


// throws Exception
func (this *CustomTextBoxIndexedAnimationFactory) GetInstance(instanceId int) IAnimation {
        var abToGBUtil *ABToGBUtil = ABToGBUtil_GetInstance(); _ = abToGBUtil
        var abCanvas IAllBinaryGameCanvas = any(abToGBUtil.AbCanvas).(IAllBinaryGameCanvas); _ = abCanvas
        var font Font = Font_GetFont(Font_FACE_SYSTEM, Font_STYLE_PLAIN, this.ScaleProperties.ScaleHeight); _ = font
        var customTextBox ICustomTextBox = NewCustomTextBox2(abCanvas.GetCustomCommandListener(), StringUtil_GetInstance().EMPTY_STRING, StringUtil_GetInstance().EMPTY_STRING, this.MaxLength, TextField_ANY, font, abCanvas.GetLayerManager().GetBackgroundBasicColor(), abCanvas.GetLayerManager().GetForegroundBasicColor()); _ = customTextBox
        return NewCustomTextBoxIndexedAnimation(customTextBox)
}


func (this *CustomTextBoxIndexedAnimationFactory) SetInitialScale(scaleProperties *ScaleProperties) {
        var fontSize int = this.ScaleProperties.ScaleHeight; _ = fontSize
        this.ScaleProperties = scaleProperties
        if this.ScaleProperties.ScaleHeight < 6 {
            this.ScaleProperties.ScaleHeight = fontSize
        }
}

