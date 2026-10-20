/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2007 AllBinary 
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

type TouchButtonsListBuilder struct {
    BaseTouchInput
    BaseTouchInputList IBasicArrayList
}

// The Java class TouchButtonsListBuilder has subclasses so its variables are this interface
type ITouchButtonsListBuilder interface {
    IBaseTouchInput
    TouchButtonsListBuilder_() *TouchButtonsListBuilder
    AddList(list IBasicArrayList)
    Add(baseTouchInput IBaseTouchInput)
}

func (this *TouchButtonsListBuilder) TouchButtonsListBuilder_() *TouchButtonsListBuilder {
    return this
}


func (this *TouchButtonsListBuilder) initFields() {
    this.BaseTouchInputList = NewBasicArrayListD()
}

func NewTouchButtonsListBuilder() *TouchButtonsListBuilder {
    this := &TouchButtonsListBuilder{}
    this.BaseTouchInput = *NewBaseTouchInput()
    this.initFields()
    return this
}


func (this *TouchButtonsListBuilder) AddList(list IBasicArrayList) {
        var size int = this.BaseTouchInputList.Size(); _ = size
        var baseTouchInput IBaseTouchInput; _ = baseTouchInput
        for index := int(0); index < size; index++ {
            baseTouchInput = any(this.BaseTouchInputList.Get(index)).(IBaseTouchInput)
            list.AddAllList(baseTouchInput.GetList())
        }
}


func (this *TouchButtonsListBuilder) Add(baseTouchInput IBaseTouchInput) {
        this.BaseTouchInputList.Add(baseTouchInput)
}

