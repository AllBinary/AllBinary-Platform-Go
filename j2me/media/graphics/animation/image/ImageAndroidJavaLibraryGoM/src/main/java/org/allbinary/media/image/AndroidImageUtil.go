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

type AndroidImageUtil struct {
}

var AndroidImageUtil_Instance *AndroidImageUtil = NewAndroidImageUtil()

func NewAndroidImageUtil() *AndroidImageUtil {
    this := &AndroidImageUtil{}
    return this
}


func AndroidImageUtil_GetInstance() *AndroidImageUtil {
        return AndroidImageUtil_Instance
}


func (this *AndroidImageUtil) PaintBitmap(bitmap *Bitmap, originalBitmap *Bitmap, paint *Paint) {
        bitmap.EraseColor(Color_TRANSPARENT)
        var canvas ICanvas = NewCanvas(); _ = canvas
        canvas.SetBitmap(bitmap)
        canvas.DrawBitmap(originalBitmap, 0.0, 0.0, paint)
}


func (this *AndroidImageUtil) Paint(image IImage, originalImage IImage, paint *Paint) {
        var androidImage AndroidImageInterface = any(image).(AndroidImageInterface); _ = androidImage
        androidImage.GetBitmap().EraseColor(Color_TRANSPARENT)
        var canvas ICanvas = androidImage.GetCanvas(); _ = canvas
        var originalAndroidImage AndroidImageInterface = any(originalImage).(AndroidImageInterface); _ = originalAndroidImage
        canvas.DrawBitmap(originalAndroidImage.GetBitmap(), 0.0, 0.0, paint)
}


func (this *AndroidImageUtil) Rotate(image IImage, originalImage IImage, matrix *Matrix, paint *Paint) {
        var androidImage AndroidImageInterface = any(image).(AndroidImageInterface); _ = androidImage
        androidImage.GetBitmap().EraseColor(Color_TRANSPARENT)
        var canvas ICanvas = androidImage.GetCanvas(); _ = canvas
        canvas.Concat(matrix)
        var originalAndroidImage AndroidImageInterface = any(originalImage).(AndroidImageInterface); _ = originalAndroidImage
        canvas.DrawBitmap(originalAndroidImage.GetBitmap(), 0.0, 0.0, paint)
}

