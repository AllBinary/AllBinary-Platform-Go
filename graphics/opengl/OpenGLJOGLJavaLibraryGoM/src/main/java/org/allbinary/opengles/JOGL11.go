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

//JOGLEToKhronosGL11
type JOGL11 struct {
    JOGL10
    Gl11 *GL2
}

// Implements GL11


func (this *JOGL11) initFields() {
    
}




// Java overloads of new JOGL11
func NewJOGL11(args ...any) *JOGL11 {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[GL](args[0], pass) {
            return NewJOGL11_1(Java_Arg[GL](args[0]))
        }
        if len(args) == 2 && Java_IsArg[GL](args[0], pass) && Java_IsArg[*GL2](args[1], pass) {
            return NewJOGL11_2(Java_Arg[GL](args[0]), Java_Arg[*GL2](args[1]))
        }
    }
    panic("No Java overload of new JOGL11 for the arguments")
}
func NewJOGL11_1(gl GL) *JOGL11 {
    this := &JOGL11{}
        *this = *NewJOGL11(gl, any(gl).(*GL2))
    return this
}


func NewJOGL11_2(gl GL, gl2 *GL2) *JOGL11 {
    this := &JOGL11{}
        this.JOGL10 = *NewJOGL10(gl, gl2)
        this.initFields()
        this.Gl11 = this.GetJOGLGL()
    return this
}


func (this *JOGL11) GlBindBuffer(target int, buffer int) {
        this.Gl11.GlBindBuffer(target, buffer)
}


func (this *JOGL11) GlBufferData(target int, size int, data *Buffer, usage int) {
        this.Gl11.GlBufferData(target, size, data, usage)
}


func (this *JOGL11) GlBufferSubData(target int, offset int, size int, data *Buffer) {
        this.Gl11.GlBufferSubData(target, offset, size, data)
}


func (this *JOGL11) GlClipPlanef_1(plane int, equation []float32, offset int) {
        this.Gl11.GlClipPlanef(plane, equation, offset)
}

// Java overloads of JOGL11.glClipPlanef
func (this *JOGL11) GlClipPlanef(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 3 && Java_IsArg[int](args[0], pass) && Java_IsArg[[]float32](args[1], pass) && Java_IsArg[int](args[2], pass) {
            this.GlClipPlanef_1(Java_Arg[int](args[0]), Java_Arg[[]float32](args[1]), Java_Arg[int](args[2]))
            return
        }
        if len(args) == 2 && Java_IsArg[int](args[0], pass) && Java_IsArg[*FloatBuffer](args[1], pass) {
            this.GlClipPlanef_2(Java_Arg[int](args[0]), Java_Arg[*FloatBuffer](args[1]))
            return
        }
    }
    panic("No Java overload of JOGL11.glClipPlanef for the arguments")
}


func (this *JOGL11) GlClipPlanef_2(plane int, equation *FloatBuffer) {
        this.Gl11.GlClipPlanef(plane, equation)
}


func (this *JOGL11) GlColor4ub(red int8, green int8, blue int8, alpha int8) {
        this.Gl11.GlColor4ub(red, green, blue, alpha)
}


func (this *JOGL11) GlColorPointer(size int, type_ int, stride int, offset int) {
        this.Gl11.GlColorPointer(size, type_, stride, offset)
}


func (this *JOGL11) GlDeleteBuffers_1(n int, buffers []int, offset int) {
        this.Gl11.GlDeleteBuffers(n, buffers, offset)
}

// Java overloads of JOGL11.glDeleteBuffers
func (this *JOGL11) GlDeleteBuffers(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 3 && Java_IsArg[int](args[0], pass) && Java_IsArg[[]int](args[1], pass) && Java_IsArg[int](args[2], pass) {
            this.GlDeleteBuffers_1(Java_Arg[int](args[0]), Java_Arg[[]int](args[1]), Java_Arg[int](args[2]))
            return
        }
        if len(args) == 2 && Java_IsArg[int](args[0], pass) && Java_IsArg[*IntBuffer](args[1], pass) {
            this.GlDeleteBuffers_2(Java_Arg[int](args[0]), Java_Arg[*IntBuffer](args[1]))
            return
        }
    }
    panic("No Java overload of JOGL11.glDeleteBuffers for the arguments")
}


func (this *JOGL11) GlDeleteBuffers_2(n int, buffers *IntBuffer) {
        this.Gl11.GlDeleteBuffers(n, buffers)
}


