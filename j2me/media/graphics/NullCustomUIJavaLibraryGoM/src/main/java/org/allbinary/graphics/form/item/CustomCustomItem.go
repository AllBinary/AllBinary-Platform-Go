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

type CustomCustomItem struct {
    ABCustomItem
    UpdateMyFontProcessor IMyFontProcessor
    MyFontProcessor IMyFontProcessor
}

// Implements UpdateMyFontInterface

// The Java class CustomCustomItem has subclasses so its variables are this interface
type ICustomCustomItem interface {
    IABCustomItem
    CustomCustomItem_() *CustomCustomItem
    UpdateMeasurement(graphics IGraphics)
}

func (this *CustomCustomItem) CustomCustomItem_() *CustomCustomItem {
    return this
}


func (this *CustomCustomItem) initFields() {
    this.UpdateMyFontProcessor = NewUpdateMyFontProcessor(this)
    this.MyFontProcessor = this.UpdateMyFontProcessor
}


func NewCustomCustomItem(label string, backgroundBasicColor *BasicColor, foregroundBasicColor *BasicColor) *CustomCustomItem {
    this := &CustomCustomItem{}
        this.ABCustomItem = *NewABCustomItem(label, backgroundBasicColor, foregroundBasicColor)
        this.initFields()
    return this
}


func (this *CustomCustomItem) UpdateMeasurement(graphics IGraphics) {
}

