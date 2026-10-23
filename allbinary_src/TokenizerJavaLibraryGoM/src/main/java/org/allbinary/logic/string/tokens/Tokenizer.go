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

type Tokenizer struct {
    Sep string
    EndSep string
}


func (this *Tokenizer) initFields() {
    
    this.EndSep = StringUtil_GetInstance().EMPTY_STRING
}


func NewTokenizer(sep string) *Tokenizer {
    this := &Tokenizer{}
    this.initFields()
        if sep == "" || TsUtil_GetInstance().CompareTo(sep, StringUtil_GetInstance().EMPTY_STRING) == 0 {
            panic(errors.New(fmt.Sprint("Sep provided is not valid")))
        }
        this.Sep = sep
    return this
}


func (this *Tokenizer) SetSep(sep string) {
        this.Sep = sep
}


func (this *Tokenizer) SetEndSep(endSep string) {
        this.EndSep = endSep
}


func (this *Tokenizer) GetInsideSep(string_ string) IBasicArrayList {
        var tokenList IBasicArrayList = NewBasicArrayListD(); _ = tokenList
        var index int = 0; _ = index
        var size int = len(string_); _ = size
        for index < size {
            index = string_.IndexOf(this.Sep, index)
            if index != -1 {
                var end int = string_.IndexOf(this.EndSep, index + len(this.Sep)); _ = end
                if end != -1 {
                    tokenList.Add(string_[index + len(this.Sep):end - (len(this.EndSep) - 1)])
                    index = end + len(this.EndSep)
                } else {
                }
            } else {
                break
            }
        }
        return tokenList
}


func (this *Tokenizer) GetTokensFromString(string_ string, tokenVector IBasicArrayList) IBasicArrayList {
        var index int = 0; _ = index
        var end int = 0; _ = end
        for index < len(string_) {
            end = string_.IndexOf(this.Sep, index)
            if end != -1 {
                tokenVector.Add(string_[index:end])
                index = end + len(this.Sep)
            } else {
                tokenVector.Add(string_[index:len(string_)])
                break
            }
        }
        return tokenVector
}


func (this *Tokenizer) GetTokens(stringVector IBasicArrayList) *ABHashtable[any, any] {
        var tokenHashtable *ABHashtable[any, any] = StdUtil_GetInstance().CreateHashtable(); _ = tokenHashtable
        var string_ string = StringUtil_GetInstance().EMPTY_STRING; _ = string_
        var end int = 0; _ = end
        var size int = stringVector.Size(); _ = size
        for index := int(0); index < size; index++ {
            string_ = any(stringVector.BasicArrayList_().ObjectArray[index]).(string)
            end = strings.Index(string_, string(this.Sep))
            if end >= 0 {
                tokenHashtable.Put(string_[0:end], string_[end + 1:len(string_)])
            }
        }
        return tokenHashtable
}

