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

type ABTextItem struct {
    CustomCustomItem
}

// Implements ABCustomItemInterface

// The Java class ABTextItem has subclasses so its variables are this interface
type IABTextItem interface {
    ICustomCustomItem
    ABTextItem_() *ABTextItem
    KeyPressed(keyCode int)
    PaintUnselected(graphics IGraphics, x int, y int)
    GetLabelStringComponent() *ABStringComponent
}

func (this *ABTextItem) ABTextItem_() *ABTextItem {
    return this
}



func NewABTextItem(label string, layout int, altText string, backgroundBasicColor *BasicColor, foregroundBasicColor *BasicColor) *ABTextItem {
    this := &ABTextItem{}
        this.CustomCustomItem = *NewCustomCustomItem(label, backgroundBasicColor, foregroundBasicColor)
    return this
}


func (this *ABTextItem) KeyPressed(keyCode int) {
}


func (this *ABTextItem) PaintUnselected(graphics IGraphics, x int, y int) {
}


func (this *ABTextItem) GetLabelStringComponent() *ABStringComponent {
        return ABStringComponent_NULL_STRING_COMPONENT
}

