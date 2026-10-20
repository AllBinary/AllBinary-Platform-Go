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

type VirtualKeyboardAPI3 struct {
    VirtualKeyboard
    VirtualKeyboard bool
    Activity *Activity
    ShowVirtualKeyboardRunnable Runnable
    HideVirtualKeyboardRunnable Runnable
}

// Implements VirtualKeyboardEventListenerInterface

var VirtualKeyboardAPI3_MAX_API int = Integer_MAX_VALUE
var VirtualKeyboardAPI3_MIN_API int = 3

func (this *VirtualKeyboardAPI3) initFields() {
    
    
    
    
}


func NewVirtualKeyboardAPI3(activity *Activity) *VirtualKeyboardAPI3 {
    this := &VirtualKeyboardAPI3{}
        this.VirtualKeyboard = *NewVirtualKeyboard()
        this.initFields()
        this.Activity = activity
        this.ShowVirtualKeyboardRunnable = NewShowVirtualKeyboardRunnable(this.Activity)
        this.HideVirtualKeyboardRunnable = NewHideVirtualKeyboardRunnable(this.Activity)
        var virtualKeyboardEventHandler *VirtualKeyboardEventHandler = VirtualKeyboardEventHandler_GetInstance(); _ = virtualKeyboardEventHandler
        virtualKeyboardEventHandler.RemoveAllListeners()
        virtualKeyboardEventHandler.AddListenerInterface(this)
    return this
}


func (this *VirtualKeyboardAPI3) OnEvent(eventObject IAllBinaryEventObject) {
        ForcedLogUtil_Log(EventStrings_GetInstance().PERFORMANCE_MESSAGE, this)
}


func (this *VirtualKeyboardAPI3) OnVirtualKeyboardEvent(virtualKeyboardEvent *VirtualKeyboardEvent) {
        var isShow bool = any(virtualKeyboardEvent.GetSource()).(bool); _ = isShow
        if isShow {
            this.ShowVirtualKeyboard()
        } else {
            this.HideVirtualKeyboard()
        }
}


func (this *VirtualKeyboardAPI3) ForceHide() {
        this.HideVirtualKeyboard()
}


func (this *VirtualKeyboardAPI3) Hide() {
        if this.VirtualKeyboard {
            this.ForceHide()
        }
}


func (this *VirtualKeyboardAPI3) ShowVirtualKeyboard() {
        this.Activity.RunOnUiThread(this.ShowVirtualKeyboardRunnable)
        this.VirtualKeyboard = true
}


func (this *VirtualKeyboardAPI3) HideVirtualKeyboard() {
        this.Activity.RunOnUiThread(this.HideVirtualKeyboardRunnable)
        this.VirtualKeyboard = false
}

