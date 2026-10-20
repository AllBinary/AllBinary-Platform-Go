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

type OpenGLGLUUtils struct {
}

var OpenGLGLUUtils_Instance *OpenGLGLUUtils = NewOpenGLGLUUtils()

func NewOpenGLGLUUtils() *OpenGLGLUUtils {
    this := &OpenGLGLUUtils{}
    return this
}


func OpenGLGLUUtils_GetInstance() *OpenGLGLUUtils {
        return OpenGLGLUUtils_Instance
}


func (this *OpenGLGLUUtils) MakeByteBuffer(bmp *OpenGLBitmap, flip bool) *ByteBuffer {
        panic(errors.New(fmt.Sprint()))
}


func (this *OpenGLGLUUtils) TexImage2D_1(gl GL10, glVersion string, target int, level int, generateMipMap bool, openGLESImage IOpenGLESImage, border int, flip bool) {
        panic(errors.New(fmt.Sprint()))
}

// Java overloads of OpenGLGLUUtils.texImage2D
func (this *OpenGLGLUUtils) TexImage2D(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 8 && Java_IsArg[GL10](args[0], pass) && Java_IsArg[string](args[1], pass) && Java_IsArg[int](args[2], pass) && Java_IsArg[int](args[3], pass) && Java_IsArg[bool](args[4], pass) && Java_IsArg[IOpenGLESImage](args[5], pass) && Java_IsArg[int](args[6], pass) && Java_IsArg[bool](args[7], pass) {
            this.TexImage2D_1(Java_Arg[GL10](args[0]), Java_Arg[string](args[1]), Java_Arg[int](args[2]), Java_Arg[int](args[3]), Java_Arg[bool](args[4]), Java_Arg[IOpenGLESImage](args[5]), Java_Arg[int](args[6]), Java_Arg[bool](args[7]))
            return
        }
        if len(args) == 8 && Java_IsArg[GL10](args[0], pass) && Java_IsArg[string](args[1], pass) && Java_IsArg[int](args[2], pass) && Java_IsArg[int](args[3], pass) && Java_IsArg[bool](args[4], pass) && Java_IsArg[*OpenGLBitmap](args[5], pass) && Java_IsArg[int](args[6], pass) && Java_IsArg[bool](args[7], pass) {
            this.TexImage2D_2(Java_Arg[GL10](args[0]), Java_Arg[string](args[1]), Java_Arg[int](args[2]), Java_Arg[int](args[3]), Java_Arg[bool](args[4]), Java_Arg[*OpenGLBitmap](args[5]), Java_Arg[int](args[6]), Java_Arg[bool](args[7]))
            return
        }
    }
    panic("No Java overload of OpenGLGLUUtils.texImage2D for the arguments")
}


func (this *OpenGLGLUUtils) TexImage2D_2(gl GL10, glVersion string, target int, level int, generateMipMap bool, bitmap *OpenGLBitmap, border int, flip bool) {
        panic(errors.New(fmt.Sprint()))
}

