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

type GL11Ext interface {
    GL
    GlTexParameterfv(target int, pname int, param []float32, offset int)
    GlCurrentPaletteMatrixOES(matrixpaletteindex int)
    GlDrawTexfOES(x float32, y float32, z float32, width float32, height float32)
    GlDrawTexfvOES(args ...any)
    GlDrawTexiOES(x int, y int, z int, width int, height int)
    GlDrawTexivOES(args ...any)
    GlDrawTexsOES(x int16, y int16, z int16, width int16, height int16)
    GlDrawTexsvOES(args ...any)
    GlDrawTexxOES(x int, y int, z int, width int, height int)
    GlDrawTexxvOES(args ...any)
    GlEnable(cap_ int)
    GlEnableClientState(array int)
    GlLoadPaletteFromModelViewMatrixOES()
    GlMatrixIndexPointerOES(args ...any)
    GlWeightPointerOES(args ...any)
}

var GL11Ext_GL_MATRIX_INDEX_ARRAY_BUFFER_BINDING_OES int = 35742
var GL11Ext_GL_MATRIX_INDEX_ARRAY_OES int = 34884
var GL11Ext_GL_MATRIX_INDEX_ARRAY_POINTER_OES int = 34889
var GL11Ext_GL_MATRIX_INDEX_ARRAY_SIZE_OES int = 34886
var GL11Ext_GL_MATRIX_INDEX_ARRAY_STRIDE_OES int = 34888
var GL11Ext_GL_MATRIX_INDEX_ARRAY_TYPE_OES int = 34887
var GL11Ext_GL_MATRIX_PALETTE_OES int = 34880
var GL11Ext_GL_MAX_PALETTE_MATRICES_OES int = 34882
var GL11Ext_GL_MAX_VERTEX_UNITS_OES int = 34468
var GL11Ext_GL_TEXTURE_CROP_RECT_OES int = 35741
var GL11Ext_GL_WEIGHT_ARRAY_BUFFER_BINDING_OES int = 34974
var GL11Ext_GL_WEIGHT_ARRAY_OES int = 34477
var GL11Ext_GL_WEIGHT_ARRAY_POINTER_OES int = 34476
var GL11Ext_GL_WEIGHT_ARRAY_SIZE_OES int = 34475
var GL11Ext_GL_WEIGHT_ARRAY_STRIDE_OES int = 34474
var GL11Ext_GL_WEIGHT_ARRAY_TYPE_OES int = 34473