func (this *JOGL11) GlDrawElements(mode int, count int, type_ int, offset int) {
        this.Gl11.GlDrawElements(mode, count, type_, offset)
}


func (this *JOGL11) GlGenBuffers_1(n int, buffers []int, offset int) {
        this.Gl11.GlGenBuffers(n, buffers, offset)
}

// Java overloads of JOGL11.glGenBuffers
func (this *JOGL11) GlGenBuffers(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 3 && Java_IsArg[int](args[0], pass) && Java_IsArg[[]int](args[1], pass) && Java_IsArg[int](args[2], pass) {
            this.GlGenBuffers_1(Java_Arg[int](args[0]), Java_Arg[[]int](args[1]), Java_Arg[int](args[2]))
            return
        }
        if len(args) == 2 && Java_IsArg[int](args[0], pass) && Java_IsArg[*IntBuffer](args[1], pass) {
            this.GlGenBuffers_2(Java_Arg[int](args[0]), Java_Arg[*IntBuffer](args[1]))
            return
        }
    }
    panic("No Java overload of JOGL11.glGenBuffers for the arguments")
}


func (this *JOGL11) GlGenBuffers_2(n int, buffers *IntBuffer) {
        this.Gl11.GlGenBuffers(n, buffers)
}


func (this *JOGL11) GlGetBufferParameteriv_1(target int, pname int, params []int, offset int) {
        this.Gl11.GlGetBufferParameteriv(target, pname, params, offset)
}

// Java overloads of JOGL11.glGetBufferParameteriv
func (this *JOGL11) GlGetBufferParameteriv(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 4 && Java_IsArg[int](args[0], pass) && Java_IsArg[int](args[1], pass) && Java_IsArg[[]int](args[2], pass) && Java_IsArg[int](args[3], pass) {
            this.GlGetBufferParameteriv_1(Java_Arg[int](args[0]), Java_Arg[int](args[1]), Java_Arg[[]int](args[2]), Java_Arg[int](args[3]))
            return
        }
        if len(args) == 3 && Java_IsArg[int](args[0], pass) && Java_IsArg[int](args[1], pass) && Java_IsArg[*IntBuffer](args[2], pass) {
            this.GlGetBufferParameteriv_2(Java_Arg[int](args[0]), Java_Arg[int](args[1]), Java_Arg[*IntBuffer](args[2]))
            return
        }
    }
    panic("No Java overload of JOGL11.glGetBufferParameteriv for the arguments")
}


func (this *JOGL11) GlGetBufferParameteriv_2(target int, pname int, params *IntBuffer) {
        this.Gl11.GlGetBufferParameteriv(target, pname, params)
}


func (this *JOGL11) GlGetClipPlanef_1(pname int, eqn []float32, offset int) {
        this.Gl11.GlGetClipPlanef(pname, eqn, offset)
}

// Java overloads of JOGL11.glGetClipPlanef
func (this *JOGL11) GlGetClipPlanef(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 3 && Java_IsArg[int](args[0], pass) && Java_IsArg[[]float32](args[1], pass) && Java_IsArg[int](args[2], pass) {
            this.GlGetClipPlanef_1(Java_Arg[int](args[0]), Java_Arg[[]float32](args[1]), Java_Arg[int](args[2]))
            return
        }
        if len(args) == 2 && Java_IsArg[int](args[0], pass) && Java_IsArg[*FloatBuffer](args[1], pass) {
            this.GlGetClipPlanef_2(Java_Arg[int](args[0]), Java_Arg[*FloatBuffer](args[1]))
            return
        }
    }
    panic("No Java overload of JOGL11.glGetClipPlanef for the arguments")
}


func (this *JOGL11) GlGetClipPlanef_2(pname int, eqn *FloatBuffer) {
        this.Gl11.GlGetClipPlanef(pname, eqn)
}


func (this *JOGL11) GlGetFloatv_1(pname int, params []float32, offset int) {
        this.Gl11.GlGetFloatv(pname, params, offset)
}

// Java overloads of JOGL11.glGetFloatv
func (this *JOGL11) GlGetFloatv(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 3 && Java_IsArg[int](args[0], pass) && Java_IsArg[[]float32](args[1], pass) && Java_IsArg[int](args[2], pass) {
            this.GlGetFloatv_1(Java_Arg[int](args[0]), Java_Arg[[]float32](args[1]), Java_Arg[int](args[2]))
            return
        }
        if len(args) == 2 && Java_IsArg[int](args[0], pass) && Java_IsArg[*FloatBuffer](args[1], pass) {
            this.GlGetFloatv_2(Java_Arg[int](args[0]), Java_Arg[*FloatBuffer](args[1]))
            return
        }
    }
    panic("No Java overload of JOGL11.glGetFloatv for the arguments")
}


