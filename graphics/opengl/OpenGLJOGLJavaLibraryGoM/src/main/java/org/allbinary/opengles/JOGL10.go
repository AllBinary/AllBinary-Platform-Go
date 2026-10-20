/*
 *
 * AllBinary Open License Version 1
 * Copyright (c) 2011 AllBinary
 * 
 * By agreeing to this license you and any business entity you represent are
 * legally bound to the AllBinary Open License Version 1 legal agreement.
 * 
 * You may obtain the AllBinary Open License Version 1 legal agreement from
 * AllBinary or the root directory of AllBinary's AllBinary Platform repository.
 * 
 * Created By: Travis Berthelot
 *   
 */
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

//JOGLEToKhronosGL10
type JOGL10 struct {
    RED string
    GREEN string
    BLUE string
    ALPHA string
    PNAME string
    _PNAME string
    PARAM string
    PARAMS string
    LIGHT string
    FACE string
    TARGET string
    Gl10 *GL2
    Glu *GLU
    GL_LINE int
    GL_QUAD_STRIP int
}

// Implements GL10

// The Java class JOGL10 has subclasses so its variables are this interface
type IJOGL10 interface {
    JOGL10_() *JOGL10
    GetJOGLGL() *GL2
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

func (this *JOGL10) JOGL10_() *JOGL10 {
    return this
}


func (this *JOGL10) initFields() {
    this.RED = "red: "
    this.GREEN = " green: "
    this.BLUE = " blue: "
    this.ALPHA = " alpha: "
    this.PNAME = "pname: "
    this._PNAME = " pname: "
    this.PARAM = " param: "
    this.PARAMS = " params: "
    this.LIGHT = "light: "
    this.FACE = "face: "
    this.TARGET = "target: "
    
    
    
    
}


func NewJOGL10(gl GL, gl2 *GL2) *JOGL10 {
    this := &JOGL10{}
    this.initFields()
        this.Gl10 = gl2
        this.Glu = GLU_CreateGLU(this.Gl10)
        this.GL_LINE = com_Jogamp.Opengl.GL2.GL_LINE
        this.GL_QUAD_STRIP = com_Jogamp.Opengl.GL2.GL_QUAD_STRIP
    return this
}


func (this *JOGL10) GetJOGLGL() *GL2 {
        return this.Gl10
}


func (this *JOGL10) GlActiveTexture(texture int) {
        this.Gl10.GlActiveTexture(texture)
}


func (this *JOGL10) GlAlphaFunc(func_ int, ref float32) {
        this.Gl10.GlAlphaFunc(func_, ref)
}


func (this *JOGL10) GlBindTexture(target int, texture int) {
        this.Gl10.GlBindTexture(target, texture)
}


func (this *JOGL10) GlBlendFunc(sfactor int, dfactor int) {
        this.Gl10.GlBlendFunc(sfactor, dfactor)
}


func (this *JOGL10) GlClear(mask int) {
        this.Gl10.GlClear(mask)
}


func (this *JOGL10) GlClearColor(red float32, green float32, blue float32, alpha float32) {
        this.Gl10.GlClearColor(red, green, blue, alpha)
}


func (this *JOGL10) GlClearDepthf(depth float32) {
        this.Gl10.GlClearDepthf(depth)
}


func (this *JOGL10) GlClearStencil(s int) {
        this.Gl10.GlClearStencil(s)
}


func (this *JOGL10) GlClientActiveTexture(texture int) {
        this.Gl10.GlClientActiveTexture(texture)
}


func (this *JOGL10) GlColor4f(red float32, green float32, blue float32, alpha float32) {
        this.Gl10.GlColor4f(red, green, blue, alpha)
}


func (this *JOGL10) GlColorMask(red bool, green bool, blue bool, alpha bool) {
        this.Gl10.GlColorMask(red, green, blue, alpha)
}


func (this *JOGL10) GlColorPointer(size int, type_ int, stride int, pointer *Buffer) {
        this.Gl10.GlColorPointer(size, type_, stride, pointer)
}


func (this *JOGL10) GlCompressedTexImage2D(target int, level int, internalformat int, width int, height int, border int, imageSize int, data *Buffer) {
        this.Gl10.GlCompressedTexImage2D(target, level, internalformat, width, height, border, imageSize, data)
}


func (this *JOGL10) GlCompressedTexSubImage2D(target int, level int, xoffset int, yoffset int, width int, height int, format int, imageSize int, data *Buffer) {
        this.Gl10.GlCompressedTexSubImage2D(target, level, xoffset, yoffset, width, height, format, imageSize, data)
}


func (this *JOGL10) GlCopyTexImage2D(target int, level int, internalformat int, x int, y int, width int, height int, border int) {
        this.Gl10.GlCopyTexImage2D(target, level, internalformat, x, y, width, height, border)
}


func (this *JOGL10) GlCopyTexSubImage2D(target int, level int, xoffset int, yoffset int, x int, y int, width int, height int) {
        this.Gl10.GlCopyTexSubImage2D(target, level, xoffset, yoffset, x, y, width, height)
}


func (this *JOGL10) GlCullFace(mode int) {
        this.Gl10.GlCullFace(mode)
}


func (this *JOGL10) GlDeleteTextures_1(n int, textures []int, offset int) {
        this.Gl10.GlDeleteTextures(n, textures, offset)
}

// Java overloads of JOGL10.glDeleteTextures
func (this *JOGL10) GlDeleteTextures(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 3 && Java_IsArg[int](args[0], pass) && Java_IsArg[[]int](args[1], pass) && Java_IsArg[int](args[2], pass) {
            this.GlDeleteTextures_1(Java_Arg[int](args[0]), Java_Arg[[]int](args[1]), Java_Arg[int](args[2]))
            return
        }
        if len(args) == 2 && Java_IsArg[int](args[0], pass) && Java_IsArg[*IntBuffer](args[1], pass) {
            this.GlDeleteTextures_2(Java_Arg[int](args[0]), Java_Arg[*IntBuffer](args[1]))
            return
        }
    }
    panic("No Java overload of JOGL10.glDeleteTextures for the arguments")
}


