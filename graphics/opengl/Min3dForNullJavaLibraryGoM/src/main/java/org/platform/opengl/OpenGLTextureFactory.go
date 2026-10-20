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

type OpenGLTextureFactory struct {
    PlatformTextureBaseFactory
    GluUtils *OpenGLGLUUtils
}

var OpenGLTextureFactory_Instance *OpenGLTextureFactory = NewOpenGLTextureFactory()

func (this *OpenGLTextureFactory) initFields() {
    this.GluUtils = OpenGLGLUUtils_GetInstance()
}


func OpenGLTextureFactory_GetInstance() *OpenGLTextureFactory {
        return OpenGLTextureFactory_Instance
}


func NewOpenGLTextureFactory() *OpenGLTextureFactory {
    this := &OpenGLTextureFactory{}
    this.PlatformTextureBaseFactory = *NewPlatformTextureBaseFactory()
    this.initFields()
    return this
}


func (this *OpenGLTextureFactory) Load(gl GL10, target int, level int, image IImage, border int, flip bool) {
        panic(errors.New(fmt.Sprint()))
}

