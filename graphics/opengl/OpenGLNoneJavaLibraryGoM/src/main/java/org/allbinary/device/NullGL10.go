/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2025 AllBinary 
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

type NullGL10 struct {
}

// Implements GL10

var NullGL10_NULL_GL10 *NullGL10 = NewNullGL10()

func NewNullGL10() *NullGL10 {
    this := &NullGL10{}
    return this
}


func (this *NullGL10) GlActiveTexture(texture int) {
        return
}


func (this *NullGL10) GlAlphaFunc(func_ int, ref float32) {
        return
}


func (this *NullGL10) GlBindTexture(target int, texture int) {
        return
}


func (this *NullGL10) GlBlendFunc(sfactor int, dfactor int) {
        return
}


func (this *NullGL10) GlClear(mask int) {
        return
}


func (this *NullGL10) GlClearColor(red float32, green float32, blue float32, alpha float32) {
        return
}


func (this *NullGL10) GlClearDepthf(depth float32) {
        return
}


func (this *NullGL10) GlClearStencil(s int) {
        return
}


func (this *NullGL10) GlClientActiveTexture(texture int) {
        return
}


func (this *NullGL10) GlColor4f(red float32, green float32, blue float32, alpha float32) {
        return
}


func (this *NullGL10) GlColorMask(red bool, green bool, blue bool, alpha bool) {
        return
}


func (this *NullGL10) GlColorPointer(size int, type_ int, stride int, pointer *Buffer) {
        return
}


func (this *NullGL10) GlCompressedTexImage2D(target int, level int, internalformat int, width int, height int, border int, imageSize int, data *Buffer) {
        return
}


func (this *NullGL10) GlCompressedTexSubImage2D(target int, level int, xoffset int, yoffset int, width int, height int, format int, imageSize int, data *Buffer) {
        return
}


func (this *NullGL10) GlCopyTexImage2D(target int, level int, internalformat int, x int, y int, width int, height int, border int) {
        return
}


func (this *NullGL10) GlCopyTexSubImage2D(target int, level int, xoffset int, yoffset int, x int, y int, width int, height int) {
        return
}


func (this *NullGL10) GlCullFace(mode int) {
        return
}


func (this *NullGL10) GlDeleteTextures_1(n int, textures []int, offset int) {
        return
}

// Java overloads of NullGL10.glDeleteTextures
func (this *NullGL10) GlDeleteTextures(args ...any) {
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
    panic("No Java overload of NullGL10.glDeleteTextures for the arguments")
}


func (this *NullGL10) GlDeleteTextures_2(n int, textures *IntBuffer) {
        return
}


func (this *NullGL10) GlDepthFunc(func_ int) {
        return
}


func (this *NullGL10) GlDepthMask(flag bool) {
        return
}


func (this *NullGL10) GlDepthRangef(zNear float32, zFar float32) {
        return
}


func (this *NullGL10) GlDisable(cap_ int) {
        return
}


func (this *NullGL10) GlDisableClientState(array int) {
        return
}


func (this *NullGL10) GlDrawArrays(mode int, first int, count int) {
        return
}


func (this *NullGL10) GlDrawElements(mode int, count int, type_ int, indices *Buffer) {
        return
}


func (this *NullGL10) GlEnable(cap_ int) {
        return
}


func (this *NullGL10) GlEnableClientState(array int) {
        return
}


func (this *NullGL10) GlFinish() {
        return
}


func (this *NullGL10) GlFlush() {
        return
}


func (this *NullGL10) GlFogf(pname int, param float32) {
        return
}


func (this *NullGL10) GlFogfv_1(pname int, params []float32, offset int) {
        return
}

// Java overloads of NullGL10.glFogfv
func (this *NullGL10) GlFogfv(args ...any) {
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
    panic("No Java overload of NullGL10.glFogfv for the arguments")
}


func (this *NullGL10) GlFogfv_2(pname int, params *FloatBuffer) {
        return
}


func (this *NullGL10) GlFrontFace(mode int) {
        return
}


func (this *NullGL10) GlFrustumf(left float32, right float32, bottom float32, top float32, zNear float32, zFar float32) {
        return
}


func (this *NullGL10) GlGenTextures_1(n int, textures []int, offset int) {
        return
}

// Java overloads of NullGL10.glGenTextures
func (this *NullGL10) GlGenTextures(args ...any) {
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
    panic("No Java overload of NullGL10.glGenTextures for the arguments")
}


func (this *NullGL10) GlGenTextures_2(n int, textures *IntBuffer) {
        return
}


func (this *NullGL10) GlGetError() int {
        return 0
}


func (this *NullGL10) GlGetIntegerv_1(pname int, params []int, offset int) {
        return
}

// Java overloads of NullGL10.glGetIntegerv
func (this *NullGL10) GlGetIntegerv(args ...any) {
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
    panic("No Java overload of NullGL10.glGetIntegerv for the arguments")
}


func (this *NullGL10) GlGetIntegerv_2(pname int, params *IntBuffer) {
        return
}


func (this *NullGL10) GlGetString(name int) string {
        return StringUtil_GetInstance().EMPTY_STRING
}


func (this *NullGL10) GlHint(target int, mode int) {
        return
}


func (this *NullGL10) GlLightModelf(pname int, param float32) {
        return
}


func (this *NullGL10) GlLightModelfv_1(pname int, params []float32, offset int) {
        return
}

// Java overloads of NullGL10.glLightModelfv
func (this *NullGL10) GlLightModelfv(args ...any) {
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
    panic("No Java overload of NullGL10.glLightModelfv for the arguments")
}


func (this *NullGL10) GlLightModelfv_2(pname int, params *FloatBuffer) {
        return
}


