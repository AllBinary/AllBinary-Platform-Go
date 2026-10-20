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

type ShakeAnimationListener struct {
}

// The Java class ShakeAnimationListener has subclasses so its variables are this interface
type IShakeAnimationListener interface {
    ShakeAnimationListener_() *ShakeAnimationListener
    OnSmallShakeEvent()
    OnMediumShakeEvent()
    OnLargeShakeEvent()
    OnShakeEvent(shakeEvent *ShakeEvent)
    GetType() int
}

func (this *ShakeAnimationListener) ShakeAnimationListener_() *ShakeAnimationListener {
    return this
}


func NewShakeAnimationListener() *ShakeAnimationListener {
    this := &ShakeAnimationListener{}
    return this
}


func (this *ShakeAnimationListener) OnSmallShakeEvent() {
}


func (this *ShakeAnimationListener) OnMediumShakeEvent() {
}


func (this *ShakeAnimationListener) OnLargeShakeEvent() {
}


func (this *ShakeAnimationListener) OnShakeEvent(shakeEvent *ShakeEvent) {
}


func (this *ShakeAnimationListener) GetType() int {
        return 0
}

