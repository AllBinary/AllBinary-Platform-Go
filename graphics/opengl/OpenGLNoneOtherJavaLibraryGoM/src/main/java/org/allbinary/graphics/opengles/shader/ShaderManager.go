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

type ShaderManager struct {
    LOAD_SHADER string
    CommonStrings *CommonStrings
    GL_VERTEX_SHADER int
    GL_FRAGMENT_SHADER int
}

// The Java class ShaderManager has subclasses so its variables are this interface
type IShaderManager interface {
    ShaderManager_() *ShaderManager
    Load(gl GL10, resource string, shaderAsStringArray []string, shaderType int) int
    LoadShader(gl GL10, resource string, stringList IBasicArrayList, shaderType int) int
}

func (this *ShaderManager) ShaderManager_() *ShaderManager {
    return this
}

var ShaderManager_Instance IShaderManager = NewShaderManager(-1, -1)

func (this *ShaderManager) initFields() {
    this.LOAD_SHADER = "loadShader"
    this.CommonStrings = CommonStrings_GetInstance()
    
    
}


func ShaderManager_GetInstance() IShaderManager {
        return ShaderManager_Instance
}


func NewShaderManager(GL_VERTEX_SHADER int, GL_FRAGMENT_SHADER int) *ShaderManager {
    this := &ShaderManager{}
    this.initFields()
        this.GL_VERTEX_SHADER = GL_VERTEX_SHADER
        this.GL_FRAGMENT_SHADER = GL_FRAGMENT_SHADER
    return this
}


func (this *ShaderManager) Load(gl GL10, resource string, shaderAsStringArray []string, shaderType int) int {
        return -1
}


func (this *ShaderManager) LoadShader(gl GL10, resource string, stringList IBasicArrayList, shaderType int) int {
        return -1
}

