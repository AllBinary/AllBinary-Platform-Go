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

type Time struct {
    TimeLong int64
}

// The Java class Time has subclasses so its variables are this interface
type ITime interface {
    Time_() *Time
    GetLong() int64
    Getlong() int64
    ToString() string
}

func (this *Time) Time_() *Time {
    return this
}


func (this *Time) initFields() {
    
}




// Java overloads of new Time
func NewTime(args ...any) *Time {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[string](args[0], pass) {
            return NewTime_1(Java_Arg[string](args[0]))
        }
        if len(args) == 1 && Java_IsArg[int64](args[0], pass) {
            return NewTime_2(Java_Arg[int64](args[0]))
        }
    }
    panic("No Java overload of new Time for the arguments")
}
func NewTime_1(time string) *Time {
    this := &Time{}
    this.initFields()
        this.TimeLong = int64(time)
    return this
}


func NewTime_2(time int64) *Time {
    this := &Time{}
    this.initFields()
        this.TimeLong = int64(time)
    return this
}


func (this *Time) GetLong() int64 {
        return this.TimeLong
}


func (this *Time) Getlong() int64 {
        return int64(this.TimeLong)
}


func (this *Time) ToString() string {
        return fmt.Sprint(this.GetLong())
}

func (this *Time) String() string {
    return this.ToString()
}

