/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2022 AllBinary 
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

type OpenGLESGraphicsCompositeFactory struct {
    OpenGLESGraphicsFactory IOpenGLESGraphicsFactory
}

var OpenGLESGraphicsCompositeFactory_Instance *OpenGLESGraphicsCompositeFactory = NewOpenGLESGraphicsCompositeFactory()

func (this *OpenGLESGraphicsCompositeFactory) initFields() {
    this.OpenGLESGraphicsFactory = NewOpenGLESGraphicsFactory()
}

func NewOpenGLESGraphicsCompositeFactory() *OpenGLESGraphicsCompositeFactory {
    this := &OpenGLESGraphicsCompositeFactory{}
    this.initFields()
    return this
}


func OpenGLESGraphicsCompositeFactory_GetInstance() *OpenGLESGraphicsCompositeFactory {
        return OpenGLESGraphicsCompositeFactory_Instance
}


func (this *OpenGLESGraphicsCompositeFactory) Set(openGLESGraphicsFactory IOpenGLESGraphicsFactory) {
        this.OpenGLESGraphicsFactory = openGLESGraphicsFactory
}


func (this *OpenGLESGraphicsCompositeFactory) Get() IOpenGLESGraphicsFactory {
        return this.OpenGLESGraphicsFactory
}

