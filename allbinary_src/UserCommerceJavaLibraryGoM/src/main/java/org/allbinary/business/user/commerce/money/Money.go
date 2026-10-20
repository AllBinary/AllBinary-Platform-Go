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


func NewMoney(money *Money) *Money {
    this := &Money{}
    this.initFields()
        this.Units = int64(money.GetUnits())
    return this
}


// Java overloaded constructor: Go needs a different name
func NewMoney() *Money {
    this := &Money{}
    this.initFields()
        this.Units = int64(0)
    return this
}


// Java overloaded constructor: Go needs a different name
func NewMoney(usDollarStr string) *Money {
    this := &Money{}
    this.initFields()
        if usDollarStr != "" && StringValidationUtil_GetInstance().IsNumber(usDollarStr) {
            this.Units = int64(this.ConvertUsdToUnits(usDollarStr))
        }
    return this
}


// Java overloaded constructor: Go needs a different name
func NewMoney(units int64) *Money {
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


func (this *Money) Add(moreMoney *Money) {
        this.Add(fmt.Sprint(moreMoney))
}


// Java overloaded method: Go needs a different name
func (this *Money) Add(usDollarStr string) {
        this.Units = int64(int64(this.Units) + this.ConvertUsdToUnits(usDollarStr))
}


func (this *Money) Subtract(usDollarStr string) {
        this.Units = int64(int64(this.Units) - this.ConvertUsdToUnits(usDollarStr))
}


func (this *Money) Multiply(multiplier float32) {
        this.Units = int64(int64(float32(float32(int64(this.Units)) * float32(multiplier))))
}


// Java overloaded method: Go needs a different name
func (this *Money) Multiply(multiplier int) {
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

