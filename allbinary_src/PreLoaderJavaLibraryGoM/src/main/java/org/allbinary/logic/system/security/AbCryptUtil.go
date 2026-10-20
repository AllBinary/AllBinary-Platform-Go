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

type AbCryptUtil struct {
    StreamUtil *StreamUtil
}

var AbCryptUtil_Instance *AbCryptUtil = NewAbCryptUtil()

func (this *AbCryptUtil) initFields() {
    this.StreamUtil = StreamUtil_GetInstance()
}


func AbCryptUtil_GetInstance() *AbCryptUtil {
        return AbCryptUtil_Instance
}


func NewAbCryptUtil() *AbCryptUtil {
    this := &AbCryptUtil{}
    this.initFields()
    return this
}


// throws Exception
func (this *AbCryptUtil) Decrypt(inputStream *InputStream, key string) []int8 {
        var outputStream *strings.Builder = nil; _ = outputStream
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 []int8) {
            defer func() {
            this.StreamUtil.Close(outputStream)
            this.StreamUtil.Close(inputStream)
            }()
            outputStream = any(this.StreamUtil.Get(inputStream, &strings.Builder{}, make([]int8, 16384))).(*strings.Builder)
            var abCrypt *AbCrypt = NewAbCrypt(KeySpecFactory_GetInstance().DESEDE); _ = abCrypt
            abCrypt.Init(key)
            tryControl1, tryResult1 = 1, abCrypt.Decrypt(outputStream.ToByteArray()); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

