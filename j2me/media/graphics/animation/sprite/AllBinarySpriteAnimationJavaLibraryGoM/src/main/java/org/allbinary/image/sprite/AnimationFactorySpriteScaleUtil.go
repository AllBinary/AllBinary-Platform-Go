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

type AnimationFactorySpriteScaleUtil struct {
    ImageCache IImageCache
    ImageScaleUtil *ImageScaleUtil
    OpenGLUtil *OpenGLUtil
    J2seMath *J2SEMath
    OpenGLESImageExclusionUtil *OpenGLESImageExclusionUtil
}

var AnimationFactorySpriteScaleUtil_Instance *AnimationFactorySpriteScaleUtil = NewAnimationFactorySpriteScaleUtil()

func (this *AnimationFactorySpriteScaleUtil) initFields() {
    this.ImageCache = ImageCacheFactory_GetInstance()
    this.ImageScaleUtil = ImageScaleUtil_GetInstance()
    this.OpenGLUtil = OpenGLUtil_GetInstance()
    this.J2seMath = J2SEMath_GetInstance()
    this.OpenGLESImageExclusionUtil = OpenGLESImageExclusionUtil_GetInstance()
}

func NewAnimationFactorySpriteScaleUtil() *AnimationFactorySpriteScaleUtil {
    this := &AnimationFactorySpriteScaleUtil{}
    this.initFields()
    return this
}


func AnimationFactorySpriteScaleUtil_GetInstance() *AnimationFactorySpriteScaleUtil {
        return AnimationFactorySpriteScaleUtil_Instance
}


// throws Exception
func (this *AnimationFactorySpriteScaleUtil) CreateImage(image IImage, width int, height int, scaleWidth int, scaleHeight int) *Sprite {
        var spriteFactory *SpriteFactory = SpriteFactory_GetInstance(); _ = spriteFactory
        var sprite *Sprite; _ = sprite
        var scaledImage IImage; _ = scaledImage
        if scaleWidth != 0 && scaleHeight != 0 {
            var scaleX float32 = (float32(scaleWidth)) / (float32(width)); _ = scaleX
            var scaleY float32 = (float32(scaleHeight)) / (float32(height)); _ = scaleY
            if (scaleX == 1.0 && scaleY == 1.0) || (scaleX == 0.0 || scaleY == 0.0) {
                scaledImage = this.OpenGLUtil.AddImage(image)
                sprite = spriteFactory.CreateSprite(scaledImage, width, height)
            } else {
                if this.OpenGLESImageExclusionUtil.IsCustomScaling(image) {
                    var width2 int = this.J2seMath.Round(float32((scaleWidth)) - 0.5); _ = width2
                    var height2 int = this.J2seMath.Round(float32((scaleHeight)) - 0.5); _ = height2
                    var multiplesOf16Width int = width2 / 16; _ = multiplesOf16Width
                    var by16Width int = multiplesOf16Width * 16; _ = by16Width
                    scaleX = (float32(by16Width)) / float32(width)
                    var multiplesOf16Height int = height2 / 16; _ = multiplesOf16Height
                    var by16Height int = multiplesOf16Height * 16; _ = by16Height
                    scaleY = (float32(by16Height)) / float32(height)
                    if scaleWidth < width {
                        scaleX = scaleX * 2.35
                    }
                    if scaleHeight < height {
                        scaleY = scaleY * 2.35
                    }
                }
                scaledImage = this.ImageScaleUtil.CreateImage2(this.ImageCache, image, scaleX, 1.0, scaleY, 1.0, true)
                sprite = spriteFactory.CreateSprite(scaledImage, int((float32(width) * scaleX)), int((float32(height) * scaleY)))
            }
        } else {
            scaledImage = this.OpenGLUtil.AddImage(image)
            sprite = spriteFactory.CreateSprite(scaledImage, width, height)
        }
        return sprite
}

