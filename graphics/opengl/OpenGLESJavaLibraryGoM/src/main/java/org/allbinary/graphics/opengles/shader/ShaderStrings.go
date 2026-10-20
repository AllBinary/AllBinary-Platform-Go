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

type ShaderStrings struct {
    NO_PROGRAM_HANDLE string
    UNIFORM_BLOCK_HANDLE string
    UNIFORM_HANDLE string
    UNIFORM_ATTRIBUTE_HANDLE string
}

var ShaderStrings_Instance *ShaderStrings = NewShaderStrings()

func (this *ShaderStrings) initFields() {
    this.NO_PROGRAM_HANDLE = "No program handle"
    this.UNIFORM_BLOCK_HANDLE = "Uniform Block Handle: "
    this.UNIFORM_HANDLE = "Uniform Handle: "
    this.UNIFORM_ATTRIBUTE_HANDLE = "Uniform Attribute Handle: "
}

func NewShaderStrings() *ShaderStrings {
    this := &ShaderStrings{}
    this.initFields()
    return this
}


func ShaderStrings_GetInstance() *ShaderStrings {
        return ShaderStrings_Instance
}

