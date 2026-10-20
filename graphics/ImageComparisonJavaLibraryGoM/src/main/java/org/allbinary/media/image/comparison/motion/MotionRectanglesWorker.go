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

type MotionRectanglesWorker struct {
    BasicEventHandler
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    ImageComparisonInfoVector IBasicArrayList
    MotionRectangleConstraintsInterface MotionRectangleConstraintsInterface
    Index int64
    Running bool
}

// Implements ImageComparisonResultsListener


func (this *MotionRectanglesWorker) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.ImageComparisonInfoVector = NewBasicArrayListD()
    
    
    
}


func NewMotionRectanglesWorker(motionRectangleConstraintsInterface MotionRectangleConstraintsInterface) *MotionRectanglesWorker {
    this := &MotionRectanglesWorker{}
    this.BasicEventHandler = *NewBasicEventHandler()
    this.initFields()
        this.MotionRectangleConstraintsInterface = motionRectangleConstraintsInterface
    return this
}


func (this *MotionRectanglesWorker) OnImageComparisonResultsEvent(imageComparisonResultsEvent *ImageComparisonResultsEvent) {
        this.ImageComparisonInfoVector.Add(imageComparisonResultsEvent)
        this.Run()
}


func (this *MotionRectanglesWorker) OnEvent(allBinaryEventObject IAllBinaryEventObject) {
        this.OnImageComparisonResultsEvent(any(allBinaryEventObject).(*ImageComparisonResultsEvent))
}


// synchronized
func (this *MotionRectanglesWorker) IsRunning() bool {
        return this.Running
}


// synchronized
func (this *MotionRectanglesWorker) SetRunning(running bool) {
        this.Running = running
}


func (this *MotionRectanglesWorker) Run() {
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
            var imageComparisonInfo *ImageComparisonResult = any(imageComparisonResultsEvent.GetImageComparisonResult()).(*ImageComparisonResult); _ = imageComparisonInfo
            this.LogUtil.PutF(fmt.Sprint(imageComparisonInfo), this, this.CommonStrings.RUN)
            var allMotionRectangles *AllMotionRectangles = NewAllMotionRectangles(imageComparisonInfo); _ = allMotionRectangles
            AllMotionRectanglesResultsCacheSingleton_GetInstance().Add(NewMotionRectanglesResultsFrameCacheable(allMotionRectangles, imageComparisonInfo.GetFrameTwo()))
            var consolidatedMotionRectangles *ConsolidateMotionRectangles = NewConsolidateMotionRectangles(allMotionRectangles); _ = consolidatedMotionRectangles
            ConsolidatedMotionRectanglesResultsCacheSingleton_GetInstance().Add(NewMotionRectanglesResultsFrameCacheable(consolidatedMotionRectangles, imageComparisonInfo.GetFrameTwo()))
            var constrainedMotionRectangles *ConstrainedMotionRectangles = NewConstrainedMotionRectangles(this.MotionRectangleConstraintsInterface, consolidatedMotionRectangles); _ = constrainedMotionRectangles
            constrainedMotionRectangles.ApplyMotionRectangleConstraints(consolidatedMotionRectangles)
            ConstrainedMotionRectanglesResultsCacheSingleton_GetInstance().Add(NewMotionRectanglesResultsFrameCacheable(constrainedMotionRectangles, imageComparisonInfo.GetFrameTwo()))
            this.FireEvent(NewMotionRectanglesResultsEvent(this, imageComparisonInfo.GetFrameTwo(), any(constrainedMotionRectangles).(IMotionRectangles)))
            this.ImageComparisonInfoVector.Remove(imageComparisonResultsEvent)
            this.Index++
            this.LogUtil.PutF(CommonLabels_GetInstance().ELAPSED + timeHelper.GetElapsedTNT(), this, this.CommonStrings.RUN)
            this.SetRunning(false)
            this.LogUtil.PutF(this.CommonStrings.END, this, this.CommonStrings.RUN)
            return
        }(); tryControl1 == 1 {
            return
        }
}

