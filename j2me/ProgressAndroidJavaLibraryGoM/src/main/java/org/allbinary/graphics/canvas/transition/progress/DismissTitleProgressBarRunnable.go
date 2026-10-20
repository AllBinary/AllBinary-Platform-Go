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

type DismissTitleProgressBarRunnable struct {
    ProgressRunnable
    LogUtil *LogUtil
    ZOOM_OUT_AND_IN []int
}


func (this *DismissTitleProgressBarRunnable) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.ZOOM_OUT_AND_IN = []int{TransistionTypes_GetInstance().ZOOM_OUT, TransistionTypes_GetInstance().ZOOM_IN}
}


func NewDismissTitleProgressBarRunnable(midletActivity *Activity, progressCanvas IProgressCanvas) *DismissTitleProgressBarRunnable {
    this := &DismissTitleProgressBarRunnable{}
        this.ProgressRunnable = *NewProgressRunnable(midletActivity, progressCanvas)
        this.initFields()
    return this
}


func (this *DismissTitleProgressBarRunnable) Run() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.Put(commonStrings.EXCEPTION, this, commonStrings.RUN, e)
                }
            }()
            this.MidletActivity.OnDismissProgress(this.ZOOM_OUT_AND_IN)
            return
        }(); tryControl1 == 1 {
            return
        }
}

