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

type OpenGLFeatureUtil struct {
    AnyThreed bool
}

var OpenGLFeatureUtil_Instance *OpenGLFeatureUtil = NewOpenGLFeatureUtil()

func (this *OpenGLFeatureUtil) initFields() {
    
}


func OpenGLFeatureUtil_GetInstance() *OpenGLFeatureUtil {
        return OpenGLFeatureUtil_Instance
}


func NewOpenGLFeatureUtil() *OpenGLFeatureUtil {
    this := &OpenGLFeatureUtil{}
    this.initFields()
        var openGLConfiguration *OpenGLConfiguration = OpenGLConfiguration_GetInstance(); _ = openGLConfiguration
        var features *Features = Features_GetInstance(); _ = features
        var openGLFeatureFactory *OpenGLFeatureFactory = OpenGLFeatureFactory_GetInstance(); _ = openGLFeatureFactory
        this.AnyThreed = openGLConfiguration.IsOpenGL() && (features.IsFeature(openGLFeatureFactory.OPENGL_2D_AND_3D) || features.IsFeature(openGLFeatureFactory.OPENGL_3D))
    return this
}


func (this *OpenGLFeatureUtil) IsAnyThreed() bool {
        return this.AnyThreed
}

