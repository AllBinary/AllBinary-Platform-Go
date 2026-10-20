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

type ABStringComponent struct {
}

var ABStringComponent_NULL_STRING_COMPONENT *ABStringComponent = NewABStringComponent(BasicColorFactory_GetInstance().BLACK, BasicColorFactory_GetInstance().WHITE)


func (this *ABStringComponent) GetBackgroundBasicColor() *BasicColor {
        return BasicColorFactory_GetInstance().WHITE
}


func NewABStringComponent(backgroundBasicColor *BasicColor, foregroundBasicColor *BasicColor) *ABStringComponent {
    this := &ABStringComponent{}
    return this
}


func (this *ABStringComponent) GetForegroundBasicColor() *BasicColor {
        return BasicColorFactory_GetInstance().WHITE
}


func (this *ABStringComponent) SetForegroundBasicColor(foregroundBasicColor *BasicColor) {
}

