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

type SaveImageComparisonWorker struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    Running bool
    ImageComparisonInfoVector IBasicArrayList
}

// Implements ImageComparisonResultsListener


func (this *SaveImageComparisonWorker) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
    this.ImageComparisonInfoVector = NewBasicArrayListD()
}


func NewSaveImageComparisonWorker() *SaveImageComparisonWorker {
    this := &SaveImageComparisonWorker{}
    this.initFields()
    return this
}


func (this *SaveImageComparisonWorker) OnImageComparisonResultsEvent(imageComparisonResultsEvent *ImageComparisonResultsEvent) {
        this.ImageComparisonInfoVector.Add(imageComparisonResultsEvent)
        this.Run()
}


func (this *SaveImageComparisonWorker) OnEvent(allBinaryEventObject IAllBinaryEventObject) {
        this.OnImageComparisonResultsEvent(any(allBinaryEventObject).(*ImageComparisonResultsEvent))
}


// synchronized
func (this *SaveImageComparisonWorker) IsRunning() bool {
        return this.Running
}


// synchronized
func (this *SaveImageComparisonWorker) SetRunning(running bool) {
        this.Running = running
}


func (this *SaveImageComparisonWorker) Run() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.RUN, e)
                }
            }()
            this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.RUN)
            this.SetRunning(true)
            var timeHelper ITimeDelayHelper = NewTimeDelayHelper(1000); _ = timeHelper
            timeHelper.SetStartTimeTNT()
            var imageComparisonResultsEvent *ImageComparisonResultsEvent = any(this.ImageComparisonInfoVector.Get(0)).(*ImageComparisonResultsEvent); _ = imageComparisonResultsEvent
            var imageComparisonInfo *ImageComparisonResult = imageComparisonResultsEvent.GetImageComparisonResult(); _ = imageComparisonInfo
            this.LogUtil.PutF(fmt.Sprint(imageComparisonInfo), this, this.CommonStrings.RUN)
            NewComparisonImageInputOutput().Save(imageComparisonInfo, imageComparisonInfo.GetFrameTwo())
            this.ImageComparisonInfoVector.Remove(imageComparisonInfo)
            this.LogUtil.PutF(CommonLabels_GetInstance().ELAPSED + timeHelper.GetElapsedTNT(), this, this.CommonStrings.RUN)
            this.LogUtil.PutF(this.CommonStrings.END, this, this.CommonStrings.RUN)
            return
        }(); tryControl1 == 1 {
            return
        }
}

