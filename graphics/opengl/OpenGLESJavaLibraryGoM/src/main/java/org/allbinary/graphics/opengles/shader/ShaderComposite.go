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

type ShaderComposite struct {
    DisableProgramShaderOpenGLProcessor IOpenGLProcessor
    RequiresOpenGLVersion string
    ShaderArray []*Shader
    ShaderInitializer IShaderInitializer
    CompositeShaderUpdater ICompositeShaderUpdater
    ModelViewProjection *ModelViewProjection
    ColorOpenGLProcessor IOpenGLProcessor
    VertexOpenGLProcessor IOpenGLProcessor
    UseProgramShaderOpenGLProcessor IOpenGLProcessor
    ShaderMatrixOpenGLProcessor IOpenGLProcessor
    ColorEnableVertexAttribArrayOpenGLProcessor IOpenGLProcessor
    NormalEnableVertexAttribArrayOpenGLProcessor IOpenGLProcessor
    VertexEnableVertexAttribArrayOpenGLProcessor IOpenGLProcessor
    TextureEnableVertexAttribArrayOpenGLProcessor IOpenGLProcessor
    ColorDisableVertexAttribArrayOpenGLProcessor IOpenGLProcessor
    NormalDisableVertexAttribArrayOpenGLProcessor IOpenGLProcessor
    VertexDisableVertexAttribArrayOpenGLProcessor IOpenGLProcessor
    TextureDisableVertexAttribArrayOpenGLProcessor IOpenGLProcessor
    UniformLightPositionOpenGLProcessor IOpenGLObject3dProcessor
    UniformLightColorOpenGLProcessor IOpenGLObject3dProcessor
    UniformCameraPositionOpenGLProcessor IOpenGLObject3dProcessor
    UniformTextureUnitOpenGLProcessor IOpenGLObject3dProcessor
    ProgramHandle int
}

// The Java class ShaderComposite has subclasses so its variables are this interface
type IShaderComposite interface {
    ShaderComposite_() *ShaderComposite
    Init(gl GL10)
}

func (this *ShaderComposite) ShaderComposite_() *ShaderComposite {
    return this
}


func (this *ShaderComposite) initFields() {
    
    
    
    
    
    
    
    
    this.UseProgramShaderOpenGLProcessor = NullOpenGLProcessorFactory_GetInstance()
    this.ShaderMatrixOpenGLProcessor = NullOpenGLProcessorFactory_GetInstance()
    this.ColorEnableVertexAttribArrayOpenGLProcessor = NullOpenGLProcessorFactory_GetInstance()
    this.NormalEnableVertexAttribArrayOpenGLProcessor = NullOpenGLProcessorFactory_GetInstance()
    this.VertexEnableVertexAttribArrayOpenGLProcessor = NullOpenGLProcessorFactory_GetInstance()
    this.TextureEnableVertexAttribArrayOpenGLProcessor = NullOpenGLProcessorFactory_GetInstance()
    this.ColorDisableVertexAttribArrayOpenGLProcessor = NullOpenGLProcessorFactory_GetInstance()
    this.NormalDisableVertexAttribArrayOpenGLProcessor = NullOpenGLProcessorFactory_GetInstance()
    this.VertexDisableVertexAttribArrayOpenGLProcessor = NullOpenGLProcessorFactory_GetInstance()
    this.TextureDisableVertexAttribArrayOpenGLProcessor = NullOpenGLProcessorFactory_GetInstance()
    this.UniformLightPositionOpenGLProcessor = NullOpenGLTextureProcessor_GetInstance()
    this.UniformLightColorOpenGLProcessor = NullOpenGLTextureProcessor_GetInstance()
    this.UniformCameraPositionOpenGLProcessor = NullOpenGLTextureProcessor_GetInstance()
    this.UniformTextureUnitOpenGLProcessor = NullOpenGLTextureProcessor_GetInstance()
    
}


func NewShaderComposite(requiresOpenGLVersion string, shaderArray []*Shader, compositeShaderUpdater ICompositeShaderUpdater, shaderInitializer IShaderInitializer, modelViewProjection *ModelViewProjection, colorOpenGLProcessor IOpenGLProcessor, vertexOpenGLProcessor IOpenGLProcessor, disableProgramShaderOpenGLProcessor IOpenGLProcessor) *ShaderComposite {
    this := &ShaderComposite{}
    this.initFields()
        this.RequiresOpenGLVersion = requiresOpenGLVersion
        this.ShaderArray = shaderArray
        this.ShaderInitializer = shaderInitializer
        this.CompositeShaderUpdater = compositeShaderUpdater
        this.ModelViewProjection = modelViewProjection
        this.ColorOpenGLProcessor = colorOpenGLProcessor
        this.VertexOpenGLProcessor = vertexOpenGLProcessor
        this.DisableProgramShaderOpenGLProcessor = disableProgramShaderOpenGLProcessor
    return this
}


func (this *ShaderComposite) Init(gl GL10) {
        this.ProgramHandle = this.ShaderInitializer.Init(gl, this.ShaderArray, this.CompositeShaderUpdater.CompositeShaderUpdater_().AttributeArray)
}

