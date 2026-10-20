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

type GL11 interface {
    GL10
    GlBindBuffer(target int, buffer int)
    GlBufferData(target int, size int, data *Buffer, usage int)
    GlBufferSubData(target int, offset int, size int, data *Buffer)
    GlClipPlanef(args ...any)
    GlColor4ub(red int8, green int8, blue int8, alpha int8)
    GlColorPointer(size int, type_ int, stride int, offset int)
    GlDeleteBuffers(args ...any)
    GlDrawElements(mode int, count int, type_ int, offset int)
    GlGenBuffers(args ...any)
    GlGetBufferParameteriv(args ...any)
    GlGetClipPlanef(args ...any)
    GlGetFloatv(args ...any)
    GlGetLightfv(args ...any)
    GlGetMaterialfv(args ...any)
    GlGetTexEnviv(args ...any)
    GlGetTexParameterfv(args ...any)
    GlGetTexParameteriv(args ...any)
    GlIsBuffer(buffer int) bool
    GlIsEnabled(cap_ int) bool
    GlIsTexture(texture int) bool
    GlNormalPointer(type_ int, stride int, offset int)
    GlPointParameterf(pname int, param float32)
    GlPointParameterfv(args ...any)
    GlTexCoordPointer(size int, type_ int, stride int, offset int)
    GlTexEnvi(target int, pname int, param int)
    GlTexEnviv(args ...any)
    GlTexParameterfv(args ...any)
    GlTexParameteriv(args ...any)
    GlVertexPointer(size int, type_ int, stride int, offset int)
}

