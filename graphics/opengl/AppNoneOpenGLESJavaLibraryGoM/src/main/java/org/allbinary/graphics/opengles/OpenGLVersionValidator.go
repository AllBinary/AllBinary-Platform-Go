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

type OpenGLVersionValidator struct {
    EGL_CONTEXT_CLIENT_VERSION_VALUE int
}

var OpenGLVersionValidator_Instance *OpenGLVersionValidator = NewOpenGLVersionValidator()

func (this *OpenGLVersionValidator) initFields() {
    this.EGL_CONTEXT_CLIENT_VERSION_VALUE = -1
}

func NewOpenGLVersionValidator() *OpenGLVersionValidator {
    this := &OpenGLVersionValidator{}
    this.initFields()
    return this
}


func OpenGLVersionValidator_GetInstance() *OpenGLVersionValidator {
        return OpenGLVersionValidator_Instance
}


func (this *OpenGLVersionValidator) IsGL31OrHigher() bool {
        return false
}


func (this *OpenGLVersionValidator) IsAvailable(glVersionRequired string) bool {
        return false
}


func (this *OpenGLVersionValidator) SetShaderComposite(shaderComposite IShaderComposite, anyType any) {
}

