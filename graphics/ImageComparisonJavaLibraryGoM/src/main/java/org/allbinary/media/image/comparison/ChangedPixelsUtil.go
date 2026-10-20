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

type ChangedPixelsUtil struct {
}

var ChangedPixelsUtil_START_X int = 0
var ChangedPixelsUtil_START_Y int = 0
var ChangedPixelsUtil_OFFSET int = 0
var ChangedPixelsUtil_SCAN_SIZE int = 0
var ChangedPixelsUtil_CLEAR_INT_ARRAY []int = make([]int, 1800 * 1600)


func NewChangedPixelsUtil() *ChangedPixelsUtil {
    this := &ChangedPixelsUtil{}
    return this
}


// throws Exception
func ChangedPixelsUtil_GenerateBufferedImageChacheables(imageComparisonInfo *ImageComparisonResult) []*BufferedImageCacheable {
        var bufferedImageCacheables []*BufferedImageCacheable; _ = bufferedImageCacheables
        var originalBufferedImage *BufferedImage = imageComparisonInfo.GetBufferedImages()[0]; _ = originalBufferedImage
        bufferedImageCacheables = make([]*BufferedImageCacheable, 2)
        var bufferedImageInfo *BufferedImageInfo = BufferedImageInfoFactory_GetInstance(imageComparisonInfo.ImageWidth, imageComparisonInfo.ImageHeight, originalBufferedImage.GetType()); _ = bufferedImageInfo
        var poolInterface PoolInterface = BufferedImagePoolSingleton_GetInstance(); _ = poolInterface
        for index := int(0); index < len(bufferedImageCacheables); index++ {
            var bufferedImageCacheable *BufferedImageCacheable = any(poolInterface.Remove(bufferedImageInfo)).(*BufferedImageCacheable); _ = bufferedImageCacheable
            bufferedImageCacheables[index] = bufferedImageCacheable
            bufferedImageCacheables[index].GetBufferedImage().SetRGB(ChangedPixelsUtil_START_X, ChangedPixelsUtil_START_Y, bufferedImageInfo.GetWidth(), bufferedImageInfo.GetHeight(), ChangedPixelsUtil_CLEAR_INT_ARRAY, ChangedPixelsUtil_OFFSET, ChangedPixelsUtil_SCAN_SIZE)
        }
        var vector IBasicArrayList = imageComparisonInfo.GetNonMatchingPixelVector(); _ = vector
        var size int = vector.Size(); _ = size
        for index := int(0); index < size; index++ {
            var pixelDelta *PixelDelta = any(vector.Get(index)).(*PixelDelta); _ = pixelDelta
            if pixelDelta.GetColorDelta() == nil {
                fmt.Print("ColorDelta")
                os.Exit(0)
            }
            bufferedImageCacheables[0].GetBufferedImage().SetRGB(pixelDelta.GetPoint().GetX(), pixelDelta.GetPoint().GetY(), pixelDelta.GetColorDelta().GetRgb1())
            bufferedImageCacheables[1].GetBufferedImage().SetRGB(pixelDelta.GetPoint().GetX(), pixelDelta.GetPoint().GetY(), pixelDelta.GetColorDelta().GetRgb2())
        }
        return bufferedImageCacheables
}

