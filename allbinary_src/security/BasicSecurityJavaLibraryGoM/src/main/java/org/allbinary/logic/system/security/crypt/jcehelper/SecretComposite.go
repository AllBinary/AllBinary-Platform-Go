/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2025 AllBinary 
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

type SecretComposite struct {
    BaseSecretComposite
    ByteUtil *ByteUtil
    Cipher *Cipher
    SecretKey *SecretKey
    Key []int8
}


func (this *SecretComposite) initFields() {
    this.ByteUtil = ByteUtil_GetInstance()
    
    
    
}


func NewSecretComposite(secretKey *SecretKey, cipher *Cipher, key []int8) *SecretComposite {
    this := &SecretComposite{}
    this.BaseSecretComposite = *NewBaseSecretComposite()
    this.initFields()
        this.SecretKey = secretKey
        this.Cipher = cipher
        this.Key = key
    return this
}


// throws Exception
func (this *SecretComposite) Encrypt(array []int8) []int8 {
        array = this.Mutilate(array)
        this.Cipher.Init(Cipher_ENCRYPT_MODE, this.SecretKey)
        return this.Cipher.DoFinal(array)
}


// throws Exception
func (this *SecretComposite) Decrypt(array []int8) []int8 {
        this.Cipher.Init(Cipher_DECRYPT_MODE, this.SecretKey)
        return this.Mutilate(this.Cipher.DoFinal(array))
}


func (this *SecretComposite) Mutilate(array []int8) []int8 {
        for index := int(0); index < len(this.Key); index++ {
            var value int8 = this.Key[index]; _ = value
            if int(value) < 8 && int(value) > 0 {
                array = this.ByteUtil.XorByte(array, int(value))
            }
        }
        return array
}

