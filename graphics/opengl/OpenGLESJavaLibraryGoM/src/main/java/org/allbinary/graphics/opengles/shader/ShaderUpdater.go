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

type ShaderUpdater struct {
    CommonStrings *CommonStrings
    RendererStrings *RendererStrings
}

// The Java class ShaderUpdater has subclasses so its variables are this interface
type IShaderUpdater interface {
    ShaderUpdater_() *ShaderUpdater
    OnSurfaceCreated(gl GL10, eglConfig *EGLConfig)
}

func (this *ShaderUpdater) ShaderUpdater_() *ShaderUpdater {
    return this
}

var ShaderUpdater_Instance IShaderUpdater = NewShaderUpdater()

func (this *ShaderUpdater) initFields() {
    this.CommonStrings = CommonStrings_GetInstance()
    this.RendererStrings = RendererStrings_GetInstance()
}

func NewShaderUpdater() *ShaderUpdater {
    this := &ShaderUpdater{}
    this.initFields()
    return this
}


func ShaderUpdater_GetInstance() IShaderUpdater {
        return ShaderUpdater_Instance
}


func (this *ShaderUpdater) OnSurfaceCreated(gl GL10, eglConfig *EGLConfig) {
}

