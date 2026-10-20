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

type TouchScreenFactory struct {
    MultiTouch bool
    Touch bool
    TouchScreenType *TouchScreenType
}

var TouchScreenFactory_SINGLETON *TouchScreenFactory = NewTouchScreenFactory()

func (this *TouchScreenFactory) initFields() {
    
    
    this.TouchScreenType = TouchScreenTypesFactory_GetInstance().NOTOUCH
}


func NewTouchScreenFactory() *TouchScreenFactory {
    this := &TouchScreenFactory{}
    this.initFields()
    return this
}


func TouchScreenFactory_GetInstance() *TouchScreenFactory {
        return TouchScreenFactory_SINGLETON
}


func (this *TouchScreenFactory) IsMultiTouch() bool {
        return this.MultiTouch
}


func (this *TouchScreenFactory) SetTouch(touch bool) {
}


func (this *TouchScreenFactory) IsTouch() bool {
        return this.Touch
}


func (this *TouchScreenFactory) GetTouchScreenType() *TouchScreenType {
        return this.TouchScreenType
}

