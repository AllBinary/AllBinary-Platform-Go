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

type TouchFeatureFactory struct {
    TOUCH_ENABLED *TouchFeature
    SHOW_SCREEN_BUTTONS *TouchFeature
    AUTO_HIDE_SHOW_SCREEN_BUTTONS *TouchFeature
    HIDE_SCREEN_BUTTONS *TouchFeature
}

var TouchFeatureFactory_Instance *TouchFeatureFactory = NewTouchFeatureFactory()

func (this *TouchFeatureFactory) initFields() {
    this.TOUCH_ENABLED = NewTouchFeature("Touch Enabled")
    this.SHOW_SCREEN_BUTTONS = NewTouchFeature("Show")
    this.AUTO_HIDE_SHOW_SCREEN_BUTTONS = NewTouchFeature("Auto Hide (One Level)")
    this.HIDE_SCREEN_BUTTONS = NewTouchFeature("Hide")
}


func TouchFeatureFactory_GetInstance() *TouchFeatureFactory {
        return TouchFeatureFactory_Instance
}


func NewTouchFeatureFactory() *TouchFeatureFactory {
    this := &TouchFeatureFactory{}
    this.initFields()
    return this
}

