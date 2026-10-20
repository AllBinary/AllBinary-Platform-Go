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

type OpenGLGLU struct {
}

var OpenGLGLU_Instance *OpenGLGLU = NewOpenGLGLU()

func NewOpenGLGLU() *OpenGLGLU {
    this := &OpenGLGLU{}
    return this
}


func OpenGLGLU_GetInstance() *OpenGLGLU {
        return OpenGLGLU_Instance
}


func (this *OpenGLGLU) GluLookAt(gl GL10, eyeX float32, eyeY float32, eyeZ float32, centerX float32, centerY float32, centerZ float32, upX float32, upY float32, upZ float32) {
        panic(errors.New(fmt.Sprint()))
}


func (this *OpenGLGLU) GluProject(gl GL10, objX float32, objY float32, objZ float32, model []float32, modelOffset int, project []float32, projectOffset int, view []int, viewOffset int, win []float32, winOffset int) int {
        panic(errors.New(fmt.Sprint()))
}


func (this *OpenGLGLU) GluUnProject(gl GL10, winX float32, winY float32, winZ float32, model []float32, modelOffset int, project []float32, projectOffset int, view []int, viewOffset int, obj []float32, objOffset int) int {
        panic(errors.New(fmt.Sprint()))
}


func (this *OpenGLGLU) GlFrustumf(gl GL10, lt float32, rt float32, btm float32, top float32, zNear float32, zFar float32) {
        panic(errors.New(fmt.Sprint()))
}

