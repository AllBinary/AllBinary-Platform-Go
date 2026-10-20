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

type OpenGLLogUtil struct {
    MAX_TEXTURE string
    LOG_ERROR string
    IMAGE string
}

var OpenGLLogUtil_Instance *OpenGLLogUtil = NewOpenGLLogUtil()

func (this *OpenGLLogUtil) initFields() {
    this.MAX_TEXTURE = " Max Texture Size: "
    this.LOG_ERROR = "logError"
    this.IMAGE = " Image: "
}

func NewOpenGLLogUtil() *OpenGLLogUtil {
    this := &OpenGLLogUtil{}
    this.initFields()
    return this
}


func OpenGLLogUtil_GetInstance() *OpenGLLogUtil {
        return OpenGLLogUtil_Instance
}


func (this *OpenGLLogUtil) LogError_1(gl GL10) {
        var error_ int = gl.GlGetError(); _ = error_
        if error_ != GL10.GL_NO_ERROR {
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append(OpenGLStrings_GetInstance().GL_ERROR_LABEL)
            stringBuffer.Appendint(error_)
            if error_ == 1281 {
                stringBuffer.Append(this.MAX_TEXTURE)
                var maxTextureSize []int = make([]int, 1); _ = maxTextureSize
                gl.GlGetIntegerv(GL10.GL_MAX_TEXTURE_SIZE, maxTextureSize, 0)
                stringBuffer.Appendint(maxTextureSize[0])
            }
            if error_ == 1280 {
                PreLogUtil_Put(fmt.Sprint(stringBuffer), this, OpenGLStrings_GetInstance().SET)
            } else {
                ForcedLogUtil_Log(fmt.Sprint(stringBuffer), this)
            }
        }
}

// Java overloads of OpenGLLogUtil.logError
func (this *OpenGLLogUtil) LogError(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[GL10](args[0], pass) {
            this.LogError_1(Java_Arg[GL10](args[0]))
            return
        }
        if len(args) == 2 && Java_IsArg[GL10](args[0], pass) && Java_IsArg[IImage](args[1], pass) {
            this.LogError_2(Java_Arg[GL10](args[0]), Java_Arg[IImage](args[1]))
            return
        }
    }
    panic("No Java overload of OpenGLLogUtil.logError for the arguments")
}


func (this *OpenGLLogUtil) LogError_2(gl GL10, image IImage) {
        var error_ int = gl.GlGetError(); _ = error_
        if error_ != GL10.GL_NO_ERROR {
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append(OpenGLStrings_GetInstance().GL_ERROR_LABEL)
            stringBuffer.Appendint(error_)
            if error_ == 1281 {
                stringBuffer.Append(this.MAX_TEXTURE)
                var maxTextureSize []int = make([]int, 1); _ = maxTextureSize
                gl.GlGetIntegerv(GL10.GL_MAX_TEXTURE_SIZE, maxTextureSize, 0)
                stringBuffer.Appendint(maxTextureSize[0])
            }
            stringBuffer.Append(this.IMAGE)
            stringBuffer.Append(fmt.Sprint(image))
            PreLogUtil_Put(fmt.Sprint(stringBuffer), this, this.LOG_ERROR)
        }
}