func (this *JOGL10) GlDeleteTextures_2(n int, textures *IntBuffer) {
        this.Gl10.GlDeleteTextures(n, textures)
}


func (this *JOGL10) GlDepthFunc(func_ int) {
        this.Gl10.GlDepthFunc(func_)
}


func (this *JOGL10) GlDepthMask(flag bool) {
        this.Gl10.GlDepthMask(flag)
}


func (this *JOGL10) GlDepthRangef(zNear float32, zFar float32) {
        this.Gl10.GlDepthRangef(zNear, zFar)
}


func (this *JOGL10) GlDisable(cap_ int) {
        this.Gl10.GlDisable(cap_)
}


func (this *JOGL10) GlDisableClientState(array int) {
        this.Gl10.GlDisableClientState(array)
}


func (this *JOGL10) GlDrawArrays(mode int, first int, count int) {
        this.Gl10.GlDrawArrays(mode, first, count)
}


func (this *JOGL10) GlDrawElements(mode int, count int, type_ int, indices *Buffer) {
        this.Gl10.GlDrawElements(mode, count, type_, indices)
}


func (this *JOGL10) GlEnable(cap_ int) {
        this.Gl10.GlEnable(cap_)
}


func (this *JOGL10) GlEnableClientState(array int) {
        this.Gl10.GlEnableClientState(array)
}


func (this *JOGL10) GlFinish() {
        this.Gl10.GlFinish()
}


func (this *JOGL10) GlFlush() {
        this.Gl10.GlFlush()
}


func (this *JOGL10) GlFogf(pname int, param float32) {
        this.Gl10.GlFogf(pname, param)
}


func (this *JOGL10) GlFogfv_1(pname int, params []float32, offset int) {
        this.Gl10.GlFogfv(pname, params, offset)
}

