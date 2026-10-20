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

type OpenGLESGL10Image struct {
    OpenGLESImage
    TextureVertexFloatBuffer *FloatBuffer
    RealOpenGLESImageDraw *OpenGLESImageDraw
    OpenGLESImageDraw *OpenGLESImageDraw
}


func (this *OpenGLESGL10Image) initFields() {
    this.TextureVertexFloatBuffer = ByteBuffer_AllocateDirect(4 * 4 * 2).Order(ByteOrder_NativeOrder()).AsFloatBuffer()
    this.RealOpenGLESImageDraw = /* Java anonymous class body: implement it as a named type
        this.GlUtil = GLUtil_GetInstance()
        this.DisplayInfoSingleton = DisplayInfoSingleton_GetInstance()
        this.TextureVertexFloatBuffer = this.TextureVertexFloatBuffer
        this.RegionTextureRectangleFloatArray = make([]float32, 8)
        this.RegionTextureVertexFloatBuffer = ByteBuffer_AllocateDirect(4 * 4 * 2).Order(ByteOrder_NativeOrder()).AsFloatBuffer()
        this.RegionRectangleFloatArray = []float32{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}
        this.RegionRectangleVertexFloatBuffer = ByteBuffer_AllocateDirect(4 * 4 * 3).Order(ByteOrder_NativeOrder()).AsFloatBuffer()
        this.OpenGLESImageProperties = this.OpenGLESImageProperties
        

    func (this *OpenGLESGL10Image) DrawRegion(gl GL10, viewHeight int, x_src float32, y_src float32, width float32, height float32, x int, y int, z int) {
            var imageWidth int = this.GetWidth(); _ = imageWidth
            var imageHeight int = this.GetHeight(); _ = imageHeight
            this.RegionRectangleFloatArray[7] = viewHeight
            this.RegionRectangleFloatArray[1] = this.RegionRectangleFloatArray[7] - height
            this.RegionRectangleFloatArray[3] = width
            this.RegionRectangleFloatArray[4] = this.RegionRectangleFloatArray[1]
            this.RegionRectangleFloatArray[9] = this.RegionRectangleFloatArray[3]
            this.RegionRectangleFloatArray[10] = this.RegionRectangleFloatArray[7]
            this.RegionTextureRectangleFloatArray[0] = x_src / imageWidth
            this.RegionTextureRectangleFloatArray[1] = ((y_src + height)) / imageHeight
            this.RegionTextureRectangleFloatArray[2] = ((x_src + width)) / imageWidth
            this.RegionTextureRectangleFloatArray[3] = this.RegionTextureRectangleFloatArray[1]
            this.RegionTextureRectangleFloatArray[4] = this.RegionTextureRectangleFloatArray[0]
            this.RegionTextureRectangleFloatArray[5] = y_src / imageHeight
            this.RegionTextureRectangleFloatArray[6] = this.RegionTextureRectangleFloatArray[2]
            this.RegionTextureRectangleFloatArray[7] = this.RegionTextureRectangleFloatArray[5]
            var u_center float32 = (this.RegionTextureRectangleFloatArray[0] + this.RegionTextureRectangleFloatArray[2]) / 2.0; _ = u_center
            var v_center float32 = (this.RegionTextureRectangleFloatArray[5] + this.RegionTextureRectangleFloatArray[1]) / 2.0; _ = v_center
            this.GlUtil.RotateUVs(this.RegionTextureRectangleFloatArray, -this.OpenGLESImageProperties.Angle, u_center, v_center)
            gl.GlPushMatrix()
            var imageProcessor IOpenGLESImageProcessor = this.ImageProcessor; _ = imageProcessor
            this.OpenGLESImageTranslate.Translate(gl, this, x, -y)
            imageProcessor.Scale(gl, this.OpenGLESImageProperties.ScaleX, this.OpenGLESImageProperties.ScaleY)
            imageProcessor.ColorMask(gl, this.OpenGLESImageProperties.Redf, this.OpenGLESImageProperties.Greenf, this.OpenGLESImageProperties.Bluef, this.OpenGLESImageProperties.Alphaf)
            this.OpenGLESImageTranslate.Translate2(gl, this)
            this.RegionRectangleVertexFloatBuffer.Put(this.RegionRectangleFloatArray)
            this.GlUtil.Position(this.RegionRectangleVertexFloatBuffer, 0)
            gl.GlVertexPointer(3, OpenGLESImage_GL10.GL_FLOAT, 0, this.RegionRectangleVertexFloatBuffer)
            gl.GlEnableClientState(OpenGLESImage_GL10.GL_VERTEX_ARRAY)
            gl.GlEnable(OpenGLESImage_GL10.GL_TEXTURE_2D)
            gl.GlBindTexture(OpenGLESImage_GL10.GL_TEXTURE_2D, this.OpenGLESImageProperties.TextureID)
            gl.GlEnableClientState(OpenGLESImage_GL10.GL_TEXTURE_COORD_ARRAY)
            this.RegionTextureVertexFloatBuffer.Put(this.RegionTextureRectangleFloatArray)
            this.GlUtil.Position(this.RegionTextureVertexFloatBuffer, 0)
            gl.GlTexCoordPointer(2, OpenGLESImage_GL10.GL_FLOAT, 0, this.RegionTextureVertexFloatBuffer)
            gl.GlDrawArrays(OpenGLESImage_GL10.GL_TRIANGLE_STRIP, 0, 4)
            gl.GlDisableClientState(OpenGLESImage_GL10.GL_VERTEX_ARRAY)
            gl.GlDisableClientState(OpenGLESImage_GL10.GL_TEXTURE_COORD_ARRAY)
            gl.GlDisable(OpenGLESImage_GL10.GL_TEXTURE_2D)
            gl.GlPopMatrix()
    }
        

    func (this *OpenGLESGL10Image) Draw(gl GL10, x int, y int, z int) {
            var width int = this.GetWidth(); _ = width
            var height int = this.GetHeight(); _ = height
            this.RegionRectangleFloatArray[7] = this.DisplayInfoSingleton.GetLastHeight()
            this.RegionRectangleFloatArray[1] = this.RegionRectangleFloatArray[7] - height
            this.RegionRectangleFloatArray[3] = width
            this.RegionRectangleFloatArray[4] = this.RegionRectangleFloatArray[1]
            this.RegionRectangleFloatArray[9] = this.RegionRectangleFloatArray[3]
            this.RegionRectangleFloatArray[10] = this.RegionRectangleFloatArray[7]
            this.RegionTextureRectangleFloatArray[0] = 0
            this.RegionTextureRectangleFloatArray[1] = 1
            this.RegionTextureRectangleFloatArray[2] = 1
            this.RegionTextureRectangleFloatArray[3] = 1
            this.RegionTextureRectangleFloatArray[4] = 0
            this.RegionTextureRectangleFloatArray[5] = 0
            this.RegionTextureRectangleFloatArray[6] = 1
            this.RegionTextureRectangleFloatArray[7] = 0
            gl.GlPushMatrix()
            this.OpenGLESImageTranslate.Translate(gl, this, x, -y)
            var imageProcessor IOpenGLESImageProcessor = this.ImageProcessor; _ = imageProcessor
            imageProcessor.Scale(gl, this.OpenGLESImageProperties.ScaleX, this.OpenGLESImageProperties.ScaleY)
            imageProcessor.ColorMask(gl, this.OpenGLESImageProperties.Redf, this.OpenGLESImageProperties.Greenf, this.OpenGLESImageProperties.Bluef, this.OpenGLESImageProperties.Alphaf)
            this.OpenGLESImageTranslate.Translate2(gl, this)
            this.RegionRectangleVertexFloatBuffer.Put(this.RegionRectangleFloatArray)
            this.GlUtil.Position(this.RegionRectangleVertexFloatBuffer, 0)
            gl.GlVertexPointer(3, OpenGLESImage_GL10.GL_FLOAT, 0, this.RegionRectangleVertexFloatBuffer)
            gl.GlEnableClientState(OpenGLESImage_GL10.GL_VERTEX_ARRAY)
            gl.GlEnable(OpenGLESImage_GL10.GL_TEXTURE_2D)
            gl.GlBindTexture(OpenGLESImage_GL10.GL_TEXTURE_2D, this.OpenGLESImageProperties.TextureID)
            gl.GlEnableClientState(OpenGLESImage_GL10.GL_TEXTURE_COORD_ARRAY)
            this.GlUtil.RotateUVs(this.RegionTextureRectangleFloatArray, -this.OpenGLESImageProperties.Angle, 0.5, 0.5)
            this.GlUtil.Position(this.TextureVertexFloatBuffer, 0)
            this.TextureVertexFloatBuffer.Put(this.RegionTextureRectangleFloatArray)
            this.GlUtil.Position(this.TextureVertexFloatBuffer, 0)
            gl.GlTexCoordPointer(2, OpenGLESImage_GL10.GL_FLOAT, 0, this.TextureVertexFloatBuffer)
            gl.GlDrawArrays(OpenGLESImage_GL10.GL_TRIANGLE_STRIP, 0, 4)
            gl.GlDisableClientState(OpenGLESImage_GL10.GL_VERTEX_ARRAY)
            gl.GlDisableClientState(OpenGLESImage_GL10.GL_TEXTURE_COORD_ARRAY)
            gl.GlDisable(OpenGLESImage_GL10.GL_TEXTURE_2D)
            gl.GlPopMatrix()
    }
*/ NewOpenGLESImageDraw()
    this.OpenGLESImageDraw = OpenGLESImageDraw_GetInstance()
}


