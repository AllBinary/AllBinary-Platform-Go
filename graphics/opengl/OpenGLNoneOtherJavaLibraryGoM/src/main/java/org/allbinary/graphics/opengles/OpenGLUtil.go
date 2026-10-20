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

type OpenGLUtil struct {
}

var OpenGLUtil_Instance *OpenGLUtil = NewOpenGLUtil()

func NewOpenGLUtil() *OpenGLUtil {
    this := &OpenGLUtil{}
    return this
}


func OpenGLUtil_GetInstance() *OpenGLUtil {
        return OpenGLUtil_Instance
}


func (this *OpenGLUtil) AddImage(image IImage) IImage {
        return image
}


func (this *OpenGLUtil) Clear() {
}


func (this *OpenGLUtil) AddRunnable(runnable Runnable) {
}

