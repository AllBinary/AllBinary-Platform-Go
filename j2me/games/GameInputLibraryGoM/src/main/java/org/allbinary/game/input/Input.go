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

type Input struct {
    Id int
    Name string
}

// The Java class Input has subclasses so its variables are this interface
type IInput interface {
    Input_() *Input
    GetId() int
    SetId(key int)
    GetName() string
    SetName(keyName string)
    ToString() string
}

func (this *Input) Input_() *Input {
    return this
}


func (this *Input) initFields() {
    
    this.Name = StringUtil_GetInstance().EMPTY_STRING
}


func NewInput(id int, name string) *Input {
    this := &Input{}
    this.initFields()
        this.SetId(id)
        this.SetName(name)
    return this
}


func (this *Input) GetId() int {
        return this.Id
}


func (this *Input) SetId(key int) {
        this.Id = key
}


func (this *Input) GetName() string {
        return this.Name
}


func (this *Input) SetName(keyName string) {
        this.Name = keyName
}


func (this *Input) ToString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append("Name: ")
        stringBuffer.Append(this.GetName())
        stringBuffer.Append(" Id: ")
        stringBuffer.Appendint(this.GetId())
        return fmt.Sprint(stringBuffer)
}

func (this *Input) String() string {
    return this.ToString()
}

