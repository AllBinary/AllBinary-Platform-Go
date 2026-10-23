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

type StreamUtil struct {
    LogUtil *LogUtil
}

var StreamUtil_Instance *StreamUtil = NewStreamUtil()

func (this *StreamUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func StreamUtil_GetInstance() *StreamUtil {
        return StreamUtil_Instance
}


func NewStreamUtil() *StreamUtil {
    this := &StreamUtil{}
    this.initFields()
    return this
}


// throws Exception
func (this *StreamUtil) Get(inputStream IInputStream, outputStream IOutputStream, buffer []int8) IOutputStream {
        var len_ int = 0; _ = len_
        for true {
            len_ = inputStream.Read(buffer, 0, len(buffer))
            if len_ != -1 {
            } else {
                break
            }
            outputStream.Write(buffer, 0, len_)
        }
        return outputStream
}


// throws Exception
func (this *StreamUtil) GetByteArray(inputStream IInputStream, outputStream2 *strings.Builder, byteArray []int8) []int8 {
        var outputStream *strings.Builder = outputStream2; _ = outputStream
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 []int8) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            this.Close(outputStream)
            }()
            this.Get(inputStream, outputStream, byteArray)
            tryControl1, tryResult1 = 1, outputStream.ToByteArray(); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *StreamUtil) Close(closeable *Closeable) bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.Put(commonStrings.EXCEPTION, this, commonStrings.CLOSE, e)
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            if closeable != nil {
                closeable.Close()
            }
            tryControl1, tryResult1 = 1, true; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

