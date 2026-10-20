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

type ImageJ2SERotationUtil struct {
    ImageUtil *ImageUtil
    ImageJ2SEUtil *ImageJ2SEUtil
    TWO_PIE float64
}

var ImageJ2SERotationUtil_Instance *ImageJ2SERotationUtil = NewImageJ2SERotationUtil()

func (this *ImageJ2SERotationUtil) initFields() {
    this.ImageUtil = ImageUtil_GetInstance()
    this.ImageJ2SEUtil = ImageJ2SEUtil_GetInstance()
    this.TWO_PIE = 2 * Math_PI
}


func ImageJ2SERotationUtil_GetInstance() *ImageJ2SERotationUtil {
        return ImageJ2SERotationUtil_Instance
}


func NewImageJ2SERotationUtil() *ImageJ2SERotationUtil {
    this := &ImageJ2SERotationUtil{}
    this.initFields()
    return this
}


func (this *ImageJ2SERotationUtil) GetRotatedImage(bufferedImage IImage, totalAngle int) *BufferedImage {
        return this.GetRotatedImageRadians(bufferedImage, this.TWO_PIE * float64(totalAngle) / 360.0)
}


func (this *ImageJ2SERotationUtil) GetRotatedImageRadians(bufferedImage IImage, radians float64) *BufferedImage {
        var newBufferedImage *BufferedImage = this.ImageUtil.Create(bufferedImage.GetWidth(nil), bufferedImage.GetHeight(nil)); _ = newBufferedImage
        return this.DrawRotatedImageRadians(bufferedImage, newBufferedImage, newBufferedImage.CreateGraphics(), radians)
}


func (this *ImageJ2SERotationUtil) DrawRotatedImageRadians(bufferedImage IImage, newBufferedImage *BufferedImage, g *Graphics2D, radians float64) *BufferedImage {
        g.Rotate(radians, newBufferedImage.GetWidth(nil) / 2, newBufferedImage.GetHeight(nil) / 2)
        g.DrawImage(bufferedImage, 0, 0, bufferedImage.GetWidth(nil), bufferedImage.GetHeight(nil), nil)
        g.Dispose()
        return newBufferedImage
}


func (this *ImageJ2SERotationUtil) RotateImage(bufferedImage IImage, newBufferedImage *BufferedImage, totalAngle int) *BufferedImage {
        var g *Graphics2D = newBufferedImage.CreateGraphics(); _ = g
        g.SetBackground(this.ImageJ2SEUtil.TRANSPARENT_COLOR)
        g.ClearRect(0, 0, newBufferedImage.GetWidth(), newBufferedImage.GetHeight())
        return this.DrawRotatedImageRadians(bufferedImage, newBufferedImage, g, this.TWO_PIE * float64(totalAngle) / 360.0)
}


func (this *ImageJ2SERotationUtil) CreateRotatedImage(bufferedImage IImage, newBufferedImage *BufferedImage, totalAngle int) *BufferedImage {
        var g *Graphics2D = newBufferedImage.CreateGraphics(); _ = g
        return this.DrawRotatedImageRadians(bufferedImage, newBufferedImage, g, this.TWO_PIE * float64(totalAngle) / 360.0)
}


func (this *ImageJ2SERotationUtil) GetRotatedImages(bufferedImage IImage, numberOfFrames int, totalAngle int) []*BufferedImage {
        var bufferedImageArray []*BufferedImage = make([]*BufferedImage, numberOfFrames); _ = bufferedImageArray
        var arc float64 = (this.TWO_PIE) * float64(totalAngle) / 360.0; _ = arc
        var size int = len(bufferedImageArray); _ = size
        for index := int(0); index < size; index++ {
            var radians float64 = (arc / float64(size)) * float64(index); _ = radians
            bufferedImageArray[index] = this.GetRotatedImageRadians(bufferedImage, radians)
        }
        return bufferedImageArray
}


func (this *ImageJ2SERotationUtil) CreateSpriteImage(bufferedImageArray []*BufferedImage) *BufferedImage {
        var columns int = 9; _ = columns
        var rows int = 0; _ = rows
        var size int = len(bufferedImageArray); _ = size
        if size < columns {
            columns = size
        }
        rows = (size / columns)
        if size % columns != 0 {
            rows++
        }
        var firstBufferedImage *BufferedImage = bufferedImageArray[0]; _ = firstBufferedImage
        var bufferedImage *BufferedImage = this.ImageUtil.Create(firstBufferedImage.GetWidth(nil) * columns, firstBufferedImage.GetHeight(nil) * rows); _ = bufferedImage
        var g *Graphics2D = bufferedImage.CreateGraphics(); _ = g
        var columnIndex int = 0; _ = columnIndex
        var rowIndex int = 0; _ = rowIndex
        var nextBufferedImage *BufferedImage; _ = nextBufferedImage
        for index := int(0); index < size; index++ {
            if index / 9 != 0 && index % 9 == 0 {
                rowIndex++
                columnIndex = 0
            }
            nextBufferedImage = bufferedImageArray[index]
            g.DrawImage(nextBufferedImage, nextBufferedImage.GetWidth(nil) * columnIndex, nextBufferedImage.GetHeight(nil) * rowIndex, nextBufferedImage.GetWidth(nil), nextBufferedImage.GetHeight(nil), nil)
            columnIndex++
        }
        g.Dispose()
        return bufferedImage
}

