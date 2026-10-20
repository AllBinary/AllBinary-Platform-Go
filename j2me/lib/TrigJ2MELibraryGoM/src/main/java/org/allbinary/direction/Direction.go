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

type Direction struct {
    Name string
    OtherName string
    Value int
    FrameFactor int
}

var Direction_DirectionArray []*Direction = make([]*Direction, 13)

func (this *Direction) initFields() {
    
    
    
    
}


func NewDirection(name string, otherName string, value int, frameFactor int) *Direction {
    this := &Direction{}
    this.initFields()
        this.Name = name
        this.OtherName = otherName
        this.Value = value
        this.FrameFactor = frameFactor
        Direction_DirectionArray[value] = this
    return this
}


func Direction_GetInstance(direction int) *Direction {
        return Direction_DirectionArray[direction]
}


func (this *Direction) GetValue() int {
        return this.Value
}


func (this *Direction) GetFrameFactor() int {
        return this.FrameFactor
}


func (this *Direction) GetOtherName() string {
        return this.OtherName
}


func (this *Direction) GetName() string {
        return this.Name
}


func (this *Direction) ToString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append("Direction: ")
        stringBuffer.Append(this.GetName())
        stringBuffer.Append(" Or: ")
        stringBuffer.Append(this.GetOtherName())
        stringBuffer.Append(" Value: ")
        stringBuffer.Appendint(this.Value)
        return fmt.Sprint(stringBuffer)
}

func (this *Direction) String() string {
    return this.ToString()
}

