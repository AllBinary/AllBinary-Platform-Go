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

type GL10 interface {
    GL
    GlActiveTexture(texture int)
    GlAlphaFunc(func_ int, ref float32)
    GlBindTexture(target int, texture int)
    GlBlendFunc(sfactor int, dfactor int)
    GlClear(mask int)
    GlClearColor(red float32, green float32, blue float32, alpha float32)
    GlClearDepthf(depth float32)
    GlClearStencil(s int)
    GlClientActiveTexture(texture int)
    GlColor4f(red float32, green float32, blue float32, alpha float32)
    GlColorMask(red bool, green bool, blue bool, alpha bool)
    GlColorPointer(size int, type_ int, stride int, pointer *Buffer)
    GlCompressedTexImage2D(target int, level int, internalformat int, width int, height int, border int, imageSize int, data *Buffer)
    GlCompressedTexSubImage2D(target int, level int, xoffset int, yoffset int, width int, height int, format int, imageSize int, data *Buffer)
    GlCopyTexImage2D(target int, level int, internalformat int, x int, y int, width int, height int, border int)
    GlCopyTexSubImage2D(target int, level int, xoffset int, yoffset int, x int, y int, width int, height int)
    GlCullFace(mode int)
    GlDeleteTextures(args ...any)
    GlDepthFunc(func_ int)
    GlDepthMask(flag bool)
    GlDepthRangef(zNear float32, zFar float32)
    GlDisable(cap_ int)
    GlDisableClientState(array int)
    GlDrawArrays(mode int, first int, count int)
    GlDrawElements(mode int, count int, type_ int, indices *Buffer)
    GlEnable(cap_ int)
    GlEnableClientState(array int)
    GlFinish()
    GlFlush()
    GlFogf(pname int, param float32)
    GlFogfv(args ...any)
    GlFrontFace(mode int)
    GlFrustumf(left float32, right float32, bottom float32, top float32, zNear float32, zFar float32)
    GlGenTextures(args ...any)
    GlGetError() int
    GlGetIntegerv(args ...any)
    GlGetString(name int) string
    GlHint(target int, mode int)
    GlLightModelf(pname int, param float32)
    GlLightModelfv(args ...any)
    GlLightf(light int, pname int, param float32)
    GlLightfv(args ...any)
    GlLineWidth(width float32)
    GlLoadIdentity()
    GlLoadMatrixf(args ...any)
    GlLogicOp(opcode int)
    GlMaterialf(face int, pname int, param float32)
    GlMaterialfv(args ...any)
    GlMatrixMode(mode int)
    GlMultMatrixf(args ...any)
    GlMultiTexCoord4f(target int, s float32, t float32, r float32, q float32)
    GlNormal3f(nx float32, ny float32, nz float32)
    GlNormalPointer(type_ int, stride int, pointer *Buffer)
    GlOrthof(left float32, right float32, bottom float32, top float32, zNear float32, zFar float32)
    GlPixelStorei(pname int, param int)
    GlPointSize(size float32)
    GlPolygonOffset(factor float32, units float32)
    GlPopMatrix()
    GlPushMatrix()
    GlReadPixels(x int, y int, width int, height int, format int, type_ int, pixels *Buffer)
    GlRotatef(angle float32, x float32, y float32, z float32)
    GlRotatex(angle int, x int, y int, z int)
    GlSampleCoverage(value float32, invert bool)
    GlScalef(x float32, y float32, z float32)
    GlScalex(x int, y int, z int)
    GlScissor(x int, y int, width int, height int)
    GlShadeModel(mode int)
    GlStencilFunc(func_ int, ref int, mask int)
    GlStencilMask(mask int)
    GlStencilOp(fail int, zfail int, zpass int)
    GlTexCoordPointer(size int, type_ int, stride int, pointer *Buffer)
    GlTexEnvf(target int, pname int, param float32)
    GlTexEnvfv(args ...any)
    GlTexImage2D(target int, level int, internalformat int, width int, height int, border int, format int, type_ int, pixels *Buffer)
    GlTexParameterf(target int, pname int, param float32)
    GlTexParameteri(target int, pname int, param int)
    GlTexSubImage2D(target int, level int, xoffset int, yoffset int, width int, height int, format int, type_ int, pixels *Buffer)
    GlTranslatef(x float32, y float32, z float32)
    GlVertexPointer(size int, type_ int, stride int, pointer *Buffer)
    GlViewport(x int, y int, width int, height int)
    GlBegin(i int)
    GlEnd()
    GlVertex3f(f float32, f2 float32, f3 float32)
    GlPolygonMode(i int, i1 int)
}

