/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2022 AllBinary 
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

type OpenGLESImageProcessor struct {
}

// The Java class OpenGLESImageProcessor has subclasses so its variables are this interface
type IOpenGLESImageProcessor interface {
    OpenGLESImageProcessor_() *OpenGLESImageProcessor
    Translate(gl GL10, translateX float32, translateY float32)
    Rotate(gl GL10, angle float32)
    Scale(gl GL10, scaleX float32, scaleY float32)
    ColorMask(gl GL10, red float32, green float32, blue float32, alpha float32)
}

func (this *OpenGLESImageProcessor) OpenGLESImageProcessor_() *OpenGLESImageProcessor {
    return this
}

var OpenGLESImageProcessor_Instance IOpenGLESImageProcessor = NewOpenGLESImageProcessor()

func NewOpenGLESImageProcessor() *OpenGLESImageProcessor {
    this := &OpenGLESImageProcessor{}
    return this
}


func OpenGLESImageProcessor_GetInstance() IOpenGLESImageProcessor {
        return OpenGLESImageProcessor_Instance
}


func (this *OpenGLESImageProcessor) Translate(gl GL10, translateX float32, translateY float32) {
        gl.GlTranslatef(translateX, translateY, 0)
}


func (this *OpenGLESImageProcessor) Rotate(gl GL10, angle float32) {
        gl.GlRotatef(-angle, 0, 0, 1)
}


func (this *OpenGLESImageProcessor) Scale(gl GL10, scaleX float32, scaleY float32) {
        gl.GlScalef(scaleX, scaleY, 1)
}


func (this *OpenGLESImageProcessor) ColorMask(gl GL10, red float32, green float32, blue float32, alpha float32) {
}

