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

//OpenGLESNullImageJavaLibrary
type OpenGLESImage struct {
    Image
    CommonStrings *CommonStrings
    TextureFactory IPlatformTextureBaseFactory
    OpenGLBitmap IPlatformBitmapBase
    OpenGLESImageProperties *OpenGLESImageProperties
    ImageProcessor IOpenGLESImageProcessor
    OpenGLESImageTranslate IOpenGLESImageTranslate
}

// Implements OpenGLSurfaceChangedInterface

// The Java class OpenGLESImage has subclasses so its variables are this interface
type IOpenGLESImage interface {
    IImage
    OpenGLESImage_() *OpenGLESImage
    Set(gl GL)
    InitTexture(gl GL10) bool
    DrawRegion(gl GL10, viewHeight int, x_src float32, y_src float32, width float32, height float32, x int, y int, z int)
    Draw(gl GL10, x int, y int, z int)
    GetType() int
}

func (this *OpenGLESImage) OpenGLESImage_() *OpenGLESImage {
    return this
}

var OpenGLESImage_NULL_OPENGL_IMAGE IOpenGLESImage = NewOpenGLESImage(NullImage_NULL_IMAGE, PlatformBitmapBaseFactory_NULL_PLATFORM_BITMAP_BASE_FACTORY, PlatformTextureBaseFactory_NULL_PLATFORM_TEXTURE_BASE_FACTORY)
var OpenGLESImage_Texture2dList IBasicArrayList = NewBasicArrayListD()
var OpenGLESImage_TYPE int = 4

func (this *OpenGLESImage) initFields() {
    this.CommonStrings = CommonStrings_GetInstance()
    
    
    this.OpenGLESImageProperties = NewOpenGLESImageProperties()
    this.ImageProcessor = OpenGLESImageProcessor_GetInstance()
    this.OpenGLESImageTranslate = OpenGLESImageTranslate_GetInstance()
}


func NewOpenGLESImage(image IImage, bitmapFactory IPlatformBitmapBaseFactory, textureFactory IPlatformTextureBaseFactory) *OpenGLESImage {
    this := &OpenGLESImage{}
    this.Image = *NewImage()
    this.initFields()
        this.OpenGLBitmap = bitmapFactory.CreateBitmap(image)
        this.TextureFactory = textureFactory
    return this
}


// throws Exception
func (this *OpenGLESImage) Set(gl GL) {
        panic(errors.New(fmt.Sprint(this.CommonStrings.NOT_IMPLEMENTED)))
}


func (this *OpenGLESImage) InitTexture(gl GL10) bool {
        return false
}


func (this *OpenGLESImage) DrawRegion(gl GL10, viewHeight int, x_src float32, y_src float32, width float32, height float32, x int, y int, z int) {
}


func (this *OpenGLESImage) Draw(gl GL10, x int, y int, z int) {
}


func (this *OpenGLESImage) GetType() int {
        return OpenGLESImage_TYPE
}

