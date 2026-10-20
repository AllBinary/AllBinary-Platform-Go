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

type V struct {
}

// The Java class V has subclasses so its variables are this interface
type IV interface {
    V_() *V
    GetId() int
}

func (this *V) V_() *V {
    return this
}


func NewV() *V {
    this := &V{}
    return this
}


func (this *V) GetId() int {
        return 0
}

