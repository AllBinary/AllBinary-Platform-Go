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

type AbFileUploadFactory struct {
    LogUtil *LogUtil
}

var AbFileUploadFactory_Instance *AbFileUploadFactory = NewAbFileUploadFactory()

func (this *AbFileUploadFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func AbFileUploadFactory_GetInstance() *AbFileUploadFactory {
        return AbFileUploadFactory_Instance
}


func NewAbFileUploadFactory() *AbFileUploadFactory {
    this := &AbFileUploadFactory{}
    this.initFields()
    return this
}


// throws Exception
func (this *AbFileUploadFactory) GetFileItemStreamList(httpServletRequest *HttpServletRequest) IList {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 IList) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().HTTPERROR) {
                var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
                this.LogUtil.Put(commonStrings.EXCEPTION, this, "get()", e)
            }
            panic(e)
                }
            }()
            var fileUpload *AbFileUpload = NewAbFileUpload(NewAbFileItemFactory()); _ = fileUpload
            var iterator *FileItemIterator = fileUpload.GetItemIterator(httpServletRequest); _ = iterator
            var items IList = NewArrayList(); _ = items
            for iterator.HasNext() {
                var fileItemStream *FileItemStream = iterator.Next(); _ = fileItemStream
                HttpFileUploadUtil_Log(fileItemStream)
                var fileItem *AbFileItem = NewAbFileItem(fileItemStream.GetName(), fileItemStream.GetFieldName(), FileItemUtil_GetInstance().GetBytes(fileItemStream)); _ = fileItem
                items.Add(fileItem)
            }
            tryControl1, tryResult1 = 1, items; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws Exception
func (this *AbFileUploadFactory) GetFileItemList(httpServletRequest *HttpServletRequest) IList {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 IList) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().HTTPERROR) {
                this.LogUtil.Put(commonStrings.EXCEPTION, this, "get()", e)
            }
            panic(e)
                }
            }()
            var abFileUpload *AbFileUpload = NewAbFileUpload(NewAbFileItemFactory()); _ = abFileUpload
            tryControl1, tryResult1 = 1, abFileUpload.ParseRequest(httpServletRequest); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