func NewOpenGLESGL10Image(image IImage, bitmapFactory IPlatformBitmapBaseFactory, textureFactory IPlatformTextureBaseFactory) *OpenGLESGL10Image {
    this := &OpenGLESGL10Image{}
        this.OpenGLESImage = *NewOpenGLESImage(image, bitmapFactory, textureFactory)
        this.initFields()
        this.InitVertices()
    return this
}


func (this *OpenGLESGL10Image) InitVertices() {
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


func (this *OpenGLESGL10Image) Set(gl GL) {
        var gl10 GL10 = any(gl).(GL10); _ = gl10
        if this.OpenGLESImage.InitTexture(gl10) {
            this.TextureFactory.Load(gl10, OpenGLESImage_GL10.GL_TEXTURE_2D, 0, this, 0, true)
            gl10.GlDisable(OpenGLESImage_GL10.GL_TEXTURE_2D)
            this.OpenGLESImageDraw = this.RealOpenGLESImageDraw
            OpenGLLogUtil_GetInstance().LogError(gl10, this)
        }
}


func (this *OpenGLESGL10Image) DrawRegion(gl GL10, viewHeight int, x_src float32, y_src float32, width float32, height float32, x int, y int, z int) {
        this.OpenGLESImageDraw.DrawRegion(gl, viewHeight, x_src, y_src, width, height, x, y, z)
}


func (this *OpenGLESGL10Image) Draw(gl GL10, x int, y int, z int) {
        this.OpenGLESImageDraw.Draw(gl, x, y, z)
}

