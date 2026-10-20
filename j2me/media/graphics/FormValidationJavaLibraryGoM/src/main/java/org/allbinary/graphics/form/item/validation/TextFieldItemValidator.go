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

type TextFieldItemValidator struct {
    ValidatorBase
    TextFieldItem *ABTextFieldItem
    Min int
    Max int
    AllowOnEmpty bool
}

// The Java class TextFieldItemValidator has subclasses so its variables are this interface
type ITextFieldItemValidator interface {
    IValidatorBase
    TextFieldItemValidator_() *TextFieldItemValidator
    IsValid() bool
    ToList() IBasicArrayList
}

func (this *TextFieldItemValidator) TextFieldItemValidator_() *TextFieldItemValidator {
    return this
}


func (this *TextFieldItemValidator) initFields() {
    
    
    
    
}


func NewTextFieldItemValidator(textFieldItem *ABTextFieldItem, min_ int, max_ int, allowOnEmpty bool) *TextFieldItemValidator {
    this := &TextFieldItemValidator{}
    this.ValidatorBase = *NewValidatorBase()
    this.initFields()
        this.TextFieldItem = textFieldItem
        this.Min = min_
        this.Max = max_
        this.AllowOnEmpty = allowOnEmpty
    return this
}


func (this *TextFieldItemValidator) IsValid() bool {
        var booleanFactory *BooleanFactory = BooleanFactory_GetInstance(); _ = booleanFactory
        var result bool = booleanFactory.TRUE; _ = result
        var string_ string = this.TextFieldItem.GetString(); _ = string_
        var textLength int = len(string_); _ = textLength
        if (textLength == 0 && this.AllowOnEmpty) || textLength > this.Min && textLength < this.Max {
        } else {
            if textLength < this.Min {
                result = booleanFactory.FALSE
            } else if textLength > this.Max {
                result = booleanFactory.FALSE
            }
        }
        return result
}


func (this *TextFieldItemValidator) ToList() IBasicArrayList {
        var list IBasicArrayList = NewBasicArrayListD(); _ = list
        var string_ string = this.TextFieldItem.GetString(); _ = string_
        var textLength int = len(string_); _ = textLength
        if textLength > this.Min && textLength < this.Max {
        } else {
            var label string = this.TextFieldItem.GetLabel(); _ = label
            var name string = label[0:len(label) - 2]; _ = name
            if textLength < this.Min {
                list.Add(fmt.Sprint(NewStringMaker().Append(name).Append(" is to short")))
            } else if textLength > this.Max {
                list.Add(fmt.Sprint(NewStringMaker().Append(name).Append(" is to long")))
            }
        }
        return list
}

