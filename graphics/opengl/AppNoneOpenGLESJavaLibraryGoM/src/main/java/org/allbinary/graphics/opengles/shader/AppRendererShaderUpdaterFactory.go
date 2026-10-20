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

type AppRendererShaderUpdaterFactory struct {
    ShaderUpdater
    NULL_SHADER_COMPOSITE_ARRAY []IShaderComposite
    ShaderCompositeArray []IShaderComposite
    ShaderOpenGLProcessorArray []IShaderOpenGLProcessor
}

var AppRendererShaderUpdaterFactory_InstanceC *AppRendererShaderUpdaterFactory = NewAppRendererShaderUpdaterFactory()

func (this *AppRendererShaderUpdaterFactory) initFields() {
    this.NULL_SHADER_COMPOSITE_ARRAY = make([]IShaderComposite, 0)
    this.ShaderCompositeArray = this.NULL_SHADER_COMPOSITE_ARRAY
    this.ShaderOpenGLProcessorArray = make([]IShaderOpenGLProcessor, 0)
}

func NewAppRendererShaderUpdaterFactory() *AppRendererShaderUpdaterFactory {
    this := &AppRendererShaderUpdaterFactory{}
    this.ShaderUpdater = *NewShaderUpdater()
    this.initFields()
    return this
}


func AppRendererShaderUpdaterFactory_GetInstance() *AppRendererShaderUpdaterFactory {
        return AppRendererShaderUpdaterFactory_InstanceC
}


func (this *AppRendererShaderUpdaterFactory) GetTestShaders() []IShaderOpenGLProcessor {
        return this.ShaderOpenGLProcessorArray
}


func (this *AppRendererShaderUpdaterFactory) SetShaderComposite(object3d any) {
}

