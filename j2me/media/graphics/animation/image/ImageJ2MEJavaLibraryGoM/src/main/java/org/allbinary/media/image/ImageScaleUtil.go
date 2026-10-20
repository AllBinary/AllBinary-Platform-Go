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

type ImageScaleUtil struct {
}

var ImageScaleUtil_Instance *ImageScaleUtil = NewImageScaleUtil()


func ImageScaleUtil_GetInstance() *ImageScaleUtil {
        return ImageScaleUtil_Instance
}


func NewImageScaleUtil() *ImageScaleUtil {
    this := &ImageScaleUtil{}
    return this
}


// throws Exception
func (this *ImageScaleUtil) CreateImage2(imageCache IImageCache, originalImage IImage, scaleNominatorX float32, scaleDenominatorX float32, scaleNominatorY float32, scaleDenominatorY float32, cached bool) IImage {
        return this.CreateImage3(imageCache, originalImage, scaleNominatorX, scaleDenominatorX, scaleNominatorY, scaleDenominatorY, cached, true)
}


// throws Exception
func (this *ImageScaleUtil) CreateImage3(imageCache IImageCache, originalImage IImage, scaleNominatorX float32, scaleDenominatorX float32, scaleNominatorY float32, scaleDenominatorY float32, cached bool, mutable bool) IImage {
        return originalImage
}

