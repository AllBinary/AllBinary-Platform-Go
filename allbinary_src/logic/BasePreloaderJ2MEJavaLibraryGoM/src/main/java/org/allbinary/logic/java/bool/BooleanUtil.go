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

type BooleanUtil struct {
    TsUtil *TsUtil
}

var BooleanUtil_Instance *BooleanUtil = NewBooleanUtil()

func (this *BooleanUtil) initFields() {
    this.TsUtil = TsUtil_GetInstance()
}


func BooleanUtil_GetInstance() *BooleanUtil {
        return BooleanUtil_Instance
}


func NewBooleanUtil() *BooleanUtil {
    this := &BooleanUtil{}
    this.initFields()
    return this
}


// throws Exception
func (this *BooleanUtil) GetFromString(booleanString string) bool {
        var stringValidationUtil *StringValidationUtil = StringValidationUtil_GetInstance(); _ = stringValidationUtil
        var booleanFactory *BooleanFactory = BooleanFactory_GetInstance(); _ = booleanFactory
        if stringValidationUtil.IsEmpty(booleanString) {
            return false
        }
        var booleanStringLowerCase string = strings.ToLower(booleanString); _ = booleanStringLowerCase
        if this.TsUtil.CompareTo(booleanStringLowerCase, booleanFactory.TRUE_STRING) == 0 {
            return true
        } else if this.TsUtil.CompareTo(booleanStringLowerCase, booleanFactory.FALSE_STRING) == 0 {
            return false
        } else {
            panic(errors.New(fmt.Sprint("Invalid booleanString - Must be true or false")))
        }
}


// throws Exception
func (this *BooleanUtil) IsStringBoolean(booleanString string) bool {
        var booleanFactory *BooleanFactory = BooleanFactory_GetInstance(); _ = booleanFactory
        var booleanStringLowerCase string = strings.ToLower(booleanString); _ = booleanStringLowerCase
        if strings.Compare(booleanStringLowerCase, booleanFactory.TRUE_STRING) == 0 || strings.Compare(booleanStringLowerCase, booleanFactory.FALSE_STRING) == 0 {
            return true
        }
        return false
}


func (this *BooleanUtil) ToStringFromBooleanArray(booleanArray []bool) string {
        var stringMaker *StringMaker = NewStringMaker(); _ = stringMaker
        var size int = len(booleanArray); _ = size
        for index := int(0); index < size; index++ {
            stringMaker.Appendboolean(booleanArray[index])
        }
        return fmt.Sprint(stringMaker)
}

