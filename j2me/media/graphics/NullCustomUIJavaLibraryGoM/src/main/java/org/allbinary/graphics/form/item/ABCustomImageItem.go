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

type ABCustomImageItem struct {
    ABCustomItem
    YOffset int
}

// The Java class ABCustomImageItem has subclasses so its variables are this interface
type IABCustomImageItem interface {
    IABCustomItem
    ABCustomImageItem_() *ABCustomImageItem
    GetImage() IImage
}

func (this *ABCustomImageItem) ABCustomImageItem_() *ABCustomImageItem {
    return this
}


func (this *ABCustomImageItem) initFields() {
    this.YOffset = 0
}


func NewABCustomImageItem(label string, image IImage, layout int, altText string, basicColor *BasicColor, yOffset int) *ABCustomImageItem {
    this := &ABCustomImageItem{}
        this.ABCustomItem = *NewABCustomItem(label, basicColor, basicColor)
        this.initFields()
    return this
}


func (this *ABCustomImageItem) GetImage() IImage {
        return NullImage_NULL_IMAGE
}

