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

type UploadFileServlet struct {
    HttpServlet
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    RequestHashMap *HashMap
    FileName string
}


func (this *UploadFileServlet) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
    
}

func NewUploadFileServlet() *UploadFileServlet {
    this := &UploadFileServlet{}
    this.HttpServlet = *NewHttpServlet()
    this.initFields()
    return this
}


// throws ServletException, IOException
func (this *UploadFileServlet) ProcessRequest(request *HttpServletRequest, response *HttpServletResponse) {
        var isError bool = false; _ = isError
        var inputStream IInputStream = nil; _ = inputStream
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEWERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "processRequest()", e)
            }
            isError = true
            response.SendError(HttpServletResponse_SC_NOT_FOUND)
                }
            if !isError {
                StreamUtil_GetInstance().Close(response.GetOutputStream())
            }
            if !StreamUtil_GetInstance().Close(inputStream) {
            }
            }()
            BlisketServletUtil_GetInstance().Init(request)
            this.RequestHashMap = NewMultipartRequestParams(request).ToHashMap()
            var filePath string = any(this.RequestHashMap.Get(FileUploadData_GetInstance().UPLOAD_TO_FILE_PATH)).(string); _ = filePath
            if filePath == "" {
                isError = true
                response.SendError(HttpServletResponse_SC_NOT_FOUND)
                tryControl1 = 1; return
            }
            var hashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = hashMap
            var authenticationHelper *AuthenticationHelper = any(NewAuthenticationHelperFactory().GetInstance(hashMap, request)).(*AuthenticationHelper); _ = authenticationHelper
            if authenticationHelper.IsAuthenticated() {
                if AuthenticationHelperUtil_GetInstance().IsAuthorized(authenticationHelper, filePath) {
                    this.SaveFile(filePath)
                    response.SetContentType("text/plain")
                    response.GetOutputStream().Write(Java_GetBytes("OK!"))
                } else {
                    isError = true
                    response.SendError(HttpServletResponse_SC_UNAUTHORIZED, "You are not Authorized")
                }
            } else {
                isError = true
                response.SendError(HttpServletResponse_SC_UNAUTHORIZED, "Please Login")
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func (this *UploadFileServlet) SaveFile(filePath string) {
        var fileOutputStream *AbFileOutputStream = nil; _ = fileOutputStream
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
            if fileOutputStream != nil {
                fileOutputStream.Close()
            }
            }()
            var fileItem *FileItem = any(this.RequestHashMap.Get(FileUploadData_GetInstance().FILE_DATA)).(*FileItem); _ = fileItem
            if fileItem != nil && fileItem.GetSize() > 1 {
                this.FileName = HttpRequestUtil_GetInstance().GenerateFileName(fileItem.GetName())
                if strings.HasSuffix(filePath, "/") || strings.HasSuffix(filePath, "\\") {
                    filePath = filePath + this.FileName
                }
                var file *AbFile = AbFile_CreateAbFile(fmt.Sprint(HttpServlet_URLGLOBALS.GetWebappPath()) + filePath); _ = file
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().HTTPREQUEST) {
                    var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                    stringBuffer.Append("Uploaded File: ")
                    stringBuffer.Append(this.FileName)
                    stringBuffer.Append(" New File: ")
                    stringBuffer.Append(file.GetPath())
                    this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "saveFile()")
                }
                HttpFileUploadUtil_Log(fileItem)
                var byteArray []int8 = fileItem.Get(); _ = byteArray
                if file.IsFile() {
                    file.Delete()
                    file.CreateNewFile()
                }
                fileOutputStream = AbFileOutputStream_CreateFromAbFile(file)
                fileOutputStream.Write(byteArray)
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws ServletException, IOException
func (this *UploadFileServlet) DoGet(request *HttpServletRequest, response *HttpServletResponse) {
        this.ProcessRequest(request, response)
}


// throws ServletException, IOException
func (this *UploadFileServlet) DoPost(request *HttpServletRequest, response *HttpServletResponse) {
        this.ProcessRequest(request, response)
}


func (this *UploadFileServlet) GetServletInfo() string {
        return "Upload Cloud Files"
}

