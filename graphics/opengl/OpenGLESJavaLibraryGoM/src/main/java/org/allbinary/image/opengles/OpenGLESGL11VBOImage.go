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

//This is not actually VBO impl
type OpenGLESGL11VBOImage struct {
    OpenGLESImage
    GlUtil *GLUtil
    TextureVertexFloatBuffer *FloatBuffer
    RegionRectangleFloatArray []float32
    RegionRectangleVertexFloatBuffer *FloatBuffer
    RegionTextureRectangleFloatArray []float32
    RegionTextureVertexFloatBuffer *FloatBuffer
}


func (this *OpenGLESGL11VBOImage) initFields() {
    this.GlUtil = GLUtil_GetInstance()
    this.TextureVertexFloatBuffer = ByteBuffer_AllocateDirect(4 * 4 * 2).Order(ByteOrder_NativeOrder()).AsFloatBuffer()
    this.RegionRectangleFloatArray = []float32{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
    this.RegionRectangleVertexFloatBuffer = ByteBuffer_AllocateDirect(4 * 4 * 3).Order(ByteOrder_NativeOrder()).AsFloatBuffer()
    this.RegionTextureRectangleFloatArray = make([]float32, 8)
    this.RegionTextureVertexFloatBuffer = ByteBuffer_AllocateDirect(4 * 4 * 2).Order(ByteOrder_NativeOrder()).AsFloatBuffer()
}


func NewOpenGLESGL11VBOImage(image IImage, bitmapFactory IPlatformBitmapBaseFactory, textureFactory IPlatformTextureBaseFactory) *OpenGLESGL11VBOImage {
    this := &OpenGLESGL11VBOImage{}
        this.OpenGLESImage = *NewOpenGLESImage(image, bitmapFactory, textureFactory)
        this.initFields()
        this.InitVertices()
        this.RegionRectangleVertexFloatBuffer.Put(FloatBuffer_Wrap(this.RegionRectangleFloatArray))
        this.RegionTextureVertexFloatBuffer.Put(FloatBuffer_Wrap(this.RegionTextureRectangleFloatArray))
    return this
}


func (this *OpenGLESGL11VBOImage) InitVertices() {
        this.TextureVertexFloatBuffer.Put(0)
        this.TextureVertexFloatBuffer.Put(1)
        this.TextureVertexFloatBuffer.Put(1)
        this.TextureVertexFloatBuffer.Put(1)
        this.TextureVertexFloatBuffer.Put(0)
        this.TextureVertexFloatBuffer.Put(0)
        this.TextureVertexFloatBuffer.Put(1)
        this.TextureVertexFloatBuffer.Put(0)
        this.TextureVertexFloatBuffer.Rewind()
}


func (this *OpenGLESGL11VBOImage) Set(gl GL) {
        var gl10 GL10 = any(gl).(GL10); _ = gl10
        if this.OpenGLESImage.InitTexture(gl10) {
            this.TextureFactory.Load(gl10, OpenGLESImage_GL10.GL_TEXTURE_2D, 0, this, 0, true)
            gl10.GlDisable(OpenGLESImage_GL10.GL_TEXTURE_2D)
            OpenGLLogUtil_GetInstance().LogError(gl10, this)
        }
}


func (this *OpenGLESGL11VBOImage) DrawRegion(gl GL10, viewHeight int, x_src float32, y_src float32, width float32, height float32, x int, y int, z int) {
        this.RegionRectangleFloatArray[0] = float32(x)
        this.RegionRectangleFloatArray[7] = float32(viewHeight - y)
        this.RegionRectangleFloatArray[1] = this.RegionRectangleFloatArray[7] - height
        this.RegionRectangleFloatArray[3] = float32(x) + width
        this.RegionRectangleFloatArray[4] = this.RegionRectangleFloatArray[1]
        this.RegionRectangleFloatArray[6] = float32(x)
        this.RegionRectangleFloatArray[9] = this.RegionRectangleFloatArray[3]
        this.RegionRectangleFloatArray[10] = this.RegionRectangleFloatArray[7]
        this.RegionTextureRectangleFloatArray[0] = x_src / this.GetWidth()
        this.RegionTextureRectangleFloatArray[1] = (y_src + height) / this.GetHeight()
        this.RegionTextureRectangleFloatArray[2] = (x_src + width) / this.GetWidth()
        this.RegionTextureRectangleFloatArray[3] = this.RegionTextureRectangleFloatArray[1]
        this.RegionTextureRectangleFloatArray[4] = this.RegionTextureRectangleFloatArray[0]
        this.RegionTextureRectangleFloatArray[5] = y_src / this.GetHeight()
        this.RegionTextureRectangleFloatArray[6] = this.RegionTextureRectangleFloatArray[2]
        this.RegionTextureRectangleFloatArray[7] = this.RegionTextureRectangleFloatArray[5]
        this.GlUtil.Position(this.RegionRectangleVertexFloatBuffer, 0)
        gl.GlVertexPointer(3, OpenGLESImage_GL10.GL_FLOAT, 0, this.RegionRectangleVertexFloatBuffer)
        gl.GlEnableClientState(OpenGLESImage_GL10.GL_VERTEX_ARRAY)
        gl.GlEnable(OpenGLESImage_GL10.GL_TEXTURE_2D)
        gl.GlBindTexture(OpenGLESImage_GL10.GL_TEXTURE_2D, this.OpenGLESImageProperties.TextureID)
        gl.GlEnableClientState(OpenGLESImage_GL10.GL_TEXTURE_COORD_ARRAY)
        this.GlUtil.Position(this.RegionTextureVertexFloatBuffer, 0)
        gl.GlTexCoordPointer(2, OpenGLESImage_GL10.GL_FLOAT, 0, this.RegionTextureVertexFloatBuffer)
        gl.GlDrawArrays(OpenGLESImage_GL10.GL_TRIANGLE_STRIP, 0, 4)
        gl.GlDisableClientState(OpenGLESImage_GL10.GL_VERTEX_ARRAY)
        gl.GlDisableClientState(OpenGLESImage_GL10.GL_TEXTURE_COORD_ARRAY)
        gl.GlDisable(OpenGLESImage_GL10.GL_TEXTURE_2D)
}


func (this *OpenGLESGL11VBOImage) Draw(gl GL10, x int, y int, z int) {
        this.RegionRectangleFloatArray[0] = float32(x)
        this.RegionRectangleFloatArray[7] = float32((DisplayInfoSingleton_GetInstance().GetLastHeight() - y))
        this.RegionRectangleFloatArray[1] = this.RegionRectangleFloatArray[7] - this.GetHeight()
        this.RegionRectangleFloatArray[3] = float32((x + this.GetWidth()))
        this.RegionRectangleFloatArray[4] = this.RegionRectangleFloatArray[1]
        this.RegionRectangleFloatArray[6] = float32(x)
        this.RegionRectangleFloatArray[9] = this.RegionRectangleFloatArray[3]
        this.RegionRectangleFloatArray[10] = this.RegionRectangleFloatArray[7]
        this.GlUtil.Position(this.RegionRectangleVertexFloatBuffer, 0)
        gl.GlVertexPointer(3, OpenGLESImage_GL10.GL_FLOAT, 0, this.RegionRectangleVertexFloatBuffer)
        gl.GlEnableClientState(OpenGLESImage_GL10.GL_VERTEX_ARRAY)
        gl.GlEnable(OpenGLESImage_GL10.GL_TEXTURE_2D)
        gl.GlBindTexture(OpenGLESImage_GL10.GL_TEXTURE_2D, this.OpenGLESImageProperties.TextureID)
        gl.GlEnableClientState(OpenGLESImage_GL10.GL_TEXTURE_COORD_ARRAY)
        this.GlUtil.Position(this.TextureVertexFloatBuffer, 0)
        gl.GlTexCoordPointer(2, OpenGLESImage_GL10.GL_FLOAT, 0, this.TextureVertexFloatBuffer)
        gl.GlDrawArrays(OpenGLESImage_GL10.GL_TRIANGLE_STRIP, 0, 4)
        gl.GlDisableClientState(OpenGLESImage_GL10.GL_VERTEX_ARRAY)
        gl.GlDisableClientState(OpenGLESImage_GL10.GL_TEXTURE_COORD_ARRAY)
        gl.GlDisable(OpenGLESImage_GL10.GL_TEXTURE_2D)
}