// Java overloads of JOGL10.glFogfv
func (this *JOGL10) GlFogfv(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 3 && Java_IsArg[int](args[0], pass) && Java_IsArg[[]float32](args[1], pass) && Java_IsArg[int](args[2], pass) {
            this.GlFogfv_1(Java_Arg[int](args[0]), Java_Arg[[]float32](args[1]), Java_Arg[int](args[2]))
            return
        }
        if len(args) == 2 && Java_IsArg[int](args[0], pass) && Java_IsArg[*FloatBuffer](args[1], pass) {
            this.GlFogfv_2(Java_Arg[int](args[0]), Java_Arg[*FloatBuffer](args[1]))
            return
        }
    }
    panic("No Java overload of JOGL10.glFogfv for the arguments")
}


func (this *JOGL10) GlFogfv_2(pname int, params *FloatBuffer) {
        this.Gl10.GlFogfv(pname, params)
}


func (this *JOGL10) GlFrontFace(mode int) {
        this.Gl10.GlFrontFace(mode)
}


func (this *JOGL10) GlFrustumf(left float32, right float32, bottom float32, top float32, zNear float32, zFar float32) {
        this.Gl10.GlFrustumf(left, right, bottom, top, zNear, zFar)
}


func (this *JOGL10) GlGenTextures_1(n int, textures []int, offset int) {
        this.Gl10.GlGenTextures(n, textures, offset)
}

// Java overloads of JOGL10.glGenTextures
func (this *JOGL10) GlGenTextures(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 3 && Java_IsArg[int](args[0], pass) && Java_IsArg[[]int](args[1], pass) && Java_IsArg[int](args[2], pass) {
            this.GlGenTextures_1(Java_Arg[int](args[0]), Java_Arg[[]int](args[1]), Java_Arg[int](args[2]))
            return
        }
        if len(args) == 2 && Java_IsArg[int](args[0], pass) && Java_IsArg[*IntBuffer](args[1], pass) {
            this.GlGenTextures_2(Java_Arg[int](args[0]), Java_Arg[*IntBuffer](args[1]))
            return
        }
    }
    panic("No Java overload of JOGL10.glGenTextures for the arguments")
}


func (this *JOGL10) GlGenTextures_2(n int, textures *IntBuffer) {
        this.Gl10.GlGenTextures(n, textures)
}


func (this *JOGL10) GlGetError() int {
        return this.Gl10.GlGetError()
}


func (this *JOGL10) GlGetIntegerv_1(pname int, params []int, offset int) {
        this.Gl10.GlGetIntegerv(pname, params, offset)
}

// Java overloads of JOGL10.glGetIntegerv
func (this *JOGL10) GlGetIntegerv(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 3 && Java_IsArg[int](args[0], pass) && Java_IsArg[[]int](args[1], pass) && Java_IsArg[int](args[2], pass) {
            this.GlGetIntegerv_1(Java_Arg[int](args[0]), Java_Arg[[]int](args[1]), Java_Arg[int](args[2]))
            return
        }
        if len(args) == 2 && Java_IsArg[int](args[0], pass) && Java_IsArg[*IntBuffer](args[1], pass) {
            this.GlGetIntegerv_2(Java_Arg[int](args[0]), Java_Arg[*IntBuffer](args[1]))
            return
        }
    }
    panic("No Java overload of JOGL10.glGetIntegerv for the arguments")
}


func (this *JOGL10) GlGetIntegerv_2(pname int, params *IntBuffer) {
        this.Gl10.GlGetIntegerv(pname, params)
}


func (this *JOGL10) GlGetString(name int) string {
        return this.Gl10.GlGetString(name)
}


func (this *JOGL10) GlHint(target int, mode int) {
        this.Gl10.GlHint(target, mode)
}


func (this *JOGL10) GlLightModelf(pname int, param float32) {
        this.Gl10.GlLightModelf(pname, param)
}


func (this *JOGL10) GlLightModelfv_1(pname int, params []float32, offset int) {
        this.Gl10.GlLightModelfv(pname, params, offset)
}

