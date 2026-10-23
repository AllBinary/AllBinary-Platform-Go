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

type DatabaseEncoder struct {
}



func NewDatabaseEncoder() *DatabaseEncoder {
    this := &DatabaseEncoder{}
    return this
}


func DatabaseEncoder_Encode(value []int8) string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            PreLogUtil_PutOE(commonStrings.EXCEPTION, "DatabaseEncoder", "decode", e)
            tryControl1, tryResult1 = 1, StringUtil_GetInstance().EMPTY_STRING; return
                }
            }()
            var array []int8 = value; _ = array
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            for index := int(0); index < len(array); index++ {
                stringBuffer.Append(fmt.Sprint(int8(array[index])))
                if index < len(array) - 1 {
                    stringBuffer.Append(CommonSeps_GetInstance().SPACE)
                }
            }
            tryControl1, tryResult1 = 1, fmt.Sprint(stringBuffer); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func DatabaseEncoder_Decode(value string) []int8 {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 []int8) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            PreLogUtil_PutSE(commonStrings.EXCEPTION, "DatabaseEncoder", "decode", e)
            tryControl1, tryResult1 = 1, NullUtil_GetInstance().NULL_BYTE_ARRAY; return
                }
            }()
            var tokenizer *Tokenizer = NewTokenizer(CommonSeps_GetInstance().SPACE); _ = tokenizer
            var vector IBasicArrayList = tokenizer.GetTokensFromString(value, NewBasicArrayListD()); _ = vector
            var byteVector IBasicArrayList = NewBasicArrayListD(); _ = byteVector
            var size int = vector.Size(); _ = size
            var byteOfData string; _ = byteOfData
            for index := int(0); index < size; index++ {
                byteOfData = any(vector.BasicArrayList_().ObjectArray[index]).(string)
                byteVector.Add(int8(byteOfData))
            }
            var decode []int8 = make([]int8, byteVector.Size()); _ = decode
            var decodeIndex int = 0; _ = decodeIndex
            var aByte int8; _ = aByte
            for index := int(0); index < size; index++ {
                aByte = any(byteVector.BasicArrayList_().ObjectArray[index]).(int8)
                decode[decodeIndex] = int8(aByte)
                decodeIndex++
            }
            tryControl1, tryResult1 = 1, decode; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

