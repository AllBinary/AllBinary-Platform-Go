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

type ImageCreationUtil struct {
}

var ImageCreationUtil_Instance *ImageCreationUtil = NewImageCreationUtil()


func ImageCreationUtil_GetInstance() *ImageCreationUtil {
        return ImageCreationUtil_Instance
}


func NewImageCreationUtil() *ImageCreationUtil {
    this := &ImageCreationUtil{}
    return this
}


// throws Exception
func (this *ImageCreationUtil) CreateImageWH(width int, height int) IImage {
        var image IImage = GameFeatureImageCacheFactory_GetInstance().Get(reflect.TypeOf(this).String(), width, height); _ = image
        return image
}


// throws Exception
func (this *ImageCreationUtil) CreateImage(width int, height int, maxScaleX float32, maxScaleY float32) IImage {
        var image IImage = GameFeatureImageCacheFactory_GetInstance().Get(reflect.TypeOf(this).String(), int((float32(width) * maxScaleX)) + 1, int((float32(height) * maxScaleY)) + 1); _ = image
        return image
}

