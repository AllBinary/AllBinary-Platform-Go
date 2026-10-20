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

type AbBasicCrypt struct {
    Cipher *Cipher
    SecretKey *SecretKey
    Algorithm string
    Key []int8
}

// Implements CryptInterface


func (this *AbBasicCrypt) initFields() {
    
    
    
    
}


func NewAbBasicCrypt(algorithm string, key string) *AbBasicCrypt {
    this := &AbBasicCrypt{}
    this.initFields()
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            PreLogUtil_PutOE(commonStrings.EXCEPTION, this, "AbCrypt(alg,key)", e)
                }
            }()
            this.Algorithm = algorithm
            this.Key = key.GetBytes()
            this.Init()
            return
        }(); tryControl1 == 1 {
            return
        }
    return this
}


func (this *AbBasicCrypt) Init() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            PreLogUtil_PutOE("init Failed", this, commonStrings.INIT, e)
                }
            }()
            var sunJce *Provider = NewSunJCE(); _ = sunJce
            Security_AddProvider(sunJce)
            var keySpec *KeySpec = KeySpecFactory_GetInstance().GetInstance(this.Algorithm, this.Key); _ = keySpec
            var keyFactory *SecretKeyFactory = SecretKeyFactory_GetInstance(this.Algorithm); _ = keyFactory
            this.SecretKey = keyFactory.GenerateSecret(keySpec)
            this.Cipher = Cipher_GetInstance(this.Algorithm)
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *AbBasicCrypt) Encrypt(array []int8) []int8 {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 []int8) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            PreLogUtil_PutOE("Encrypt Failed", this, "encrypt", e)
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            this.Cipher.Init(Cipher_ENCRYPT_MODE, this.SecretKey)
            var ivArray []int8 = this.SecretKey.GetEncoded(); _ = ivArray
            var encrypted []int8 = this.Cipher.DoFinal(array); _ = encrypted
            var result []int8 = make([]int8, len(ivArray) + len(encrypted)); _ = result
            PreLogUtil_Put("ivArray Length: " + fmt.Sprint(len(ivArray)), this, "encrypt")
            for index := int(0); index < len(ivArray); index++ {
                result[index] = ivArray[index]
            }
            for index := int(0); index < len(encrypted); index++ {
                result[index + len(ivArray)] = encrypted[index]
            }
            tryControl1, tryResult1 = 1, result; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *AbBasicCrypt) Decrypt(array []int8) []int8 {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 []int8) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            PreLogUtil_PutOE("decrypt Failed", this, "decrypt", e)
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            this.Cipher.Init(Cipher_DECRYPT_MODE, this.SecretKey)
            var ivArray []int8 = make([]int8, 8); _ = ivArray
            for index := int(0); index < 8; index++ {
                ivArray[index] = array[index]
            }
            PreLogUtil_Put("ivArray Length: " + fmt.Sprint(len(ivArray)), this, "encrypt")
            var result []int8 = make([]int8, len(array) - len(ivArray)); _ = result
            for index := int(len(ivArray)); index < len(array); index++ {
                result[index - len(ivArray)] = array[index]
            }
            var decrypted []int8 = this.Cipher.DoFinal(result); _ = decrypted
            tryControl1, tryResult1 = 1, result; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

