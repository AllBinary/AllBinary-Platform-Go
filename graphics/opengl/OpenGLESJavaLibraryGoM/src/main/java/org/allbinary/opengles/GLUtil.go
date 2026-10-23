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

type GLUtil struct {
    BYTES_PER_FLOAT int
    NULL_FLOAT_BUFFER *FloatBuffer
}

var GLUtil_Instance *GLUtil = NewGLUtil()

func (this *GLUtil) initFields() {
    this.BYTES_PER_FLOAT = 4
    this.NULL_FLOAT_BUFFER = ByteBuffer_AllocateDirect(0).AsFloatBuffer()
}

func NewGLUtil() *GLUtil {
    this := &GLUtil{}
    this.initFields()
    return this
}


func GLUtil_GetInstance() *GLUtil {
        return GLUtil_Instance
}


func (this *GLUtil) MakeFloatBuffer3(p_a float32, p_b float32, p_c float32) *FloatBuffer {
        var b *ByteBuffer = ByteBuffer_AllocateDirect(3 * this.BYTES_PER_FLOAT); _ = b
        b.Order(ByteOrder_NativeOrder())
        var buffer *FloatBuffer = b.AsFloatBuffer(); _ = buffer
        buffer.Put(p_a)
        buffer.Put(p_b)
        buffer.Put(p_c)
        this.Position(buffer, 0)
        return buffer
}


func (this *GLUtil) MakeFloatBuffer4(p_a float32, p_b float32, p_c float32, p_d float32) *FloatBuffer {
        var b *ByteBuffer = ByteBuffer_AllocateDirect(4 * this.BYTES_PER_FLOAT); _ = b
        b.Order(ByteOrder_NativeOrder())
        var buffer *FloatBuffer = b.AsFloatBuffer(); _ = buffer
        buffer.Put(p_a)
        buffer.Put(p_b)
        buffer.Put(p_c)
        buffer.Put(p_d)
        this.Position(buffer, 0)
        return buffer
}


func (this *GLUtil) MakeFloatBuffer(byteBuffer *ByteBuffer) *FloatBuffer {
        var b *ByteBuffer = ByteBuffer_AllocateDirect(byteBuffer.Limit() * this.BYTES_PER_FLOAT); _ = b
        b.Order(ByteOrder_NativeOrder())
        var buffer *FloatBuffer = b.AsFloatBuffer(); _ = buffer
        var unsigned int; _ = unsigned
        var v float32; _ = v
        for byteBuffer.HasRemaining() {
            unsigned = (int(byteBuffer.Get())) & 0xFF
            v = (float32(unsigned)) / 255
            buffer.Put(v)
        }
        this.Position(buffer, 0)
        return buffer
}


func (this *GLUtil) Position(buffer *Buffer, newPosition int) *Buffer {
        return buffer.Position(newPosition)
}


func (this *GLUtil) RotateUVs(uvs []float32, angleDegrees float32, pivotU float32, pivotV float32) {
        var angleRad float64 = (float64(float64(angleDegrees)) * math.Pi / 180); _ = angleRad
        var cosA float64 = math.Cos(angleRad); _ = cosA
        var sinA float64 = math.Sin(angleRad); _ = sinA
        for i := int(0); i < len(uvs); i += 2 {
            var u float32 = uvs[i]; _ = u
            var v float32 = uvs[i + 1]; _ = v
            var du float32 = u - pivotU; _ = du
            var dv float32 = v - pivotV; _ = dv
            var ru float32 = float32((float64(du) * cosA - float64(dv) * sinA)) + pivotU; _ = ru
            var rv float32 = float32((float64(du) * sinA + float64(dv) * cosA)) + pivotV; _ = rv
            uvs[i] = ru
            uvs[i + 1] = rv
        }
}

