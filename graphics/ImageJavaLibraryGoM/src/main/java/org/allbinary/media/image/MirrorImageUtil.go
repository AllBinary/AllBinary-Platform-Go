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

type MirrorImageUtil struct {
    LogUtil *LogUtil
    ImageUtil *ImageUtil
}

var MirrorImageUtil_Instance *MirrorImageUtil = NewMirrorImageUtil()

func (this *MirrorImageUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.ImageUtil = ImageUtil_GetInstance()
}


func MirrorImageUtil_GetInstance() *MirrorImageUtil {
        return MirrorImageUtil_Instance
}


func NewMirrorImageUtil() *MirrorImageUtil {
    this := &MirrorImageUtil{}
    this.initFields()
    return this
}


func (this *MirrorImageUtil) GetImage(bufferedImage *BufferedImage, verticle bool, horizontal bool) *BufferedImage {
        this.LogUtil.PutF("Starting", this, "getImage")
        var newBufferedImage *BufferedImage = this.ImageUtil.Create(bufferedImage.GetWidth(nil), bufferedImage.GetHeight(nil)); _ = newBufferedImage
        var g *Graphics2D = newBufferedImage.CreateGraphics(); _ = g
        if verticle {
            var tx *AffineTransform = AffineTransform_GetScaleInstance(1, -1); _ = tx
            tx.Translate(0, -bufferedImage.GetHeight(nil))
            var op *AffineTransformOp = NewAffineTransformOp(tx, AffineTransformOp_TYPE_NEAREST_NEIGHBOR); _ = op
            bufferedImage = op.Filter(bufferedImage, nil)
        }
        if horizontal {
            var tx *AffineTransform = AffineTransform_GetScaleInstance(-1, 1); _ = tx
            tx.Translate(-bufferedImage.GetWidth(nil), 0)
            var op *AffineTransformOp = NewAffineTransformOp(tx, AffineTransformOp_TYPE_NEAREST_NEIGHBOR); _ = op
            bufferedImage = op.Filter(bufferedImage, nil)
        }
        g.DrawImage(bufferedImage, 0, 0, bufferedImage.GetWidth(nil), bufferedImage.GetHeight(nil), nil)
        g.Dispose()
        return newBufferedImage
}


func (this *MirrorImageUtil) GetImages(bufferedImage *BufferedImage, verticle bool, horizontal bool) []*BufferedImage {
        var width int = bufferedImage.GetWidth(); _ = width
        var height int = bufferedImage.GetHeight(); _ = height
        var cellHeight int = height; _ = cellHeight
        var cellWidth int = height; _ = cellWidth
        var numberOfFramesPerOrientation int = width / cellWidth; _ = numberOfFramesPerOrientation
        var numberOfFrames int = numberOfFramesPerOrientation; _ = numberOfFrames
        if verticle {
            numberOfFrames *= 2
        }
        if horizontal {
            numberOfFrames *= 2
        }
        this.LogUtil.PutF("numberOfFramesPerOrientation: " + fmt.Sprint(numberOfFramesPerOrientation) + " numberOfFrames: " + fmt.Sprint(numberOfFrames), this, "getImages")
        var bufferedImageArray []*BufferedImage = make([]*BufferedImage, numberOfFrames); _ = bufferedImageArray
        var y int = 0; _ = y
        for index := int(0); index < numberOfFramesPerOrientation; index++ {
            var x int = index * cellWidth; _ = x
            bufferedImageArray[index] = bufferedImage.GetSubimage(x, y, cellWidth, cellHeight)
        }
        for index := int(0); index < numberOfFramesPerOrientation; index++ {
            bufferedImageArray[index + numberOfFramesPerOrientation] = this.GetImage(bufferedImageArray[index], verticle, horizontal)
        }
        return bufferedImageArray
}

