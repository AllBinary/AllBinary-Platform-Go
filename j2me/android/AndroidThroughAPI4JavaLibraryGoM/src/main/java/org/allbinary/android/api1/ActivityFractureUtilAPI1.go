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

type ActivityFractureUtilAPI1 struct {
    ActivityFractureUtil
}

// The Java class ActivityFractureUtilAPI1 has subclasses so its variables are this interface
type IActivityFractureUtilAPI1 interface {
    IActivityFractureUtil
    ActivityFractureUtilAPI1_() *ActivityFractureUtilAPI1
    Process(activity *Activity)
    SetFullScreen(activity *Activity)
}

func (this *ActivityFractureUtilAPI1) ActivityFractureUtilAPI1_() *ActivityFractureUtilAPI1 {
    return this
}

var ActivityFractureUtilAPI1_Instance IActivityFractureUtilAPI1 = NewActivityFractureUtilAPI1()

func NewActivityFractureUtilAPI1() *ActivityFractureUtilAPI1 {
    this := &ActivityFractureUtilAPI1{}
    this.ActivityFractureUtil = *NewActivityFractureUtil()
    return this
}


func ActivityFractureUtilAPI1_GetInstance() IActivityFractureUtilAPI1 {
        return ActivityFractureUtilAPI1_Instance
}


// throws Exception
func (this *ActivityFractureUtilAPI1) Process(activity *Activity) {
}


func (this *ActivityFractureUtilAPI1) SetFullScreen(activity *Activity) {
        activity.GetWindow().SetFlags(WindowManager_LayoutParams.FLAG_FULLSCREEN, WindowManager_LayoutParams.FLAG_FULLSCREEN)
}

