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

type ActivityFractureUtilAPI30 struct {
    ActivityFractureUtilAPI5
}

var ActivityFractureUtilAPI30_Instance30 *ActivityFractureUtilAPI30 = NewActivityFractureUtilAPI30()

func NewActivityFractureUtilAPI30() *ActivityFractureUtilAPI30 {
    this := &ActivityFractureUtilAPI30{}
    this.ActivityFractureUtilAPI5 = *NewActivityFractureUtilAPI5()
    return this
}


func ActivityFractureUtilAPI30_GetInstance() *ActivityFractureUtilAPI30 {
        return ActivityFractureUtilAPI30_Instance30
}


func (this *ActivityFractureUtilAPI30) SetFullScreen(activity *Activity) {
        var window *Window = activity.GetWindow(); _ = window
        var windowInsetsController *WindowInsetsControllerCompat = WindowCompat_GetInsetsController(window, window.GetDecorView()); _ = windowInsetsController
        windowInsetsController.Hide(WindowInsetsCompat_Type.SystemBars())
        windowInsetsController.SetSystemBarsBehavior(WindowInsetsControllerCompat_BEHAVIOR_SHOW_TRANSIENT_BARS_BY_SWIPE)
}

