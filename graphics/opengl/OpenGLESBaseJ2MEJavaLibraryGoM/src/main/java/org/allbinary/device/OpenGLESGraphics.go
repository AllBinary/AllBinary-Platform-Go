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

type OpenGLESGraphics struct {
    CommonStrings *CommonStrings
    Gl GL10
}

// Implements OpenGLSurfaceChangedInterface, GraphicsInterface

// The Java class OpenGLESGraphics has subclasses so its variables are this interface
type IOpenGLESGraphics interface {
    OpenGLESGraphics_() *OpenGLESGraphics
    UpdateGL(gl GL)
    Set(gl GL)
    Init()
    Update()
    SetOrthoginalMode()
    SetCameraMode()
    GetGlP() GL10
    SetFontDebug(font Font, fontDebug IFontDebugBase)
}

func (this *OpenGLESGraphics) OpenGLESGraphics_() *OpenGLESGraphics {
    return this
}


func (this *OpenGLESGraphics) initFields() {
    this.CommonStrings = CommonStrings_GetInstance()
    this.Gl = NullGL10_NULL_GL10
}


func NewOpenGLESGraphics() *OpenGLESGraphics {
    this := &OpenGLESGraphics{}
    this.initFields()
    return this
}


func (this *OpenGLESGraphics) UpdateGL(gl GL) {
}


func (this *OpenGLESGraphics) Set(gl GL) {
}


// throws Exception
func (this *OpenGLESGraphics) Init() {
}


// throws Exception
func (this *OpenGLESGraphics) Update() {
}


func (this *OpenGLESGraphics) SetOrthoginalMode() {
}


func (this *OpenGLESGraphics) SetCameraMode() {
}


func (this *OpenGLESGraphics) GetGlP() GL10 {
        return this.Gl
}


func (this *OpenGLESGraphics) SetFontDebug(font Font, fontDebug IFontDebugBase) {
        panic(errors.New(fmt.Sprint()))
}

