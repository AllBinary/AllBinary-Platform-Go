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

type AbCrypt struct {
    TsUtil *TsUtil
    SecretComposite *BaseSecretComposite
    Algorithm string
}

// Implements CryptInterface


func (this *AbCrypt) initFields() {
    this.TsUtil = TsUtil_GetInstance()
    this.SecretComposite = BaseSecretComposite_NULL_SECRET_COMPOSITE
    
}


func NewAbCrypt(algorithm string) *AbCrypt {
    this := &AbCrypt{}
    this.initFields()
        this.Algorithm = algorithm
    return this
}


func (this *AbCrypt) Init(keyAsString string) {
        var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            PreLogUtil_PutOE("init Failed", this, commonStrings.INIT, e)
                }
            }()
            if tryControl2 := func() (tryControl2 int) {
                defer func() {
                    if recovered := recover(); recovered != nil {
                        e := Java_ToError(recovered)
                        _ = e
                PreLogUtil_PutOE(commonStrings.EXCEPTION, this, commonStrings.INIT, e)
                    }
                }()
                Security_AddProvider(NewBouncyCastleProvider())
                return
            }(); tryControl2 == 1 {
                tryControl1 = 1; return
            }
            var key []int8 = this.TsUtil.GetByteArray(keyAsString); _ = key
            var keySpec *KeySpec = KeySpecFactory_GetInstance().GetInstance(this.Algorithm, key); _ = keySpec
            var keyFactory *SecretKeyFactory = SecretKeyFactory_GetInstance(this.Algorithm); _ = keyFactory
            var secretKey *SecretKey = keyFactory.GenerateSecret(keySpec); _ = secretKey
            var cipher *Cipher = Cipher_GetInstance(this.Algorithm); _ = cipher
            this.SecretComposite = NewSecretComposite(secretKey, cipher, key)
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *AbCrypt) Encrypt(array []int8) []int8 {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 []int8) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            PreLogUtil_PutOE("Encrypt Failed", this, "encrypt", e)
            tryControl1, tryResult1 = 1, NullUtil_GetInstance().NULL_BYTE_ARRAY; return
                }
            }()
            tryControl1, tryResult1 = 1, this.SecretComposite.Encrypt(array); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *AbCrypt) Decrypt(array []int8) []int8 {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 []int8) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            PreLogUtil_PutOE("decrypt Failed", this, "decrypt", e)
            tryControl1, tryResult1 = 1, NullUtil_GetInstance().NULL_BYTE_ARRAY; return
                }
            }()
            tryControl1, tryResult1 = 1, this.SecretComposite.Decrypt(array); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

