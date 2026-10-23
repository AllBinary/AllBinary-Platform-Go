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

type BasicCryptUtil struct {
    LogUtil *LogUtil
    TsUtil *TsUtil
    XML_START string
}

var BasicCryptUtil_Instance *BasicCryptUtil = NewBasicCryptUtil()

func (this *BasicCryptUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.TsUtil = TsUtil_GetInstance()
    this.XML_START = "<?xml"
}

func NewBasicCryptUtil() *BasicCryptUtil {
    this := &BasicCryptUtil{}
    this.initFields()
    return this
}


func BasicCryptUtil_GetInstance() *BasicCryptUtil {
        return BasicCryptUtil_Instance
}


// throws Exception
func (this *BasicCryptUtil) GetDecryptedInputStream(inputStream IInputStream, cryptInterface CryptInterface) IInputStream {
        var buffer *strings.Builder = &strings.Builder{}; _ = buffer
        var ch int; _ = ch
        for true {
            ch = inputStream.Read()
            if ch == -1 {
                break
            }
            buffer.Write(ch)
        }
        var decrypted []int8 = cryptInterface.Decrypt(buffer.ToByteArray()); _ = decrypted
        var decryptedString string = string(decrypted); _ = decryptedString
        var index int = strings.Index(decryptedString, string(this.XML_START)); _ = index
        if index > 0 {
            decryptedString = decryptedString[index:]
        }
        return NewByteArrayInputStream(this.TsUtil.GetByteArray(decryptedString))
}

