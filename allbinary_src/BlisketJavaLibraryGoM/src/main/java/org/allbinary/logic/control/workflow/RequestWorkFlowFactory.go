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

type RequestWorkFlowFactory struct {
}



func NewRequestWorkFlowFactory() *RequestWorkFlowFactory {
    this := &RequestWorkFlowFactory{}
    return this
}


// throws Exception, LicensingException
func RequestWorkFlowFactory_GetInstance_1(request *ServletRequest, response *ServletResponse, servletConfig *ServletConfig, servletContext *ServletContext) WorkFlowInterface {
        var httpRequest *HttpServletRequest = any(request).(*HttpServletRequest); _ = httpRequest
        var requestCommand string = httpRequest.GetPathInfo(); _ = requestCommand
        var index int = strings.Index(requestCommand, string(AbPathData_GetInstance().SEPARATOR)); _ = index
        var storeName string = requestCommand[0:index]; _ = storeName
        var requestName string = requestCommand[index:len(requestCommand)]; _ = requestName
        return WorkFlowEntityFactory_GetInstance().Create2().Get(requestName, storeName).(WorkFlowInterface)
}

// Java overloads of RequestWorkFlowFactory.getInstance
func RequestWorkFlowFactory_GetInstance(args ...any) WorkFlowInterface {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 4 && Java_IsArg[*ServletRequest](args[0], pass) && Java_IsArg[*ServletResponse](args[1], pass) && Java_IsArg[*ServletConfig](args[2], pass) && Java_IsArg[*ServletContext](args[3], pass) {
            return RequestWorkFlowFactory_GetInstance_1(Java_Arg[*ServletRequest](args[0]), Java_Arg[*ServletResponse](args[1]), Java_Arg[*ServletConfig](args[2]), Java_Arg[*ServletContext](args[3]))
        }
        if len(args) == 2 && Java_IsArg[*HashMap](args[0], pass) && Java_IsArg[*PageContext](args[1], pass) {
            return RequestWorkFlowFactory_GetInstance_2(Java_Arg[*HashMap](args[0]), Java_Arg[*PageContext](args[1]))
        }
    }
    panic("No Java overload of RequestWorkFlowFactory.getInstance for the arguments")
}


// throws Exception, LicensingException
func RequestWorkFlowFactory_GetInstance_2(hashMap *HashMap, pageContext *PageContext) WorkFlowInterface {
        var workFlowName string = any(pageContext.GetRequest().GetParameter(WorkFlowData_GetInstance().NAME)).(string); _ = workFlowName
        var weblisketSession *WeblisketSession = NewWeblisketSession(hashMap, pageContext); _ = weblisketSession
        var storeName string = weblisketSession.GetStoreName(); _ = storeName
        return WorkFlowEntityFactory_GetInstance().Create2().Get(workFlowName, storeName).(WorkFlowInterface)
}

