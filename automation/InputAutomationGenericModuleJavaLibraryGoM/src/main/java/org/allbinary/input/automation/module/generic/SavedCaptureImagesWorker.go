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

type SavedCaptureImagesWorker struct {
    BasicEventHandler
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    Running bool
    ScreenScavangerRobot *ScreenScavangerRobot
    SavedCaptureGenericProfileDataWorkerType *SavedCaptureGenericProfileDataWorkerType
}

// Implements CaptureWorkerInterface

var SavedCaptureImagesWorker_Index int64

func (this *SavedCaptureImagesWorker) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
    
    
}


func NewSavedCaptureImagesWorker(savedCaptureGenericProfileDataWorkerType *SavedCaptureGenericProfileDataWorkerType) *SavedCaptureImagesWorker {
    this := &SavedCaptureImagesWorker{}
    this.BasicEventHandler = *NewBasicEventHandler()
    this.initFields()
        this.SavedCaptureGenericProfileDataWorkerType = savedCaptureGenericProfileDataWorkerType
        this.ScreenScavangerRobot = NewScreenScavangerRobot()
        SavedCaptureImagesWorker_Index = ProcessingFrameIndexFactory_Next()
    return this
}


// throws Exception
func (this *SavedCaptureImagesWorker) SetThread(thread IThread) {
}


// synchronized
func (this *SavedCaptureImagesWorker) IsRunning() bool {
        return this.Running
}


// synchronized
func (this *SavedCaptureImagesWorker) SetRunning(running bool) {
        this.Running = running
}


func (this *SavedCaptureImagesWorker) Run() {
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
                var frame int64 = int64(SavedCaptureImagesWorker_Index); _ = frame
                var filePathStringBuffer *StringMaker = NewStringMaker(); _ = filePathStringBuffer
                filePathStringBuffer.Append(this.SavedCaptureGenericProfileDataWorkerType.GetPath())
                filePathStringBuffer.Append(LongUtil_FillIn(fmt.Sprint(frame)))
                filePathStringBuffer.Append(MediaDataFactory_GetInstance().JPG.GetExtension())
                var filePath string = fmt.Sprint(filePathStringBuffer); _ = filePath
                this.LogUtil.PutF("Loading Image File Path: " + filePath, this, this.CommonStrings.RUN)
                var file *File = NewFile(filePath); _ = file
                if file.IsFile() {
                    var bufferedImage *BufferedImage = ImageIO_Read(file); _ = bufferedImage
                    SavedCaptureImagesWorker_Index++
                    CapturedBufferedImagesCacheSingleton_GetInstance().Add(NewBufferedImageFrameCacheable(bufferedImage, frame))
                    var capturedImageEvent *CapturedImageWorkerResultsEvent = NewCapturedImageWorkerResultsEvent(this, frame, bufferedImage); _ = capturedImageEvent
                    this.FireEvent(capturedImageEvent)
                } else {
                    this.LogUtil.PutF("Could Not Load File: " + filePath, this, this.CommonStrings.RUN)
                }
                this.LogUtil.PutF(CommonLabels_GetInstance().ELAPSED + timeHelper.GetElapsedTNT(), this, this.CommonStrings.RUN)
                this.SetRunning(false)
            }
            this.LogUtil.PutF(this.CommonStrings.END, this, this.CommonStrings.RUN)
            return
        }(); tryControl1 == 1 {
            return
        }
}

