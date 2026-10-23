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

type OpenGLProcessor struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
}

// The Java class OpenGLProcessor has subclasses so its variables are this interface
type IOpenGLProcessor interface {
    OpenGLProcessor_() *OpenGLProcessor
    Process(gl GL10)
}

func (this *OpenGLProcessor) OpenGLProcessor_() *OpenGLProcessor {
    return this
}


func (this *OpenGLProcessor) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
}

func NewOpenGLProcessor() *OpenGLProcessor {
    this := &OpenGLProcessor{}
    this.initFields()
    return this
}


func (this *OpenGLProcessor) Process(gl GL10) {
}

