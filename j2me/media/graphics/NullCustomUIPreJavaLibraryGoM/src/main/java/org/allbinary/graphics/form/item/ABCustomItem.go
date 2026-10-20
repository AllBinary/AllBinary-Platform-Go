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

type ABCustomItem struct {
}

// The Java class ABCustomItem has subclasses so its variables are this interface
type IABCustomItem interface {
    ABCustomItem_() *ABCustomItem
    SetOwner(owner IScreen)
    PaintXY(graphics IGraphics, x int, y int)
    PaintUnselected(graphics IGraphics, x int, y int)
    IsFocusable() bool
    GetLabel() string
    SetLabel(label string)
    GetHeight() int
    SetHeight(height int)
    GetValue() float32
    SetValue(value float32)
    GetMinimumWidth() int
    GetMinimumHeight() int
    PreMeasurement(graphics IGraphics)
}

func (this *ABCustomItem) ABCustomItem_() *ABCustomItem {
    return this
}

var ABCustomItem_NULL_CUSTOM_ITEM any = NullUtil_GetInstance().NULL_OBJECT


func ABCustomItem_GetNullInstance() IABCustomItem {
        if ABCustomItem_NULL_CUSTOM_ITEM == NullUtil_GetInstance().NULL_OBJECT {
            ABCustomItem_NULL_CUSTOM_ITEM = NewABCustomItem(StringUtil_GetInstance().EMPTY_STRING, BasicColorFactory_GetInstance().BLACK, BasicColorFactory_GetInstance().WHITE)
        }
        return any(ABCustomItem_NULL_CUSTOM_ITEM).(IABCustomItem)
}


func NewABCustomItem(label string, backgroundBasicColor *BasicColor, foregroundBasicColor *BasicColor) *ABCustomItem {
    this := &ABCustomItem{}
    return this
}


func (this *ABCustomItem) SetOwner(owner IScreen) {
}


func (this *ABCustomItem) PaintXY(graphics IGraphics, x int, y int) {
}


func (this *ABCustomItem) PaintUnselected(graphics IGraphics, x int, y int) {
}


func (this *ABCustomItem) IsFocusable() bool {
        return false
}


func (this *ABCustomItem) GetLabel() string {
        return StringUtil_GetInstance().EMPTY_STRING
}


func (this *ABCustomItem) SetLabel(label string) {
}


func (this *ABCustomItem) GetHeight() int {
        return 0
}


func (this *ABCustomItem) SetHeight(height int) {
}


func (this *ABCustomItem) GetValue() float32 {
        return 0.0
}


func (this *ABCustomItem) SetValue(value float32) {
}


func (this *ABCustomItem) GetMinimumWidth() int {
        return 0
}


func (this *ABCustomItem) GetMinimumHeight() int {
        return 0
}


func (this *ABCustomItem) PreMeasurement(graphics IGraphics) {
}

