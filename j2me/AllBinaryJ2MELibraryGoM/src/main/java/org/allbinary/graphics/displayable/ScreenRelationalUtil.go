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

type ScreenRelationalUtil struct {
}

var ScreenRelationalUtil_Instance *ScreenRelationalUtil = NewScreenRelationalUtil()

func NewScreenRelationalUtil() *ScreenRelationalUtil {
    this := &ScreenRelationalUtil{}
    return this
}


func ScreenRelationalUtil_GetInstance() *ScreenRelationalUtil {
        return ScreenRelationalUtil_Instance
}


func (this *ScreenRelationalUtil) GetScaleImage(image IImage) float32 {
        return this.GetScaleWH(image.GetWidth(), image.GetHeight())
}


func (this *ScreenRelationalUtil) GetScaleWH(aWidth int, aHeight int) float32 {
        return this.GetScaleWHA(aWidth, aHeight, 20)
}


func (this *ScreenRelationalUtil) GetScaleWHA(aWidth int, aHeight int, add int) float32 {
        var displayInfoSingleton *DisplayInfoSingleton = DisplayInfoSingleton_GetInstance(); _ = displayInfoSingleton
        var last []int = displayInfoSingleton.GetLast(); _ = last
        var width float32 = float32(last[displayInfoSingleton.WIDTH]); _ = width
        var height float32 = float32(last[displayInfoSingleton.HEIGHT]); _ = height
        var largestSize float32 = width; _ = largestSize
        if largestSize < height {
            largestSize = height
        }
        var scale float32 = (largestSize) / float32(aWidth); _ = scale
        return scale
}


func (this *ScreenRelationalUtil) GetScaledHeight(aWidth int, aHeight int) int {
        return int((this.GetScaleWH(aWidth, aHeight) * float32(aHeight)))
}


func (this *ScreenRelationalUtil) GetScaledHeightAdd(aWidth int, aHeight int, add int) int {
        return int((this.GetScaleWHA(aWidth, aHeight, add) * float32(aHeight)))
}

