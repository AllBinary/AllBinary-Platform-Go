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

type OpenGLObject3dProcessor struct {
    CommonStrings *CommonStrings
}

// The Java class OpenGLObject3dProcessor has subclasses so its variables are this interface
type IOpenGLObject3dProcessor interface {
    OpenGLObject3dProcessor_() *OpenGLObject3dProcessor
    Process(args ...any)
}

func (this *OpenGLObject3dProcessor) OpenGLObject3dProcessor_() *OpenGLObject3dProcessor {
    return this
}


func (this *OpenGLObject3dProcessor) initFields() {
    this.CommonStrings = CommonStrings_GetInstance()
}

func NewOpenGLObject3dProcessor() *OpenGLObject3dProcessor {
    this := &OpenGLObject3dProcessor{}
    this.initFields()
    return this
}


func (this *OpenGLObject3dProcessor) Process_1(gl GL10, object3d any) {
}

// Java overloads of OpenGLObject3dProcessor.process
func (this *OpenGLObject3dProcessor) Process(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 2 && Java_IsArg[GL10](args[0], pass) && Java_IsArg[any](args[1], pass) {
            this.Process_1(Java_Arg[GL10](args[0]), Java_Arg[any](args[1]))
            return
        }
        if len(args) == 3 && Java_IsArg[GL10](args[0], pass) && Java_IsArg[any](args[1], pass) && Java_IsArg[int](args[2], pass) {
            this.Process_2(Java_Arg[GL10](args[0]), Java_Arg[any](args[1]), Java_Arg[int](args[2]))
            return
        }
    }
    panic("No Java overload of OpenGLObject3dProcessor.process for the arguments")
}


func (this *OpenGLObject3dProcessor) Process_2(gl GL10, object3d any, index int) {
}

