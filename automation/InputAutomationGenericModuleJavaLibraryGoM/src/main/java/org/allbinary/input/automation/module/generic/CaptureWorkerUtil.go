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

type CaptureWorkerUtil struct {
    LogUtil *LogUtil
}


func (this *CaptureWorkerUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func NewCaptureWorkerUtil() *CaptureWorkerUtil {
    this := &CaptureWorkerUtil{}
    this.initFields()
    return this
}


// throws Exception
func CaptureWorkerUtil_ProcessProfileActionConditions(vector IBasicArrayList, frame int64) {
        var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
        var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
        logUtil.PutF(commonStrings.START, "CaptureWorkerUtil", "processProfileActionConditions")
        var size int = vector.Size(); _ = size
        for index := int(0); index < size; index++ {
            var profileActionConditionInterface ProfileActionScriptConditionInterface = any(vector.Get(index)).(ProfileActionScriptConditionInterface); _ = profileActionConditionInterface
            if profileActionConditionInterface.ShouldProcess(frame) {
                profileActionConditionInterface.Process(frame)
            } else {
            }
        }
}

