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

type CustomDisplayable struct {
}

// The Java class CustomDisplayable has subclasses so its variables are this interface
type ICustomDisplayable interface {
    CustomDisplayable_() *CustomDisplayable
    SetTitle(title string)
}

func (this *CustomDisplayable) CustomDisplayable_() *CustomDisplayable {
    return this
}


func NewCustomDisplayable() *CustomDisplayable {
    this := &CustomDisplayable{}
    return this
}


func (this *CustomDisplayable) SetTitle(title string) {
}

