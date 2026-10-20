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

type SaveCapturedImageWorker struct {
    BasicEventHandler
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    Running bool
    CapturedImageWorkerResultsEventVector IBasicArrayList
}

// Implements CapturedImageWorkerResultsListener


func (this *SaveCapturedImageWorker) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
    this.CapturedImageWorkerResultsEventVector = NewBasicArrayListD()
}


func NewSaveCapturedImageWorker() *SaveCapturedImageWorker {
    this := &SaveCapturedImageWorker{}
    this.BasicEventHandler = *NewBasicEventHandler()
    this.initFields()
    return this
}


// synchronized
func (this *SaveCapturedImageWorker) IsRunning() bool {
        return this.Running
}


// synchronized
func (this *SaveCapturedImageWorker) SetRunning(running bool) {
        this.Running = running
}


func (this *SaveCapturedImageWorker) OnCaptureEvent(capturedImageEvent *CapturedImageWorkerResultsEvent) {
        this.CapturedImageWorkerResultsEventVector.Add(capturedImageEvent)
        this.Run()
}


func (this *SaveCapturedImageWorker) OnEvent(allBinaryEventObject IAllBinaryEventObject) {
        this.OnCaptureEvent(any(allBinaryEventObject).(*CapturedImageWorkerResultsEvent))
}


func (this *SaveCapturedImageWorker) Run() {
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
            var capturedImageWorkerResultsEvent *CapturedImageWorkerResultsEvent = (any(this.CapturedImageWorkerResultsEventVector.Get(0)).(*CapturedImageWorkerResultsEvent)); _ = capturedImageWorkerResultsEvent
            var screenBufferedImage *BufferedImage = capturedImageWorkerResultsEvent.GetBufferedImage(); _ = screenBufferedImage
            NewCapturedImageInputOutput().Save(screenBufferedImage, capturedImageWorkerResultsEvent.GetFrame())
            this.CapturedImageWorkerResultsEventVector.Remove(capturedImageWorkerResultsEvent)
            this.LogUtil.PutF(CommonLabels_GetInstance().ELAPSED + timeHelper.GetElapsedTNT(), this, this.CommonStrings.RUN)
            this.SetRunning(false)
            this.LogUtil.PutF(this.CommonStrings.END, this, this.CommonStrings.RUN)
            return
        }(); tryControl1 == 1 {
            return
        }
}

