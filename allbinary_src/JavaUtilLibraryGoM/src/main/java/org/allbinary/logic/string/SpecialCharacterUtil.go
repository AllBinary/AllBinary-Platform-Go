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

type SpecialCharacterUtil struct {
}


func NewSpecialCharacterUtil() *SpecialCharacterUtil {
    this := &SpecialCharacterUtil{}
    return this
}


func SpecialCharacterUtil_GetHashMap() *HashMap {
        var commonSeps *CommonSeps = CommonSeps_GetInstance(); _ = commonSeps
        var commonPhoneStrings *CommonPhoneStrings = CommonPhoneStrings_GetInstance(); _ = commonPhoneStrings
        var stringUtil *StringUtil = StringUtil_GetInstance(); _ = stringUtil
        var hashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = hashMap
        hashMap.Put("!", stringUtil.EMPTY_STRING)
        hashMap.Put("@", stringUtil.EMPTY_STRING)
        hashMap.Put(commonPhoneStrings.POUND, stringUtil.EMPTY_STRING)
        hashMap.Put("$", stringUtil.EMPTY_STRING)
        hashMap.Put("%", stringUtil.EMPTY_STRING)
        hashMap.Put("^", stringUtil.EMPTY_STRING)
        hashMap.Put(commonSeps.AMPERSAND, stringUtil.EMPTY_STRING)
        hashMap.Put(commonPhoneStrings.STAR, stringUtil.EMPTY_STRING)
        hashMap.Put(commonSeps.PARENTHESIS_OPEN, stringUtil.EMPTY_STRING)
        hashMap.Put(commonSeps.PARENTHESIS_CLOSE, stringUtil.EMPTY_STRING)
        hashMap.Put(commonSeps.DASH, stringUtil.EMPTY_STRING)
        hashMap.Put(commonSeps.UNDERSCORE, stringUtil.EMPTY_STRING)
        hashMap.Put(MathData_GetInstance().PLUS, stringUtil.EMPTY_STRING)
        hashMap.Put(commonSeps.EQUALS, stringUtil.EMPTY_STRING)
        hashMap.Put("\\", stringUtil.EMPTY_STRING)
        hashMap.Put("|", stringUtil.EMPTY_STRING)
        hashMap.Put(AbPathData_GetInstance().EXTENSION_SEP, stringUtil.EMPTY_STRING)
        hashMap.Put(",", stringUtil.EMPTY_STRING)
        hashMap.Put("<", stringUtil.EMPTY_STRING)
        hashMap.Put(">", stringUtil.EMPTY_STRING)
        hashMap.Put("?", stringUtil.EMPTY_STRING)
        hashMap.Put(AbPathData_GetInstance().SEPARATOR, stringUtil.EMPTY_STRING)
        hashMap.Put("~", stringUtil.EMPTY_STRING)
        hashMap.Put("`", stringUtil.EMPTY_STRING)
        return hashMap
}

