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

type LoggingInitInfo struct {
}

var LoggingInitInfo_INITFILENAME string = "logConfig.xml"
var LoggingInitInfo_PACKAGE string = PATH_GLOBALS.GetInstance().INIT_PATH
var LoggingInitInfo_LogConfigInfoList IBasicArrayList = nil
var LoggingInitInfo_HasRead bool = false


func NewLoggingInitInfo() *LoggingInitInfo {
    this := &LoggingInitInfo{}
    return this
}


// throws Exception
// synchronized
func LoggingInitInfo_GetDoc() *Document {
        var path string = URLGLOBALS.GetWebappPath() + LoggingInitInfo_PACKAGE; _ = path
        var FILEABPATH IAbPath = NewAbPath(path, LoggingInitInfo_INITFILENAME); _ = FILEABPATH
        var file *AbFile = AbFile_CreateAbFileFromAbPath(FILEABPATH); _ = file
        var document *Document = DomDocumentHelper_Create(NewAbFileLocalInputStream(file)); _ = document
        return document
}


// throws Exception
// synchronized
func LoggingInitInfo_Write() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            }()
            LoggingInitInfo_HasRead = false
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
// synchronized
func LoggingInitInfo_Read() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            }()
            var document *Document = LoggingInitInfo_GetDoc(); _ = document
            var logConfigsNodeList *NodeList = document.GetElementsByTagName(LogConfigsData_GetInstance().NAME); _ = logConfigsNodeList
            var logConfigsNode *Node = logConfigsNodeList.Item(0); _ = logConfigsNode
            var logConfigNodeVector IBasicArrayList = DomSearchHelper_GetAllNodes(LogConfigData_GetInstance().NAME, logConfigsNode.GetChildNodes()); _ = logConfigNodeVector
            LoggingInitInfo_LogConfigInfoList = NewBasicArrayListD()
            var size int = logConfigNodeVector.Size(); _ = size
            for i := int(0); i < size; i++ {
                var node *Node = any(logConfigNodeVector.Get(i)).(*Node); _ = node
                var logConfigInfo *LogConfig = NewLogConfig(node); _ = logConfigInfo
                LoggingInitInfo_LogConfigInfoList.Add(logConfigInfo)
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


// synchronized
func LoggingInitInfo_SetHasRead(value bool) {
        LoggingInitInfo_HasRead = value
}


// throws Exception
// synchronized
func LoggingInitInfo_UpdateIfNeeded() {
        if !LoggingInitInfo_HasRead {
            LoggingInitInfo_Read()
            LoggingInitInfo_HasRead = true
            if LoggingInitInfo_LogConfigInfoList == nil {
                panic(errors.New(fmt.Sprint("Read Failed")))
            }
        }
}


func (this *LoggingInitInfo) IsValid() bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            tryControl1, tryResult1 = 1, true; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// synchronized
func LoggingInitInfo_Set(logConfigInfoVector IBasicArrayList) {
        LoggingInitInfo_LogConfigInfoList = logConfigInfoVector
}


// throws Exception
func LoggingInitInfo_Get() IBasicArrayList {
        LoggingInitInfo_UpdateIfNeeded()
        return LoggingInitInfo_LogConfigInfoList
}


// throws Exception
func LoggingInitInfo_GetTypeNameList() IBasicArrayList {
        LoggingInitInfo_UpdateIfNeeded()
        var allLogTypeVector IBasicArrayList = NewBasicArrayListD(); _ = allLogTypeVector
        var size int = LoggingInitInfo_LogConfigInfoList.Size(); _ = size
        for index := int(0); index < size; index++ {
            var logConfigInfo *LogConfig = any(LoggingInitInfo_LogConfigInfoList.BasicArrayList_().ObjectArray[index]).(*LogConfig); _ = logConfigInfo
            var logTypeVector IBasicArrayList = logConfigInfo.GetTypeVector(); _ = logTypeVector
            if logConfigInfo.IsEnabled() {
                allLogTypeVector.AddAllList(logTypeVector)
            }
        }
        return allLogTypeVector
}


// throws Exception
func LoggingInitInfo_GetTypeList() IBasicArrayList {
        LoggingInitInfo_UpdateIfNeeded()
        var allLogTypeVector IBasicArrayList = NewBasicArrayListD(); _ = allLogTypeVector
        var size int = LoggingInitInfo_LogConfigInfoList.Size(); _ = size
        for index := int(0); index < size; index++ {
            var logConfigInfo *LogConfig = any(LoggingInitInfo_LogConfigInfoList.BasicArrayList_().ObjectArray[index]).(*LogConfig); _ = logConfigInfo
            var logTypeVector IBasicArrayList = logConfigInfo.GetTypeVector(); _ = logTypeVector
            if logConfigInfo.IsEnabled() {
                allLogTypeVector.AddAllList(logTypeVector)
            }
        }
        return allLogTypeVector
}


// throws Exception
func (this *LoggingInitInfo) GetNumberOfLogConfigs() int {
        return LoggingInitInfo_LogConfigInfoList.Size()
}


func (this *LoggingInitInfo) ToString() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            tryControl1, tryResult1 = 1, "Error"; return
                }
            }()
            var document *Document = LoggingInitInfo_GetDoc(); _ = document
            tryControl1, tryResult1 = 1, XmlDocumentHelper_ToString(document); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

func (this *LoggingInitInfo) String() string {
    return this.ToString()
}

