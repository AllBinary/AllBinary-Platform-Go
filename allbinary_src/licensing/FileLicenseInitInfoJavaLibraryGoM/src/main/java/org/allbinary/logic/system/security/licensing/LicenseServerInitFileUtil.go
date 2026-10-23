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

type LicenseServerInitFileUtil struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    NULL_OUTPUT_STREAM IOutputStream
}


func (this *LicenseServerInitFileUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.NULL_OUTPUT_STREAM = SimpleFileUtil_NullOutputStream()
}

func NewLicenseServerInitFileUtil() *LicenseServerInitFileUtil {
    this := &LicenseServerInitFileUtil{}
    this.initFields()
    return this
}


func (this *LicenseServerInitFileUtil) Init() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.INIT, e)
                }
            }()
            var filePath string = LicenseInitInfoUtil_GetInstance().INITFILENAME; _ = filePath
            LicenseInitInfoUtil_GetInstance().SetFilePath(StringUtil_GetInstance().EMPTY_STRING)
            if FileFactory_GetInstance().IsFile(filePath) {
                this.LogUtil.PutF("Using Existing License File", this, this.CommonStrings.INIT)
            } else {
                this.Write()
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *LicenseServerInitFileUtil) Write() {
        var fileOutputStream IOutputStream = this.NULL_OUTPUT_STREAM; _ = fileOutputStream
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.INIT, e)
                }
            StreamUtil_GetInstance().Close(fileOutputStream)
            }()
            var resourceUtil *ResourceUtil = ResourceUtil_GetInstance(); _ = resourceUtil
            var filePath string = LicenseInitInfoUtil_GetInstance().INITFILENAME; _ = filePath
            var inputStream IInputStream = resourceUtil.GetResourceAsStream(filePath); _ = inputStream
            this.LogUtil.PutF("Writing Default License File", this, this.CommonStrings.INIT)
            var fileStreamFactory *FileStreamFactory = FileStreamFactory_GetInstance(); _ = fileStreamFactory
            fileOutputStream = fileStreamFactory.GetFileOutputStreamInstance(StringUtil_GetInstance().EMPTY_STRING, filePath)
            var b int; _ = b
            var index int = 0; _ = index
            for true {
                b = inputStream.Read()
                if b == -1 {
                    break
                }
                fileOutputStream.Write(b)
                index++
            }
            this.LogUtil.PutF("Wrote Bytes: " + fmt.Sprint(index), this, this.CommonStrings.INIT)
            fileOutputStream.Flush()
            return
        }(); tryControl1 == 1 {
            return
        }
}