func (this *JOGL11) GlGetFloatv_2(pname int, params *FloatBuffer) {
        this.Gl11.GlGetFloatv(pname, params)
}


func (this *JOGL11) GlGetLightfv_1(light int, pname int, params []float32, offset int) {
        this.Gl11.GlGetLightfv(light, pname, params, offset)
}

// Java overloads of JOGL11.glGetLightfv
func (this *JOGL11) GlGetLightfv(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 4 && Java_IsArg[int](args[0], pass) && Java_IsArg[int](args[1], pass) && Java_IsArg[[]float32](args[2], pass) && Java_IsArg[int](args[3], pass) {
            this.GlGetLightfv_1(Java_Arg[int](args[0]), Java_Arg[int](args[1]), Java_Arg[[]float32](args[2]), Java_Arg[int](args[3]))
            return
        }
        if len(args) == 3 && Java_IsArg[int](args[0], pass) && Java_IsArg[int](args[1], pass) && Java_IsArg[*FloatBuffer](args[2], pass) {
            this.GlGetLightfv_2(Java_Arg[int](args[0]), Java_Arg[int](args[1]), Java_Arg[*FloatBuffer](args[2]))
            return
        }
    }
    panic("No Java overload of JOGL11.glGetLightfv for the arguments")
}


func (this *JOGL11) GlGetLightfv_2(light int, pname int, params *FloatBuffer) {
        this.Gl11.GlGetLightfv(light, pname, params)
}


func (this *JOGL11) GlGetMaterialfv_1(face int, pname int, params []float32, offset int) {
        this.Gl11.GlGetMaterialfv(face, pname, params, offset)
}

// Java overloads of JOGL11.glGetMaterialfv
func (this *JOGL11) GlGetMaterialfv(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 4 && Java_IsArg[int](args[0], pass) && Java_IsArg[int](args[1], pass) && Java_IsArg[[]float32](args[2], pass) && Java_IsArg[int](args[3], pass) {
            this.GlGetMaterialfv_1(Java_Arg[int](args[0]), Java_Arg[int](args[1]), Java_Arg[[]float32](args[2]), Java_Arg[int](args[3]))
            return
        }
        if len(args) == 3 && Java_IsArg[int](args[0], pass) && Java_IsArg[int](args[1], pass) && Java_IsArg[*FloatBuffer](args[2], pass) {
            this.GlGetMaterialfv_2(Java_Arg[int](args[0]), Java_Arg[int](args[1]), Java_Arg[*FloatBuffer](args[2]))
            return
        }
    }
    panic("No Java overload of JOGL11.glGetMaterialfv for the arguments")
}


func (this *JOGL11) GlGetMaterialfv_2(face int, pname int, params *FloatBuffer) {
        this.Gl11.GlGetMaterialfv(face, pname, params)
}


func (this *JOGL11) GlGetTexEnviv_1(env int, pname int, params []int, offset int) {
        this.Gl11.GlGetTexEnviv(env, pname, params, offset)
}

// Java overloads of JOGL11.glGetTexEnviv
func (this *JOGL11) GlGetTexEnviv(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 4 && Java_IsArg[int](args[0], pass) && Java_IsArg[int](args[1], pass) && Java_IsArg[[]int](args[2], pass) && Java_IsArg[int](args[3], pass) {
            this.GlGetTexEnviv_1(Java_Arg[int](args[0]), Java_Arg[int](args[1]), Java_Arg[[]int](args[2]), Java_Arg[int](args[3]))
            return
        }
        if len(args) == 3 && Java_IsArg[int](args[0], pass) && Java_IsArg[int](args[1], pass) && Java_IsArg[*IntBuffer](args[2], pass) {
            this.GlGetTexEnviv_2(Java_Arg[int](args[0]), Java_Arg[int](args[1]), Java_Arg[*IntBuffer](args[2]))
            return
        }
    }
    panic("No Java overload of JOGL11.glGetTexEnviv for the arguments")
}


func (this *JOGL11) GlGetTexEnviv_2(env int, pname int, params *IntBuffer) {
        this.Gl11.GlGetTexEnviv(env, pname, params)
}


func (this *JOGL11) GlGetTexParameterfv_1(target int, pname int, params []float32, offset int) {
        this.Gl11.GlGetTexParameterfv(target, pname, params, offset)
}

