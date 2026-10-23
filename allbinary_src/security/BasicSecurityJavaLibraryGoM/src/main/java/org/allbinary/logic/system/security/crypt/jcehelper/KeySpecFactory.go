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

type KeySpecFactory struct {
    DES string
    DESEDE string
    BLOWFISH string
}

var KeySpecFactory_Instance *KeySpecFactory = NewKeySpecFactory()

func (this *KeySpecFactory) initFields() {
    this.DES = "DES"
    this.DESEDE = "DESede"
    this.BLOWFISH = "Blowfish"
}


func KeySpecFactory_GetInstance() *KeySpecFactory {
        return KeySpecFactory_Instance
}


func NewKeySpecFactory() *KeySpecFactory {
    this := &KeySpecFactory{}
    this.initFields()
    return this
}


func (this *KeySpecFactory) GetInstance(algorithm string, keyData []int8) *KeySpec {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 *KeySpec) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            PreLogUtil_PutOE(commonStrings.EXCEPTION, this, commonStrings.GET_INSTANCE, e)
            panic(errors.New(fmt.Sprint()))
                }
            }()
            var tsUtil *TsUtil = TsUtil_GetInstance(); _ = tsUtil
            if tsUtil.CompareTo(algorithm, this.DES) == 0 {
                tryControl1, tryResult1 = 1, NewDESKeySpec(keyData); return
            } else if tsUtil.CompareTo(algorithm, this.DESEDE) == 0 {
                tryControl1, tryResult1 = 1, NewDESedeKeySpec(keyData); return
            } else {
                panic(errors.New(fmt.Sprint()))
            }
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

