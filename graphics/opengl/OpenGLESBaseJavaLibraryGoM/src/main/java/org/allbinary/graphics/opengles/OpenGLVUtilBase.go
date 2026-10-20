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

type OpenGLVUtilBase struct {
}

// The Java class OpenGLVUtilBase has subclasses so its variables are this interface
type IOpenGLVUtilBase interface {
    OpenGLVUtilBase_() *OpenGLVUtilBase
    GlTexEnvix(gl GL10, target int, pname int, param int)
    GlTexParameterix(gl GL10, target int, pname int, param int)
    GlGenerateMipmap(gl GL10, textureId int, target int, generateMipMap bool)
}

func (this *OpenGLVUtilBase) OpenGLVUtilBase_() *OpenGLVUtilBase {
    return this
}


func NewOpenGLVUtilBase() *OpenGLVUtilBase {
    this := &OpenGLVUtilBase{}
    return this
}


func (this *OpenGLVUtilBase) GlTexEnvix(gl GL10, target int, pname int, param int) {
}


func (this *OpenGLVUtilBase) GlTexParameterix(gl GL10, target int, pname int, param int) {
}


func (this *OpenGLVUtilBase) GlGenerateMipmap(gl GL10, textureId int, target int, generateMipMap bool) {
}