// Java overloads of JOGL11.glGetTexParameterfv
func (this *JOGL11) GlGetTexParameterfv(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 4 && Java_IsArg[int](args[0], pass) && Java_IsArg[int](args[1], pass) && Java_IsArg[[]float32](args[2], pass) && Java_IsArg[int](args[3], pass) {
            this.GlGetTexParameterfv_1(Java_Arg[int](args[0]), Java_Arg[int](args[1]), Java_Arg[[]float32](args[2]), Java_Arg[int](args[3]))
            return
        }
        if len(args) == 3 && Java_IsArg[int](args[0], pass) && Java_IsArg[int](args[1], pass) && Java_IsArg[*FloatBuffer](args[2], pass) {
            this.GlGetTexParameterfv_2(Java_Arg[int](args[0]), Java_Arg[int](args[1]), Java_Arg[*FloatBuffer](args[2]))
            return
        }
    }
    panic("No Java overload of JOGL11.glGetTexParameterfv for the arguments")
}


func (this *JOGL11) GlGetTexParameterfv_2(target int, pname int, params *FloatBuffer) {
        this.Gl11.GlGetTexParameterfv(target, pname, params)
}


func (this *JOGL11) GlGetTexParameteriv_1(target int, pname int, params []int, offset int) {
        this.Gl11.GlGetTexParameteriv(target, pname, params, offset)
}

// Java overloads of JOGL11.glGetTexParameteriv
func (this *JOGL11) GlGetTexParameteriv(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 4 && Java_IsArg[int](args[0], pass) && Java_IsArg[int](args[1], pass) && Java_IsArg[[]int](args[2], pass) && Java_IsArg[int](args[3], pass) {
            this.GlGetTexParameteriv_1(Java_Arg[int](args[0]), Java_Arg[int](args[1]), Java_Arg[[]int](args[2]), Java_Arg[int](args[3]))
            return
        }
        if len(args) == 3 && Java_IsArg[int](args[0], pass) && Java_IsArg[int](args[1], pass) && Java_IsArg[*IntBuffer](args[2], pass) {
            this.GlGetTexParameteriv_2(Java_Arg[int](args[0]), Java_Arg[int](args[1]), Java_Arg[*IntBuffer](args[2]))
            return
        }
    }
    panic("No Java overload of JOGL11.glGetTexParameteriv for the arguments")
}


func (this *JOGL11) GlGetTexParameteriv_2(target int, pname int, params *IntBuffer) {
        this.Gl11.GlGetTexParameteriv(target, pname, params)
}


func (this *JOGL11) GlIsBuffer(buffer int) bool {
        return this.Gl11.GlIsBuffer(buffer)
}


func (this *JOGL11) GlIsEnabled(cap_ int) bool {
        return this.Gl11.GlIsEnabled(cap_)
}


func (this *JOGL11) GlIsTexture(texture int) bool {
        return this.Gl11.GlIsTexture(texture)
}


func (this *JOGL11) GlNormalPointer(type_ int, stride int, offset int) {
        this.Gl11.GlNormalPointer(type_, stride, offset)
}


func (this *JOGL11) GlPointParameterf(pname int, param float32) {
        this.Gl11.GlPointParameterf(pname, param)
}


func (this *JOGL11) GlPointParameterfv_1(pname int, params []float32, offset int) {
        this.Gl11.GlPointParameterfv(pname, params, offset)
}

// Java overloads of JOGL11.glPointParameterfv
func (this *JOGL11) GlPointParameterfv(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 3 && Java_IsArg[int](args[0], pass) && Java_IsArg[[]float32](args[1], pass) && Java_IsArg[int](args[2], pass) {
            this.GlPointParameterfv_1(Java_Arg[int](args[0]), Java_Arg[[]float32](args[1]), Java_Arg[int](args[2]))
            return
        }
        if len(args) == 2 && Java_IsArg[int](args[0], pass) && Java_IsArg[*FloatBuffer](args[1], pass) {
            this.GlPointParameterfv_2(Java_Arg[int](args[0]), Java_Arg[*FloatBuffer](args[1]))
            return
        }
    }
    panic("No Java overload of JOGL11.glPointParameterfv for the arguments")
}


func (this *JOGL11) GlPointParameterfv_2(pname int, params *FloatBuffer) {
        this.Gl11.GlPointParameterfv(pname, params)
}


func (this *JOGL11) GlTexCoordPointer(size int, type_ int, stride int, offset int) {
        this.Gl11.GlTexCoordPointer(size, type_, stride, offset)
}


