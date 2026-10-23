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

type Money struct {
    Units int64
    DEFAULT_CURRENCY string
}


func (this *Money) initFields() {
    
    this.DEFAULT_CURRENCY = "USD"
}




// Java overloads of new Money
func NewMoney(args ...any) *Money {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[*Money](args[0], pass) {
            return NewMoney_1(Java_Arg[*Money](args[0]))
        }
        if len(args) == 0 {
            return NewMoney_2()
        }
        if len(args) == 1 && Java_IsArg[string](args[0], pass) {
            return NewMoney_3(Java_Arg[string](args[0]))
        }
        if len(args) == 1 && Java_IsArg[int64](args[0], pass) {
            return NewMoney_4(Java_Arg[int64](args[0]))
        }
    }
    panic("No Java overload of new Money for the arguments")
}
func NewMoney_1(money *Money) *Money {
    this := &Money{}
    this.initFields()
        this.Units = int64(money.GetUnits())
    return this
}


func NewMoney_2() *Money {
    this := &Money{}
    this.initFields()
        this.Units = int64(0)
    return this
}


func NewMoney_3(usDollarStr string) *Money {
    this := &Money{}
    this.initFields()
        if usDollarStr != "" && StringValidationUtil_GetInstance().IsNumber(usDollarStr) {
            this.Units = int64(this.ConvertUsdToUnits(usDollarStr))
        }
    return this
}


func NewMoney_4(units int64) *Money {
    this := &Money{}
    this.initFields()
        this.Units = int64(units)
    return this
}


func (this *Money) IsValid() bool {
        if this.Units != 0 && int64(this.Units) > 0 && int64(this.Units) < Long_MAX_VALUE {
            return true
        } else {
            return false
        }
}


func (this *Money) Add_1(moreMoney *Money) {
        this.Add(fmt.Sprint(moreMoney))
}

// Java overloads of Money.add
func (this *Money) Add(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[*Money](args[0], pass) {
            this.Add_1(Java_Arg[*Money](args[0]))
            return
        }
        if len(args) == 1 && Java_IsArg[string](args[0], pass) {
            this.Add_2(Java_Arg[string](args[0]))
            return
        }
    }
    panic("No Java overload of Money.add for the arguments")
}


func (this *Money) Add_2(usDollarStr string) {
        this.Units = int64(int64(this.Units) + this.ConvertUsdToUnits(usDollarStr))
}


func (this *Money) Subtract(usDollarStr string) {
        this.Units = int64(int64(this.Units) - this.ConvertUsdToUnits(usDollarStr))
}


func (this *Money) Multiply_1(multiplier float32) {
        this.Units = int64(int64(float32(float32(int64(this.Units)) * float32(multiplier))))
}

// Java overloads of Money.multiply
func (this *Money) Multiply(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[float32](args[0], pass) {
            this.Multiply_1(Java_Arg[float32](args[0]))
            return
        }
        if len(args) == 1 && Java_IsArg[int](args[0], pass) {
            this.Multiply_2(Java_Arg[int](args[0]))
            return
        }
    }
    panic("No Java overload of Money.multiply for the arguments")
}


func (this *Money) Multiply_2(multiplier int) {
        this.Units = int64(int64(this.Units) * int64(multiplier))
}


func (this *Money) ConvertUnitsTo(currency string) string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        if strings.Compare(currency, this.DEFAULT_CURRENCY) == 0 {
            var dollar int64 = int64(this.Units) / 100; _ = dollar
            var cents int64 = int64(this.Units) - (dollar * 100); _ = cents
            stringBuffer.Append(fmt.Sprint(int64(dollar)))
            stringBuffer.Append(AbPathData_GetInstance().EXTENSION_SEP)
            if cents < 10 {
                stringBuffer.Append(CommonPhoneStrings_GetInstance().ZERO)
            }
            stringBuffer.Append(fmt.Sprint(int64(cents)))
        }
        return fmt.Sprint(stringBuffer)
}


func (this *Money) ConvertUsdToUnits(usDollarStr string) int64 {
        var localUnit int64 = 0; _ = localUnit
        var index int = strings.Index(usDollarStr, string(AbPathData_GetInstance().EXTENSION_SEP)); _ = index
        var stringUtil *StringUtil = StringUtil_GetInstance(); _ = stringUtil
        var dollarStr string = stringUtil.EMPTY_STRING; _ = dollarStr
        var centsStr string = stringUtil.EMPTY_STRING; _ = centsStr
        if index < 1 {
            dollarStr = usDollarStr
            centsStr = "00"
        } else {
            dollarStr = usDollarStr[0:index]
            centsStr = usDollarStr[index + 1:]
        }
        localUnit = int64(int64(int64(int64(dollarStr)) * 100))
        localUnit += int64(int64(centsStr))
        return localUnit
}


func (this *Money) GetUnits() int64 {
        return int64(this.Units)
}


func (this *Money) GetUnitsLong() int64 {
        return this.Units
}


func (this *Money) ToString() string {
        return this.ConvertUnitsTo(this.DEFAULT_CURRENCY)
}

func (this *Money) String() string {
    return this.ToString()
}

