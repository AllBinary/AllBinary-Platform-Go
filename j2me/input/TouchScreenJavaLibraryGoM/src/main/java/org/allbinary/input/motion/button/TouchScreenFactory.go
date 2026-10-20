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
    LogUtil *LogUtil
    Touch bool
    MultiTouch bool
    MultiTouchDistinct bool
    TouchScreenType *TouchScreenType
}

var TouchScreenFactory_SINGLETON *TouchScreenFactory = NewTouchScreenFactory()

func (this *TouchScreenFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
    
    this.TouchScreenType = TouchScreenTypesFactory_GetInstance().NOTOUCH
}


func NewTouchScreenFactory() *TouchScreenFactory {
    this := &TouchScreenFactory{}
    this.initFields()
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.Put(commonStrings.EXCEPTION, TouchScreenFactory_SINGLETON, commonStrings.GET_INSTANCE, e)
                }
            }()
            this.SetTouch(true)
            this.SetMultiTouch(false)
            return
        }(); tryControl1 == 1 {
            return
        }
    return this
}


func TouchScreenFactory_GetInstance() *TouchScreenFactory {
        return TouchScreenFactory_SINGLETON
}


func (this *TouchScreenFactory) SetMultiTouch(multiTouch bool) {
        this.MultiTouch = multiTouch
}


func (this *TouchScreenFactory) IsMultiTouch() bool {
        return this.MultiTouch
}


func (this *TouchScreenFactory) SetTouch(touch bool) {
        this.Touch = touch
}


func (this *TouchScreenFactory) IsTouch() bool {
        return this.Touch
}


func (this *TouchScreenFactory) SetMultiTouchDistinct(multiTouchDistinct bool) {
        this.MultiTouchDistinct = multiTouchDistinct
}


func (this *TouchScreenFactory) IsMultiTouchDistinct() bool {
        return this.MultiTouchDistinct
}


func (this *TouchScreenFactory) SetTouchScreenType(touchScreenType *TouchScreenType) {
        this.TouchScreenType = touchScreenType
}


func (this *TouchScreenFactory) GetTouchScreenType() *TouchScreenType {
        return this.TouchScreenType
}


func (this *TouchScreenFactory) ToString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append("TouchScreen: ")
        stringBuffer.Appendboolean(this.IsTouch())
        stringBuffer.Append(" MultiTouch: ")
        stringBuffer.Appendboolean(this.IsMultiTouch())
        stringBuffer.Append(" Distinct MultiTouch: ")
        stringBuffer.Appendboolean(this.IsMultiTouchDistinct())
        stringBuffer.Append(" TouchScreenType: ")
        stringBuffer.Append(StringUtil_GetInstance().ToString(this.GetTouchScreenType()))
        return fmt.Sprint(stringBuffer)
}

func (this *TouchScreenFactory) String() string {
    return this.ToString()
}

