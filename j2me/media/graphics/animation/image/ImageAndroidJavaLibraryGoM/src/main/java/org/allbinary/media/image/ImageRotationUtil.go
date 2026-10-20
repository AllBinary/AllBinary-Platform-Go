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

type ImageRotationUtil struct {
    Matrix *Matrix
    Anchor int
}

var ImageRotationUtil_Instance *ImageRotationUtil = NewImageRotationUtil()

func (this *ImageRotationUtil) initFields() {
    this.Matrix = NewMatrix()
    this.Anchor = Anchor_TOP_LEFT
}


func ImageRotationUtil_GetInstance() *ImageRotationUtil {
        return ImageRotationUtil_Instance
}


func NewImageRotationUtil() *ImageRotationUtil {
    this := &ImageRotationUtil{}
    this.initFields()
    return this
}


// throws Exception
func (this *ImageRotationUtil) CreateRotatedImage(originalImage IImage, rotationInDegrees int) IImage {
        var image IImage = ImageCreationUtil_GetInstance().CreateImageWH(originalImage.GetWidth(), originalImage.GetHeight()); _ = image
        if image.IsMutable() {
            var bitmap *Bitmap = (any(image).(AndroidImageInterface)).GetBitmap(); _ = bitmap
            var width int = bitmap.GetWidth(); _ = width
            var height int = bitmap.GetHeight(); _ = height
            this.Matrix.SetRotate(float32(rotationInDegrees), float32((width >> 1)), float32((height >> 1)))
            var canvas ICanvas = (any(image).(AndroidImageInterface)).GetCanvas(); _ = canvas
            canvas.Concat(this.Matrix)
            image.GetGraphics().DrawImage(originalImage, 0, 0, this.Anchor)
            return image
        } else {
            panic(errors.New(fmt.Sprint("Not Mutable")))
        }
}

