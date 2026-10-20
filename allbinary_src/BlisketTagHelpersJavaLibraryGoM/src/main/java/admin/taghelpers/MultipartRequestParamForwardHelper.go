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

type MultipartRequestParamForwardHelper struct {
    LogUtil *LogUtil
    PageContext *PageContext
    RequestMapInterface RequestMapInterface
    Page string
    Command string
    ParamVector IBasicArrayList
}


func (this *MultipartRequestParamForwardHelper) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
    
    
    
}


func NewMultipartRequestParamForwardHelper(parentHelperTag IHelperTag, page string, command string, paramVector IBasicArrayList, pageContext *PageContext) *MultipartRequestParamForwardHelper {
    this := &MultipartRequestParamForwardHelper{}
    this.initFields()
        this.PageContext = pageContext
        this.Page = page
        this.Command = command
        this.ParamVector = paramVector
        this.RequestMapInterface = InventoryItemViewParentTagHelper_GetInstance().GetRequestMapInterface(parentHelperTag)
    return this
}


// throws Exception
func (this *MultipartRequestParamForwardHelper) Forward() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(commonStrings.EXCEPTION, this, "forward()", e)
            }
                }
            }()
            if this.ShouldForward() {
                var params string = this.GetParams(); _ = params
                var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                stringBuffer.Append(this.Page)
                stringBuffer.Append(params)
                this.PageContext.Forward(fmt.Sprint(stringBuffer))
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *MultipartRequestParamForwardHelper) GetParams() string {
        var commonSeps *CommonSeps = CommonSeps_GetInstance(); _ = commonSeps
        var AMP string = commonSeps.AMPERSAND; _ = AMP
        var QUESTION string = commonSeps.QUESTION; _ = QUESTION
        var EQUALS string = commonSeps.EQUALS; _ = EQUALS
        var hashMap *HashMap = this.RequestMapInterface.GetRequestHashMap(); _ = hashMap
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        var size int = this.ParamVector.Size(); _ = size
        for index := int(0); index < size; index++ {
            var key string = any(this.ParamVector.Get(index)).(string); _ = key
            var value string = any(hashMap.Get(key)).(string); _ = value
            if index != 0 {
                stringBuffer.Append(AMP)
            } else {
                stringBuffer.Append(QUESTION)
            }
            stringBuffer.Append(key)
            stringBuffer.Append(EQUALS)
            stringBuffer.Append(value)
        }
        return fmt.Sprint(stringBuffer)
}


func (this *MultipartRequestParamForwardHelper) ShouldForward() bool {
        var hashMap *HashMap = this.RequestMapInterface.GetRequestHashMap(); _ = hashMap
        var requestCommand string = any(hashMap.Get(GLOBALS2.ADMINCOMMAND)).(string); _ = requestCommand
        if StringValidationUtil_GetInstance().IsEmpty(this.Command) || (!StringValidationUtil_GetInstance().IsEmpty(requestCommand) && strings.Compare(this.Command, requestCommand) == 0) {
            return true
        }
        return false
}

