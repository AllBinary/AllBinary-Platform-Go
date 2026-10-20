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

type TouchScreenTypeFactory struct {
}

var TouchScreenTypeFactory_Instance *TouchScreenTypeFactory = NewTouchScreenTypeFactory()

func NewTouchScreenTypeFactory() *TouchScreenTypeFactory {
    this := &TouchScreenTypeFactory{}
    return this
}


func TouchScreenTypeFactory_GetInstance() *TouchScreenTypeFactory {
        return TouchScreenTypeFactory_Instance
}


func (this *TouchScreenTypeFactory) Update(activity *Activity) {
        var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
        var touchScreenFactory *TouchScreenFactory = TouchScreenFactory_GetInstance(); _ = touchScreenFactory
        var touchScreenTypesFactory *TouchScreenTypesFactory = TouchScreenTypesFactory_GetInstance(); _ = touchScreenTypesFactory
        var configuration *Configuration = activity.GetResources().GetConfiguration(); _ = configuration
        var touchScreen int = configuration.Touchscreen; _ = touchScreen
        if touchScreen == Configuration_TOUCHSCREEN_FINGER {
            touchScreenFactory.SetTouchScreenType(touchScreenTypesFactory.FINGER)
            if !touchScreenFactory.IsTouch() {
                PreLogUtil_Put("Not Really Exception: This indicates that a touch screen does exist but is not API8 so we will try it", this, commonStrings.PROCESS)
                touchScreenFactory.SetTouch(true)
            }
        } else if touchScreen == Configuration_TOUCHSCREEN_STYLUS {
            touchScreenFactory.SetTouchScreenType(touchScreenTypesFactory.STYLUS)
            if !touchScreenFactory.IsTouch() {
                PreLogUtil_Put("Not Really Exception: This indicates that a touch screen does exist but is not API8 so we will try it", this, commonStrings.PROCESS)
                touchScreenFactory.SetTouch(true)
            }
        } else if touchScreen == Configuration_TOUCHSCREEN_NOTOUCH {
            touchScreenFactory.SetTouchScreenType(touchScreenTypesFactory.NOTOUCH)
        } else if touchScreen == Configuration_TOUCHSCREEN_UNDEFINED {
            touchScreenFactory.SetTouchScreenType(touchScreenTypesFactory.UNDEFINED)
        }
}