func (this *JOGL11) GlTexEnvi(target int, pname int, param int) {
        this.Gl11.GlTexEnvi(target, pname, param)
}


func (this *JOGL11) GlTexEnviv_1(target int, pname int, params []int, offset int) {
        this.Gl11.GlTexEnviv(target, pname, params, offset)
}

// Java overloads of JOGL11.glTexEnviv
func (this *JOGL11) GlTexEnviv(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 4 && Java_IsArg[int](args[0], pass) && Java_IsArg[int](args[1], pass) && Java_IsArg[[]int](args[2], pass) && Java_IsArg[int](args[3], pass) {
            this.GlTexEnviv_1(Java_Arg[int](args[0]), Java_Arg[int](args[1]), Java_Arg[[]int](args[2]), Java_Arg[int](args[3]))
            return
        }
        if len(args) == 3 && Java_IsArg[int](args[0], pass) && Java_IsArg[int](args[1], pass) && Java_IsArg[*IntBuffer](args[2], pass) {
            this.GlTexEnviv_2(Java_Arg[int](args[0]), Java_Arg[int](args[1]), Java_Arg[*IntBuffer](args[2]))
            return
        }
    }
    panic("No Java overload of JOGL11.glTexEnviv for the arguments")
}


func (this *JOGL11) GlTexEnviv_2(target int, pname int, params *IntBuffer) {
        this.Gl11.GlTexEnviv(target, pname, params)
}


func (this *JOGL11) GlTexParameterfv_1(target int, pname int, params []float32, offset int) {
        this.Gl11.GlTexParameterfv(target, pname, params, offset)
}

// Java overloads of JOGL11.glTexParameterfv
func (this *JOGL11) GlTexParameterfv(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 4 && Java_IsArg[int](args[0], pass) && Java_IsArg[int](args[1], pass) && Java_IsArg[[]float32](args[2], pass) && Java_IsArg[int](args[3], pass) {
            this.GlTexParameterfv_1(Java_Arg[int](args[0]), Java_Arg[int](args[1]), Java_Arg[[]float32](args[2]), Java_Arg[int](args[3]))
            return
        }
        if len(args) == 3 && Java_IsArg[int](args[0], pass) && Java_IsArg[int](args[1], pass) && Java_IsArg[*FloatBuffer](args[2], pass) {
            this.GlTexParameterfv_2(Java_Arg[int](args[0]), Java_Arg[int](args[1]), Java_Arg[*FloatBuffer](args[2]))
            return
        }
    }
    panic("No Java overload of JOGL11.glTexParameterfv for the arguments")
}


func (this *JOGL11) GlTexParameterfv_2(target int, pname int, params *FloatBuffer) {
        this.Gl11.GlTexParameterfv(target, pname, params)
}


func (this *JOGL11) GlTexParameteri(target int, pname int, param int) {
        this.Gl11.GlTexParameteri(target, pname, param)
}


func (this *JOGL11) GlTexParameteriv_1(target int, pname int, params []int, offset int) {
        this.Gl11.GlTexParameteriv(target, pname, params, offset)
}

// Java overloads of JOGL11.glTexParameteriv
func (this *JOGL11) GlTexParameteriv(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 4 && Java_IsArg[int](args[0], pass) && Java_IsArg[int](args[1], pass) && Java_IsArg[[]int](args[2], pass) && Java_IsArg[int](args[3], pass) {
            this.GlTexParameteriv_1(Java_Arg[int](args[0]), Java_Arg[int](args[1]), Java_Arg[[]int](args[2]), Java_Arg[int](args[3]))
            return
        }
        if len(args) == 3 && Java_IsArg[int](args[0], pass) && Java_IsArg[int](args[1], pass) && Java_IsArg[*IntBuffer](args[2], pass) {
            this.GlTexParameteriv_2(Java_Arg[int](args[0]), Java_Arg[int](args[1]), Java_Arg[*IntBuffer](args[2]))
            return
        }
    }
    panic("No Java overload of JOGL11.glTexParameteriv for the arguments")
}


func (this *JOGL11) GlTexParameteriv_2(target int, pname int, params *IntBuffer) {
        this.Gl11.GlTexParameteriv(target, pname, params)
}


func (this *JOGL11) GlTexParameterxv(target int, pname int, params []int, offset int) {
        this.Gl11.GlTexParameteriv(target, pname, params, offset)
}


func (this *JOGL11) GlVertexPointer(size int, type_ int, stride int, offset int) {
        this.Gl11.GlVertexPointer(size, type_, stride, offset)
}

