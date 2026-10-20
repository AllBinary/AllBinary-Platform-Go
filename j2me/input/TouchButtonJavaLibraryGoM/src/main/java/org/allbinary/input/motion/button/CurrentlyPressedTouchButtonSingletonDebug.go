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

type CurrentlyPressedTouchButtonSingletonDebug struct {
    CurrentlyPressedTouchButtonSingleton
    String string
    ListString string
    Append string
    PRESSED_AND_FIRED string
    PRESSED_AND_NOT_FIRED string
    ASSOCIATED_RELEASED_AND_FIRED string
    RELEASED_AND_FIRED string
    RELEASED_AND_FIRED_2 string
    RELEASED_AND_NOT_FIRED string
}

var CurrentlyPressedTouchButtonSingletonDebug_InstanceC *CurrentlyPressedTouchButtonSingletonDebug = NewCurrentlyPressedTouchButtonSingletonDebug()

func (this *CurrentlyPressedTouchButtonSingletonDebug) initFields() {
    this.String = StringUtil_GetInstance().EMPTY_STRING
    this.ListString = StringUtil_GetInstance().EMPTY_STRING
    this.Append = StringUtil_GetInstance().EMPTY_STRING
    this.PRESSED_AND_FIRED = "pressed & fired"
    this.PRESSED_AND_NOT_FIRED = "pressed in button but not fired"
    this.ASSOCIATED_RELEASED_AND_FIRED = "assoc. rel & fired"
    this.RELEASED_AND_FIRED = "rel & fired 1"
    this.RELEASED_AND_FIRED_2 = "rel & fired 2"
    this.RELEASED_AND_NOT_FIRED = "rel & not fired"
}

func NewCurrentlyPressedTouchButtonSingletonDebug() *CurrentlyPressedTouchButtonSingletonDebug {
    this := &CurrentlyPressedTouchButtonSingletonDebug{}
    this.CurrentlyPressedTouchButtonSingleton = *NewCurrentlyPressedTouchButtonSingleton()
    this.initFields()
    return this
}


func CurrentlyPressedTouchButtonSingletonDebug_GetInstance() ICurrentlyPressedTouchButtonSingleton {
        return CurrentlyPressedTouchButtonSingletonDebug_InstanceC
}


func (this *CurrentlyPressedTouchButtonSingletonDebug) AppendTouchButtonInput(append_ string, touchButtonInput *TouchButtonInput) {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(fmt.Sprint(touchButtonInput))
        stringBuffer.Append(CommonSeps_GetInstance().SPACE)
        stringBuffer.Append(append_)
        this.Append = fmt.Sprint(stringBuffer)
        this.String = fmt.Sprint(NewStringMaker().Append(this.ListString).Append(this.Append))
}


func (this *CurrentlyPressedTouchButtonSingletonDebug) Append(append_ string) {
        this.Append = fmt.Sprint(NewStringMaker().Append(this.Append).Append(append_))
        this.String = fmt.Sprint(NewStringMaker().Append(this.ListString).Append(this.Append))
}


func (this *CurrentlyPressedTouchButtonSingletonDebug) ClearLog(x int, y int) {
        this.Append = fmt.Sprint(NewStringMaker().Append(GPoint_ToStringStatic(x, y, 0)).Append(CommonSeps_GetInstance().SPACE))
        this.String = fmt.Sprint(NewStringMaker().Append(this.ListString).Append(this.Append))
}


func (this *CurrentlyPressedTouchButtonSingletonDebug) ReleaseAndNotFired() {
        this.Append(this.RELEASED_AND_NOT_FIRED)
}


func (this *CurrentlyPressedTouchButtonSingletonDebug) ReleaseAndFired(touchButtonInput *TouchButtonInput) {
        this.AppendTouchButtonInput(this.RELEASED_AND_FIRED, touchButtonInput)
}


func (this *CurrentlyPressedTouchButtonSingletonDebug) ReleaseAndFired2(touchButtonInput *TouchButtonInput) {
        this.AppendTouchButtonInput(this.RELEASED_AND_FIRED_2, touchButtonInput)
}


func (this *CurrentlyPressedTouchButtonSingletonDebug) ReleaseAndFiredAssociated(touchButtonInput *TouchButtonInput) {
        this.AppendTouchButtonInput(this.ASSOCIATED_RELEASED_AND_FIRED, touchButtonInput)
}


func (this *CurrentlyPressedTouchButtonSingletonDebug) PressedAndFired(touchButtonInput *TouchButtonInput) {
        this.AppendTouchButtonInput(this.PRESSED_AND_FIRED, touchButtonInput)
}


func (this *CurrentlyPressedTouchButtonSingletonDebug) PressedAndNotFired(touchButtonInput *TouchButtonInput) {
        this.AppendTouchButtonInput(this.PRESSED_AND_NOT_FIRED, touchButtonInput)
}


func (this *CurrentlyPressedTouchButtonSingletonDebug) RemoveAt(index int) *TouchButtonInput {
        var touchButtonInput *TouchButtonInput = this.CurrentlyPressedTouchButtonSingleton.RemoveAt(index); _ = touchButtonInput
        this.ListString = fmt.Sprint(this.List)
        this.String = fmt.Sprint(NewStringMaker().Append(this.ListString).Append(this.Append))
        return touchButtonInput
}


func (this *CurrentlyPressedTouchButtonSingletonDebug) Remove(touchButtonInput *TouchButtonInput) bool {
        var isRemoved bool = this.CurrentlyPressedTouchButtonSingleton.Remove(touchButtonInput); _ = isRemoved
        this.ListString = fmt.Sprint(this.List)
        this.String = fmt.Sprint(NewStringMaker().Append(this.ListString).Append(this.Append))
        return isRemoved
}


func (this *CurrentlyPressedTouchButtonSingletonDebug) Add(touchButtonInput *TouchButtonInput) {
        this.CurrentlyPressedTouchButtonSingleton.Add(touchButtonInput)
        this.ListString = fmt.Sprint(this.List)
        this.String = fmt.Sprint(NewStringMaker().Append(this.ListString).Append(this.Append))
}


func (this *CurrentlyPressedTouchButtonSingletonDebug) ToString() string {
        return this.String
}

func (this *CurrentlyPressedTouchButtonSingletonDebug) String() string {
    return this.ToString()
}