var GL11_GL_ACTIVE_TEXTURE int = 34016
var GL11_GL_ADD_SIGNED int = 34164
var GL11_GL_ALPHA_SCALE int = 3356
var GL11_GL_ALPHA_TEST_FUNC int = 3009
var GL11_GL_ALPHA_TEST_REF int = 3010
var GL11_GL_ARRAY_BUFFER int = 34962
var GL11_GL_ARRAY_BUFFER_BINDING int = 34964
var GL11_GL_BLEND_DST int = 3040
var GL11_GL_BLEND_SRC int = 3041
var GL11_GL_BUFFER_ACCESS int = 35003
var GL11_GL_BUFFER_SIZE int = 34660
var GL11_GL_BUFFER_USAGE int = 34661
var GL11_GL_CLIENT_ACTIVE_TEXTURE int = 34017
var GL11_GL_CLIP_PLANE0 int = 12288
var GL11_GL_CLIP_PLANE1 int = 12289
var GL11_GL_CLIP_PLANE2 int = 12290
var GL11_GL_CLIP_PLANE3 int = 12291
var GL11_GL_CLIP_PLANE4 int = 12292
var GL11_GL_CLIP_PLANE5 int = 12293
var GL11_GL_COLOR_ARRAY_BUFFER_BINDING int = 34968
var GL11_GL_COLOR_ARRAY_POINTER int = 32912
var GL11_GL_COLOR_ARRAY_SIZE int = 32897
var GL11_GL_COLOR_ARRAY_STRIDE int = 32899
var GL11_GL_COLOR_ARRAY_TYPE int = 32898
var GL11_GL_COLOR_CLEAR_VALUE int = 3106
var GL11_GL_COLOR_WRITEMASK int = 3107
var GL11_GL_COMBINE int = 34160
var GL11_GL_COMBINE_ALPHA int = 34162
var GL11_GL_COMBINE_RGB int = 34161
var GL11_GL_CONSTANT int = 34166
var GL11_GL_COORD_REPLACE_OES int = 34914
var GL11_GL_CULL_FACE_MODE int = 2885
var GL11_GL_CURRENT_COLOR int = 2816
var GL11_GL_CURRENT_NORMAL int = 2818
var GL11_GL_CURRENT_TEXTURE_COORDS int = 2819
var GL11_GL_DEPTH_CLEAR_VALUE int = 2931
var GL11_GL_DEPTH_FUNC int = 2932
var GL11_GL_DEPTH_RANGE int = 2928
var GL11_GL_DEPTH_WRITEMASK int = 2930
var GL11_GL_DOT3_RGB int = 34478
var GL11_GL_DOT3_RGBA int = 34479
var GL11_GL_DYNAMIC_DRAW int = 35048
var GL11_GL_ELEMENT_ARRAY_BUFFER int = 34963
var GL11_GL_ELEMENT_ARRAY_BUFFER_BINDING int = 34965
var GL11_GL_FRONT_FACE int = 2886
var GL11_GL_GENERATE_MIPMAP int = 33169
var GL11_GL_GENERATE_MIPMAP_HINT int = 33170
var GL11_GL_INTERPOLATE int = 34165
var GL11_GL_LINE_WIDTH int = 2849
var GL11_GL_LOGIC_OP_MODE int = 3056
var GL11_GL_MATRIX_MODE int = 2976
var GL11_GL_MAX_CLIP_PLANES int = 3378
var GL11_GL_MODELVIEW_MATRIX int = 2982
var GL11_GL_MODELVIEW_MATRIX_FLOAT_AS_INT_BITS_OES int = 35213
var GL11_GL_MODELVIEW_STACK_DEPTH int = 2979
var GL11_GL_NORMAL_ARRAY_BUFFER_BINDING int = 34967
var GL11_GL_NORMAL_ARRAY_POINTER int = 32911
var GL11_GL_NORMAL_ARRAY_STRIDE int = 32895
var GL11_GL_NORMAL_ARRAY_TYPE int = 32894
var GL11_GL_OPERAND0_ALPHA int = 34200
var GL11_GL_OPERAND0_RGB int = 34192
var GL11_GL_OPERAND1_ALPHA int = 34201
var GL11_GL_OPERAND1_RGB int = 34193
var GL11_GL_OPERAND2_ALPHA int = 34202
var GL11_GL_OPERAND2_RGB int = 34194
var GL11_GL_POINT_DISTANCE_ATTENUATION int = 33065
var GL11_GL_POINT_FADE_THRESHOLD_SIZE int = 33064
var GL11_GL_POINT_SIZE int = 2833
var GL11_GL_POINT_SIZE_ARRAY_BUFFER_BINDING_OES int = 35743
var GL11_GL_POINT_SIZE_ARRAY_OES int = 35740
var GL11_GL_POINT_SIZE_ARRAY_POINTER_OES int = 35212
var GL11_GL_POINT_SIZE_ARRAY_STRIDE_OES int = 35211
var GL11_GL_POINT_SIZE_ARRAY_TYPE_OES int = 35210
var GL11_GL_POINT_SIZE_MAX int = 33063
var GL11_GL_POINT_SIZE_MIN int = 33062
var GL11_GL_POINT_SPRITE_OES int = 34913
var GL11_GL_POLYGON_OFFSET_FACTOR int = 32824
var GL11_GL_POLYGON_OFFSET_UNITS int = 10752
var GL11_GL_PREVIOUS int = 34168
var GL11_GL_PRIMARY_COLOR int = 34167
var GL11_GL_PROJECTION_MATRIX int = 2983
var GL11_GL_PROJECTION_MATRIX_FLOAT_AS_INT_BITS_OES int = 35214
var GL11_GL_PROJECTION_STACK_DEPTH int = 2980
var GL11_GL_RGB_SCALE int = 34163
var GL11_GL_SAMPLE_BUFFERS int = 32936
var GL11_GL_SAMPLE_COVERAGE_INVERT int = 32939
var GL11_GL_SAMPLE_COVERAGE_VALUE int = 32938
var GL11_GL_SAMPLES int = 32937
var GL11_GL_SCISSOR_BOX int = 3088
var GL11_GL_SHADE_MODEL int = 2900
var GL11_GL_SRC0_ALPHA int = 34184
var GL11_GL_SRC0_RGB int = 34176
var GL11_GL_SRC1_ALPHA int = 34185
var GL11_GL_SRC1_RGB int = 34177
var GL11_GL_SRC2_ALPHA int = 34186
var GL11_GL_SRC2_RGB int = 34178
var GL11_GL_STATIC_DRAW int = 35044
var GL11_GL_STENCIL_CLEAR_VALUE int = 2961
var GL11_GL_STENCIL_FAIL int = 2964
var GL11_GL_STENCIL_FUNC int = 2962
var GL11_GL_STENCIL_PASS_DEPTH_FAIL int = 2965
var GL11_GL_STENCIL_PASS_DEPTH_PASS int = 2966
var GL11_GL_STENCIL_REF int = 2967
var GL11_GL_STENCIL_VALUE_MASK int = 2963
var GL11_GL_STENCIL_WRITEMASK int = 2968
var GL11_GL_SUBTRACT int = 34023
var GL11_GL_TEXTURE_BINDING_2D int = 32873
var GL11_GL_TEXTURE_COORD_ARRAY_BUFFER_BINDING int = 34970
var GL11_GL_TEXTURE_COORD_ARRAY_POINTER int = 32914
var GL11_GL_TEXTURE_COORD_ARRAY_SIZE int = 32904
var GL11_GL_TEXTURE_COORD_ARRAY_STRIDE int = 32906
var GL11_GL_TEXTURE_COORD_ARRAY_TYPE int = 32905
var GL11_GL_TEXTURE_MATRIX int = 2984
var GL11_GL_TEXTURE_MATRIX_FLOAT_AS_INT_BITS_OES int = 35215
var GL11_GL_TEXTURE_STACK_DEPTH int = 2981
var GL11_GL_VERTEX_ARRAY_BUFFER_BINDING int = 34966
var GL11_GL_VERTEX_ARRAY_POINTER int = 32910
var GL11_GL_VERTEX_ARRAY_SIZE int = 32890
var GL11_GL_VERTEX_ARRAY_STRIDE int = 32892
var GL11_GL_VERTEX_ARRAY_TYPE int = 32891
var GL11_GL_VIEWPORT int = 2978
var GL11_GL_WRITE_ONLY int = 35001

