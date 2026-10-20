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

type DownloadFileServlet struct {
    HttpServlet
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    DOWNLOAD string
    DEFAULT_BUFFER_SIZE int
}


func (this *DownloadFileServlet) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.DOWNLOAD = "download"
    this.DEFAULT_BUFFER_SIZE = 16384
}

func NewDownloadFileServlet() *DownloadFileServlet {
    this := &DownloadFileServlet{}
    this.HttpServlet = *NewHttpServlet()
    this.initFields()
    return this
}


// throws ServletException, IOException
func (this *DownloadFileServlet) ProcessRequest(request *HttpServletRequest, response *HttpServletResponse) {
        var inputStream IInputStream = nil; _ = inputStream
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEWERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "processRequest()", e)
            }
            response.SendError(HttpServletResponse_SC_NOT_FOUND)
                }
            StreamUtil_GetInstance().Close(response.GetOutputStream())
            if !StreamUtil_GetInstance().Close(inputStream) {
            }
            }()
            BlisketServletUtil_GetInstance().Init(request)
            var requestPath string = request.GetRequestURI(); _ = requestPath
            if requestPath == "" {
                response.SendError(HttpServletResponse_SC_NOT_FOUND)
                tryControl1 = 1; return
            }
            var beginIndex int = strings.Index(requestPath, string(this.DOWNLOAD)); _ = beginIndex
            var filePath string = StringUtil_GetInstance().EMPTY_STRING; _ = filePath
            if beginIndex >= 0 {
                filePath = requestPath[beginIndex + len(this.DOWNLOAD):]
            } else {
                response.SendError(HttpServletResponse_SC_NOT_FOUND)
            }
            var file *AbFile = AbFile_CreateAbFile(fmt.Sprint(HttpServlet_URLGLOBALS.GetWebappPath()) + filePath); _ = file
            if !file.Exists() {
                response.SendError(HttpServletResponse_SC_NOT_FOUND)
                tryControl1 = 1; return
            }
            var hashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = hashMap
            var authenticationHelper *AuthenticationHelper = any(NewAuthenticationHelperFactory().GetInstance(hashMap, request)).(*AuthenticationHelper); _ = authenticationHelper
            if authenticationHelper.IsAuthenticated() {
                if AuthenticationHelperUtil_GetInstance().IsAuthorized(authenticationHelper, filePath) {
                    inputStream = CloudStreamUtil_GetInstance().GetFile(file)
                    var contentType string = this.GetServletContext().GetMimeType(file.GetName()); _ = contentType
                    if contentType == "" {
                        contentType = "application/octet-stream"
                    }
                    response.Reset()
                    response.SetBufferSize(this.DEFAULT_BUFFER_SIZE)
                    response.SetContentType(contentType)
                    response.SetHeader("Content-Length", fmt.Sprint(file.Length()))
                    var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                    stringBuffer.Append("attachment; filename=\"")
                    stringBuffer.Append(file.GetName())
                    stringBuffer.Append(CommonSeps_GetInstance().QUOTE)
                    response.SetHeader("Content-Disposition", fmt.Sprint(stringBuffer))
                    StreamUtil_GetInstance().Get(inputStream, response.GetOutputStream(), make([]int8, 16348))
                } else {
                    response.SendError(HttpServletResponse_SC_UNAUTHORIZED, "You are not Authorized")
                }
            } else {
                response.SendError(HttpServletResponse_SC_UNAUTHORIZED, "Please Login")
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws ServletException, IOException
func (this *DownloadFileServlet) DoGet(request *HttpServletRequest, response *HttpServletResponse) {
        this.ProcessRequest(request, response)
}


// throws ServletException, IOException
func (this *DownloadFileServlet) DoPost(request *HttpServletRequest, response *HttpServletResponse) {
        this.ProcessRequest(request, response)
}


func (this *DownloadFileServlet) GetServletInfo() string {
        return "Serves Up Cloud Files"
}

