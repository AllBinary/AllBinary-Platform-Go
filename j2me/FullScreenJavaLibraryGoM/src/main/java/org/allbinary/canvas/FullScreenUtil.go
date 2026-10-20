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

type FullScreenUtil struct {
}

var FullScreenUtil_Instance *FullScreenUtil = NewFullScreenUtil()

func NewFullScreenUtil() *FullScreenUtil {
    this := &FullScreenUtil{}
    return this
}


func FullScreenUtil_GetInstance() *FullScreenUtil {
        return FullScreenUtil_Instance
}


// throws Exception
func (this *FullScreenUtil) InitOnRun(fullScreenInterface ICanvas, commandListener CommandListener) {
        (any(fullScreenInterface).(IRunnableCanvas)).WaitOnNotify(0)
        if commandListener != nil {
            this.Init(fullScreenInterface)
        }
}


func (this *FullScreenUtil) Init(fullScreenInterface ICanvas) {
        var mainFeatureFactory *MainFeatureFactory = MainFeatureFactory_GetInstance(); _ = mainFeatureFactory
        var features *Features = Features_GetInstance(); _ = features
        if features.IsFeature(mainFeatureFactory.FULL_SCREEN) {
            fullScreenInterface.SetFullScreenMode(true)
        } else {
            fullScreenInterface.SetFullScreenMode(false)
        }
}


func (this *FullScreenUtil) IsScreenChange(isFullScreen bool) bool {
        var mainFeatureFactory *MainFeatureFactory = MainFeatureFactory_GetInstance(); _ = mainFeatureFactory
        var features *Features = Features_GetInstance(); _ = features
        if features.IsFeature(mainFeatureFactory.FULL_SCREEN) != isFullScreen {
            return true
        }
        return false
}


func (this *FullScreenUtil) IsScreenChangeCanvas(fullScreenInterface ICanvas) bool {
        return this.IsScreenChange(fullScreenInterface.IsFullScreenMode())
}

