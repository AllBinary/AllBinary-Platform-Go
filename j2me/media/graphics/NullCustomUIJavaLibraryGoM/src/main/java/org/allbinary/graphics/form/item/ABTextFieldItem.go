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

type ABTextFieldItem struct {
    ABTextItem
    Font Font
}


func (this *ABTextFieldItem) initFields() {
    
}


func NewABTextFieldItem(canvas ICanvas, visitor IVisitor, label string, value string, maxSize int, layout int, altText string, font Font, backgroundBasicColor *BasicColor, foregroundBasicColor *BasicColor) *ABTextFieldItem {
    this := &ABTextFieldItem{}
        this.ABTextItem = *NewABTextItem(label, layout, altText, backgroundBasicColor, foregroundBasicColor)
        this.initFields()
        this.Font = font
    return this
}


func (this *ABTextFieldItem) GetLabel() string {
        return StringUtil_GetInstance().EMPTY_STRING
}


func (this *ABTextFieldItem) SetLabel(label string) {
}


func (this *ABTextFieldItem) GetString() string {
        return StringUtil_GetInstance().EMPTY_STRING
}


func (this *ABTextFieldItem) SetString(string_ string) {
}


func (this *ABTextFieldItem) SetFocus(state bool) {
}


func (this *ABTextFieldItem) GetFontHeight() int {
        return this.Font.GetHeight()
}


func (this *ABTextFieldItem) SetBackgroundBasicColorP(basicColor *BasicColor) {
}


func (this *ABTextFieldItem) SetForegroundBasicColorP(basicColor *BasicColor) {
}

