/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2022 AllBinary 
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

type ConcurrentImageLoadingProcessor struct {
    BaseImageLoadingProcessor
    LogUtil *LogUtil
    Runnable IABRunnable
}


func (this *ConcurrentImageLoadingProcessor) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
}


func NewConcurrentImageLoadingProcessor(imageCache IImageCache) *ConcurrentImageLoadingProcessor {
    this := &ConcurrentImageLoadingProcessor{}
    this.BaseImageLoadingProcessor = *NewBaseImageLoadingProcessor()
    this.initFields()
        this.Runnable = NewImageCacheRunnable(imageCache)
    return this
}


func (this *ConcurrentImageLoadingProcessor) RunTask() {
        if !this.Runnable.IsRunning() {
            ImageThreadPool_GetInstance().RunTask(this.Runnable)
        }
}

type ImageCacheRunnable struct {
    ABRunnable
    ImageCache IImageCache
}


func (this *ImageCacheRunnable) initFields() {
        
}
    

    func NewConcurrentImageLoadingProcessor(imageCache IImageCache) *ConcurrentImageLoadingProcessor {
        this := &ConcurrentImageLoadingProcessor{}
        this.ABRunnable = *NewABRunnable()
        this.initFields()
            this.ImageCache = imageCache
        return this
    }
    

    func (this *ConcurrentImageLoadingProcessor) Run() {
            var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
            if tryControl1 := func() (tryControl1 int) {
                defer func() {
                    if recovered := recover(); recovered != nil {
                        e := Java_ToError(recovered)
                        _ = e
                this.SetRunning(false)
                var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
                logUtil.Put(commonStrings.EXCEPTION, this, commonStrings.RUN, e)
                    }
                }()
                this.SetRunning(true)
                this.ImageCache.WaitForLoadNow()
                this.ImageCache.LoadImages()
                this.ImageCache.LoadRemainingAnimations()
                this.SetRunning(false)
                var progressCanvas IProgressCanvas = ProgressCanvasFactory_GetInstance(); _ = progressCanvas
                if !progressCanvas.ProgressCanvas_().InProgress {
                    progressCanvas.EndFromInitialLazyLoadingComplete()
                }
                return
            }(); tryControl1 == 1 {
                return
            }
    }

