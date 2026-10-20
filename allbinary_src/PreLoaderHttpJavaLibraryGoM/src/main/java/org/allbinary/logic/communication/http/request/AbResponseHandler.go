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

type AbResponseHandler struct {
    LogUtil *LogUtil
}


func (this *AbResponseHandler) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func NewAbResponseHandler() *AbResponseHandler {
    this := &AbResponseHandler{}
    this.initFields()
    return this
}


// throws JspTagException
func AbResponseHandler_SendJspTagRedirect(pageContext *PageContext, e error) {
        var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e2 := Java_ToError(recovered)
                    _ = e2
            if LogConfigTypes_LOGGING.Contains(LogConfigTypeFactory_GetInstance().LICENSINGERROR) {
                logUtil.Put("Exception in Redirect Handling", "AbResponseHandler", "sendJspRedirect", e)
            }
            panic(errors.New(fmt.Sprint()))
                }
            }()
            if LogConfigTypes_LOGGING.Contains(LogConfigTypeFactory_GetInstance().JSPTAGERROR) {
                logUtil.Put("Jsp Tag Exception", "AbResponseHandler", "sendJspRedirect()", e)
            }
            AbResponseHandler_SendRedirect(pageContext, FREEBLISKET_PATH_GLOBALS.GetInstance().ERRORPAGE)
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws ServletException
func AbResponseHandler_SendServletRedirect(pageContext *PageContext, e error) {
        var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e2 := Java_ToError(recovered)
                    _ = e2
            if LogConfigTypes_LOGGING.Contains(LogConfigTypeFactory_GetInstance().LICENSINGERROR) {
                logUtil.Put("Exception in Redirect Handling", "AbResponseHandler", "sendJspRedirect", e)
            }
            panic(errors.New(fmt.Sprint()))
                }
            }()
            if LogConfigTypes_LOGGING.Contains(LogConfigTypeFactory_GetInstance().JSPTAGERROR) {
                logUtil.Put("Jsp Tag Exception", "AbResponseHandler", "sendJspRedirect()", e)
            }
            AbResponseHandler_SendRedirect(pageContext, FREEBLISKET_PATH_GLOBALS.GetInstance().ERRORPAGE)
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws JspTagException
func AbResponseHandler_SendJspTagLicensingRedirect(pageContext *PageContext, e error) {
        var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e2 := Java_ToError(recovered)
                    _ = e2
            if LogConfigTypes_LOGGING.Contains(LogConfigTypeFactory_GetInstance().LICENSINGERROR) {
                logUtil.Put("Exception in Redirect Handling", "AbResponseHandler", "sendJspTagLicensingRedirect", e)
            }
            panic(errors.New(fmt.Sprint()))
                }
            }()
            if LogConfigTypes_LOGGING.Contains(LogConfigTypeFactory_GetInstance().JSPTAGERROR) {
                logUtil.Put("Licensing Exception", "AbResponseHandler", "sendJspTagLicensingRedirect()", e)
            }
            AbResponseHandler_SendRedirect(pageContext, FREEBLISKET_PATH_GLOBALS.GetInstance().LICENSEERRORPAGE)
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws ServletException
func AbResponseHandler_SendServletLicensingRedirect(pageContext *PageContext, e error) {
        var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e2 := Java_ToError(recovered)
                    _ = e2
            if LogConfigTypes_LOGGING.Contains(LogConfigTypeFactory_GetInstance().LICENSINGERROR) {
                logUtil.Put("Exception in Redirect Handling", "AbResponseHandler", "sendServletLicensingRedirect", e)
            }
            panic(errors.New(fmt.Sprint()))
                }
            }()
            if LogConfigTypes_LOGGING.Contains(LogConfigTypeFactory_GetInstance().JSPTAGERROR) {
                logUtil.Put("Licensing Exception", "AbResponseHandler", "sendServletLicensingRedirect()", e)
            }
            AbResponseHandler_SendRedirect(pageContext, FREEBLISKET_PATH_GLOBALS.GetInstance().LICENSEERRORPAGE)
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func AbResponseHandler_SendRedirect(pageContext *PageContext, page string) {
        var request *HttpServletRequest = any(pageContext.GetRequest()).(*HttpServletRequest); _ = request
        var response *HttpServletResponse = any(pageContext.GetResponse()).(*HttpServletResponse); _ = response
        var requestDispatcher *RequestDispatcher = request.GetRequestDispatcher(page); _ = requestDispatcher
        if response.IsCommitted() {
            response.SendRedirect(page)
        } else {
            if tryControl1 := func() (tryControl1 int) {
                defer func() {
                    if recovered := recover(); recovered != nil {
                        e := Java_ToError(recovered)
                        _ = e
                pageContext.GetOut().Print("Error: " + page)
                    }
                }()
                requestDispatcher.Forward(request, response)
                return
            }(); tryControl1 == 1 {
                return
            }
        }
}

