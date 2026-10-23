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

type CamelCaseUtil struct {
    CommonSeps *CommonSeps
}

var CamelCaseUtil_Instance *CamelCaseUtil = NewCamelCaseUtil()

func (this *CamelCaseUtil) initFields() {
    this.CommonSeps = CommonSeps_GetInstance()
}

func NewCamelCaseUtil() *CamelCaseUtil {
    this := &CamelCaseUtil{}
    this.initFields()
    return this
}


func CamelCaseUtil_GetInstance() *CamelCaseUtil {
        return CamelCaseUtil_Instance
}


// throws Exception
func (this *CamelCaseUtil) GetAsCamelCase(string_ string, stringBuilder *StringMaker) string {
        if string_ == "" {
            return StringUtil_GetInstance().EMPTY_STRING
        }
        var stringValidationUtil *StringValidationUtil = StringValidationUtil_GetInstance(); _ = stringValidationUtil
        stringBuilder.Delete(0, stringBuilder.Length())
        var list IBasicArrayList = NewBasicArrayListD(); _ = list
        var tokenizer *Tokenizer = NewTokenizer(this.CommonSeps.UNDERSCORE); _ = tokenizer
        tokenizer.GetTokensFromString(string_, list)
        var size int = list.Size(); _ = size
        var word string; _ = word
        for i := int(0); i < size; i++ {
            word = any(list.Get(i)).(string)
            word = func() string { if stringValidationUtil.IsEmpty(word) { return word }; return fmt.Sprint(NewStringMaker().Appendchar(unicode.ToUpper(rune(word[0]))).Append(strings.ToLower(word[1:]))) }()
            stringBuilder.Append(word)
        }
        return fmt.Sprint(stringBuilder)
}

