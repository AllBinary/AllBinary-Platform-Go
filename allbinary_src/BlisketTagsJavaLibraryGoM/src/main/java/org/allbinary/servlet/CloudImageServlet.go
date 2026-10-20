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

type CloudImageServlet struct {
    HttpServlet
    LogUtil *LogUtil
    CommonStrings *CommonStrings
}


func (this *CloudImageServlet) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
}

func NewCloudImageServlet() *CloudImageServlet {
    this := &CloudImageServlet{}
    this.HttpServlet = *NewHttpServlet()
    this.initFields()
    return this
}


// throws ServletException, IOException
func (this *CloudImageServlet) ProcessRequest(request *HttpServletRequest, response *HttpServletResponse) {
        var inputStream IInputStream = nil; _ = inputStream
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEWERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "processRequest()", e)
            }
                }
            StreamUtil_GetInstance().Close(response.GetOutputStream())
            if !StreamUtil_GetInstance().Close(inputStream) {
                response.SendError(HttpServletResponse_SC_NOT_FOUND)
            }
            }()
            BlisketServletUtil_GetInstance().Init(request)
            var requestURI string = request.GetRequestURI(); _ = requestURI
            var file *AbFile = AbFile_CreateAbFile(fmt.Sprint(HttpServlet_URLGLOBALS.GetWebappPath()) + requestURI); _ = file
            inputStream = CloudStreamUtil_GetInstance().GetFile(file)
            var outputStream *strings.Builder = &strings.Builder{}; _ = outputStream
            var byteArray []int8 = StreamUtil_GetInstance().GetByteArray(inputStream, outputStream, make([]int8, 16384)); _ = byteArray
            response.SetContentType("image/jpeg;charset=utf-8")
            response.GetOutputStream().Write(byteArray)
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws ServletException, IOException
func (this *CloudImageServlet) DoGet(request *HttpServletRequest, response *HttpServletResponse) {
        this.ProcessRequest(request, response)
}


// throws ServletException, IOException
func (this *CloudImageServlet) DoPost(request *HttpServletRequest, response *HttpServletResponse) {
        this.ProcessRequest(request, response)
}


func (this *CloudImageServlet) GetServletInfo() string {
        return "Serves Up Cloud Image"
}

