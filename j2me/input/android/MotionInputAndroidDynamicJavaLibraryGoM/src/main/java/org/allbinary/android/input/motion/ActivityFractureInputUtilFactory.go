/*
 *
 * AllBinary Open License Version 1
 * Copyright (c) 2009 AllBinary
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

type ActivityFractureInputUtilFactory struct {
}


func NewActivityFractureInputUtilFactory() *ActivityFractureInputUtilFactory {
    this := &ActivityFractureInputUtilFactory{}
    return this
}


func ActivityFractureInputUtilFactory_GetInstance() IActivityFractureUtil {
        var SDK_VERSION int = AndroidInfoFactory_GetInstance().GetVersion(); _ = SDK_VERSION
        if SDK_VERSION <= 4 {
            return ActivityFractureUtilAPI1_GetInstance()
        } else if SDK_VERSION <= 15 {
            return ActivityFractureUtilAPI5_GetInstance()
        } else if SDK_VERSION <= 16 {
            return ActivityFractureUtilAPI16_GetInstance()
        } else {
            return ActivityFractureUtilAPI19_GetInstance()
        }
}

