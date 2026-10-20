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

type OpenGLESDeviceImageTranslate struct {
    OpenGLESImageTranslate
    TranslateX float32
    TranslateY float32
    TranslateX2 float32
    TranslateY2 float32
}


func (this *OpenGLESDeviceImageTranslate) initFields() {
    this.TranslateX = 0.0
    this.TranslateY = 0.0
    this.TranslateX2 = 0.0
    this.TranslateY2 = 0.0
}

func NewOpenGLESDeviceImageTranslate() *OpenGLESDeviceImageTranslate {
    this := &OpenGLESDeviceImageTranslate{}
    this.OpenGLESImageTranslate = *NewOpenGLESImageTranslate()
    this.initFields()
    return this
}


func (this *OpenGLESDeviceImageTranslate) Translate(gl GL10, openGLESImage IOpenGLESImage, x float32, y float32) {
        openGLESImage.OpenGLESImage_().ImageProcessor.Translate(gl, this.TranslateX + (openGLESImage.OpenGLESImage_().OpenGLESImageProperties.ScaleX * x), this.TranslateY + (openGLESImage.OpenGLESImage_().OpenGLESImageProperties.ScaleY * y))
}


func (this *OpenGLESDeviceImageTranslate) Translate2(gl GL10, openGLESImage IOpenGLESImage) {
        openGLESImage.OpenGLESImage_().ImageProcessor.Translate(gl, this.TranslateX2, this.TranslateY2)
}

