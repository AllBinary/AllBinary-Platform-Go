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

type PrimitiveLongUtil struct {
    LogUtil *LogUtil
    PowerOfTen int
    MaxValue int
    MaxDigits int
    CharArray []rune
    CurrentTotalDigits int
    PrimitiveLongSingleton *PrimitiveLongSingleton
}


func (this *PrimitiveLongUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
    
    
    
    this.PrimitiveLongSingleton = PrimitiveLongSingleton_GetInstance()
}


func PrimitiveLongUtil_CreatePowerOfTen(powerOfTen int) *PrimitiveLongUtil {
        var primitiveLongUtil *PrimitiveLongUtil = NewPrimitiveLongUtil(powerOfTen, NullUtil_GetInstance().NULL_OBJECT); _ = primitiveLongUtil
        if powerOfTen % 10 != 0 {
            var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            logUtil.Put(commonStrings.EXCEPTION, primitiveLongUtil, commonStrings.CONSTRUCTOR, errors.New(fmt.Sprint("Max must be power of 10")))
        }
        return primitiveLongUtil
}


func NewPrimitiveLongUtil(powerOfTen int, unused any) *PrimitiveLongUtil {
    this := &PrimitiveLongUtil{}
    this.initFields()
        this.MaxDigits = MathUtil_GetInstance().GetTotalDigits(powerOfTen)
        this.PowerOfTen = powerOfTen
        this.MaxValue = (powerOfTen * 10) - 1
        this.CharArray = make([]rune, this.MaxDigits)
    return this
}


func (this *PrimitiveLongUtil) GetCharArray(value int) []rune {
        if value < 10 {
            this.SetCurrentTotalDigits(1)
            return this.PrimitiveLongSingleton.NUMBER_CHAR_ARRAYS[value]
        }
        if value > this.MaxValue {
            this.SetCurrentTotalDigits(3)
            return this.PrimitiveLongSingleton.UNK
        }
        var index int = 0; _ = index
        var div int = this.PowerOfTen; _ = div
        var lastValue int = 0; _ = lastValue
        for div >= 1 {
            var digit int = (value - lastValue) / div; _ = digit
            if digit >= 0 && digit < 10 {
                if digit == 0 && index == 0 {
                } else {
                    this.CharArray[Java_PostIncrement(&index)] = this.PrimitiveLongSingleton.NUMBER_CHAR_ARRAY[digit]
                }
                lastValue += digit * div
            }
            div = div / 10
        }
        this.SetCurrentTotalDigits(index)
        return this.CharArray
}


func (this *PrimitiveLongUtil) SetCurrentTotalDigits(currentTotalDigits int) {
        this.CurrentTotalDigits = currentTotalDigits
}


func (this *PrimitiveLongUtil) GetCurrentTotalDigits() int {
        return this.CurrentTotalDigits
}


func (this *PrimitiveLongUtil) GetMaxDigits() int {
        return this.MaxDigits
}
