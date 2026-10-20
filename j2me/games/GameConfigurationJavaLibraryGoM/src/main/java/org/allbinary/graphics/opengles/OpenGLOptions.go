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

type OpenGLOptions struct {
}


func NewOpenGLOptions() *OpenGLOptions {
    this := &OpenGLOptions{}
    return this
}


func (this *OpenGLOptions) Init() {
        var openGLFeatureFactory *OpenGLFeatureFactory = OpenGLFeatureFactory_GetInstance(); _ = openGLFeatureFactory
        if Features_GetInstance().IsFeature(openGLFeatureFactory.OPENGL_OPTIONS) {
            var openGLMultipleList IBasicArrayList = NewBasicArrayListD(); _ = openGLMultipleList
            openGLMultipleList.Add(openGLFeatureFactory.OPENGL)
            GameFeatureChoiceGroups_GetMultipleInstance().Add("OpenGL (Complete Restart)", openGLMultipleList)
            var openGLVersionSelectorMultipleList IBasicArrayList = NewBasicArrayListD(); _ = openGLVersionSelectorMultipleList
            openGLVersionSelectorMultipleList.Add(openGLFeatureFactory.OPENGL_AUTO_SELECT)
            openGLVersionSelectorMultipleList.Add(openGLFeatureFactory.OPENGL_MINIMUM)
            GameFeatureChoiceGroups_GetExclusiveInstance().Add("OpenGL Version Selector", openGLVersionSelectorMultipleList)
            var openGLImageColorMultipleList IBasicArrayList = NewBasicArrayListD(); _ = openGLImageColorMultipleList
            openGLImageColorMultipleList.Add(openGLFeatureFactory.IMAGE_COLOR_DEPTH_4444)
            openGLImageColorMultipleList.Add(openGLFeatureFactory.IMAGE_COLOR_DEPTH_8888)
            GameFeatureChoiceGroups_GetExclusiveInstance().Add("OpenGL Image Color RGBA", openGLImageColorMultipleList)
        }
}

