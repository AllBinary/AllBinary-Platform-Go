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

type ActivityFractureUtilAPI5 struct {
    ActivityFractureUtilAPI1
    FEATURE_TOUCHSCREEN_MULTITOUCH string
    FEATURE_TOUCHSCREEN string
    FEATURE_TOUCHSCREEN_MULTITOUCH_DISTINCT string
}

// The Java class ActivityFractureUtilAPI5 has subclasses so its variables are this interface
type IActivityFractureUtilAPI5 interface {
    IActivityFractureUtilAPI1
    ActivityFractureUtilAPI5_() *ActivityFractureUtilAPI5
    Process(activity *Activity)
    IsTouch(activity *Activity) bool
    IsMultiTouch(activity *Activity) bool
    IsMultiTouchDistinct(activity *Activity) bool
    IsFeature(activity *Activity, property string) bool
}

func (this *ActivityFractureUtilAPI5) ActivityFractureUtilAPI5_() *ActivityFractureUtilAPI5 {
    return this
}

var ActivityFractureUtilAPI5_Instance5 IActivityFractureUtilAPI5 = NewActivityFractureUtilAPI5()

func (this *ActivityFractureUtilAPI5) initFields() {
    this.FEATURE_TOUCHSCREEN_MULTITOUCH = "android.hardware.touchscreen.multitouch"
    this.FEATURE_TOUCHSCREEN = "android.hardware.touchscreen"
    this.FEATURE_TOUCHSCREEN_MULTITOUCH_DISTINCT = "android.hardware.touchscreen.multitouch.distinct"
}

func NewActivityFractureUtilAPI5() *ActivityFractureUtilAPI5 {
    this := &ActivityFractureUtilAPI5{}
    this.ActivityFractureUtilAPI1 = *NewActivityFractureUtilAPI1()
    this.initFields()
    return this
}


func ActivityFractureUtilAPI5_GetInstance() IActivityFractureUtilAPI5 {
        return ActivityFractureUtilAPI5_Instance5
}


func (this *ActivityFractureUtilAPI5) Process(activity *Activity) {
        var touchScreenFactory *TouchScreenFactory = TouchScreenFactory_GetInstance(); _ = touchScreenFactory
        var touchScreenTypeFactory *TouchScreenTypeFactory = TouchScreenTypeFactory_GetInstance(); _ = touchScreenTypeFactory
        touchScreenFactory.SetTouch(this.IsTouch(activity))
        touchScreenFactory.SetMultiTouch(this.IsMultiTouch(activity))
        touchScreenFactory.SetMultiTouchDistinct(this.IsMultiTouchDistinct(activity))
        touchScreenTypeFactory.Update(activity)
        var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
        PreLogUtil_Put(fmt.Sprint(touchScreenFactory), this, commonStrings.PROCESS)
}


func (this *ActivityFractureUtilAPI5) IsTouch(activity *Activity) bool {
        return this.IsFeature(activity, this.FEATURE_TOUCHSCREEN)
}


func (this *ActivityFractureUtilAPI5) IsMultiTouch(activity *Activity) bool {
        return this.IsFeature(activity, this.FEATURE_TOUCHSCREEN_MULTITOUCH)
}


func (this *ActivityFractureUtilAPI5) IsMultiTouchDistinct(activity *Activity) bool {
        return this.IsFeature(activity, this.FEATURE_TOUCHSCREEN_MULTITOUCH_DISTINCT)
}


func (this *ActivityFractureUtilAPI5) IsFeature(activity *Activity, property string) bool {
        return activity.GetPackageManager().HasSystemFeature(property)
}

