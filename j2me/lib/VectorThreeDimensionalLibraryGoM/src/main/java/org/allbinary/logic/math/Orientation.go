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

type Orientation struct {
    Yaw *BasicDecimal
    Pitch *BasicDecimal
    Roll *BasicDecimal
}

var Orientation_NONE *Orientation = NewOrientation(0, 0, 0)
var Orientation_SINGLETON *Orientation = NewOrientation(0, 0, 0)

func (this *Orientation) initFields() {
    this.Yaw = NewBasicDecimal(0)
    this.Pitch = NewBasicDecimal(0)
    this.Roll = NewBasicDecimal(0)
}


func Orientation_GetInstance(yaw int, pitch int, roll int) *Orientation {
        Orientation_SINGLETON.Yaw.Setint(yaw)
        Orientation_SINGLETON.Pitch.Setint(pitch)
        Orientation_SINGLETON.Roll.Setint(roll)
        return Orientation_SINGLETON
}


func NewOrientation(yaw int, pitch int, roll int) *Orientation {
    this := &Orientation{}
    this.initFields()
        this.Yaw.Setint(yaw)
        this.Pitch.Setint(pitch)
        this.Roll.Setint(roll)
    return this
}


func (this *Orientation) ToString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append("Y: ")
        stringBuffer.Append(fmt.Sprint(this.Yaw))
        stringBuffer.Append(" P: ")
        stringBuffer.Append(fmt.Sprint(this.Pitch))
        stringBuffer.Append(" R: ")
        stringBuffer.Append(fmt.Sprint(this.Roll))
        return fmt.Sprint(stringBuffer)
}

func (this *Orientation) String() string {
    return this.ToString()
}

