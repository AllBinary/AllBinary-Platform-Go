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

type ImageComparator struct {
    ImageComparatorConstraintsInterface ImageComparatorConstraintsInterface
}


func (this *ImageComparator) initFields() {
    
}


func NewImageComparator(imageComparatorConstraintsInterface ImageComparatorConstraintsInterface) *ImageComparator {
    this := &ImageComparator{}
    this.initFields()
        this.ImageComparatorConstraintsInterface = imageComparatorConstraintsInterface
    return this
}


// throws Exception
func (this *ImageComparator) Process(imageComparisonInfo *ImageComparisonResult, point IGPoint) {
        var isCollsionWithAvoidRectangles bool = this.ImageComparatorConstraintsInterface.IsCollisionWithAvoidRectangles(point); _ = isCollsionWithAvoidRectangles
        if !isCollsionWithAvoidRectangles {
            var rgb1 int = imageComparisonInfo.GetBufferedImages()[0].GetRGB(point.GetX(), point.GetY()); _ = rgb1
            var rgb2 int = imageComparisonInfo.GetBufferedImages()[1].GetRGB(point.GetX(), point.GetY()); _ = rgb2
            var automaticCacheInterface AutomaticCacheInterface = ColorCacheFactory_GetInstance(); _ = automaticCacheInterface
            var colorInteger int = fmt.Sprint(rgb1); _ = colorInteger
            var colorCacheable *ColorCacheable = any(automaticCacheInterface.Get(colorInteger)).(*ColorCacheable); _ = colorCacheable
            var color *Color = colorCacheable.GetColor(); _ = color
            var colorInteger2 int = fmt.Sprint(rgb2); _ = colorInteger2
            var colorCacheable2 *ColorCacheable = any(automaticCacheInterface.Get(colorInteger)).(*ColorCacheable); _ = colorCacheable2
            var color2 *Color = colorCacheable.GetColor(); _ = color2
            if this.ImageComparatorConstraintsInterface.IsColorAllowed(0, point, color) || this.ImageComparatorConstraintsInterface.IsColorAllowed(1, point, color2) {
                if max(rgb1 - rgb2, -(rgb1 - rgb2)) <= imageComparisonInfo.GetTolerance() {
                    imageComparisonInfo.PixelsThatMatch++
                } else {
                    imageComparisonInfo.Add(PixelDeltaFactory_GetInstance(point.GetX(), point.GetY(), rgb1, rgb2))
                }
            }
        } else {
            imageComparisonInfo.PixelsIgnored++
        }
}


// throws Exception
func (this *ImageComparator) Compare_1(bufferedImage1 *BufferedImage, bufferedImage2 *BufferedImage, frameOne int64, frameTwo int64, tolerance int) *ImageComparisonResult {
        return this.Compare(bufferedImage1, bufferedImage2, frameOne, frameTwo, tolerance, StringUtil_GetInstance().EMPTY_STRING)
}

// Java overloads of ImageComparator.compare
func (this *ImageComparator) Compare(args ...any) *ImageComparisonResult {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 5 && Java_IsArg[*BufferedImage](args[0], pass) && Java_IsArg[*BufferedImage](args[1], pass) && Java_IsArg[int64](args[2], pass) && Java_IsArg[int64](args[3], pass) && Java_IsArg[int](args[4], pass) {
            return this.Compare_1(Java_Arg[*BufferedImage](args[0]), Java_Arg[*BufferedImage](args[1]), Java_Arg[int64](args[2]), Java_Arg[int64](args[3]), Java_Arg[int](args[4]))
        }
        if len(args) == 6 && Java_IsArg[*BufferedImage](args[0], pass) && Java_IsArg[*BufferedImage](args[1], pass) && Java_IsArg[int64](args[2], pass) && Java_IsArg[int64](args[3], pass) && Java_IsArg[int](args[4], pass) && Java_IsArg[string](args[5], pass) {
            return this.Compare_2(Java_Arg[*BufferedImage](args[0]), Java_Arg[*BufferedImage](args[1]), Java_Arg[int64](args[2]), Java_Arg[int64](args[3]), Java_Arg[int](args[4]), Java_Arg[string](args[5]))
        }
    }
    panic("No Java overload of ImageComparator.compare for the arguments")
}


// throws Exception
func (this *ImageComparator) Compare_2(bufferedImage1 *BufferedImage, bufferedImage2 *BufferedImage, frameOne int64, frameTwo int64, tolerance int, name string) *ImageComparisonResult {
        if bufferedImage1 == nil || bufferedImage2 == nil {
            panic(errors.New(fmt.Sprint("Input images must not be null.")))
        }
        var imageComparisonInfo *ImageComparisonResult = NewImageComparisonResult(name, bufferedImage1, bufferedImage2, frameOne, frameTwo, tolerance); _ = imageComparisonInfo
        for indexY := int(0); indexY < imageComparisonInfo.ImageHeight; indexY++ {
            for indexX := int(0); indexX < imageComparisonInfo.ImageWidth; indexX++ {
                this.Process(imageComparisonInfo, PointFactory_GetInstance().CreateXY(indexX, indexY))
                if imageComparisonInfo.GetNonMatchingPixelVector().Size() > this.ImageComparatorConstraintsInterface.GetMaxNonMatchingPixelDeltas() {
                    break
                }
            }
        }
        return imageComparisonInfo
}

