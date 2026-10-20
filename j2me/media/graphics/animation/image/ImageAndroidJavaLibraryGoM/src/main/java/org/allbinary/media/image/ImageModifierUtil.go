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

type ImageModifierUtil struct {
    Paint *Paint
}


func (this *ImageModifierUtil) initFields() {
    this.Paint = NewPaint(Paint_ANTI_ALIAS_FLAG)
}

func NewImageModifierUtil() *ImageModifierUtil {
    this := &ImageModifierUtil{}
    this.initFields()
    return this
}


func ImageModifierUtil_GetInstanceOrCreate() *ImageModifierUtil {
        return NewImageModifierUtil()
}


func (this *ImageModifierUtil) SetColor(unusedOriginalImage IImage, image IImage, imageIndex int, basicColor *BasicColor) {
        this.Paint.SetColorFilter(NewPorterDuffColorFilter(basicColor.IntValue(), PorterDuff_Mode.SRC_IN))
}


func (this *ImageModifierUtil) ChangeColor(unusedOriginalImage IImage, image IImage, imageIndex int, basicColor *BasicColor) {
        this.Paint.SetColorFilter(NewPorterDuffColorFilter(basicColor.IntValue(), PorterDuff_Mode.SRC_IN))
}


func (this *ImageModifierUtil) SetAlpha(originalImage IImage, image IImage, imageIndex int, alphaInt int) {
        if alphaInt != this.Paint.GetAlpha() {
            this.Paint.SetAlpha(alphaInt)
        }
}


func (this *ImageModifierUtil) GetImageArray(originalImageArray []IImage) []IImage {
        return originalImageArray
}


func (this *ImageModifierUtil) HandleImage(imageArray []IImage, index int, image IImage) {
}


func (this *ImageModifierUtil) Reset() {
}

