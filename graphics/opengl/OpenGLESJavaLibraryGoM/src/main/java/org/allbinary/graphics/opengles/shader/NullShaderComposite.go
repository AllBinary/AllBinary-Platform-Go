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

type NullShaderComposite struct {
    ShaderComposite
}

var NullShaderComposite_Instance *NullShaderComposite = /* Java anonymous class body: implement it as a named type
        

    func (this *NullShaderComposite) Init(gl GL10) {
    }
*/ NewNullShaderComposite(OpenGLCapabilities_GetInstance().VERSION_2_0, []*Shader{NewShader(), NewShader()}, CompositeShaderUpdater_GetInstance(), ShaderInitializer_GetInstance(), ModelViewProjection_GetInstance(), nil, nil, NullOpenGLProcessorFactory_GetInstance())


func NullShaderComposite_GetInstance() IShaderComposite {
        return NullShaderComposite_Instance
}


func NewNullShaderComposite(requiresOpenGLVersion string, shaderArray []*Shader, compositeShaderUpdater ICompositeShaderUpdater, shaderInitializer IShaderInitializer, modelViewProjection *ModelViewProjection, colorOpenGLProcessor IOpenGLProcessor, vertexOpenGLProcessor IOpenGLProcessor, disableProgramShaderOpenGLProcessor IOpenGLProcessor) *NullShaderComposite {
    this := &NullShaderComposite{}
        this.ShaderComposite = *NewShaderComposite(requiresOpenGLVersion, shaderArray, compositeShaderUpdater, shaderInitializer, modelViewProjection, colorOpenGLProcessor, vertexOpenGLProcessor, disableProgramShaderOpenGLProcessor)
        this.UseProgramShaderOpenGLProcessor = NullOpenGLProcessorFactory_GetInstance()
    return this
}

