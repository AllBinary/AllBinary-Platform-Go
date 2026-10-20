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

type HttpFileUploadUtil struct {
    LogUtil *LogUtil
}

var HttpFileUploadUtil_Instance *HttpFileUploadUtil = NewHttpFileUploadUtil()

func (this *HttpFileUploadUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func HttpFileUploadUtil_GetInstance() *HttpFileUploadUtil {
        return HttpFileUploadUtil_Instance
}


func NewHttpFileUploadUtil() *HttpFileUploadUtil {
    this := &HttpFileUploadUtil{}
    this.initFields()
    return this
}


// throws Exception
func (this *HttpFileUploadUtil) IsValid(anyType any) bool {
        if anyType != nil && !(func() bool { _, ok := any(anyType).(string); return ok }()) {
            if func() bool { _, ok := any(anyType).(*FileItem); return ok }() {
                return true
            } else {
                panic(errors.New(fmt.Sprint("Object Not Instance Of FileItem but is: " + fmt.Sprint(ClassUtil_ViewAll(anyType, CommonSeps_GetInstance().NEW_LINE)))))
            }
        }
        return false
}


// throws Exception
func HttpFileUploadUtil_Log_1(fileItem *FileItemStream) {
        var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        var commonSeps *CommonSeps = CommonSeps_GetInstance(); _ = commonSeps
        stringBuffer.Append("FileItem Log:")
        stringBuffer.Append(commonSeps.NEW_LINE)
        stringBuffer.Append("Content Type:")
        stringBuffer.Append(fileItem.GetContentType())
        stringBuffer.Append(commonSeps.NEW_LINE)
        stringBuffer.Append("Field Name:")
        stringBuffer.Append(fileItem.GetFieldName())
        stringBuffer.Append(commonSeps.NEW_LINE)
        stringBuffer.Append("Name:")
        stringBuffer.Append(fileItem.GetName())
        stringBuffer.Append(commonSeps.NEW_LINE)
        stringBuffer.Append("Size:")
        stringBuffer.Append(" Not Calculated Yet")
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().HTTP) {
            logUtil.PutF(fmt.Sprint(stringBuffer), HttpFileUploadUtil_GetInstance(), "log()")
        }
}

// Java overloads of HttpFileUploadUtil.log
func HttpFileUploadUtil_Log(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[*FileItemStream](args[0], pass) {
            HttpFileUploadUtil_Log_1(Java_Arg[*FileItemStream](args[0]))
            return
        }
        if len(args) == 1 && Java_IsArg[*FileItem](args[0], pass) {
            HttpFileUploadUtil_Log_2(Java_Arg[*FileItem](args[0]))
            return
        }
    }
    panic("No Java overload of HttpFileUploadUtil.log for the arguments")
}


func HttpFileUploadUtil_Log_2(fileItem *FileItem) {
        var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        var commonSeps *CommonSeps = CommonSeps_GetInstance(); _ = commonSeps
        stringBuffer.Append("FileItem Log:")
        stringBuffer.Append(commonSeps.NEW_LINE)
        stringBuffer.Append("Content Type:")
        stringBuffer.Append(fileItem.GetContentType())
        stringBuffer.Append(commonSeps.NEW_LINE)
        stringBuffer.Append("Field Name:")
        stringBuffer.Append(fileItem.GetFieldName())
        stringBuffer.Append(commonSeps.NEW_LINE)
        stringBuffer.Append("Name:")
        stringBuffer.Append(fileItem.GetName())
        stringBuffer.Append(commonSeps.NEW_LINE)
        stringBuffer.Append("Size:")
        stringBuffer.Appendlong(int64(fileItem.GetSize()))
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().HTTPREQUEST) {
            logUtil.PutF(fmt.Sprint(stringBuffer), HttpFileUploadUtil_GetInstance(), "log()")
        }
}

