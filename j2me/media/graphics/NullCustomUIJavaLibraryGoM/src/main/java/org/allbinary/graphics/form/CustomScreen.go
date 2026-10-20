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

type CustomScreen struct {
    CustomDisplayable
}

// The Java class CustomScreen has subclasses so its variables are this interface
type ICustomScreen interface {
    ICustomDisplayable
    CustomScreen_() *CustomScreen
}

func (this *CustomScreen) CustomScreen_() *CustomScreen {
    return this
}


func NewCustomScreen() *CustomScreen {
    this := &CustomScreen{}
    this.CustomDisplayable = *NewCustomDisplayable()
    return this
}

