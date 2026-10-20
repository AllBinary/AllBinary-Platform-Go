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

type OpenGLESImageProperties struct {
    TextureID int
    ScaleX float32
    ScaleY float32
    ScaleX2 float32
    ScaleY2 float32
    Angle int
    Redf float32
    Greenf float32
    Bluef float32
    Alphaf float32
}


func (this *OpenGLESImageProperties) initFields() {
    this.TextureID = -1
    this.ScaleX = 1.0
    this.ScaleY = 1.0
    this.ScaleX2 = 1.0
    this.ScaleY2 = 1.0
    
    this.Redf = 1.0
    this.Greenf = 1.0
    this.Bluef = 1.0
    this.Alphaf = 1.0
}

func NewOpenGLESImageProperties() *OpenGLESImageProperties {
    this := &OpenGLESImageProperties{}
    this.initFields()
    return this
}

