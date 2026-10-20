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

type StringValidationUtil struct {
    StringUtil *StringUtil
    TsUtil *TsUtil
}

var StringValidationUtil_Instance *StringValidationUtil = NewStringValidationUtil()

func (this *StringValidationUtil) initFields() {
    this.StringUtil = StringUtil_GetInstance()
    this.TsUtil = TsUtil_GetInstance()
}


func StringValidationUtil_GetInstance() *StringValidationUtil {
        return StringValidationUtil_Instance
}


func NewStringValidationUtil() *StringValidationUtil {
    this := &StringValidationUtil{}
    this.initFields()
    return this
}


func (this *StringValidationUtil) ContainsSpaces(value string) bool {
        if strings.Index(value, string(' ')) >= 0 {
            return true
        }
        return false
}


func (this *StringValidationUtil) IsNumber(value string) bool {
        var numberOfDecimalPoints int = 0; _ = numberOfDecimalPoints
        for index := int(0); index < len(value); index++ {
            var digit rune = rune(value[index]); _ = digit
            if !this.IsNumberFromChar(digit) && digit != '.' {
                return false
            }
            if digit == '.' {
                numberOfDecimalPoints++
                if numberOfDecimalPoints > 1 {
                    return false
                }
            }
        }
        return true
}


func (this *StringValidationUtil) IsNumberFromChar(digit rune) bool {
        if digit != '0' && digit != '1' && digit != '2' && digit != '3' && digit != '4' && digit != '5' && digit != '6' && digit != '7' && digit != '8' && digit != '9' {
            return false
        }
        return true
}


func (this *StringValidationUtil) IsValidRequired(value string, min_ int, max_ int) bool {
        if value == "" || len(value) < min_ || len(value) > max_ {
            return false
        }
        return true
}


func (this *StringValidationUtil) IsValidRequiredNumber(value string, min_ int, max_ int) bool {
        if this.IsEmpty(value) || len(value) < min_ || len(value) > max_ {
            return false
        }
        if !this.IsNumber(value) {
            return false
        }
        return true
}


func (this *StringValidationUtil) IsValidNotRequired(value string, min_ int, max_ int) bool {
        if value != "" {
            if len(value) < min_ || len(value) > max_ {
                return false
            }
        }
        return true
}


func (this *StringValidationUtil) IsValidNotRequiredNumber(value string, min_ int, max_ int) bool {
        if value != "" {
            if this.TsUtil.CompareTo(value, this.StringUtil.NULL_STRING) == 0 || len(value) < min_ || len(value) > max_ {
                return false
            }
            if !this.IsNumber(value) {
                return false
            }
        }
        return true
}


func (this *StringValidationUtil) IsEmpty(string_ string) bool {
        if string_ != "" && this.TsUtil.CompareTo(string_, this.StringUtil.NULL_STRING) != 0 && this.TsUtil.CompareTo(string_, this.StringUtil.EMPTY_STRING) != 0 {
            return false
        } else {
            return true
        }
}
