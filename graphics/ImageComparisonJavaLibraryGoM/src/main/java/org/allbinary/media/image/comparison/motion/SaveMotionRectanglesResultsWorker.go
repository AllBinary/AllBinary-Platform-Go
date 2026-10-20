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

type SaveMotionRectanglesResultsWorker struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    Running bool
    MotionRectanglesVector IBasicArrayList
}

// Implements MotionRectanglesResultsListener


func (this *SaveMotionRectanglesResultsWorker) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
    
}


func NewSaveMotionRectanglesResultsWorker() *SaveMotionRectanglesResultsWorker {
    this := &SaveMotionRectanglesResultsWorker{}
    this.initFields()
        this.MotionRectanglesVector = NewBasicArrayListD()
    return this
}


func (this *SaveMotionRectanglesResultsWorker) GetMotionRectanglesVector() IBasicArrayList {
        return this.MotionRectanglesVector
}


func (this *SaveMotionRectanglesResultsWorker) OnMotionRectanglesImageComparisonResultsEvent(motionRectanglesResultsEvent *MotionRectanglesResultsEvent) {
        this.GetMotionRectanglesVector().Add(motionRectanglesResultsEvent)
        this.Run()
}


func (this *SaveMotionRectanglesResultsWorker) OnEvent(allBinaryEventObject IAllBinaryEventObject) {
        this.OnMotionRectanglesImageComparisonResultsEvent(any(allBinaryEventObject).(*MotionRectanglesResultsEvent))
}


// synchronized
func (this *SaveMotionRectanglesResultsWorker) IsRunning() bool {
        return this.Running
}


// synchronized
func (this *SaveMotionRectanglesResultsWorker) SetRunning(running bool) {
        this.Running = running
}


func (this *SaveMotionRectanglesResultsWorker) Run() {
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
            var motionRectanglesResultsEvent *MotionRectanglesResultsEvent = any(this.GetMotionRectanglesVector().Get(0)).(*MotionRectanglesResultsEvent); _ = motionRectanglesResultsEvent
            var motionRectangles IMotionRectangles = motionRectanglesResultsEvent.GetMotionRectangles(); _ = motionRectangles
            NewMotionRectanglesImageInputOutput().Save(motionRectangles, motionRectanglesResultsEvent.GetFrame())
            this.GetMotionRectanglesVector().Remove(motionRectangles)
            this.LogUtil.PutF(CommonLabels_GetInstance().ELAPSED + timeHelper.GetElapsedTNT(), this, this.CommonStrings.RUN)
            this.LogUtil.PutF(this.CommonStrings.END, this, this.CommonStrings.RUN)
            return
        }(); tryControl1 == 1 {
            return
        }
}

