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

type ImageComparisonSearchWorker struct {
    BasicEventHandler
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    ImageComparisonInfoVector IBasicArrayList
    Index int64
    Running bool
    ImageSearchConstraintsInterface ImageComparisonSearchConstraintsInterface
}

// Implements ImageComparisonResultsListener


func (this *ImageComparisonSearchWorker) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
    
    
    
}


func NewImageComparisonSearchWorker(imageSearchConstraintsInterface ImageComparisonSearchConstraintsInterface) *ImageComparisonSearchWorker {
    this := &ImageComparisonSearchWorker{}
    this.BasicEventHandler = *NewBasicEventHandler()
    this.initFields()
        this.ImageComparisonInfoVector = NewBasicArrayListD()
        this.ImageSearchConstraintsInterface = imageSearchConstraintsInterface
    return this
}


func (this *ImageComparisonSearchWorker) OnImageComparisonResultsEvent(imageComparisonResultsEvent *ImageComparisonResultsEvent) {
        this.ImageComparisonInfoVector.Add(imageComparisonResultsEvent.GetImageComparisonResult())
        this.Run()
}


func (this *ImageComparisonSearchWorker) OnEvent(allBinaryEventObject IAllBinaryEventObject) {
        this.OnImageComparisonResultsEvent(any(allBinaryEventObject).(*ImageComparisonResultsEvent))
}


// synchronized
func (this *ImageComparisonSearchWorker) IsRunning() bool {
        return this.Running
}


// synchronized
func (this *ImageComparisonSearchWorker) SetRunning(running bool) {
        this.Running = running
}


func (this *ImageComparisonSearchWorker) Run() {
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
            var imageComparisonInfo *ImageComparisonResult = any(this.ImageComparisonInfoVector.Get(0)).(*ImageComparisonResult); _ = imageComparisonInfo
            this.LogUtil.PutF(fmt.Sprint(imageComparisonInfo), this, this.CommonStrings.RUN)
            var latestBufferedImage *BufferedImage = imageComparisonInfo.GetBufferedImages()[1]; _ = latestBufferedImage
            this.ImageComparisonInfoVector.Remove(imageComparisonInfo)
            this.Index++
            this.LogUtil.PutF(CommonLabels_GetInstance().ELAPSED + timeHelper.GetElapsedTNT(), this, this.CommonStrings.RUN)
            this.SetRunning(false)
            this.LogUtil.PutF(this.CommonStrings.END, this, this.CommonStrings.RUN)
            return
        }(); tryControl1 == 1 {
            return
        }
}