// Java overloads of JOGL10.glLightModelfv
func (this *JOGL10) GlLightModelfv(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 3 && Java_IsArg[int](args[0], pass) && Java_IsArg[[]float32](args[1], pass) && Java_IsArg[int](args[2], pass) {
            this.GlLightModelfv_1(Java_Arg[int](args[0]), Java_Arg[[]float32](args[1]), Java_Arg[int](args[2]))
            return
        }
        if len(args) == 2 && Java_IsArg[int](args[0], pass) && Java_IsArg[*FloatBuffer](args[1], pass) {
            this.GlLightModelfv_2(Java_Arg[int](args[0]), Java_Arg[*FloatBuffer](args[1]))
            return
        }
    }
    panic("No Java overload of JOGL10.glLightModelfv for the arguments")
}


func (this *JOGL10) GlLightModelfv_2(pname int, params *FloatBuffer) {
        this.Gl10.GlLightModelfv(pname, params)
}


func (this *JOGL10) GlLightf(light int, pname int, param float32) {
        this.Gl10.GlLightf(light, pname, param)
}


func (this *JOGL10) GlLightfv_1(light int, pname int, params []float32, offset int) {
        this.Gl10.GlLightfv(light, pname, params, offset)
}

// Java overloads of JOGL10.glLightfv
func (this *JOGL10) GlLightfv(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 4 && Java_IsArg[int](args[0], pass) && Java_IsArg[int](args[1], pass) && Java_IsArg[[]float32](args[2], pass) && Java_IsArg[int](args[3], pass) {
            this.GlLightfv_1(Java_Arg[int](args[0]), Java_Arg[int](args[1]), Java_Arg[[]float32](args[2]), Java_Arg[int](args[3]))
            return
        }
        if len(args) == 3 && Java_IsArg[int](args[0], pass) && Java_IsArg[int](args[1], pass) && Java_IsArg[*FloatBuffer](args[2], pass) {
            this.GlLightfv_2(Java_Arg[int](args[0]), Java_Arg[int](args[1]), Java_Arg[*FloatBuffer](args[2]))
            return
        }
    }
    panic("No Java overload of JOGL10.glLightfv for the arguments")
}


func (this *JOGL10) GlLightfv_2(light int, pname int, params *FloatBuffer) {
        this.Gl10.GlLightfv(light, pname, params)
}


func (this *JOGL10) GlLineWidth(width float32) {
        this.Gl10.GlLineWidth(width)
}


func (this *JOGL10) GlLoadIdentity() {
        this.Gl10.GlLoadIdentity()
}


func (this *JOGL10) GlLoadMatrixf_1(m []float32, offset int) {
        this.Gl10.GlLoadMatrixf(m, offset)
}

// Java overloads of JOGL10.glLoadMatrixf
func (this *JOGL10) GlLoadMatrixf(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 2 && Java_IsArg[[]float32](args[0], pass) && Java_IsArg[int](args[1], pass) {
            this.GlLoadMatrixf_1(Java_Arg[[]float32](args[0]), Java_Arg[int](args[1]))
            return
        }
        if len(args) == 1 && Java_IsArg[*FloatBuffer](args[0], pass) {
            this.GlLoadMatrixf_2(Java_Arg[*FloatBuffer](args[0]))
            return
        }
    }
    panic("No Java overload of JOGL10.glLoadMatrixf for the arguments")
}


func (this *JOGL10) GlLoadMatrixf_2(m *FloatBuffer) {
        this.Gl10.GlLoadMatrixf(m)
}


func (this *JOGL10) GlLogicOp(opcode int) {
        this.Gl10.GlLogicOp(opcode)
}


func (this *JOGL10) GlMaterialf(face int, pname int, param float32) {
        this.Gl10.GlMaterialf(face, pname, param)
}


func (this *JOGL10) GlMaterialfv_1(face int, pname int, params []float32, offset int) {
        this.Gl10.GlMaterialfv(face, pname, params, offset)
}

