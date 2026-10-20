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

type ImagesRatioUtil struct {
    LogUtil *LogUtil
    ImageUtil *ImageUtil
}

var ImagesRatioUtil_Instance *ImagesRatioUtil = NewImagesRatioUtil()

func (this *ImagesRatioUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.ImageUtil = ImageUtil_GetInstance()
}


func ImagesRatioUtil_GetInstance() *ImagesRatioUtil {
        return ImagesRatioUtil_Instance
}


func NewImagesRatioUtil() *ImagesRatioUtil {
    this := &ImagesRatioUtil{}
    this.initFields()
    return this
}


func (this *ImagesRatioUtil) IsEqual(bufferedImageArray []*BufferedImage, totalImages int) bool {
        var end int = totalImages; _ = end
        if len(bufferedImageArray) < totalImages {
            end = len(bufferedImageArray)
        }
        var ratio float64 = (float64(bufferedImageArray[0].GetWidth()) / bufferedImageArray[0].GetHeight()); _ = ratio
        for index := int(1); index < end; index++ {
            if ratio != (float64(bufferedImageArray[index].GetWidth()) / bufferedImageArray[index].GetHeight()) {
                return false
            }
        }
        return true
}


func (this *ImagesRatioUtil) GetAverage(bufferedImageArray []*BufferedImage, totalImages int) float64 {
        var ratio float64 = 0.0; _ = ratio
        var end int = totalImages; _ = end
        if len(bufferedImageArray) < totalImages {
            end = len(bufferedImageArray)
        }
        for index := int(0); index < end; index++ {
            ratio += (float64(bufferedImageArray[index].GetWidth()) / bufferedImageArray[index].GetHeight())
        }
        return (ratio / float64(end))
}


// throws Exception
func (this *ImagesRatioUtil) Fudge_1(bufferedImage *BufferedImage, ratio float64) *BufferedImage {
        var imageRatio float64 = (float64(bufferedImage.GetWidth()) / bufferedImage.GetHeight()); _ = imageRatio
        var point *Point = NewPoint(); _ = point
        var newHeight int = bufferedImage.GetHeight(); _ = newHeight
        var newWidth int = bufferedImage.GetWidth(); _ = newWidth
        if ratio > imageRatio {
            newWidth = int((float64(bufferedImage.GetHeight()) * ratio))
            point.X = (newWidth - bufferedImage.GetWidth()) / 2
        } else {
            newHeight = int((float64(bufferedImage.GetWidth()) / ratio))
            point.Y = (newHeight - bufferedImage.GetHeight()) / 2
        }
        var commonLabels *CommonLabels = CommonLabels_GetInstance(); _ = commonLabels
        this.LogUtil.PutF(fmt.Sprint(commonLabels.WIDTH_LABEL + bufferedImage.GetWidth()) + " newWidth: " + fmt.Sprint(newWidth) + fmt.Sprint(commonLabels.HEIGHT_LABEL) + fmt.Sprint(bufferedImage.GetHeight()) + " newHeight: " + fmt.Sprint(newHeight) + " needed ratio: " + fmt.Sprint((float64(newWidth) / float64(newHeight))), this, "fudge")
        var newBufferedImage *BufferedImage = this.ImageUtil.Create(newWidth, newHeight); _ = newBufferedImage
        var g *Graphics2D = newBufferedImage.CreateGraphics(); _ = g
        if point.X > 0 {
            var firstColumnBufferedImage *BufferedImage = bufferedImage.GetSubimage(0, 0, 1, bufferedImage.GetHeight()); _ = firstColumnBufferedImage
            var lastColumnBufferedImage *BufferedImage = bufferedImage.GetSubimage(bufferedImage.GetWidth() - 1, 0, 1, bufferedImage.GetHeight()); _ = lastColumnBufferedImage
            this.LogUtil.PutF("Draw some columns to fill in gap", this, "fudge")
            for index := int(0); index < point.X; index++ {
                g.DrawImage(firstColumnBufferedImage, index, 0, nil)
                g.DrawImage(lastColumnBufferedImage, newWidth - index, 0, nil)
            }
        }
        if point.Y > 0 {
            var firstRowBufferedImage *BufferedImage = bufferedImage.GetSubimage(0, 0, bufferedImage.GetWidth(), 1); _ = firstRowBufferedImage
            var lastRowBufferedImage *BufferedImage = bufferedImage.GetSubimage(0, bufferedImage.GetHeight() - 1, bufferedImage.GetWidth(), 1); _ = lastRowBufferedImage
            this.LogUtil.PutF("Draw some rows to fill in gap", this, "fudge")
            for index := int(0); index < point.Y; index++ {
                g.DrawImage(firstRowBufferedImage, 0, index, nil)
                g.DrawImage(lastRowBufferedImage, 0, newHeight - index, nil)
            }
        }
        g.DrawImage(bufferedImage, point.X, point.Y, bufferedImage.GetWidth(), bufferedImage.GetHeight(), nil)
        return newBufferedImage
}

// Java overloads of ImagesRatioUtil.fudge
func (this *ImagesRatioUtil) Fudge(args ...any) any {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 2 && Java_IsArg[*BufferedImage](args[0], pass) && Java_IsArg[float64](args[1], pass) {
            return this.Fudge_1(Java_Arg[*BufferedImage](args[0]), Java_Arg[float64](args[1]))
        }
        if len(args) == 3 && Java_IsArg[[]*BufferedImage](args[0], pass) && Java_IsArg[int](args[1], pass) && Java_IsArg[float64](args[2], pass) {
            return this.Fudge_2(Java_Arg[[]*BufferedImage](args[0]), Java_Arg[int](args[1]), Java_Arg[float64](args[2]))
        }
    }
    panic("No Java overload of ImagesRatioUtil.fudge for the arguments")
}


// throws Exception
func (this *ImagesRatioUtil) Fudge_2(bufferedImageArray []*BufferedImage, totalImages int, ratio float64) []*BufferedImage {
        var end int = totalImages; _ = end
        if len(bufferedImageArray) < totalImages {
            end = len(bufferedImageArray)
        }
        var fudgedBufferedImageArray []*BufferedImage = make([]*BufferedImage, end); _ = fudgedBufferedImageArray
        for index := int(0); index < end; index++ {
            fudgedBufferedImageArray[index] = this.Fudge(bufferedImageArray[index], ratio).(*BufferedImage)
        }
        return fudgedBufferedImageArray
}

