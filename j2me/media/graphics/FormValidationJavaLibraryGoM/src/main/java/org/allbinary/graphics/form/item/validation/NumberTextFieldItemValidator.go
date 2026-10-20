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

type NumberTextFieldItemValidator struct {
    ValidatorBase
    TextFieldItem *ABTextFieldItem
    Min int
    Max int
    MaxChars int
}


func (this *NumberTextFieldItemValidator) initFields() {
    
    
    
    
}


func NewNumberTextFieldItemValidator(textFieldItem *ABTextFieldItem, maxChars int, min_ int, max_ int) *NumberTextFieldItemValidator {
    this := &NumberTextFieldItemValidator{}
    this.ValidatorBase = *NewValidatorBase()
    this.initFields()
        this.TextFieldItem = textFieldItem
        this.Min = min_
        this.Max = max_
        this.MaxChars = maxChars
    return this
}


func (this *NumberTextFieldItemValidator) IsNumberValid(value int) bool {
        return BooleanFactory_GetInstance().TRUE
}


func (this *NumberTextFieldItemValidator) ToNumberVector(value int) IBasicArrayList {
        return NewBasicArrayListD()
}


func (this *NumberTextFieldItemValidator) IsValid() bool {
        var result bool = BooleanFactory_GetInstance().TRUE; _ = result
        var string_ string = this.TextFieldItem.GetString(); _ = string_
        var textLength int = len(string_); _ = textLength
        if textLength > 0 && textLength < this.MaxChars {
            if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
                defer func() {
                    if recovered := recover(); recovered != nil {
                        e := Java_ToError(recovered)
                        _ = e
                result = BooleanFactory_GetInstance().FALSE
                    }
                }()
                func() int { value, _ := strconv.Atoi(this.TextFieldItem.GetString()); return value }()
                return
            }(); tryControl1 == 1 {
                return tryResult1
            }
        } else {
            if textLength < 1 {
                result = BooleanFactory_GetInstance().FALSE
            } else if textLength > this.MaxChars {
                result = BooleanFactory_GetInstance().FALSE
            }
        }
        return result
}


func (this *NumberTextFieldItemValidator) ToList() IBasicArrayList {
        var list IBasicArrayList = NewBasicArrayListD(); _ = list
        var string_ string = this.TextFieldItem.GetString(); _ = string_
        var textLength int = len(string_); _ = textLength
        var label string = this.TextFieldItem.GetLabel(); _ = label
        var name string = label[0:len(label) - 2]; _ = name
        var smallIntegerSingletonFactory *SmallIntegerSingletonFactory = SmallIntegerSingletonFactory_GetInstance(); _ = smallIntegerSingletonFactory
        var stringMaker *StringMaker = NewStringMaker(); _ = stringMaker
        if textLength > 0 && textLength < this.MaxChars {
            if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 IBasicArrayList) {
                defer func() {
                    if recovered := recover(); recovered != nil {
                        e := Java_ToError(recovered)
                        _ = e
                stringMaker.Delete(0, stringMaker.Length())
                list.Add(fmt.Sprint(stringMaker.Append(name).Append(" is not a number")))
                    }
                }()
                var number int = smallIntegerSingletonFactory.CreateInstance(func() int { value, _ := strconv.Atoi(this.TextFieldItem.GetString()); return value }()); _ = number
                if int(number) > this.Min {
                    stringMaker.Delete(0, stringMaker.Length())
                    list.Add(fmt.Sprint(stringMaker.Append(name).Append(" is to small")))
                } else if int(number) > this.Max {
                    stringMaker.Delete(0, stringMaker.Length())
                    list.Add(fmt.Sprint(stringMaker.Append(name).Append(" is to large")))
                }
                return
            }(); tryControl1 == 1 {
                return tryResult1
            }
        } else {
            if textLength < 1 {
                stringMaker.Delete(0, stringMaker.Length())
                list.Add(fmt.Sprint(stringMaker.Append(name).Append(" is to short")))
            } else if textLength > this.MaxChars {
                stringMaker.Delete(0, stringMaker.Length())
                list.Add(fmt.Sprint(stringMaker.Append(name).Append(" is to long")))
            }
        }
        return list
}

