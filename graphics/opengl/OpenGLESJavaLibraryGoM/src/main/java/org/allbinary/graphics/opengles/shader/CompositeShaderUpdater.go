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

type CompositeShaderUpdater struct {
    ShaderUpdater
    CommonSeps *CommonSeps
    ShaderStrings *ShaderStrings
    UniformBlockHandleArray []int
    UniformBlockArray []string
    UniformHandleArray []int
    UniformArray []string
    AttributeHandleArray []int
    AttributeArray []string
}

// The Java class CompositeShaderUpdater has subclasses so its variables are this interface
type ICompositeShaderUpdater interface {
    IShaderUpdater
    CompositeShaderUpdater_() *CompositeShaderUpdater
    OnSurfaceCreated(gl GL10, eglConfig *EGLConfig, programHandle int)
    OnSurfaceDestroyed(gl GL10)
}

func (this *CompositeShaderUpdater) CompositeShaderUpdater_() *CompositeShaderUpdater {
    return this
}

var CompositeShaderUpdater_InstanceC ICompositeShaderUpdater = NewCompositeShaderUpdater(StringUtil_GetInstance().GetArrayInstance(), StringUtil_GetInstance().GetArrayInstance(), StringUtil_GetInstance().GetArrayInstance(), NullUtil_GetInstance().NULL_INT_ARRAY)

func (this *CompositeShaderUpdater) initFields() {
    this.CommonSeps = CommonSeps_GetInstance()
    this.ShaderStrings = ShaderStrings_GetInstance()
    
    
    
    
    
    
}


func CompositeShaderUpdater_GetInstance() ICompositeShaderUpdater {
        return CompositeShaderUpdater_InstanceC
}


func NewCompositeShaderUpdater(uniformBlockArray []string, uniformArray []string, attributeArray []string, attributeHandleArray []int) *CompositeShaderUpdater {
    this := &CompositeShaderUpdater{}
    this.ShaderUpdater = *NewShaderUpdater()
    this.initFields()
        this.UniformBlockArray = uniformBlockArray
        this.UniformBlockHandleArray = make([]int, len(this.UniformBlockArray))
        this.UniformArray = uniformArray
        this.UniformHandleArray = make([]int, len(this.UniformArray))
        this.AttributeArray = attributeArray
        this.AttributeHandleArray = attributeHandleArray
    return this
}


func (this *CompositeShaderUpdater) OnSurfaceCreated(gl GL10, eglConfig *EGLConfig, programHandle int) {
}


func (this *CompositeShaderUpdater) OnSurfaceDestroyed(gl GL10) {
}

