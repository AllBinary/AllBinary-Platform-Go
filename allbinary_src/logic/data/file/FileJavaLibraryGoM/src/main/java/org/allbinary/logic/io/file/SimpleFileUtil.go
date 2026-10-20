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

type SimpleFileUtil struct {
    LogUtil *LogUtil
    NULL_WRITER *Writer
    NULL_READER *Reader
    CommonStrings *CommonStrings
    NullUtil *NullUtil
    StreamUtil *StreamUtil
}

var SimpleFileUtil_Instance *SimpleFileUtil = NewSimpleFileUtil()

func (this *SimpleFileUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.NULL_WRITER = SimpleFileUtil_NullWriter()
    this.NULL_READER = SimpleFileUtil_NullReader()
    this.CommonStrings = CommonStrings_GetInstance()
    this.NullUtil = NullUtil_GetInstance()
    this.StreamUtil = StreamUtil_GetInstance()
}

func NewSimpleFileUtil() *SimpleFileUtil {
    this := &SimpleFileUtil{}
    this.initFields()
    return this
}


func SimpleFileUtil_GetInstance() *SimpleFileUtil {
        return SimpleFileUtil_Instance
}


func SimpleFileUtil_NullOutputStream() *OutputStream {
        return /* Java anonymous class body: implement it as a named type
            

        // throws IOException
        func (this *SimpleFileUtil) Write(b int) {
        }
            

        // throws IOException
        func (this *SimpleFileUtil) Write(b []int8, off int, len_ int) {
        }
            

        func (this *SimpleFileUtil) Close() {
        }
*/ NewOutputStream()
}


func SimpleFileUtil_NullWriter() *Writer {
        return /* Java anonymous class body: implement it as a named type
            
            

        // throws IOException
        func (this *SimpleFileUtil) Append(c rune) *Writer {
                return this
        }
            

        // throws IOException
        func (this *SimpleFileUtil) Append(csq CharSequence) *Writer {
                return this
        }
            

        // throws IOException
        func (this *SimpleFileUtil) Append(csq CharSequence, start int, end int) *Writer {
                return this
        }
            

        // throws IOException
        func (this *SimpleFileUtil) Write(c int) {
        }
            

        // throws IOException
        func (this *SimpleFileUtil) Write(cbuf []rune, off int, len_ int) {
        }
            

        // throws IOException
        func (this *SimpleFileUtil) Write(str string) {
        }
            

        // throws IOException
        func (this *SimpleFileUtil) Write(str string, off int, len_ int) {
        }
            

        // throws IOException
        func (this *SimpleFileUtil) Flush() {
        }
            

        // throws IOException
        func (this *SimpleFileUtil) Close() {
        }
*/ NewWriter()
}


func SimpleFileUtil_NullReader() *Reader {
        return /* Java anonymous class body: implement it as a named type
            

        // throws IOException
        func (this *SimpleFileUtil) Read() int {
                return -1
        }
            

        // throws IOException
        func (this *SimpleFileUtil) Read(cbuf []rune, off int, len_ int) int {
                return -1
        }
            

        // throws IOException
        func (this *SimpleFileUtil) Read(target *CharBuffer) int {
                return -1
        }
            

        // throws IOException
        func (this *SimpleFileUtil) Ready() bool {
                return false
        }
            

        // throws IOException
        func (this *SimpleFileUtil) Skip(n int64) int64 {
                return 0
        }
            

        func (this *SimpleFileUtil) Close() {
        }
*/ NewReader()
}


func (this *SimpleFileUtil) LoadFileAsList(inputStream *InputStream, max_ int, byteArray1 []int8) *BasicArrayList {
        return this.LoadFileAsListReturnLine(inputStream, max_, byteArray1, 0)
}


func (this *SimpleFileUtil) LoadFileAsListReturnLine(inputStream *InputStream, max_ int, byteArray1 []int8, includeReturnLine int) *BasicArrayList {
        var stringList *BasicArrayList = NewBasicArrayListD(); _ = stringList
        if inputStream == nil {
            return stringList
        }
        var byteArray []int8 = this.NullUtil.NULL_BYTE_ARRAY; _ = byteArray
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 *BasicArrayList) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.StreamUtil.Close(inputStream)
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.ADD, e)
                }
            }()
            var outputStream2 *strings.Builder = &strings.Builder{}; _ = outputStream2
            byteArray = this.StreamUtil.GetByteArray(inputStream, outputStream2, byteArray1)
            this.StreamUtil.Close(inputStream)
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        if byteArray == nil {
            return stringList
        }
        var size int = len(byteArray); _ = size
        var index int = 0; _ = index
        var startIndex int; _ = startIndex
        var returnLine int = 0; _ = returnLine
        for index < size {
            startIndex = index
            for index < size - 1 && int(byteArray[index]) != int(int8('\n')) {
                index++
            }
            if includeReturnLine == 0 {
                if index > 0 && int(byteArray[index - 1]) == int(int8('\r')) {
                    returnLine = 1
                } else {
                    returnLine = 0
                }
            }
            var s string = string(byteArray[startIndex:startIndex + (index + includeReturnLine - returnLine - startIndex)]); _ = s
            stringList.Add(s)
            index++
        }
        return stringList
}


func (this *SimpleFileUtil) CreateStringFromArrayOfStrings(stringArray []string) string {
        var stringMaker *StringMaker = NewStringMaker(); _ = stringMaker
        var size int = len(stringArray); _ = size
        for index := int(0); index < size; index++ {
            stringMaker.Append(stringArray[index])
        }
        return fmt.Sprint(stringMaker)
}

