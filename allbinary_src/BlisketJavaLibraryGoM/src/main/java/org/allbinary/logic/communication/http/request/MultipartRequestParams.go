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

type MultipartRequestParams struct {
    RequestParams
    LogUtil *LogUtil
    Special bool
}


func (this *MultipartRequestParams) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.Special = false
}




// Java overloads of new MultipartRequestParams
func NewMultipartRequestParams(args ...any) *MultipartRequestParams {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[*HttpServletRequest](args[0], pass) {
            return NewMultipartRequestParams_1(Java_Arg[*HttpServletRequest](args[0]))
        }
        if len(args) == 1 && Java_IsArg[*PageContext](args[0], pass) {
            return NewMultipartRequestParams_2(Java_Arg[*PageContext](args[0]))
        }
    }
    panic("No Java overload of new MultipartRequestParams for the arguments")
}
func NewMultipartRequestParams_1(request *HttpServletRequest) *MultipartRequestParams {
    this := &MultipartRequestParams{}
        this.RequestParams = *NewRequestParams()
        this.initFields()
        this.ProcessMultipartRequest(request)
    return this
}


func NewMultipartRequestParams_2(pageContext *PageContext) *MultipartRequestParams {
    this := &MultipartRequestParams{}
        this.RequestParams = *NewRequestParams()
        this.initFields()
        this.ProcessMultipartRequest(any(pageContext.GetRequest()).(*HttpServletRequest))
    return this
}


func (this *MultipartRequestParams) SetSpecial() {
        this.Special = true
}


func (this *MultipartRequestParams) ProcessMultipartRequest(request *HttpServletRequest) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().HTTPREQUEST) {
                this.LogUtil.PutF("Using Normal RequestParams", this, "processMultipartRequest()")
            }
            this.RequestParams.SetMap(request.GetParameterMap())
                }
            }()
            var specialRequest *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = specialRequest
            var multipartRequestList IList = AbFileUploadFactory_GetInstance().GetFileItemStreamList(request); _ = multipartRequestList
            if multipartRequestList != nil {
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().HTTPREQUEST) {
                    this.LogUtil.PutF("FileItem List Size: " + fmt.Sprint(multipartRequestList.Size()), this, "processMultipartRequest()")
                }
                var fileItemArray []any = multipartRequestList.ToArray(); _ = fileItemArray
                var size int = len(fileItemArray); _ = size
                for index := int(0); index < size; index++ {
                    var fileItem *FileItem = any(fileItemArray[index]).(*FileItem); _ = fileItem
                    var name string = fileItem.GetName(); _ = name
                    if StringValidationUtil_GetInstance().IsEmpty(name) {
                        specialRequest.Put(fileItem.GetFieldName(), fileItem.GetString())
                    } else {
                        specialRequest.Put(fileItem.GetFieldName(), fileItem)
                    }
                }
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().HTTPREQUEST) {
                    this.LogUtil.PutF("Special Request Data: " + fmt.Sprint(specialRequest), this, "processMultipartRequest()")
                }
                this.SetSpecial()
                this.SetMap(specialRequest)
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func (this *MultipartRequestParams) ToHashMap() *HashMap {
        if this.Special {
            return any(this.GetMap()).(*HashMap)
        } else {
            return this.RequestParams.ToHashMap()
        }
}

