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

//OpenGLESBaseJavaLibrary
type FontDebugBase struct {
    SET_FONT string
}

// The Java class FontDebugBase has subclasses so its variables are this interface
type IFontDebugBase interface {
    FontDebugBase_() *FontDebugBase
    DebugMetrics(currentFont Font, convertedTextSize float32, displayMetrics any, stringBuffer *StringMaker)
    DebugDpi(densityDpi int, stringBuffer *StringMaker)
    DebugDimension(currentFont Font, convertedTextSize float32, longestDimensionTotalPixels int, stringBuffer *StringMaker)
}

func (this *FontDebugBase) FontDebugBase_() *FontDebugBase {
    return this
}

var FontDebugBase_InstanceBase IFontDebugBase = NewFontDebugBase()

func (this *FontDebugBase) initFields() {
    this.SET_FONT = "setFont"
}

func NewFontDebugBase() *FontDebugBase {
    this := &FontDebugBase{}
    this.initFields()
    return this
}


func FontDebugBase_GetInstance() IFontDebugBase {
        return FontDebugBase_InstanceBase
}


func (this *FontDebugBase) DebugMetrics(currentFont Font, convertedTextSize float32, displayMetrics any, stringBuffer *StringMaker) {
}


func (this *FontDebugBase) DebugDpi(densityDpi int, stringBuffer *StringMaker) {
}


func (this *FontDebugBase) DebugDimension(currentFont Font, convertedTextSize float32, longestDimensionTotalPixels int, stringBuffer *StringMaker) {
}

