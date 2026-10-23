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

type StringUtil struct {
    INIT_STRING string
    NULL_STRING string
    EMPTY_STRING string
    ONE_EMPTY_STRING_ARRAY []string
    StringArray []string
}

var StringUtil_Instance *StringUtil = NewStringUtil()

func (this *StringUtil) initFields() {
    this.INIT_STRING = ""
    this.NULL_STRING = "null"
    this.EMPTY_STRING = ""
    this.ONE_EMPTY_STRING_ARRAY = []string{this.EMPTY_STRING}
    this.StringArray = make([]string, 0)
}


func StringUtil_GetInstance() *StringUtil {
        return StringUtil_Instance
}


func NewStringUtil() *StringUtil {
    this := &StringUtil{}
    this.initFields()
    return this
}


func (this *StringUtil) GetNonNull(string_ string) string {
        if string_ == "" {
            return this.EMPTY_STRING
        } else {
            return string_
        }
}


func (this *StringUtil) GetArrayInstance() []string {
        return this.StringArray
}


func (this *StringUtil) Count(string_ string, aChar rune) int {
        var count int = 0; _ = count
        var size int = len(string_); _ = size
        for index := int(0); index < size; index++ {
            if rune(string_[index]) == aChar {
                count++
            }
        }
        return count
}


func (this *StringUtil) ToString(anyType any) string {
        if anyType != nil {
            return fmt.Sprint(anyType)
        } else {
            return this.NULL_STRING
        }
}

