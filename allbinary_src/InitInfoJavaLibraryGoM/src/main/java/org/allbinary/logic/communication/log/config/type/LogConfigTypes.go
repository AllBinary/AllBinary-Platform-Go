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

type LogConfigTypes struct {
}

var LogConfigTypes_JSP_ERROR_NO_FUNCTION string = "in JSP Not In Function"
var LogConfigTypes_LOGGING IBasicArrayList = NewBasicArrayListD()


func init() {
        LogConfigTypes_Init()
}


func NewLogConfigTypes() *LogConfigTypes {
    this := &LogConfigTypes{}
    return this
}


// synchronized
func LogConfigTypes_Init() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            PreLogUtil_PutSE("Unable to initialize LogConfigTypes", "LogConfigTypes", "init()", e)
                }
            }()
            PreLogUtil_Put("Initialize LogconfigTypes", "LogConfigTypes", "init()")
            LogConfigTypeFactory_GetInstance()
            var loggingInitInfo *LoggingInitInfo = NewLoggingInitInfo(); _ = loggingInitInfo
            var logConfigTypeVector IBasicArrayList = LoggingInitInfo_GetTypeList(); _ = logConfigTypeVector
            PreLogUtil_Put("Number Of Log Configs: " + fmt.Sprint(loggingInitInfo.GetNumberOfLogConfigs()), "LogConfigTypes", "init()")
            PreLogUtil_Put("Number Of Log Config Type Names: " + fmt.Sprint(logConfigTypeVector.Size()), "LogConfigTypes", "init()")
            LogConfigTypes_LOGGING.AddAllList(logConfigTypeVector)
            LogConfigTypes_LOGGING.Add(LogConfigTypeFactory_GetInstance().NETBEANS_MODULE)
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func LogConfigTypes_GetInstance(node *Node) *LogConfigType {
        var nameValueNode *Node = DomSearchHelper_GetNode(LogConfigTypeData_GetInstance().NAME, node.GetChildNodes()); _ = nameValueNode
        var name string = DomNodeHelper_GetTextNodeValue(nameValueNode); _ = name
        var descriptionValueNode *Node = DomSearchHelper_GetNodeNoThrow(LogConfigTypeData_GetInstance().DESCRIPTION, node.GetChildNodes()); _ = descriptionValueNode
        if descriptionValueNode != nil {
            var description string = DomNodeHelper_GetTextNodeValue(descriptionValueNode); _ = description
        }
        var availableLogConfigTypes IBasicArrayList = LogConfigType_AvailableLogConfigTypes; _ = availableLogConfigTypes
        var size int = availableLogConfigTypes.Size(); _ = size
        var logConfigType *LogConfigType; _ = logConfigType
        for index := int(0); index < size; index++ {
            logConfigType = any(availableLogConfigTypes.Get(index)).(*LogConfigType)
            if strings.Compare(logConfigType.GetName(), name) == 0 {
                return logConfigType
            }
        }
        panic(errors.New(fmt.Sprint("No Such LogConfigType: " + name)))
}

