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

type OpenGLThreadUtil struct {
}

var OpenGLThreadUtil_Instance *OpenGLThreadUtil = NewOpenGLThreadUtil()

func NewOpenGLThreadUtil() *OpenGLThreadUtil {
    this := &OpenGLThreadUtil{}
    return this
}


func OpenGLThreadUtil_GetInstance() *OpenGLThreadUtil {
        return OpenGLThreadUtil_Instance
}


func (this *OpenGLThreadUtil) Set(glSurfaceView any) {
}


func (this *OpenGLThreadUtil) OnPause() {
}


func (this *OpenGLThreadUtil) OnResume() {
}

