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
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    ResourceCallbackStrings *ResourceCallbackStrings
    PlaynUtil *HTMLPlaynUtil
    AlphaArray []bool
}


func (this *ImageModifierUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.ResourceCallbackStrings = ResourceCallbackStrings_GetInstance()
    this.PlaynUtil = HTMLPlaynUtil_GetInstance()
    
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
}


func (this *ImageModifierUtil) ChangeColor(unusedOriginalImage IImage, image IImage, imageIndex int, basicColor *BasicColor) {
}


func (this *ImageModifierUtil) SetAlpha(originalImage IImage, image IImage, imageIndex int, alpha int) {
        var alphaF float32 = float32(alpha); _ = alphaF
        var alphaFloat float32 = alphaF / 255.0; _ = alphaFloat
        this.SetAlphaF(originalImage, image, imageIndex, alphaFloat)
}


func (this *ImageModifierUtil) SetAlphaF(originalImage IImage, image IImage, imageIndex int, alpha float32) {
        if this.AlphaArray[imageIndex] {
            this.AlphaArray[imageIndex] = false
            this.SetAlpha2(originalImage, image, imageIndex, alpha)
        }
}


func (this *ImageModifierUtil) SetAlpha2(originalImage IImage, image IImage, imageIndex int, alpha float32) {
        var htmlImage *PlaynImage = any(image).(*PlaynImage); _ = htmlImage
        var canvasImage *ImageImpl = any(htmlImage.GetImage()).(*ImageImpl); _ = canvasImage
        var originalPlaynImage *CoreImage = any((any(originalImage).(*PlaynImage)).GetImage()).(*CoreImage); _ = originalPlaynImage
        var canvas ICanvas = this.PlaynUtil.GetCanvas(canvasImage); _ = canvas
        canvas.Clear()
        canvas.SetAlpha(alpha)
        canvas.Draw(originalPlaynImage, 0, 0)
}


func (this *ImageModifierUtil) SetAlpha3(image IImage, alpha float32) {
        var alphaF float32 = alpha; _ = alphaF
        var alphaFloat float32 = alphaF / 255.0; _ = alphaFloat
        var htmlImage *PlaynImage = any(image).(*PlaynImage); _ = htmlImage
        var canvasImage *ImageImpl = any(htmlImage.GetImage()).(*ImageImpl); _ = canvasImage
        var canvas ICanvas = this.PlaynUtil.GetCanvas(canvasImage); _ = canvas
        canvas.SetAlpha(alphaFloat)
}


func (this *ImageModifierUtil) GetImageArray(originalImageArray []IImage) []IImage {
        var size int = len(originalImageArray); _ = size
        this.AlphaArray = make([]bool, size)
        var imageArray []IImage = make([]IImage, size); _ = imageArray
        for index := int(0); index < size; index++ {
            imageArray[index] = originalImageArray[index]
            this.HandleImage(imageArray, index, originalImageArray[index])
        }
        return imageArray
}


func (this *ImageModifierUtil) HandleImage(imageArray []IImage, index int, image IImage) {
        var image3 *CoreImage = any((any(image).(*PlaynImage)).GetImage()).(*CoreImage); _ = image3
        if image3 != nil {
            if image3.IsReady() || image3.Width() + image3.Height() <= 0 || image.GetName() == this.ResourceCallbackStrings.FROM_DATA {
                this.Copy(imageArray, index, image, image3)
            } else {
                var callback *Callback = /* Java anonymous class body: implement it as a named type
                    

                func (this *ImageModifierUtil) OnSuccess(resource any) {
                        var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
                        logUtil.PutF(fmt.Sprint(this.ResourceCallbackStrings.DONE) + image.GetName(), this, this.ResourceCallbackStrings.HANDLE_IMAGE)
                        this.Copy(imageArray, index, image, image3)
                }
                    

                func (this *ImageModifierUtil) OnFailure(e error) {
                        var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
                        logUtil.Put(fmt.Sprint(NewStringMaker().Append(this.CommonStrings.EXCEPTION_LABEL).Append(this.ResourceCallbackStrings.ERROR).Append(image.GetName())), this, this.ResourceCallbackStrings.HANDLE_IMAGE, e)
                }
*/ NewCallback(); _ = callback
                image3.AddCallback(callback)
            }
        } else {
            this.LogUtil.PutF(this.ResourceCallbackStrings.NULL + image.IsMutable(), this, this.ResourceCallbackStrings.HANDLE_IMAGE)
        }
}


func (this *ImageModifierUtil) Copy(imageArray []IImage, index int, image IImage, image3 *CoreImage) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.PutF(this.CommonStrings.EXCEPTION_LABEL + this.ResourceCallbackStrings.DONE, this, this.ResourceCallbackStrings.HANDLE_IMAGE)
                }
            }()
            var image2 IImage = Image_CreateImageWH(int(image3.Width()), int(image3.Height())); _ = image2
            var graphics IGraphics = image2.GetGraphics(); _ = graphics
            graphics.DrawImage(image, 0, 0, Anchor_TOP_LEFT)
            imageArray[index] = image2
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *ImageModifierUtil) Reset() {
        var size int = len(this.AlphaArray); _ = size
        for index := int(0); index < size; index++ {
            this.AlphaArray[index] = true
        }
}

