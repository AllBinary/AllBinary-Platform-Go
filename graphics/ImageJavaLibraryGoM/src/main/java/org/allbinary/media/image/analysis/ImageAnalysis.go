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

type ImageAnalysis struct {
    LogUtil *LogUtil
}

var ImageAnalysis_Instance *ImageAnalysis = NewImageAnalysis()

func (this *ImageAnalysis) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func ImageAnalysis_GetInstance() *ImageAnalysis {
        return ImageAnalysis_Instance
}


func NewImageAnalysis() *ImageAnalysis {
    this := &ImageAnalysis{}
    this.initFields()
    return this
}


// throws Exception
func (this *ImageAnalysis) Process(bufferedImageArray []*BufferedImage, colorRangeInterface ColorRangeInterface) []*ImageAnalysisResults {
        var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
        this.LogUtil.PutF(fmt.Sprint(CommonLabels_GetInstance().START) + fmt.Sprint(colorRangeInterface), this, commonStrings.PROCESS)
        var imageAnalysisResultsArray []*ImageAnalysisResults = make([]*ImageAnalysisResults, len(bufferedImageArray)); _ = imageAnalysisResultsArray
        for index := int(0); index < len(bufferedImageArray); index++ {
            imageAnalysisResultsArray[index] = ImageAnalysis_Process(bufferedImageArray[index], colorRangeInterface)
        }
        return imageAnalysisResultsArray
}


// throws Exception
func ImageAnalysis_Process(bufferedImage *BufferedImage, colorRangeInterface ColorRangeInterface) *ImageAnalysisResults {
        var imageAnalysisResults *ImageAnalysisResults = NewImageAnalysisResults(); _ = imageAnalysisResults
        var redTotal int64 = 0; _ = redTotal
        var greenTotal int64 = 0; _ = greenTotal
        var blueTotal int64 = 0; _ = blueTotal
        for indexY := int(0); indexY < bufferedImage.GetHeight(); indexY++ {
            for indexX := int(0); indexX < bufferedImage.GetWidth(); indexX++ {
                var keyInteger int = fmt.Sprint(bufferedImage.GetRGB(indexX, indexY)); _ = keyInteger
                var colorCacheable *ColorCacheable = any(ColorCacheFactory_GetInstance().Get(keyInteger)).(*ColorCacheable); _ = colorCacheable
                var color *Color = colorCacheable.GetColor(); _ = color
                ImageAnalysis_ProcessColorRangeResults(imageAnalysisResults, colorRangeInterface, color)
                ImageAnalysis_ProcessImageColorResults(imageAnalysisResults.GetImageColorResults(), colorRangeInterface, color)
                redTotal += int64(color.GetRed())
                greenTotal += int64(color.GetGreen())
                blueTotal += int64(color.GetBlue())
            }
        }
        var totalPixels int64 = imageAnalysisResults.GetImageColorRangeResults().GetTotalPixelsChecked(); _ = totalPixels
        var colorAverage *ColorAverage = imageAnalysisResults.GetImageColorResults().GetColorAverage(); _ = colorAverage
        colorAverage.SetAvgRed(float32(redTotal) / float32(totalPixels))
        colorAverage.SetAvgGreen(float32(greenTotal) / float32(totalPixels))
        colorAverage.SetAvgBlue(float32(blueTotal) / float32(totalPixels))
        return imageAnalysisResults
}


func ImageAnalysis_ProcessColorRangeResults(imageAnalysisResults *ImageAnalysisResults, colorRangeInterface ColorRangeInterface, color *Color) {
        if colorRangeInterface.IsInRange(color) {
            imageAnalysisResults.GetImageColorRangeResults().AddMatchingPixelsChecked()
        } else {
        }
        imageAnalysisResults.GetImageColorRangeResults().AddTotalPixelsChecked()
}


func ImageAnalysis_ProcessImageColorResults(imageColorResults *ImageColorResults, colorRangeInterface ColorRangeInterface, color *Color) {
        if color.GetRed() < imageColorResults.GetColorRange().GetMinRed() {
            imageColorResults.GetColorRange().SetMinRed(color.GetRed())
        }
        if color.GetGreen() < imageColorResults.GetColorRange().GetMinGreen() {
            imageColorResults.GetColorRange().SetMinGreen(color.GetGreen())
        }
        if color.GetBlue() < imageColorResults.GetColorRange().GetMinBlue() {
            imageColorResults.GetColorRange().SetMinBlue(color.GetBlue())
        }
        if color.GetRed() > imageColorResults.GetColorRange().GetMaxRed() {
            imageColorResults.GetColorRange().SetMaxRed(color.GetRed())
        }
        if color.GetGreen() > imageColorResults.GetColorRange().GetMaxGreen() {
            imageColorResults.GetColorRange().SetMaxGreen(color.GetGreen())
        }
        if color.GetBlue() > imageColorResults.GetColorRange().GetMaxBlue() {
            imageColorResults.GetColorRange().SetMaxBlue(color.GetBlue())
        }
}

