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

type CurrentlyPressedTouchButtonSingleton struct {
    List IBasicArrayList
}

// The Java class CurrentlyPressedTouchButtonSingleton has subclasses so its variables are this interface
type ICurrentlyPressedTouchButtonSingleton interface {
    CurrentlyPressedTouchButtonSingleton_() *CurrentlyPressedTouchButtonSingleton
    Size() int
    RemoveAt(index int) *TouchButtonInput
    Remove(touchButtonInput *TouchButtonInput) bool
    Get(index int) *TouchButtonInput
    Contains(touchButtonInput *TouchButtonInput) bool
    Add(touchButtonInput *TouchButtonInput)
}

func (this *CurrentlyPressedTouchButtonSingleton) CurrentlyPressedTouchButtonSingleton_() *CurrentlyPressedTouchButtonSingleton {
    return this
}

var CurrentlyPressedTouchButtonSingleton_Instance ICurrentlyPressedTouchButtonSingleton = NewCurrentlyPressedTouchButtonSingleton()

func (this *CurrentlyPressedTouchButtonSingleton) initFields() {
    this.List = NewBasicArrayListD()
}


func CurrentlyPressedTouchButtonSingleton_GetInstance() ICurrentlyPressedTouchButtonSingleton {
        return CurrentlyPressedTouchButtonSingleton_Instance
}


func NewCurrentlyPressedTouchButtonSingleton() *CurrentlyPressedTouchButtonSingleton {
    this := &CurrentlyPressedTouchButtonSingleton{}
    this.initFields()
    return this
}


func (this *CurrentlyPressedTouchButtonSingleton) Size() int {
        return this.List.Size()
}


func (this *CurrentlyPressedTouchButtonSingleton) RemoveAt(index int) *TouchButtonInput {
        return any(this.List.RemoveAt(index)).(*TouchButtonInput)
}


func (this *CurrentlyPressedTouchButtonSingleton) Remove(touchButtonInput *TouchButtonInput) bool {
        return this.List.Remove(touchButtonInput)
}


func (this *CurrentlyPressedTouchButtonSingleton) Get(index int) *TouchButtonInput {
        return any(this.List.BasicArrayList_().ObjectArray[index]).(*TouchButtonInput)
}


func (this *CurrentlyPressedTouchButtonSingleton) Contains(touchButtonInput *TouchButtonInput) bool {
        return this.List.Contains(touchButtonInput)
}


func (this *CurrentlyPressedTouchButtonSingleton) Add(touchButtonInput *TouchButtonInput) {
        this.List.Add(touchButtonInput)
}