func (this *NullGL10) GlLightf(light int, pname int, param float32) {
        return
}


func (this *NullGL10) GlLightfv_1(light int, pname int, params []float32, offset int) {
        return
}

// Java overloads of NullGL10.glLightfv
func (this *NullGL10) GlLightfv(args ...any) {
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
    panic("No Java overload of NullGL10.glLightfv for the arguments")
}


func (this *NullGL10) GlLightfv_2(light int, pname int, params *FloatBuffer) {
        return
}


func (this *NullGL10) GlLineWidth(width float32) {
        return
}


func (this *NullGL10) GlLoadIdentity() {
        return
}


func (this *NullGL10) GlLoadMatrixf_1(m []float32, offset int) {
        return
}

// Java overloads of NullGL10.glLoadMatrixf
func (this *NullGL10) GlLoadMatrixf(args ...any) {
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
    panic("No Java overload of NullGL10.glLoadMatrixf for the arguments")
}


func (this *NullGL10) GlLoadMatrixf_2(m *FloatBuffer) {
        return
}


func (this *NullGL10) GlLogicOp(opcode int) {
        return
}


func (this *NullGL10) GlMaterialf(face int, pname int, param float32) {
        return
}


func (this *NullGL10) GlMaterialfv_1(face int, pname int, params []float32, offset int) {
        return
}

// Java overloads of NullGL10.glMaterialfv
func (this *NullGL10) GlMaterialfv(args ...any) {
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
    panic("No Java overload of NullGL10.glMaterialfv for the arguments")
}


func (this *NullGL10) GlMaterialfv_2(face int, pname int, params *FloatBuffer) {
        return
}


func (this *NullGL10) GlMatrixMode(mode int) {
        return
}


func (this *NullGL10) GlMultMatrixf_1(m []float32, offset int) {
        return
}

// Java overloads of NullGL10.glMultMatrixf
func (this *NullGL10) GlMultMatrixf(args ...any) {
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
    panic("No Java overload of NullGL10.glMultMatrixf for the arguments")
}


func (this *NullGL10) GlMultMatrixf_2(m *FloatBuffer) {
        return
}


func (this *NullGL10) GlMultiTexCoord4f(target int, s float32, t float32, r float32, q float32) {
        return
}


func (this *NullGL10) GlNormal3f(nx float32, ny float32, nz float32) {
        return
}


func (this *NullGL10) GlNormalPointer(type_ int, stride int, pointer *Buffer) {
        return
}


func (this *NullGL10) GlOrthof(left float32, right float32, bottom float32, top float32, zNear float32, zFar float32) {
        return
}


func (this *NullGL10) GlPixelStorei(pname int, param int) {
        return
}


func (this *NullGL10) GlPointSize(size float32) {
        return
}


func (this *NullGL10) GlPolygonOffset(factor float32, units float32) {
        return
}


func (this *NullGL10) GlPopMatrix() {
        return
}


func (this *NullGL10) GlPushMatrix() {
        return
}


func (this *NullGL10) GlReadPixels(x int, y int, width int, height int, format int, type_ int, pixels *Buffer) {
        return
}


func (this *NullGL10) GlRotatef(angle float32, x float32, y float32, z float32) {
        return
}


func (this *NullGL10) GlRotatex(angle int, x int, y int, z int) {
        return
}


func (this *NullGL10) GlSampleCoverage(value float32, invert bool) {
        return
}


func (this *NullGL10) GlScalef(x float32, y float32, z float32) {
        return
}


func (this *NullGL10) GlScalex(x int, y int, z int) {
        return
}


func (this *NullGL10) GlScissor(x int, y int, width int, height int) {
        return
}


func (this *NullGL10) GlShadeModel(mode int) {
        return
}


func (this *NullGL10) GlStencilFunc(func_ int, ref int, mask int) {
        return
}


func (this *NullGL10) GlStencilMask(mask int) {
        return
}


func (this *NullGL10) GlStencilOp(fail int, zfail int, zpass int) {
        return
}


func (this *NullGL10) GlTexCoordPointer(size int, type_ int, stride int, pointer *Buffer) {
        return
}


func (this *NullGL10) GlTexEnvf(target int, pname int, param float32) {
        return
}


func (this *NullGL10) GlTexEnvfv_1(target int, pname int, params []float32, offset int) {
        return
}

// Java overloads of NullGL10.glTexEnvfv
func (this *NullGL10) GlTexEnvfv(args ...any) {
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
    panic("No Java overload of NullGL10.glTexEnvfv for the arguments")
}


func (this *NullGL10) GlTexEnvfv_2(target int, pname int, params *FloatBuffer) {
        return
}


func (this *NullGL10) GlTexImage2D(target int, level int, internalformat int, width int, height int, border int, format int, type_ int, pixels *Buffer) {
        return
}


func (this *NullGL10) GlTexParameterf(target int, pname int, param float32) {
        return
}


func (this *NullGL10) GlTexParameteri(target int, pname int, param int) {
        return
}


func (this *NullGL10) GlTexSubImage2D(target int, level int, xoffset int, yoffset int, width int, height int, format int, type_ int, pixels *Buffer) {
        return
}


func (this *NullGL10) GlTranslatef(x float32, y float32, z float32) {
        return
}


func (this *NullGL10) GlVertexPointer(size int, type_ int, stride int, pointer *Buffer) {
        return
}


func (this *NullGL10) GlViewport(x int, y int, width int, height int) {
        return
}


func (this *NullGL10) GlBegin(i int) {
        return
}


func (this *NullGL10) GlEnd() {
        return
}


func (this *NullGL10) GlVertex3f(f float32, f2 float32, f3 float32) {
        return
}


func (this *NullGL10) GlPolygonMode(i int, i1 int) {
        return
}

