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

type ImageComparisonSearch struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    IsSameHeight bool
    IsSameWidth bool
    PixelsThatMatch int
    PixelsIgnored int
    ImageHeight int
    ImageWidth int
    Tolerance int
    NonMatchingPixelVector IBasicArrayList
    BufferedImages []*BufferedImage
    MatchingPercent float32
}


func (this *ImageComparisonSearch) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
    
    
    
    
    
    
    
    
    this.MatchingPercent = float32(-1)
}


func NewImageComparisonSearch(bufferedImage *BufferedImage, bufferedImage2 *BufferedImage, tolerance int) *ImageComparisonSearch {
    this := &ImageComparisonSearch{}
    this.initFields()
        this.NonMatchingPixelVector = NewBasicArrayListD()
        this.SetBufferedImages(make([]*BufferedImage, 2))
        this.GetBufferedImages()[0] = bufferedImage
        this.GetBufferedImages()[1] = bufferedImage2
        this.ImageHeight = bufferedImage.GetHeight()
        this.ImageWidth = bufferedImage.GetWidth()
        if bufferedImage.GetHeight() != bufferedImage2.GetHeight() {
            this.IsSameHeight = false
            if this.ImageHeight > bufferedImage2.GetHeight() {
                this.ImageHeight = bufferedImage2.GetHeight()
            }
        } else {
            this.IsSameHeight = true
        }
        if bufferedImage.GetWidth() != bufferedImage2.GetWidth() {
            this.IsSameWidth = false
            if this.ImageWidth > bufferedImage2.GetWidth() {
                this.ImageWidth = bufferedImage2.GetWidth()
            }
        } else {
            this.IsSameWidth = true
        }
        if !this.IsSameWidth || !this.IsSameHeight {
            var imageUtil *ImageUtil = ImageUtil_GetInstance(); _ = imageUtil
            this.LogUtil.PutF("Images were not the same size? Most likely a resolution change.", this, this.CommonStrings.CONSTRUCTOR)
            this.LogUtil.PutF("1: " + imageUtil.ToString(bufferedImage), this, this.CommonStrings.CONSTRUCTOR)
            this.LogUtil.PutF("2: " + imageUtil.ToString(bufferedImage2), this, this.CommonStrings.CONSTRUCTOR)
        }
    return this
}


func (this *ImageComparisonSearch) GetMatchingPercent() float32 {
        if this.MatchingPercent == float32(-1) {
            this.MatchingPercent = float32(this.PixelsThatMatch) / float32((this.ImageWidth * this.ImageHeight))
        }
        return this.MatchingPercent
}


func (this *ImageComparisonSearch) Add(pixel *PixelDelta) {
        this.NonMatchingPixelVector.Add(pixel)
}


func (this *ImageComparisonSearch) GetNonMatchingPixelVector() IBasicArrayList {
        return this.NonMatchingPixelVector
}


func (this *ImageComparisonSearch) ToString() string {
        return "ImageComparisonInfo: " + " Number Of Matching Pixels: " + fmt.Sprint(this.PixelsThatMatch) + "\nNumber Of Non-Matching Pixels: " + fmt.Sprint(this.GetNonMatchingPixelVector().Size()) + "\nNumber Of Pixels Ignored: " + fmt.Sprint(this.PixelsIgnored) + "\nMatching Percentage: " + fmt.Sprint(this.GetMatchingPercent())
}

func (this *ImageComparisonSearch) String() string {
    return this.ToString()
}


func (this *ImageComparisonSearch) GetBufferedImages() []*BufferedImage {
        return this.BufferedImages
}


func (this *ImageComparisonSearch) SetBufferedImages(bufferedImages []*BufferedImage) {
        this.BufferedImages = bufferedImages
}


func (this *ImageComparisonSearch) GetTolerance() int {
        return this.Tolerance
}


func (this *ImageComparisonSearch) SetTolerance(tolerance int) {
        this.Tolerance = tolerance
}

