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

type LongUtil struct {
}



func NewLongUtil() *LongUtil {
    this := &LongUtil{}
    return this
}


func LongUtil_FillIn(end string) string {
        var frameStringBuffer *StringMaker = NewStringMaker(); _ = frameStringBuffer
        if frameStringBuffer.Length() < LongData_MAX_LONG_LENGTH {
            var commonPhoneStrings *CommonPhoneStrings = CommonPhoneStrings_GetInstance(); _ = commonPhoneStrings
            var size int = LongData_MAX_LONG_LENGTH - frameStringBuffer.Length(); _ = size
            for index := int(0); index < size; index++ {
                frameStringBuffer.Append(commonPhoneStrings.ZERO)
            }
        }
        frameStringBuffer.Append(end)
        return fmt.Sprint(frameStringBuffer)
}

