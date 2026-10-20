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

type ActivityFractureUtilAPI16 struct {
    ActivityFractureUtilAPI5
}

var ActivityFractureUtilAPI16_Instance16 *ActivityFractureUtilAPI16 = NewActivityFractureUtilAPI16()

func NewActivityFractureUtilAPI16() *ActivityFractureUtilAPI16 {
    this := &ActivityFractureUtilAPI16{}
    this.ActivityFractureUtilAPI5 = *NewActivityFractureUtilAPI5()
    return this
}


func ActivityFractureUtilAPI16_GetInstance() *ActivityFractureUtilAPI16 {
        return ActivityFractureUtilAPI16_Instance16
}


func (this *ActivityFractureUtilAPI16) SetFullScreen(activity *Activity) {
        var decorView *View = activity.GetWindow().GetDecorView(); _ = decorView
        var visibility int = View_SYSTEM_UI_FLAG_HIDE_NAVIGATION | View_SYSTEM_UI_FLAG_FULLSCREEN; _ = visibility
        decorView.SetSystemUiVisibility(visibility)
}

