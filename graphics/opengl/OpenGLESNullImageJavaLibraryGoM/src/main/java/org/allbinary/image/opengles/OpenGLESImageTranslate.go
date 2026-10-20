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

type OpenGLESImageTranslate struct {
}

// The Java class OpenGLESImageTranslate has subclasses so its variables are this interface
type IOpenGLESImageTranslate interface {
    OpenGLESImageTranslate_() *OpenGLESImageTranslate
    Translate(gl GL10, openGLESImage IOpenGLESImage, x float32, y float32)
    Translate2(gl GL10, openGLESImage IOpenGLESImage)
}

func (this *OpenGLESImageTranslate) OpenGLESImageTranslate_() *OpenGLESImageTranslate {
    return this
}

var OpenGLESImageTranslate_Instance IOpenGLESImageTranslate = NewOpenGLESImageTranslate()

func NewOpenGLESImageTranslate() *OpenGLESImageTranslate {
    this := &OpenGLESImageTranslate{}
    return this
}


func OpenGLESImageTranslate_GetInstance() IOpenGLESImageTranslate {
        return OpenGLESImageTranslate_Instance
}


func (this *OpenGLESImageTranslate) Translate(gl GL10, openGLESImage IOpenGLESImage, x float32, y float32) {
        openGLESImage.OpenGLESImage_().ImageProcessor.Translate(gl, x, y)
}


func (this *OpenGLESImageTranslate) Translate2(gl GL10, openGLESImage IOpenGLESImage) {
}

