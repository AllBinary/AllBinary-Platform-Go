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

type RequestHandler struct {
    HttpServlet
    LogUtil *LogUtil
}


func (this *RequestHandler) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}

func NewRequestHandler() *RequestHandler {
    this := &RequestHandler{}
    this.HttpServlet = *NewHttpServlet()
    this.initFields()
    return this
}


// throws ServletException
func (this *RequestHandler) Init(config *ServletConfig) {
        this.HttpServlet.Init(config)
}


func (this *RequestHandler) Destroy() {
}


// throws ServletException, IOException
func (this *RequestHandler) ProcessRequest(request *HttpServletRequest, response *HttpServletResponse) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SERVLETERROR) {
                this.LogUtil.Put("Servlet LicensingException", this, "processRequest()", e)
            }
            response.SendRedirect(this.Abcs.Globals.URLGLOBALS.LICENSEERRORPAGE)
                }
            }()
            response.SetContentType("text/html")
            var out *PrintWriter = response.GetWriter(); _ = out
            var workflowInterface WorkFlowInterface = RequestWorkFlowFactory_GetInstance(any(request).(*ServletRequest), any(response).(*ServletResponse), this.GetServletConfig(), this.GetServletContext()); _ = workflowInterface
            if workflowInterface != nil {
                out.Print(workflowInterface.Process())
            }
            out.Flush()
            out.Close()
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws ServletException, IOException
func (this *RequestHandler) DoGet(request *HttpServletRequest, response *HttpServletResponse) {
        this.ProcessRequest(request, response)
}


// throws ServletException, IOException
func (this *RequestHandler) DoPost(request *HttpServletRequest, response *HttpServletResponse) {
        this.ProcessRequest(request, response)
}


func (this *RequestHandler) GetServletInfo() string {
        return "WorkFlow Request Handler"
}

