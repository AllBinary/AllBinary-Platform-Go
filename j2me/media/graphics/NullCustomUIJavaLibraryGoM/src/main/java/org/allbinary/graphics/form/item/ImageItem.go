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

type ImageItem struct {
}

// The Java class ImageItem has subclasses so its variables are this interface
type IImageItem interface {
    ImageItem_() *ImageItem
}

func (this *ImageItem) ImageItem_() *ImageItem {
    return this
}


func NewImageItem() *ImageItem {
    this := &ImageItem{}
    return this
}

