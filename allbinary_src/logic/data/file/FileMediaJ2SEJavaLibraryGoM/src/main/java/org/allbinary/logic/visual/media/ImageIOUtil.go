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

type ImageIOUtil struct {
}


func NewImageIOUtil() *ImageIOUtil {
    this := &ImageIOUtil{}
    return this
}


// throws Exception
func ImageIOUtil_Read(file *AbFile) *BufferedImage {
        return ImageIO_Read(AbFileNativeUtil_Get(file))
}


// throws Exception
func ImageIOUtil_Write(renderedImage *RenderedImage, name string, file *AbFile) bool {
        return ImageIO_Write(renderedImage, name, AbFileNativeUtil_Get(file))
}

