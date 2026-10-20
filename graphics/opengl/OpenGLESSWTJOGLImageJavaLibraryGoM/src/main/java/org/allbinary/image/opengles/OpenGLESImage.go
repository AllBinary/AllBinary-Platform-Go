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

//OpenGLESSWTJOGLImageJavaLibrary
type OpenGLESImage struct {
    Image
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    TextureFactory IPlatformTextureBaseFactory
    OpenGLBitmap IPlatformBitmapBase
    ImageProcessor IOpenGLESImageProcessor
    OpenGLESImageTranslate IOpenGLESImageTranslate
    OpenGLESImageProperties *OpenGLESImageProperties
}

// Implements OpenGLSurfaceChangedInterface

// The Java class OpenGLESImage has subclasses so its variables are this interface
type IOpenGLESImage interface {
    IImage
    OpenGLESImage_() *OpenGLESImage
    GetHeight() int
    GetWidth() int
    Set(gl GL)
    InitTexture(gl GL10) bool
    DrawRegion(gl GL10, viewHeight int, x_src float32, y_src float32, width float32, height float32, x int, y int, z int)
    Draw(gl GL10, x int, y int, z int)
    GetType() int
    ToString() string
}

func (this *OpenGLESImage) OpenGLESImage_() *OpenGLESImage {
    return this
}

var OpenGLESImage_NULL_OPENGL_IMAGE IOpenGLESImage = NewOpenGLESImage(NullImage_NULL_IMAGE, PlatformBitmapBaseFactory_NULL_PLATFORM_BITMAP_BASE_FACTORY, PlatformTextureBaseFactory_NULL_PLATFORM_TEXTURE_BASE_FACTORY)
var OpenGLESImage_Texture2dList IBasicArrayList = NewBasicArrayListD()
var OpenGLESImage_TYPE int = 4

func (this *OpenGLESImage) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
    
    this.ImageProcessor = OpenGLESImageProcessor_GetInstance()
    this.OpenGLESImageTranslate = OpenGLESImageTranslate_GetInstance()
    this.OpenGLESImageProperties = NewOpenGLESImageProperties()
}


func NewOpenGLESImage(image IImage, bitmapFactory IPlatformBitmapBaseFactory, textureFactory IPlatformTextureBaseFactory) *OpenGLESImage {
    this := &OpenGLESImage{}
        this.Image = *NewImage(image.GetName(), PostLoadImageProcessor_NULL_POST_LOAD_IMAGE_PROCESSOR)
        this.initFields()
        this.OpenGLBitmap = bitmapFactory.CreateBitmap(image)
        this.TextureFactory = textureFactory
        OpenGLImageCacheFactory_GetInstance().Init(this)
        this.PlatformImage = OpenGLESPostLoadPlatformImage_GetInstance()
    return this
}


func (this *OpenGLESImage) GetHeight() int {
        return int((float32(this.OpenGLBitmap.GetHeight()) * this.OpenGLESImageProperties.ScaleY2))
}


func (this *OpenGLESImage) GetWidth() int {
        return int((float32(this.OpenGLBitmap.GetWidth()) * this.OpenGLESImageProperties.ScaleX2))
}


// throws Exception
func (this *OpenGLESImage) Set(gl GL) {
        panic(errors.New(fmt.Sprint(this.CommonStrings.NOT_IMPLEMENTED)))
}


func (this *OpenGLESImage) InitTexture(gl GL10) bool {
        if !OpenGLESImage_Texture2dList.Contains(this) {
            OpenGLESImage_Texture2dList.Add(this)
            var textures []int = make([]int, 1); _ = textures
            gl.GlEnable(Image_GL10.GL_TEXTURE_2D)
            if this.OpenGLESImageProperties.TextureID != -1 {
                textures[0] = this.OpenGLESImageProperties.TextureID
                gl.GlDeleteTextures(1, textures, 0)
            }
            gl.GlGenTextures(1, textures, 0)
            this.OpenGLESImageProperties.TextureID = textures[0]
            gl.GlBindTexture(Image_GL10.GL_TEXTURE_2D, this.OpenGLESImageProperties.TextureID)
            gl.GlTexParameterf(Image_GL10.GL_TEXTURE_2D, Image_GL10.GL_TEXTURE_MIN_FILTER, Image_GL10.GL_NEAREST)
            gl.GlTexParameterf(Image_GL10.GL_TEXTURE_2D, Image_GL10.GL_TEXTURE_WRAP_S, Image_GL10.GL_CLAMP_TO_EDGE)
            gl.GlTexParameterf(Image_GL10.GL_TEXTURE_2D, Image_GL10.GL_TEXTURE_WRAP_T, Image_GL10.GL_CLAMP_TO_EDGE)
            return true
        }
        return false
}


func (this *OpenGLESImage) DrawRegion(gl GL10, viewHeight int, x_src float32, y_src float32, width float32, height float32, x int, y int, z int) {
}


func (this *OpenGLESImage) Draw(gl GL10, x int, y int, z int) {
}


func (this *OpenGLESImage) GetType() int {
        return OpenGLESImage_TYPE
}


func (this *OpenGLESImage) ToString() string {
        var graphicsStrings *GraphicsStrings = GraphicsStrings_GetInstance(); _ = graphicsStrings
        return fmt.Sprint(NewStringMaker().Append(this.GetName()).Append(this.Image.ToString()).Append(graphicsStrings.OPACITY).Appendfloat(this.OpenGLESImageProperties.Alphaf))
}

func (this *OpenGLESImage) String() string {
    return this.ToString()
}

