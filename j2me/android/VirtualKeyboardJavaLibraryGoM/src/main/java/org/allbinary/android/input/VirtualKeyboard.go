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

type VirtualKeyboard struct {
}

// The Java class VirtualKeyboard has subclasses so its variables are this interface
type IVirtualKeyboard interface {
    VirtualKeyboard_() *VirtualKeyboard
    OnEvent(eventObject IAllBinaryEventObject)
    OnVirtualKeyboardEvent(virtualKeyboardEvent *VirtualKeyboardEvent)
    ForceHide()
    Hide()
}

func (this *VirtualKeyboard) VirtualKeyboard_() *VirtualKeyboard {
    return this
}

var VirtualKeyboard_NULL_VIRTUAL_KEYBOARD IVirtualKeyboard = NewVirtualKeyboard()

func NewVirtualKeyboard() *VirtualKeyboard {
    this := &VirtualKeyboard{}
    return this
}


func (this *VirtualKeyboard) OnEvent(eventObject IAllBinaryEventObject) {
        ForcedLogUtil_Log(EventStrings_GetInstance().PERFORMANCE_MESSAGE, this)
}


func (this *VirtualKeyboard) OnVirtualKeyboardEvent(virtualKeyboardEvent *VirtualKeyboardEvent) {
}


func (this *VirtualKeyboard) ForceHide() {
}


func (this *VirtualKeyboard) Hide() {
}

