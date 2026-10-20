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

type ShaderOpenGLProcessor struct {
    OpenGLProcessor
}

// The Java class ShaderOpenGLProcessor has subclasses so its variables are this interface
type IShaderOpenGLProcessor interface {
    IOpenGLProcessor
    ShaderOpenGLProcessor_() *ShaderOpenGLProcessor
    Process0(gl GL10)
    Process1(gl GL10)
    Process(gl GL10)
    Process2(gl GL10)
    Process3(gl GL10)
}

func (this *ShaderOpenGLProcessor) ShaderOpenGLProcessor_() *ShaderOpenGLProcessor {
    return this
}



func NewShaderOpenGLProcessor() *ShaderOpenGLProcessor {
    this := &ShaderOpenGLProcessor{}
    this.OpenGLProcessor = *NewOpenGLProcessor()
    return this
}


func (this *ShaderOpenGLProcessor) Process0(gl GL10) {
}


func (this *ShaderOpenGLProcessor) Process1(gl GL10) {
}


func (this *ShaderOpenGLProcessor) Process(gl GL10) {
}


func (this *ShaderOpenGLProcessor) Process2(gl GL10) {
}


func (this *ShaderOpenGLProcessor) Process3(gl GL10) {
}

