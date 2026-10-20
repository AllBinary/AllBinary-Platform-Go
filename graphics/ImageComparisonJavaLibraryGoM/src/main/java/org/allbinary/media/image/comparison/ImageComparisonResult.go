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

type ImageComparisonResult struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    BufferedImages []*BufferedImage
    Name string
    IsSameHeight bool
    IsSameWidth bool
    ImageHeight int
    ImageWidth int
    NonMatchingPixelVector IBasicArrayList
    FrameOne int64
    FrameTwo int64
    PixelsThatMatch int
    PixelsIgnored int
    Tolerance int
    MatchingPercent float32
}


func (this *ImageComparisonResult) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.BufferedImages = make([]*BufferedImage, 2)
    
    
    
    
    
    
    
    
    
    
    
    this.MatchingPercent = float32(-1)
}


func NewImageComparisonResult(name string, bufferedImage *BufferedImage, bufferedImage2 *BufferedImage, frameOne int64, frameTwo int64, tolerance int) *ImageComparisonResult {
    this := &ImageComparisonResult{}
    this.initFields()
        this.Name = name
        this.NonMatchingPixelVector = NewBasicArrayListD()
        this.BufferedImages[0] = bufferedImage
        this.BufferedImages[1] = bufferedImage2
        this.FrameOne = frameOne
        this.FrameTwo = frameTwo
        this.SetTolerance(0)
        var imageHeight int = bufferedImage.GetHeight(); _ = imageHeight
        var imageWidth int = bufferedImage.GetWidth(); _ = imageWidth
        if bufferedImage.GetHeight() != bufferedImage2.GetHeight() {
            this.IsSameHeight = false
            if imageHeight > bufferedImage2.GetHeight() {
                imageHeight = bufferedImage2.GetHeight()
            }
        } else {
            this.IsSameHeight = true
        }
        if bufferedImage.GetWidth() != bufferedImage2.GetWidth() {
            this.IsSameWidth = false
            if imageWidth > bufferedImage2.GetWidth() {
                imageWidth = bufferedImage2.GetWidth()
            }
        } else {
            this.IsSameWidth = true
        }
        this.ImageWidth = imageWidth
        this.ImageHeight = imageHeight
        if !this.IsSameWidth || !this.IsSameHeight {
            var imageUtil *ImageUtil = ImageUtil_GetInstance(); _ = imageUtil
            this.LogUtil.PutF("Images were not the same size? Most likely a resolution change.", this, this.CommonStrings.CONSTRUCTOR)
            this.LogUtil.PutF("1: " + imageUtil.ToString(bufferedImage), this, this.CommonStrings.CONSTRUCTOR)
            this.LogUtil.PutF("2: " + imageUtil.ToString(bufferedImage2), this, this.CommonStrings.CONSTRUCTOR)
        }
    return this
}


func (this *ImageComparisonResult) GetMatchingPercent() float32 {
        if this.MatchingPercent == float32(-1) {
            this.MatchingPercent = float32(this.PixelsThatMatch) / float32((this.ImageWidth * this.ImageHeight))
        }
        return this.MatchingPercent
}


func (this *ImageComparisonResult) Add(pixel *PixelDelta) {
        this.NonMatchingPixelVector.Add(pixel)
}


func (this *ImageComparisonResult) GetNonMatchingPixelVector() IBasicArrayList {
        return this.NonMatchingPixelVector
}


func (this *ImageComparisonResult) ToString() string {
        return fmt.Sprint(NewStringMaker().Append("ImageComparisonInfo: ").Append(this.Name).Append(" Number Of Matching Pixels: ").Appendint(this.PixelsThatMatch).Append("\nNumber Of Non-Matching Pixels: ").Appendint(this.GetNonMatchingPixelVector().Size()).Append("\nNumber Of Pixels Ignored: ").Appendint(this.PixelsIgnored).Append("\nMatching Percentage: ").Appendfloat(this.GetMatchingPercent()))
}

func (this *ImageComparisonResult) String() string {
    return this.ToString()
}


func (this *ImageComparisonResult) GetBufferedImages() []*BufferedImage {
        return this.BufferedImages
}


func (this *ImageComparisonResult) GetTolerance() int {
        return this.Tolerance
}


func (this *ImageComparisonResult) SetTolerance(tolerance int) {
        this.Tolerance = tolerance
}


func (this *ImageComparisonResult) GetFrameOne() int64 {
        return this.FrameOne
}


func (this *ImageComparisonResult) GetFrameTwo() int64 {
        return this.FrameTwo
}

