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

type WeakCrypt struct {
    TsUtil *TsUtil
    AbCrypt *AbCrypt
}

var WeakCrypt_Keys []string = []string{"qwres6qewrupoi1132790jfahleavi89398", "143uk32lngf8675gt1tv9214uvt924ut9u9", "12kvjlvlkv11445v12951v512v9m51295v1", "gampipoiupg4g43y5u238tu2485u20u0932", "a93u5uu02s2d2g3259u4395u409u093u9vu", "nlkxzj90je3mln9l8l89qjf0j3v50j23r9j", "2lk5fn5j23nr8f0233vq35n23jn58vhy843", "9q24334m598t092tkn2iv0982456n425j9g", "5m9uvmt9436243m5v34g09d43i52ki35g99", "n123lr1iubiu2br32lv32rbu23b32ub2rlb", "va0u0iovji2ojpvj2oij22poiu4ov2uv28v", "bkk80gn2v9uc902340ut209pcup0u09utc9", "mn32m32neee23c42fn32nf3mf2n3fn2m3fm", "zcxpozo3423voizupovsijevnv3n33o33n3", "v0345324thv93h03893893vh893v8h3vhzz", "v8na898v8wyef9ty39yv98yaw39yeeyf8ye", "45zd36d34jg09434hfshp8uj9sy8s7cyr78", "b098yy0908vdvh983b2lb327980a76s5e5e", "v90y9v9p8y3r98pagsdgsagyr6y3985yhfh", "vblkzhhzp8yvzi8u4t9856456zy3rv8a388", "da7696023rgf28dssfa98asd879hg32hewe", "afo7d6ffkjb3jr3i34gb24ytuvu3jr3uryf", "4gijvqo8vq984hv984hj5j3511h832hf833"}
var WeakCrypt_KEYMAX int = len(WeakCrypt_Keys)

func (this *WeakCrypt) initFields() {
    this.TsUtil = TsUtil_GetInstance()
    this.AbCrypt = NewAbCrypt(KeySpecFactory_GetInstance().DES)
}


func NewWeakCrypt(key int) *WeakCrypt {
    this := &WeakCrypt{}
    this.initFields()
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            PreLogUtil_PutOE(commonStrings.EXCEPTION, this, "SuperCrypt(key)", e)
                }
            }()
            this.AbCrypt.Init(WeakCrypt_Keys[key])
            return
        }(); tryControl1 == 1 {
            return
        }
    return this
}


func (this *WeakCrypt) Encrypt(value string) string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            PreLogUtil_PutOE("Encrypt Failed", this, "encrypt", e)
            tryControl1, tryResult1 = 1, StringUtil_GetInstance().EMPTY_STRING; return
                }
            }()
            var crypted []int8 = this.AbCrypt.Encrypt(this.TsUtil.GetByteArray(value)); _ = crypted
            tryControl1, tryResult1 = 1, DatabaseEncoder_Encode(crypted); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *WeakCrypt) Decrypt(value string) string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            PreLogUtil_PutOE("decrypt Failed", this, "decrypt", e)
            tryControl1, tryResult1 = 1, StringUtil_GetInstance().EMPTY_STRING; return
                }
            }()
            var decrypted []int8 = this.AbCrypt.Decrypt(DatabaseEncoder_Decode(value)); _ = decrypted
            tryControl1, tryResult1 = 1, string(decrypted); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