// Java overloads of JOGL10.glMaterialfv
func (this *JOGL10) GlMaterialfv(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 4 && Java_IsArg[int](args[0], pass) && Java_IsArg[int](args[1], pass) && Java_IsArg[[]float32](args[2], pass) && Java_IsArg[int](args[3], pass) {
            this.GlMaterialfv_1(Java_Arg[int](args[0]), Java_Arg[int](args[1]), Java_Arg[[]float32](args[2]), Java_Arg[int](args[3]))
            return
        }
        if len(args) == 3 && Java_IsArg[int](args[0], pass) && Java_IsArg[int](args[1], pass) && Java_IsArg[*FloatBuffer](args[2], pass) {
            this.GlMaterialfv_2(Java_Arg[int](args[0]), Java_Arg[int](args[1]), Java_Arg[*FloatBuffer](args[2]))
            return
        }
    }
    panic("No Java overload of JOGL10.glMaterialfv for the arguments")
}


func (this *JOGL10) GlMaterialfv_2(face int, pname int, params *FloatBuffer) {
        this.Gl10.GlMaterialfv(face, pname, params)
}


func (this *JOGL10) GlMatrixMode(mode int) {
        this.Gl10.GlMatrixMode(mode)
}


func (this *JOGL10) GlMultMatrixf_1(m []float32, offset int) {
        this.Gl10.GlMultMatrixf(m, offset)
}

// Java overloads of JOGL10.glMultMatrixf
func (this *JOGL10) GlMultMatrixf(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 2 && Java_IsArg[[]float32](args[0], pass) && Java_IsArg[int](args[1], pass) {
            this.GlMultMatrixf_1(Java_Arg[[]float32](args[0]), Java_Arg[int](args[1]))
            return
        }
        if len(args) == 1 && Java_IsArg[*FloatBuffer](args[0], pass) {
            this.GlMultMatrixf_2(Java_Arg[*FloatBuffer](args[0]))
            return
        }
    }
    panic("No Java overload of JOGL10.glMultMatrixf for the arguments")
}


func (this *JOGL10) GlMultMatrixf_2(m *FloatBuffer) {
        this.Gl10.GlMultMatrixf(m)
}


func (this *JOGL10) GlMultiTexCoord4f(target int, s float32, t float32, r float32, q float32) {
        this.Gl10.GlMultiTexCoord4f(target, s, t, r, q)
}


func (this *JOGL10) GlNormal3f(nx float32, ny float32, nz float32) {
        this.Gl10.GlNormal3f(nx, ny, nz)
}


func (this *JOGL10) GlNormalPointer(type_ int, stride int, pointer *Buffer) {
        this.Gl10.GlNormalPointer(type_, stride, pointer)
}


func (this *JOGL10) GlOrthof(left float32, right float32, bottom float32, top float32, zNear float32, zFar float32) {
        this.Gl10.GlOrthof(left, right, bottom, top, zNear, zFar)
}


func (this *JOGL10) GlPixelStorei(pname int, param int) {
        this.Gl10.GlPixelStorei(pname, param)
}


func (this *JOGL10) GlPointSize(size float32) {
        this.Gl10.GlPointSize(size)
}


func (this *JOGL10) GlPolygonOffset(factor float32, units float32) {
        this.Gl10.GlPolygonOffset(factor, units)
}


func (this *JOGL10) GlPopMatrix() {
        this.Gl10.GlPopMatrix()
}


func (this *JOGL10) GlPushMatrix() {
        this.Gl10.GlPushMatrix()
}


func (this *JOGL10) GlReadPixels(x int, y int, width int, height int, format int, type_ int, pixels *Buffer) {
        this.Gl10.GlReadPixels(x, y, width, height, format, type_, pixels)
}


func (this *JOGL10) GlRotatef(angle float32, x float32, y float32, z float32) {
        this.Gl10.GlRotatef(angle, x, y, z)
}


func (this *JOGL10) GlRotatex(angle int, x int, y int, z int) {
        this.Gl10.GlRotatef(angle, x, y, z)
}


func (this *JOGL10) GlSampleCoverage(value float32, invert bool) {
        this.Gl10.GlSampleCoverage(value, invert)
}


func (this *JOGL10) GlScalef(x float32, y float32, z float32) {
        this.Gl10.GlScalef(x, y, z)
}


func (this *JOGL10) GlScalex(x int, y int, z int) {
        this.Gl10.GlScalef(x, y, z)
}


