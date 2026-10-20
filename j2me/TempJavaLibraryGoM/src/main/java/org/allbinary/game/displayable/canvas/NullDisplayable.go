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

//Temp should not be in any build
type NullDisplayable struct {
    Canvas
}

var NullDisplayable_SINGLETON *NullDisplayable = NewNullDisplayable()


func NullDisplayable_GetInstance() *NullDisplayable {
        return NullDisplayable_SINGLETON
}


func NewNullDisplayable() *NullDisplayable {
    this := &NullDisplayable{}
    this.Canvas = *NewCanvas()
    return this
}


func (this *NullDisplayable) Paint(graphics IGraphics) {
}

