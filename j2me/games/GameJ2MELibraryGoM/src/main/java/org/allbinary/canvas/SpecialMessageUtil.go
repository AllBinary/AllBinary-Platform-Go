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

type SpecialMessageUtil struct {
}

var SpecialMessageUtil_SINGLETON *SpecialMessageUtil = NewSpecialMessageUtil()

func NewSpecialMessageUtil() *SpecialMessageUtil {
    this := &SpecialMessageUtil{}
    return this
}


func SpecialMessageUtil_GetInstance() *SpecialMessageUtil {
        return SpecialMessageUtil_SINGLETON
}


func (this *SpecialMessageUtil) Get() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(Memory_GetInfo())
        stringBuffer.Append(CommonSeps_GetInstance().SPACE)
        stringBuffer.Append(fmt.Sprint(GameStatisticsFactory_GetInstance()))
        stringBuffer.Append(CommonSeps_GetInstance().SPACE)
        stringBuffer.Append(" Option: ")
        if Features_GetInstance().IsDefault(OpenGLFeatureFactory_GetInstance().OPENGL_AUTO_SELECT) {
            stringBuffer.Append(OpenGLFeatureFactory_GetInstance().OPENGL_AUTO_SELECT.GetName())
        } else {
            stringBuffer.Append(OpenGLFeatureFactory_GetInstance().OPENGL_MINIMUM.GetName())
        }
        return fmt.Sprint(stringBuffer)
}