func (this *JOGL10) GlScissor(x int, y int, width int, height int) {
        this.Gl10.GlScissor(x, y, width, height)
}


func (this *JOGL10) GlShadeModel(mode int) {
        this.Gl10.GlShadeModel(mode)
}


func (this *JOGL10) GlStencilFunc(func_ int, ref int, mask int) {
        this.Gl10.GlStencilFunc(func_, ref, mask)
}


func (this *JOGL10) GlStencilMask(mask int) {
        this.Gl10.GlStencilMask(mask)
}


func (this *JOGL10) GlStencilOp(fail int, zfail int, zpass int) {
        this.Gl10.GlStencilOp(fail, zfail, zpass)
}


func (this *JOGL10) GlTexCoordPointer(size int, type_ int, stride int, pointer *Buffer) {
        this.Gl10.GlTexCoordPointer(size, type_, stride, pointer)
}


func (this *JOGL10) GlTexEnvf(target int, pname int, param float32) {
        this.Gl10.GlTexEnvf(target, pname, param)
}


func (this *JOGL10) GlTexEnvfv_1(target int, pname int, params []float32, offset int) {
        this.Gl10.GlTexEnvfv(target, pname, params, offset)
}

// Java overloads of JOGL10.glTexEnvfv
func (this *JOGL10) GlTexEnvfv(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 4 && Java_IsArg[int](args[0], pass) && Java_IsArg[int](args[1], pass) && Java_IsArg[[]float32](args[2], pass) && Java_IsArg[int](args[3], pass) {
            this.GlTexEnvfv_1(Java_Arg[int](args[0]), Java_Arg[int](args[1]), Java_Arg[[]float32](args[2]), Java_Arg[int](args[3]))
            return
        }
        if len(args) == 3 && Java_IsArg[int](args[0], pass) && Java_IsArg[int](args[1], pass) && Java_IsArg[*FloatBuffer](args[2], pass) {
            this.GlTexEnvfv_2(Java_Arg[int](args[0]), Java_Arg[int](args[1]), Java_Arg[*FloatBuffer](args[2]))
            return
        }
    }
    panic("No Java overload of JOGL10.glTexEnvfv for the arguments")
}


func (this *JOGL10) GlTexEnvfv_2(target int, pname int, params *FloatBuffer) {
        this.Gl10.GlTexEnvfv(target, pname, params)
}


func (this *JOGL10) GlTexImage2D(target int, level int, internalformat int, width int, height int, border int, format int, type_ int, pixels *Buffer) {
        this.Gl10.GlTexImage2D(target, level, internalformat, width, height, border, format, type_, pixels)
}


func (this *JOGL10) GlTexParameterf(target int, pname int, param float32) {
        this.Gl10.GlTexParameterf(target, pname, param)
}


func (this *JOGL10) GlTexParameteri(target int, pname int, param int) {
        this.Gl10.GlTexParameteri(target, pname, param)
}


func (this *JOGL10) GlTexSubImage2D(target int, level int, xoffset int, yoffset int, width int, height int, format int, type_ int, pixels *Buffer) {
        this.Gl10.GlTexSubImage2D(target, level, xoffset, yoffset, width, height, format, type_, pixels)
}


func (this *JOGL10) GlTranslatef(x float32, y float32, z float32) {
        this.Gl10.GlTranslatef(x, y, z)
}


func (this *JOGL10) GlVertexPointer(size int, type_ int, stride int, pointer *Buffer) {
        this.Gl10.GlVertexPointer(size, type_, stride, pointer)
}


func (this *JOGL10) GlViewport(x int, y int, width int, height int) {
        this.Gl10.GlViewport(x, y, width, height)
}


func (this *JOGL10) GlBegin(i int) {
        this.Gl10.GlBegin(i)
}


func (this *JOGL10) GlEnd() {
        this.Gl10.GlEnd()
}


func (this *JOGL10) GlVertex3f(f float32, f2 float32, f3 float32) {
        this.Gl10.GlVertex3f(f, f2, f3)
}


func (this *JOGL10) GlPolygonMode(i int, i1 int) {
        this.Gl10.GlPolygonMode(i, i1)
}

