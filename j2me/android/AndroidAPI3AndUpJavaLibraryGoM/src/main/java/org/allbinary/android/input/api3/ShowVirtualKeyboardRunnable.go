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

type ShowVirtualKeyboardRunnable struct {
    LogUtil *LogUtil
    Activity *Activity
}

// Implements Runnable


func (this *ShowVirtualKeyboardRunnable) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
}


func NewShowVirtualKeyboardRunnable(activity *Activity) *ShowVirtualKeyboardRunnable {
    this := &ShowVirtualKeyboardRunnable{}
    this.initFields()
        this.Activity = activity
    return this
}


func (this *ShowVirtualKeyboardRunnable) Run() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.Put(commonStrings.EXCEPTION, this, commonStrings.RUN, e)
                }
            }()
            var inputMethodManager *InputMethodManager = any(this.Activity.GetSystemService(Context_INPUT_METHOD_SERVICE)).(*InputMethodManager); _ = inputMethodManager
            var viewCompositeInterface ViewCompositeInterface = any(this.Activity).(ViewCompositeInterface); _ = viewCompositeInterface
            if viewCompositeInterface == nil {
                ForcedLogUtil_Log("Activity Null", this)
            }
            var view *View = viewCompositeInterface.GetView(); _ = view
            inputMethodManager.ToggleSoftInputFromWindow(view.GetWindowToken(), InputMethodManager_SHOW_FORCED, 0)
            return
        }(); tryControl1 == 1 {
            return
        }
}