var GL10_GL_ADD int = 260
var GL10_GL_ALIASED_LINE_WIDTH_RANGE int = 33902
var GL10_GL_ALIASED_POINT_SIZE_RANGE int = 33901
var GL10_GL_ALPHA int = 6406
var GL10_GL_ALPHA_BITS int = 3413
var GL10_GL_ALPHA_TEST int = 3008
var GL10_GL_ALWAYS int = 519
var GL10_GL_AMBIENT int = 4608
var GL10_GL_AMBIENT_AND_DIFFUSE int = 5634
var GL10_GL_AND int = 5377
var GL10_GL_AND_INVERTED int = 5380
var GL10_GL_AND_REVERSE int = 5378
var GL10_GL_BACK int = 1029
var GL10_GL_BLEND int = 3042
var GL10_GL_BLUE_BITS int = 3412
var GL10_GL_BYTE int = 5120
var GL10_GL_CCW int = 2305
var GL10_GL_CLAMP_TO_EDGE int = 33071
var GL10_GL_CLEAR int = 5376
var GL10_GL_COLOR_ARRAY int = 32886
var GL10_GL_COLOR_BUFFER_BIT int = 16384
var GL10_GL_COLOR_LOGIC_OP int = 3058
var GL10_GL_COLOR_MATERIAL int = 2903
var GL10_GL_COMPRESSED_TEXTURE_FORMATS int = 34467
var GL10_GL_CONSTANT_ATTENUATION int = 4615
var GL10_GL_COPY int = 5379
var GL10_GL_COPY_INVERTED int = 5388
var GL10_GL_CULL_FACE int = 2884
var GL10_GL_CW int = 2304
var GL10_GL_DECAL int = 8449
var GL10_GL_DECR int = 7683
var GL10_GL_DEPTH_BITS int = 3414
var GL10_GL_DEPTH_BUFFER_BIT int = 256
var GL10_GL_DEPTH_TEST int = 2929
var GL10_GL_DIFFUSE int = 4609
var GL10_GL_DITHER int = 3024
var GL10_GL_DONT_CARE int = 4352
var GL10_GL_DST_ALPHA int = 772
var GL10_GL_DST_COLOR int = 774
var GL10_GL_EMISSION int = 5632
var GL10_GL_EQUAL int = 514
var GL10_GL_EQUIV int = 5385
var GL10_GL_EXP int = 2048
var GL10_GL_EXP2 int = 2049
var GL10_GL_EXTENSIONS int = 7939
var GL10_GL_FALSE int = 0
var GL10_GL_FASTEST int = 4353
var GL10_GL_FIXED int = 5132
var GL10_GL_FLAT int = 7424
var GL10_GL_FLOAT int = 5126
var GL10_GL_FOG int = 2912
var GL10_GL_FOG_COLOR int = 2918
var GL10_GL_FOG_DENSITY int = 2914
var GL10_GL_FOG_END int = 2916
var GL10_GL_FOG_HINT int = 3156
var GL10_GL_FOG_MODE int = 2917
var GL10_GL_FOG_START int = 2915
var GL10_GL_FRONT int = 1028
var GL10_GL_FRONT_AND_BACK int = 1032
var GL10_GL_GEQUAL int = 518
var GL10_GL_GREATER int = 516
var GL10_GL_GREEN_BITS int = 3411
var GL10_GL_IMPLEMENTATION_COLOR_READ_FORMAT_OES int = 35739
var GL10_GL_IMPLEMENTATION_COLOR_READ_TYPE_OES int = 35738
var GL10_GL_INCR int = 7682
var GL10_GL_INVALID_ENUM int = 1280
var GL10_GL_INVALID_OPERATION int = 1282
var GL10_GL_INVALID_VALUE int = 1281
var GL10_GL_INVERT int = 5386
var GL10_GL_KEEP int = 7680
var GL10_GL_LEQUAL int = 515
var GL10_GL_LESS int = 513
var GL10_GL_LIGHT_MODEL_AMBIENT int = 2899
var GL10_GL_LIGHT_MODEL_TWO_SIDE int = 2898
var GL10_GL_LIGHT0 int = 16384
var GL10_GL_LIGHT1 int = 16385
var GL10_GL_LIGHT2 int = 16386
var GL10_GL_LIGHT3 int = 16387
var GL10_GL_LIGHT4 int = 16388
var GL10_GL_LIGHT5 int = 16389
var GL10_GL_LIGHT6 int = 16390
var GL10_GL_LIGHT7 int = 16391
var GL10_GL_LIGHTING int = 2896
var GL10_GL_LINE_LOOP int = 2
var GL10_GL_LINE_SMOOTH int = 2848
var GL10_GL_LINE_SMOOTH_HINT int = 3154
var GL10_GL_LINE_STRIP int = 3
var GL10_GL_LINEAR int = 9729
var GL10_GL_LINEAR_ATTENUATION int = 4616
var GL10_GL_LINEAR_MIPMAP_LINEAR int = 9987
var GL10_GL_LINEAR_MIPMAP_NEAREST int = 9985
var GL10_GL_LINES int = 1
var GL10_GL_LUMINANCE int = 6409
var GL10_GL_LUMINANCE_ALPHA int = 6410
var GL10_GL_MAX_ELEMENTS_INDICES int = 33001
var GL10_GL_MAX_ELEMENTS_VERTICES int = 33000
var GL10_GL_MAX_LIGHTS int = 3377
var GL10_GL_MAX_MODELVIEW_STACK_DEPTH int = 3382
var GL10_GL_MAX_PROJECTION_STACK_DEPTH int = 3384
var GL10_GL_MAX_TEXTURE_SIZE int = 3379
var GL10_GL_MAX_TEXTURE_STACK_DEPTH int = 3385
var GL10_GL_MAX_TEXTURE_UNITS int = 34018
var GL10_GL_MAX_VIEWPORT_DIMS int = 3386
var GL10_GL_MODELVIEW int = 5888
var GL10_GL_MODULATE int = 8448
var GL10_GL_MULTISAMPLE int = 32925
var GL10_GL_NAND int = 5390
var GL10_GL_NEAREST int = 9728
var GL10_GL_NEAREST_MIPMAP_LINEAR int = 9986
var GL10_GL_NEAREST_MIPMAP_NEAREST int = 9984
var GL10_GL_NEVER int = 512
var GL10_GL_NICEST int = 4354
var GL10_GL_NO_ERROR int = 0
var GL10_GL_NOOP int = 5381
var GL10_GL_NOR int = 5384
var GL10_GL_NORMAL_ARRAY int = 32885
var GL10_GL_NORMALIZE int = 2977
var GL10_GL_NOTEQUAL int = 517
var GL10_GL_NUM_COMPRESSED_TEXTURE_FORMATS int = 34466
var GL10_GL_ONE int = 1
var GL10_GL_ONE_MINUS_DST_ALPHA int = 773
var GL10_GL_ONE_MINUS_DST_COLOR int = 775
var GL10_GL_ONE_MINUS_SRC_ALPHA int = 771
var GL10_GL_ONE_MINUS_SRC_COLOR int = 769
var GL10_GL_OR int = 5383
var GL10_GL_OR_INVERTED int = 5389
var GL10_GL_OR_REVERSE int = 5387
var GL10_GL_OUT_OF_MEMORY int = 1285
var GL10_GL_PACK_ALIGNMENT int = 3333
var GL10_GL_PALETTE4_R5_G6_B5_OES int = 35730
var GL10_GL_PALETTE4_RGB5_A1_OES int = 35732
var GL10_GL_PALETTE4_RGB8_OES int = 35728
var GL10_GL_PALETTE4_RGBA4_OES int = 35731
var GL10_GL_PALETTE4_RGBA8_OES int = 35729
var GL10_GL_PALETTE8_R5_G6_B5_OES int = 35735
var GL10_GL_PALETTE8_RGB5_A1_OES int = 35737
var GL10_GL_PALETTE8_RGB8_OES int = 35733
var GL10_GL_PALETTE8_RGBA4_OES int = 35736
var GL10_GL_PALETTE8_RGBA8_OES int = 35734
var GL10_GL_PERSPECTIVE_CORRECTION_HINT int = 3152
var GL10_GL_POINT_SMOOTH int = 2832
var GL10_GL_POINT_SMOOTH_HINT int = 3153
var GL10_GL_POINTS int = 0
var GL10_GL_POINT_FADE_THRESHOLD_SIZE int = 33064
var GL10_GL_POINT_SIZE int = 2833
var GL10_GL_POLYGON_OFFSET_FILL int = 32823
var GL10_GL_POLYGON_SMOOTH_HINT int = 3155
var GL10_GL_POSITION int = 4611
var GL10_GL_PROJECTION int = 5889
var GL10_GL_QUADRATIC_ATTENUATION int = 4617
var GL10_GL_RED_BITS int = 3410
var GL10_GL_RENDERER int = 7937
var GL10_GL_REPEAT int = 10497
var GL10_GL_REPLACE int = 7681
var GL10_GL_RESCALE_NORMAL int = 32826
var GL10_GL_RGB int = 6407
var GL10_GL_RGBA int = 6408
var GL10_GL_SAMPLE_ALPHA_TO_COVERAGE int = 32926
var GL10_GL_SAMPLE_ALPHA_TO_ONE int = 32927
var GL10_GL_SAMPLE_COVERAGE int = 32928
var GL10_GL_SCISSOR_TEST int = 3089
var GL10_GL_SET int = 5391
var GL10_GL_SHININESS int = 5633
var GL10_GL_SHORT int = 5122
var GL10_GL_SMOOTH int = 7425
var GL10_GL_SMOOTH_LINE_WIDTH_RANGE int = 2850
var GL10_GL_SMOOTH_POINT_SIZE_RANGE int = 2834
var GL10_GL_SPECULAR int = 4610
var GL10_GL_SPOT_CUTOFF int = 4614
var GL10_GL_SPOT_DIRECTION int = 4612
var GL10_GL_SPOT_EXPONENT int = 4613
var GL10_GL_SRC_ALPHA int = 770
var GL10_GL_SRC_ALPHA_SATURATE int = 776
var GL10_GL_SRC_COLOR int = 768
var GL10_GL_STACK_OVERFLOW int = 1283
var GL10_GL_STACK_UNDERFLOW int = 1284
var GL10_GL_STENCIL_BITS int = 3415
var GL10_GL_STENCIL_BUFFER_BIT int = 1024
var GL10_GL_STENCIL_TEST int = 2960
var GL10_GL_SUBPIXEL_BITS int = 3408
var GL10_GL_TEXTURE int = 5890
var GL10_GL_TEXTURE_2D int = 3553
var GL10_GL_TEXTURE_COORD_ARRAY int = 32888
var GL10_GL_TEXTURE_ENV int = 8960
var GL10_GL_TEXTURE_ENV_COLOR int = 8705
var GL10_GL_TEXTURE_ENV_MODE int = 8704
var GL10_GL_TEXTURE_MAG_FILTER int = 10240
var GL10_GL_TEXTURE_MIN_FILTER int = 10241
var GL10_GL_TEXTURE_WRAP_S int = 10242
var GL10_GL_TEXTURE_WRAP_T int = 10243
var GL10_GL_TEXTURE0 int = 33984
var GL10_GL_TEXTURE1 int = 33985
var GL10_GL_TEXTURE2 int = 33986
var GL10_GL_TEXTURE3 int = 33987
var GL10_GL_TEXTURE4 int = 33988
var GL10_GL_TEXTURE5 int = 33989
var GL10_GL_TEXTURE6 int = 33990
var GL10_GL_TEXTURE7 int = 33991
var GL10_GL_TEXTURE8 int = 33992
var GL10_GL_TEXTURE9 int = 33993
var GL10_GL_TEXTURE10 int = 33994
var GL10_GL_TEXTURE11 int = 33995
var GL10_GL_TEXTURE12 int = 33996
var GL10_GL_TEXTURE13 int = 33997
var GL10_GL_TEXTURE14 int = 33998
var GL10_GL_TEXTURE15 int = 33999
var GL10_GL_TEXTURE16 int = 34000
var GL10_GL_TEXTURE17 int = 34001
var GL10_GL_TEXTURE18 int = 34002
var GL10_GL_TEXTURE19 int = 34003
var GL10_GL_TEXTURE20 int = 34004
var GL10_GL_TEXTURE21 int = 34005
var GL10_GL_TEXTURE22 int = 34006
var GL10_GL_TEXTURE23 int = 34007
var GL10_GL_TEXTURE24 int = 34008
var GL10_GL_TEXTURE25 int = 34009
var GL10_GL_TEXTURE26 int = 34010
var GL10_GL_TEXTURE27 int = 34011
var GL10_GL_TEXTURE28 int = 34012
var GL10_GL_TEXTURE29 int = 34013
var GL10_GL_TEXTURE30 int = 34014
var GL10_GL_TEXTURE31 int = 34015
var GL10_GL_TRIANGLE_FAN int = 6
var GL10_GL_TRIANGLE_STRIP int = 5
var GL10_GL_TRIANGLES int = 4
var GL10_GL_TRUE int = 1
var GL10_GL_UNPACK_ALIGNMENT int = 3317
var GL10_GL_UNSIGNED_BYTE int = 5121
var GL10_GL_UNSIGNED_SHORT int = 5123
var GL10_GL_UNSIGNED_SHORT_4_4_4_4 int = 32819
var GL10_GL_UNSIGNED_SHORT_5_5_5_1 int = 32820
var GL10_GL_UNSIGNED_SHORT_5_6_5 int = 33635
var GL10_GL_VENDOR int = 7936
var GL10_GL_VERSION int = 7938
var GL10_GL_VERTEX_ARRAY int = 32884
var GL10_GL_XOR int = 5382
var GL10_GL_ZERO int = 0

