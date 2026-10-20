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

type BasicCrypt struct {
    ByteUtil *ByteUtil
    Key []int8
}

// Implements CryptInterface


func (this *BasicCrypt) initFields() {
    this.ByteUtil = ByteUtil_GetInstance()
    
}


func NewBasicCrypt(keyAsString string) *BasicCrypt {
    this := &BasicCrypt{}
    this.initFields()
        var key []int8 = NullUtil_GetInstance().NULL_BYTE_ARRAY; _ = key
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            PreLogUtil_PutOE(commonStrings.EXCEPTION, this, "AbCrypt(alg,key)", e)
                }
            }()
            key = TsUtil_GetInstance().GetByteArray(keyAsString)
            return
        }(); tryControl1 == 1 {
            return
        }
        this.Key = key
    return this
}


func (this *BasicCrypt) Encrypt(array []int8) []int8 {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 []int8) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            PreLogUtil_PutOE("Encrypt Failed", this, "encrypt", e)
            tryControl1, tryResult1 = 1, NullUtil_GetInstance().NULL_BYTE_ARRAY; return
                }
            }()
            tryControl1, tryResult1 = 1, this.Mutilate(array); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *BasicCrypt) Decrypt(array []int8) []int8 {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 []int8) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            PreLogUtil_PutOE("decrypt Failed", this, "decrypt", e)
            tryControl1, tryResult1 = 1, NullUtil_GetInstance().NULL_BYTE_ARRAY; return
                }
            }()
            tryControl1, tryResult1 = 1, this.Mutilate(array); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *BasicCrypt) Mutilate(array []int8) []int8 {
        var value int8; _ = value
        for index := int(0); index < len(this.Key); index++ {
            value = this.Key[index]
            array = this.ByteUtil.Xor(array, value)
        }
        return array
}

