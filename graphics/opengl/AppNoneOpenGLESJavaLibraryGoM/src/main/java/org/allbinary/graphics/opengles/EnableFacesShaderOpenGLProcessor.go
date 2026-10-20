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

type EnableFacesShaderOpenGLProcessor struct {
    ObjectThreedOpenGLProcessor
}

var EnableFacesShaderOpenGLProcessor_Instance *EnableFacesShaderOpenGLProcessor = NewEnableFacesShaderOpenGLProcessor()


func EnableFacesShaderOpenGLProcessor_GetInstance() *EnableFacesShaderOpenGLProcessor {
        return EnableFacesShaderOpenGLProcessor_Instance
}


func NewEnableFacesShaderOpenGLProcessor() *EnableFacesShaderOpenGLProcessor {
    this := &EnableFacesShaderOpenGLProcessor{}
    this.ObjectThreedOpenGLProcessor = *NewObjectThreedOpenGLProcessor()
    return this
}


func (this *EnableFacesShaderOpenGLProcessor) Process(gl GL10, anyType any) {
}

