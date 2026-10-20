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

type OpenGLESGraphicsFactory struct {
}

// The Java class OpenGLESGraphicsFactory has subclasses so its variables are this interface
type IOpenGLESGraphicsFactory interface {
    OpenGLESGraphicsFactory_() *OpenGLESGraphicsFactory
    GetOpenGLESGraphicsInstance() IOpenGLESGraphics
}

func (this *OpenGLESGraphicsFactory) OpenGLESGraphicsFactory_() *OpenGLESGraphicsFactory {
    return this
}


func NewOpenGLESGraphicsFactory() *OpenGLESGraphicsFactory {
    this := &OpenGLESGraphicsFactory{}
    return this
}


func (this *OpenGLESGraphicsFactory) GetOpenGLESGraphicsInstance() IOpenGLESGraphics {
        return nil
}

