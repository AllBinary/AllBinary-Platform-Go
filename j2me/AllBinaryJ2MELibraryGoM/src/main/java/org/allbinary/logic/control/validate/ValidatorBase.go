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

type ValidatorBase struct {
}

// Implements ValidatorBaseInterface

// The Java class ValidatorBase has subclasses so its variables are this interface
type IValidatorBase interface {
    ValidatorBase_() *ValidatorBase
    IsValid() bool
    ValidationInfo() string
    ToHashtable() *ABHashtable[any, any]
    ToList() IBasicArrayList
}

func (this *ValidatorBase) ValidatorBase_() *ValidatorBase {
    return this
}

var ValidatorBase_Instance IValidatorBase = NewValidatorBase()


func ValidatorBase_GetInstance() IValidatorBase {
        return ValidatorBase_Instance
}


func NewValidatorBase() *ValidatorBase {
    this := &ValidatorBase{}
    return this
}


func (this *ValidatorBase) IsValid() bool {
        return BooleanFactory_GetInstance().TRUE
}


func (this *ValidatorBase) ValidationInfo() string {
        return StringUtil_GetInstance().EMPTY_STRING
}


func (this *ValidatorBase) ToHashtable() *ABHashtable[any, any] {
        return StdUtil_GetInstance().NULL_TABLE
}


func (this *ValidatorBase) ToList() IBasicArrayList {
        return BasicArrayListUtil_GetInstance().GetImmutableInstance()
}

