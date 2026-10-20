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

type OpenGLCapabilities struct {
    VERSION_1_0 string
    VERSION_1_1 string
    VERSION_2_0 string
    VERSION_3_0 string
    VERSION_3_1 string
    VERSION_3_2 string
    VERSION_HIGHER_THAN_EXISTS string
    GlVersion string
    ShaderVersion int
    GlInstanceVersion string
    MaxTextureSize int
}

var OpenGLCapabilities_Instance *OpenGLCapabilities = NewOpenGLCapabilities()

func (this *OpenGLCapabilities) initFields() {
    this.VERSION_1_0 = "1.0"
    this.VERSION_1_1 = "1.1"
    this.VERSION_2_0 = "2.0"
    this.VERSION_3_0 = "3.0"
    this.VERSION_3_1 = "3.1"
    this.VERSION_3_2 = "3.2"
    this.VERSION_HIGHER_THAN_EXISTS = "999.999"
    this.GlVersion = StringUtil_GetInstance().EMPTY_STRING
    this.ShaderVersion = 0
    this.GlInstanceVersion = StringUtil_GetInstance().EMPTY_STRING
    this.MaxTextureSize = 64
}


func OpenGLCapabilities_GetInstance() *OpenGLCapabilities {
        return OpenGLCapabilities_Instance
}


func NewOpenGLCapabilities() *OpenGLCapabilities {
    this := &OpenGLCapabilities{}
    this.initFields()
    return this
}


func (this *OpenGLCapabilities) InitCapabilities(gl any) {
}


func (this *OpenGLCapabilities) IsGlExtensionDrawTexture() bool {
        return false
}


func (this *OpenGLCapabilities) GetGlVersion() string {
        return this.GlVersion
}


func (this *OpenGLCapabilities) IsVertexBufferObjectSupport() bool {
        return false
}


func (this *OpenGLCapabilities) IsTextureSizeValid(widthAndHeight int) bool {
        return true
}

