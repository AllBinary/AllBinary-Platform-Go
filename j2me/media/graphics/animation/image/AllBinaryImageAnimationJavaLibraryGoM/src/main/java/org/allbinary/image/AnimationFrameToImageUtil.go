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

type AnimationFrameToImageUtil struct {
}

var AnimationFrameToImageUtil_Instance *AnimationFrameToImageUtil = NewAnimationFrameToImageUtil()

func NewAnimationFrameToImageUtil() *AnimationFrameToImageUtil {
    this := &AnimationFrameToImageUtil{}
    return this
}


func AnimationFrameToImageUtil_GetInstance() *AnimationFrameToImageUtil {
        return AnimationFrameToImageUtil_Instance
}


// throws Exception
func (this *AnimationFrameToImageUtil) GetInstanceTranslate(width int, height int, animationInterface IAnimation) IImage {
        var image IImage = GameFeatureImageCacheFactory_GetInstance().Get(reflect.TypeOf(AnimationFrameToImageUtil_Instance).String(), ((width * 3) >> 1), ((height * 3) >> 1)); _ = image
        var graphics IGraphics = image.GetGraphics(); _ = graphics
        var tranlateX int = (width >> 2); _ = tranlateX
        var tranlateY int = (height >> 2); _ = tranlateY
        graphics.Translate(tranlateX, tranlateY)
        animationInterface.PaintXY(graphics, 0, 0)
        graphics.Translate(-tranlateX, -tranlateY)
        return image
}


// throws Exception
func (this *AnimationFrameToImageUtil) GetInstanceWH(width int, height int, animationInterface IAnimation) IImage {
        var image IImage = GameFeatureImageCacheFactory_GetInstance().Get(reflect.TypeOf(AnimationFrameToImageUtil_Instance).String(), width, height); _ = image
        var graphics IGraphics = image.GetGraphics(); _ = graphics
        animationInterface.PaintXY(graphics, 0, 0)
        return image
}

