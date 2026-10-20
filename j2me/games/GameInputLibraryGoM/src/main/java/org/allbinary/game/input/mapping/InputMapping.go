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

type InputMapping struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    Hashtable *ABHashtable[any, any]
    MappedList IBasicArrayList
}

// The Java class InputMapping has subclasses so its variables are this interface
type IInputMapping interface {
    InputMapping_() *InputMapping
    GetTotalMapped() int
    IsDefaultNew() bool
    GetDefault() *InputToGameKeyMapping
    RemoveAll()
    Remove(input IInput, mappedToInput IInput)
    AddMapping(inputToGameKeyMapping *InputToGameKeyMapping)
    Add(input IInput, mappedToInput IInput)
    IsMapped(input IInput) bool
    GetMappedInput(id IInput) IBasicArrayList
    GetHashtable() *ABHashtable[any, any]
}

func (this *InputMapping) InputMapping_() *InputMapping {
    return this
}


func (this *InputMapping) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.Hashtable = StdUtil_GetInstance().CreateHashtable()
    this.MappedList = NewBasicArrayListD()
}


func NewInputMapping() *InputMapping {
    this := &InputMapping{}
    this.initFields()
    return this
}


func (this *InputMapping) GetTotalMapped() int {
        return this.Hashtable.Size()
}


func (this *InputMapping) IsDefaultNew() bool {
        return false
}


func (this *InputMapping) GetDefault() *InputToGameKeyMapping {
        return InputToGameKeyMapping_GetNullInstance()
}


func (this *InputMapping) RemoveAll() {
        this.Hashtable.Clear()
        this.MappedList.Clear()
}


func (this *InputMapping) Remove(input IInput, mappedToInput IInput) {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        var stringUtil *StringUtil = StringUtil_GetInstance(); _ = stringUtil
        stringBuffer.Append(CommonLabels_GetInstance().START_LABEL)
        stringBuffer.Append(stringUtil.ToString(input))
        stringBuffer.Append(" == ")
        stringBuffer.Append(stringUtil.ToString(mappedToInput))
        this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "InputMapping::remove")
        var list IBasicArrayList = this.GetMappedInput(input); _ = list
        if list.Contains(mappedToInput) {
            list.Remove(mappedToInput)
            this.MappedList.Remove(mappedToInput)
            stringBuffer.Delete(0, stringBuffer.Length())
            stringBuffer.Append("Removed: ")
            stringBuffer.Append(stringUtil.ToString(input))
            stringBuffer.Append(" == ")
            stringBuffer.Append(stringUtil.ToString(mappedToInput))
            this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "InputMapping::remove")
        } else {
            stringBuffer.Delete(0, stringBuffer.Length())
            stringBuffer.Append("Not Deleted: ")
            stringBuffer.Append(stringUtil.ToString(input))
            stringBuffer.Append(" == ")
            stringBuffer.Append(stringUtil.ToString(mappedToInput))
            this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "InputMapping::remove")
        }
}


// throws Exception
func (this *InputMapping) AddMapping(inputToGameKeyMapping *InputToGameKeyMapping) {
        panic(errors.New(fmt.Sprint(this.CommonStrings.NOT_IMPLEMENTED)))
}


func (this *InputMapping) Add(input IInput, mappedToInput IInput) {
        var list IBasicArrayList = this.GetMappedInput(input); _ = list
        if !list.Contains(mappedToInput) {
            list.Add(mappedToInput)
            this.MappedList.Add(mappedToInput)
        }
}


func (this *InputMapping) IsMapped(input IInput) bool {
        return this.MappedList.Contains(input)
}


func (this *InputMapping) GetMappedInput(id IInput) IBasicArrayList {
        var mappingInputCanBeNullList any = this.Hashtable.Get(id); _ = mappingInputCanBeNullList
        if mappingInputCanBeNullList == nil {
            var list IBasicArrayList = NewBasicArrayListD(); _ = list
            this.Hashtable.Put(id, list)
            return list
        }
        return any(mappingInputCanBeNullList).(IBasicArrayList)
}


func (this *InputMapping) GetHashtable() *ABHashtable[any, any] {
        return this.Hashtable
}

