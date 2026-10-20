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

type OpenGLImageSpecificFactory struct {
    ImageFactory IOpenGLImageFactory
    Updating bool
}

var OpenGLImageSpecificFactory_Instance *OpenGLImageSpecificFactory = NewOpenGLImageSpecificFactory()

func (this *OpenGLImageSpecificFactory) initFields() {
    this.ImageFactory = OpenGLImageFactory_NULL_OPENGL_IMAGE_FACTORY
    this.Updating = false
}

func NewOpenGLImageSpecificFactory() *OpenGLImageSpecificFactory {
    this := &OpenGLImageSpecificFactory{}
    this.initFields()
    return this
}


func OpenGLImageSpecificFactory_GetInstance() *OpenGLImageSpecificFactory {
        return OpenGLImageSpecificFactory_Instance
}


func (this *OpenGLImageSpecificFactory) SetImageFactory(imageFactory IOpenGLImageFactory) {
        this.ImageFactory = imageFactory
        PreLogUtil_Put(fmt.Sprint(NewStringMaker().Append("Selected Image Factory: ").Append(StringUtil_GetInstance().ToString(this.ImageFactory))), this, "setImageFactory")
}


func (this *OpenGLImageSpecificFactory) GetImageFactory() IOpenGLImageFactory {
        return this.ImageFactory
}

