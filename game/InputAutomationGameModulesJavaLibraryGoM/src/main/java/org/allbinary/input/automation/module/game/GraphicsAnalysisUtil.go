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

type GraphicsAnalysisUtil struct {
}

var GraphicsAnalysisUtil_MAX int = 122


func NewGraphicsAnalysisUtil() *GraphicsAnalysisUtil {
    this := &GraphicsAnalysisUtil{}
    return this
}


// throws Exception
func GraphicsAnalysisUtil_GetNominator(bufferedImage *BufferedImage, min_x int, max_x int, y int) float64 {
        for index := int(int(max_x)); index > int(min_x); index -= 1 {
            var colorInteger int = fmt.Sprint(bufferedImage.GetRGB(index, int(y))); _ = colorInteger
            var automaticCacheInterface AutomaticCacheInterface = ColorCacheFactory_GetInstance(); _ = automaticCacheInterface
            var colorCacheable *ColorCacheable = any(automaticCacheInterface.Get(colorInteger)).(*ColorCacheable); _ = colorCacheable
            var color *Color = colorCacheable.GetColor(); _ = color
            if color.GetRed() < GraphicsAnalysisUtil_MAX && color.GetGreen() < GraphicsAnalysisUtil_MAX && color.GetBlue() < GraphicsAnalysisUtil_MAX {
            } else {
                return float64(index - int(min_x))
            }
        }
        return 0.0
}

