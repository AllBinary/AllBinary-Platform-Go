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

type ImageComparisonWorker struct {
    BasicEventHandler
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    BufferedImageVector IBasicArrayList
    ImageComparatorConstraintsInterface ImageComparatorConstraintsInterface
    ImageComparator *ImageComparator
    Running bool
    Index2 int
}

// Implements CapturedImageWorkerResultsListener

// The Java class ImageComparisonWorker has subclasses so its variables are this interface
type IImageComparisonWorker interface {
    IBasicEventHandler
    ImageComparisonWorker_() *ImageComparisonWorker
    IsRunning() bool
    SetRunning(running bool)
    OnCaptureEvent(capturedImageWorkerResultsEvent *CapturedImageWorkerResultsEvent)
    OnEvent(allBinaryEventObject IAllBinaryEventObject)
    Run()
}

func (this *ImageComparisonWorker) ImageComparisonWorker_() *ImageComparisonWorker {
    return this
}


func (this *ImageComparisonWorker) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.BufferedImageVector = NewBasicArrayListD()
    
    
    
    
}


func NewImageComparisonWorker(imageComparatorConstraintsInterface ImageComparatorConstraintsInterface) *ImageComparisonWorker {
    this := &ImageComparisonWorker{}
    this.BasicEventHandler = *NewBasicEventHandler()
    this.initFields()
        this.ImageComparatorConstraintsInterface = imageComparatorConstraintsInterface
        this.ImageComparator = NewImageComparator(imageComparatorConstraintsInterface)
    return this
}


// synchronized
func (this *ImageComparisonWorker) IsRunning() bool {
        return this.Running
}


// synchronized
func (this *ImageComparisonWorker) SetRunning(running bool) {
        this.Running = running
}


// synchronized
func (this *ImageComparisonWorker) OnCaptureEvent(capturedImageWorkerResultsEvent *CapturedImageWorkerResultsEvent) {
        this.BufferedImageVector.Add(capturedImageWorkerResultsEvent)
        if this.BufferedImageVector.Size() > 1 {
            this.Run()
        }
}


// synchronized
func (this *ImageComparisonWorker) OnEvent(allBinaryEventObject IAllBinaryEventObject) {
        this.OnCaptureEvent(any(allBinaryEventObject).(*CapturedImageWorkerResultsEvent))
}


func (this *ImageComparisonWorker) Run() {
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
            if this.ImageComparatorConstraintsInterface.IsFrameAllowed(this.Index2) {
                var capturedImageWorkerResultsEvent []*CapturedImageWorkerResultsEvent = make([]*CapturedImageWorkerResultsEvent, 2); _ = capturedImageWorkerResultsEvent
                capturedImageWorkerResultsEvent[0] = any(this.BufferedImageVector.Get(0)).(*CapturedImageWorkerResultsEvent)
                capturedImageWorkerResultsEvent[1] = any(this.BufferedImageVector.Get(1)).(*CapturedImageWorkerResultsEvent)
                if this.ImageComparatorConstraintsInterface.IsImageValid(capturedImageWorkerResultsEvent[0].GetBufferedImage()) && this.ImageComparatorConstraintsInterface.IsImageValid(capturedImageWorkerResultsEvent[1].GetBufferedImage()) {
                    var imageComparisonResult *ImageComparisonResult = this.ImageComparator.Compare(capturedImageWorkerResultsEvent[0].GetBufferedImage(), capturedImageWorkerResultsEvent[1].GetBufferedImage(), capturedImageWorkerResultsEvent[0].GetFrame(), capturedImageWorkerResultsEvent[1].GetFrame(), 0); _ = imageComparisonResult
                    var frame int64 = capturedImageWorkerResultsEvent[1].GetFrame(); _ = frame
                    var imageComparisonResultFrameCacheable *ImageComparisonResultFrameCacheable = NewImageComparisonResultFrameCacheable(imageComparisonResult, frame); _ = imageComparisonResultFrameCacheable
                    ImageComparisonResultCacheSingleton_GetInstance().Add(imageComparisonResultFrameCacheable)
                    this.FireEvent(NewImageComparisonResultsEvent(this, imageComparisonResult))
                    this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append("Image Comparison Result: ").Append(fmt.Sprint(imageComparisonResult)).Append(" for frame: ").Appendlong(frame)), this, this.CommonStrings.RUN)
                } else {
                    this.LogUtil.PutF("An Image Was Not Valid: Image Worker Event Processing terminated", this, this.CommonStrings.RUN)
                }
            }
            this.BufferedImageVector.Remove(0)
            this.Index2++
            var message string = fmt.Sprint(NewStringMaker().Append("Frame: ").Appendint(this.Index2).Append(CommonLabels_GetInstance().ELAPSED).Appendlong(timeHelper.GetElapsedTNT())); _ = message
            this.LogUtil.PutF(message, this, this.CommonStrings.RUN)
            this.SetRunning(false)
            this.LogUtil.PutF(this.CommonStrings.END, this, this.CommonStrings.RUN)
            return
        }(); tryControl1 == 1 {
            return
        }
}

