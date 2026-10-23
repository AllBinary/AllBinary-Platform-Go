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

type TimeLastModified struct {
    Time
}





// Java overloads of new TimeLastModified
func NewTimeLastModified(args ...any) *TimeLastModified {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[string](args[0], pass) {
            return NewTimeLastModified_1(Java_Arg[string](args[0]))
        }
        if len(args) == 1 && Java_IsArg[int64](args[0], pass) {
            return NewTimeLastModified_2(Java_Arg[int64](args[0]))
        }
    }
    panic("No Java overload of new TimeLastModified for the arguments")
}
func NewTimeLastModified_1(time string) *TimeLastModified {
    this := &TimeLastModified{}
        this.Time = *NewTime(time)
    return this
}


func NewTimeLastModified_2(time int64) *TimeLastModified {
    this := &TimeLastModified{}
        this.Time = *NewTime(time)
    return this
}

