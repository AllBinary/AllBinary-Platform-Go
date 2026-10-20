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

type ImageUnifierUtil struct {
    LogUtil *LogUtil
    ImageUtil *ImageUtil
}

var ImageUnifierUtil_Instance *ImageUnifierUtil = NewImageUnifierUtil()

func (this *ImageUnifierUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.ImageUtil = ImageUtil_GetInstance()
}


func ImageUnifierUtil_GetInstance() *ImageUnifierUtil {
        return ImageUnifierUtil_Instance
}


func NewImageUnifierUtil() *ImageUnifierUtil {
    this := &ImageUnifierUtil{}
    this.initFields()
    return this
}


func (this *ImageUnifierUtil) GetDefaultConfiguration() *GraphicsConfiguration {
        var ge *GraphicsEnvironment = GraphicsEnvironment_GetLocalGraphicsEnvironment(); _ = ge
        var gd *GraphicsDevice = ge.GetDefaultScreenDevice(); _ = gd
        return gd.GetDefaultConfiguration()
}


func (this *ImageUnifierUtil) GetImage(bufferedImageArray []*BufferedImage, imageUnifierProperties *ImageUnifierProperties) *BufferedImage {
        var newBufferedImage *BufferedImage = this.ImageUtil.Create(imageUnifierProperties.GetWidth(), imageUnifierProperties.GetHeight()); _ = newBufferedImage
        var commonLabels *CommonLabels = CommonLabels_GetInstance(); _ = commonLabels
        this.LogUtil.PutF("Setting Image - " + fmt.Sprint(commonLabels.WIDTH_LABEL) + fmt.Sprint(newBufferedImage.GetWidth()) + fmt.Sprint(commonLabels.HEIGHT_LABEL) + fmt.Sprint(newBufferedImage.GetHeight()), this, "getImage")
        var g *Graphics2D = newBufferedImage.CreateGraphics(); _ = g
        var columnIndex int = 0; _ = columnIndex
        var rowIndex int = 0; _ = rowIndex
        for index := int(0); index < len(bufferedImageArray); index++ {
            var x int = int(imageUnifierProperties.GetImageUnifierCell().GetWidth()) * Java_PostIncrement(&columnIndex); _ = x
            var y int = int(imageUnifierProperties.GetImageUnifierCell().GetHeight()) * rowIndex; _ = y
            this.LogUtil.PutF("Adding Image: " + fmt.Sprint(index) + " x: " + fmt.Sprint(x) + " y: " + fmt.Sprint(y), this, "getImage")
            g.DrawImage(bufferedImageArray[index], x, y, int(imageUnifierProperties.GetImageUnifierCell().GetWidth()), int(imageUnifierProperties.GetImageUnifierCell().GetHeight()), nil)
            var totalColumnsMinusOne int = (int(imageUnifierProperties.GetColumns()) - 1); _ = totalColumnsMinusOne
            if columnIndex > totalColumnsMinusOne {
                rowIndex++
                columnIndex = 0
            }
        }
        g.Dispose()
        return newBufferedImage
}

