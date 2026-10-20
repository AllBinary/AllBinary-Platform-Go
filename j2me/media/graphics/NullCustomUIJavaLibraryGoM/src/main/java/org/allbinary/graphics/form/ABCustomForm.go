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

type ABCustomForm struct {
    CustomScreen
    LogUtil *LogUtil
    CommonStrings *CommonStrings
}

// The Java class ABCustomForm has subclasses so its variables are this interface
type IABCustomForm interface {
    ICustomScreen
    ABCustomForm_() *ABCustomForm
    GetAllitems() []IABCustomItem
    Size() int
    SetSelectedIndex(index int)
    GetSelectedIndex() int
    DeleteAll()
    Delete(itemNum int)
    Append(item IABCustomItem) int
    Insert(itemNum int, item IABCustomItem)
    Set(itemNum int, item IABCustomItem)
    Get(itemNum int) IABCustomItem
    GetWidth() int
    Paint(graphics IGraphics)
    GetTitle() string
}

func (this *ABCustomForm) ABCustomForm_() *ABCustomForm {
    return this
}


func (this *ABCustomForm) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
}


func NewABCustomForm(title string, items []IABCustomItem, backgroundBasicColor *BasicColor, foregroundBasicColor *BasicColor) *ABCustomForm {
    this := &ABCustomForm{}
    this.CustomScreen = *NewCustomScreen()
    this.initFields()
    return this
}


func (this *ABCustomForm) GetAllitems() []IABCustomItem {
        return make([]IABCustomItem, 0)
}


func (this *ABCustomForm) Size() int {
        return 0
}


func (this *ABCustomForm) SetSelectedIndex(index int) {
}


func (this *ABCustomForm) GetSelectedIndex() int {
        return 0
}


func (this *ABCustomForm) DeleteAll() {
}


func (this *ABCustomForm) Delete(itemNum int) {
}


func (this *ABCustomForm) Append(item IABCustomItem) int {
        return -1
}


func (this *ABCustomForm) Insert(itemNum int, item IABCustomItem) {
}


func (this *ABCustomForm) Set(itemNum int, item IABCustomItem) {
}


func (this *ABCustomForm) Get(itemNum int) IABCustomItem {
        return ABCustomItem_GetNullInstance()
}


func (this *ABCustomForm) GetWidth() int {
        panic(errors.New(fmt.Sprint()))
}


func (this *ABCustomForm) Paint(graphics IGraphics) {
}


func (this *ABCustomForm) GetTitle() string {
        return StringUtil_GetInstance().EMPTY_STRING
}

