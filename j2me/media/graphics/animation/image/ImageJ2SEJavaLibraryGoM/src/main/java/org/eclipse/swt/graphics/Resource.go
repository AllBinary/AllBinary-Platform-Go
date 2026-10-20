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

type Resource struct {
}

// The Java class Resource has subclasses so its variables are this interface
type IResource interface {
    Resource_() *Resource
    Dispose()
    IsDisposed() bool
}

func (this *Resource) Resource_() *Resource {
    return this
}


func NewResource() *Resource {
    this := &Resource{}
    return this
}


func (this *Resource) Dispose() {
}


func (this *Resource) IsDisposed() bool {
        return false
}

