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

type ScreenCaptureImagesWorker struct {
    BasicEventHandler
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    Running bool
    ScreenScavangerRobot *ScreenScavangerRobot
}

// Implements CaptureWorkerInterface

var ScreenCaptureImagesWorker_Index int64

func (this *ScreenCaptureImagesWorker) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
    
}


func NewScreenCaptureImagesWorker() *ScreenCaptureImagesWorker {
    this := &ScreenCaptureImagesWorker{}
    this.BasicEventHandler = *NewBasicEventHandler()
    this.initFields()
        this.ScreenScavangerRobot = NewScreenScavangerRobot()
        ScreenCaptureImagesWorker_Index = ProcessingFrameIndexFactory_Next()
    return this
}


// synchronized
func (this *ScreenCaptureImagesWorker) IsRunning() bool {
        return this.Running
}


// synchronized
func (this *ScreenCaptureImagesWorker) SetRunning(running bool) {
        this.Running = running
}


// throws Exception
func (this *ScreenCaptureImagesWorker) SetThread(thread IThread) {
}


func (this *ScreenCaptureImagesWorker) Run() {
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
                var bufferedImage *BufferedImage = this.ScreenScavangerRobot.GetScreenAsBufferedImages()[0]; _ = bufferedImage
                var frame int64 = int64(ScreenCaptureImagesWorker_Index); _ = frame
                ScreenCaptureImagesWorker_Index++
                CapturedBufferedImagesCacheSingleton_GetInstance().Add(NewBufferedImageFrameCacheable(bufferedImage, frame))
                var capturedImageEvent *CapturedImageWorkerResultsEvent = NewCapturedImageWorkerResultsEvent(this, frame, bufferedImage); _ = capturedImageEvent
                this.FireEvent(capturedImageEvent)
                this.LogUtil.PutF(CommonLabels_GetInstance().ELAPSED + timeHelper.GetElapsedTNT(), this, this.CommonStrings.RUN)
                this.SetRunning(false)
            }
            this.LogUtil.PutF(this.CommonStrings.END, this, this.CommonStrings.RUN)
            return
        }(); tryControl1 == 1 {
            return
        }
}

