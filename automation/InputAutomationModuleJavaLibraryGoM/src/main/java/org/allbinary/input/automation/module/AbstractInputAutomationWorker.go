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

type AbstractInputAutomationWorker struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    Index int64
    Running bool
    InputAutomationActionInterface InputAutomationActionInterface
    CaptureWorkerInterface CaptureWorkerInterface
    ImageComparisonWorker IImageComparisonWorker
    MotionRectanglesWorker *MotionRectanglesWorker
    CaptureThread IThread
}

// Implements RunnableInterface

// The Java class AbstractInputAutomationWorker has subclasses so its variables are this interface
type IAbstractInputAutomationWorker interface {
    AbstractInputAutomationWorker_() *AbstractInputAutomationWorker
    SetThread(thread IThread)
    IsRunning() bool
    SetRunning(running bool)
    IsAnyDataWorkerRunning() bool
    StartDataWorkers()
    WaitForDataWorkers()
    StopDataWorkers()
    Process()
    Run()
    GetInputAutomationActionInterface() InputAutomationActionInterface
    SetInputAutomationActionInterface(inputAutomationActionInterface InputAutomationActionInterface)
    GetCaptureWorker() CaptureWorkerInterface
    SetCaptureWorker(captureWorkerInterface CaptureWorkerInterface)
    GetImageComparisonWorker() IImageComparisonWorker
    SetImageComparisonWorker(imageComparisonWorker IImageComparisonWorker)
    GetMotionRectanglesWorker() *MotionRectanglesWorker
    SetMotionRectanglesWorker(motionRectanglesWorker *MotionRectanglesWorker)
}

func (this *AbstractInputAutomationWorker) AbstractInputAutomationWorker_() *AbstractInputAutomationWorker {
    return this
}


func (this *AbstractInputAutomationWorker) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
    
    
    
    
    
    
}


func NewAbstractInputAutomationWorker(inputAutomationActionInterface InputAutomationActionInterface) *AbstractInputAutomationWorker {
    this := &AbstractInputAutomationWorker{}
    this.initFields()
    return this
}


// throws Exception
func (this *AbstractInputAutomationWorker) SetThread(thread IThread) {
}


// synchronized
func (this *AbstractInputAutomationWorker) IsRunning() bool {
        return this.Running
}


// synchronized
func (this *AbstractInputAutomationWorker) SetRunning(running bool) {
        this.Running = running
}


// synchronized
func (this *AbstractInputAutomationWorker) IsAnyDataWorkerRunning() bool {
        if this.CaptureThread != nil && (this.CaptureThread.IsAlive() || this.GetCaptureWorker().IsRunning() || this.GetMotionRectanglesWorker().IsRunning() || this.GetImageComparisonWorker().IsRunning()) {
            return true
        }
        return false
}


// throws Exception
// synchronized
func (this *AbstractInputAutomationWorker) StartDataWorkers() {
        if !this.IsAnyDataWorkerRunning() {
            this.CaptureThread = NewThread(this.GetCaptureWorker())
            this.LogUtil.PutF("Starting CaptureWorkers - Need more images - Thread State: " + fmt.Sprint(this.CaptureThread.GetState()), this, "startCaptureWorkers")
            this.CaptureThread.Start()
        }
}


// throws Exception
// synchronized
func (this *AbstractInputAutomationWorker) WaitForDataWorkers() {
        for this.IsAnyDataWorkerRunning() {
            this.LogUtil.PutF("Waiting", this, this.CommonStrings.RUN)
            time.Sleep(time.Duration(250) * time.Millisecond)
        }
}


// throws Exception
func (this *AbstractInputAutomationWorker) StopDataWorkers() {
        this.GetCaptureWorker().SetRunning(false)
}


// throws Exception
func (this *AbstractInputAutomationWorker) Process() {
        panic(errors.New(fmt.Sprint()))
}


func (this *AbstractInputAutomationWorker) Run() {
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
            for this.IsRunning() {
                timeHelper.SetStartTimeTNT()
                this.Process()
                this.Index++
                this.LogUtil.PutF(fmt.Sprint(CommonLabels_GetInstance().ELAPSED + timeHelper.GetElapsedTNT()) + " Index: " + fmt.Sprint(this.Index), this, this.CommonStrings.RUN)
            }
            this.StopDataWorkers()
            this.WaitForDataWorkers()
            this.LogUtil.PutF(this.CommonStrings.END, this, this.CommonStrings.RUN)
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *AbstractInputAutomationWorker) GetInputAutomationActionInterface() InputAutomationActionInterface {
        return this.InputAutomationActionInterface
}


func (this *AbstractInputAutomationWorker) SetInputAutomationActionInterface(inputAutomationActionInterface InputAutomationActionInterface) {
        this.InputAutomationActionInterface = inputAutomationActionInterface
}


func (this *AbstractInputAutomationWorker) GetCaptureWorker() CaptureWorkerInterface {
        return this.CaptureWorkerInterface
}


func (this *AbstractInputAutomationWorker) SetCaptureWorker(captureWorkerInterface CaptureWorkerInterface) {
        this.CaptureWorkerInterface = captureWorkerInterface
}


func (this *AbstractInputAutomationWorker) GetImageComparisonWorker() IImageComparisonWorker {
        return this.ImageComparisonWorker
}


func (this *AbstractInputAutomationWorker) SetImageComparisonWorker(imageComparisonWorker IImageComparisonWorker) {
        this.ImageComparisonWorker = imageComparisonWorker
}


func (this *AbstractInputAutomationWorker) GetMotionRectanglesWorker() *MotionRectanglesWorker {
        return this.MotionRectanglesWorker
}


func (this *AbstractInputAutomationWorker) SetMotionRectanglesWorker(motionRectanglesWorker *MotionRectanglesWorker) {
        this.MotionRectanglesWorker = motionRectanglesWorker
}

