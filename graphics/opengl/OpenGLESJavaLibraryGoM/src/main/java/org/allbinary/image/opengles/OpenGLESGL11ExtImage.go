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

//Many devices don't support this even though it is supposed to
type OpenGLESGL11ExtImage struct {
    OpenGLESImage
    LogUtil *LogUtil
    A int
    Rectangle []int
}


func (this *OpenGLESGL11ExtImage) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
}


func NewOpenGLESGL11ExtImage(image IImage, bitmapFactory IPlatformBitmapBaseFactory, textureFactory IPlatformTextureBaseFactory) *OpenGLESGL11ExtImage {
    this := &OpenGLESGL11ExtImage{}
        this.OpenGLESImage = *NewOpenGLESImage(image, bitmapFactory, textureFactory)
        this.initFields()
        this.OnDisplayChangeEvent(DisplayInfoSingleton_GetInstance().DisplayChangeEvent)
        this.Rectangle = []int{0, this.GetHeight(), this.GetWidth(), -this.GetHeight()}
    return this
}


func (this *OpenGLESGL11ExtImage) OnDisplayChangeEvent(displayChangeEvent *DisplayChangeEvent) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "onResize", e)
                }
            }()
            this.LogUtil.PutF(this.CommonStrings.START, this, "onResize")
            this.A = DisplayInfoSingleton_GetInstance().GetLastHeight() - this.GetHeight()
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *OpenGLESGL11ExtImage) Set(gl GL) {
        this.OnDisplayChangeEvent(DisplayInfoSingleton_GetInstance().DisplayChangeEvent)
        var gl11 GL11 = any(gl).(GL11); _ = gl11
        if this.OpenGLESImage.InitTexture(gl11) {
            this.TextureFactory.Load(gl11, OpenGLESImage_GL10.GL_TEXTURE_2D, 0, this, 0, true)
            gl11.GlTexParameteriv(OpenGLESImage_GL10.GL_TEXTURE_2D, GL11Ext_GL_TEXTURE_CROP_RECT_OES, this.Rectangle, 0)
            gl11.GlDisable(OpenGLESImage_GL10.GL_TEXTURE_2D)
            OpenGLLogUtil_GetInstance().LogError(gl11, this)
        }
}


func (this *OpenGLESGL11ExtImage) Draw(gl GL10, x int, y int, z int) {
        gl.GlEnable(OpenGLESImage_GL10.GL_TEXTURE_2D)
        gl.GlBindTexture(OpenGLESImage_GL10.GL_TEXTURE_2D, this.OpenGLESImageProperties.TextureID)
        var gl11Ext GL11Ext = any(gl).(GL11Ext); _ = gl11Ext
        gl11Ext.GlDrawTexfOES(float32(x), float32((this.A - y)), float32(z), float32(this.GetWidth()), float32(this.GetHeight()))
        gl.GlDisable(OpenGLESImage_GL10.GL_TEXTURE_2D)
}

