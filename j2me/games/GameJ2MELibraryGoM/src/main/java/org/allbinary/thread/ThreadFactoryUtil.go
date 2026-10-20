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

type ThreadFactoryUtil struct {
    LogUtil *LogUtil
}

var ThreadFactoryUtil_Instance *ThreadFactoryUtil = NewThreadFactoryUtil()

func (this *ThreadFactoryUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}

func NewThreadFactoryUtil() *ThreadFactoryUtil {
    this := &ThreadFactoryUtil{}
    this.initFields()
    return this
}


func ThreadFactoryUtil_GetInstance() *ThreadFactoryUtil {
        return ThreadFactoryUtil_Instance
}


func (this *ThreadFactoryUtil) GetInstanceGameCanvasRunnable(runnable GameCanvasRunnableInterface) IThread {
        return this.GetInstanceType(runnable, runnable.GetType())
}


func (this *ThreadFactoryUtil) GetInstanceForRunnable(runnable IABRunnable) IThread {
        return this.GetInstanceType(runnable, runnable.GetType())
}


func (this *ThreadFactoryUtil) GetInstanceType(runnable Runnable, type_ int) IThread {
        var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
        if J2MEUtil_IsHTML() {
            if type_ == DemoCanvas_TYPE || type_ == AllBinaryGameCanvas_TYPE {
                var demoGameSingleThreadStartRunnable *RunnableCanvasSingleThreadStartRunnable = NewRunnableCanvasSingleThreadStartRunnable(any(runnable).(IRunnableCanvas)); _ = demoGameSingleThreadStartRunnable
                PreLogUtil_Put(fmt.Sprint(NewStringMaker().Append("Using Pseudo Thread for DemoCanvas/AllBinaryGameCanvas under PlayN/HTML5: ").Append(StringUtil_GetInstance().ToString(runnable))), this, commonStrings.CONSTRUCTOR)
                var primaryThreadPool IThreadPool = PrimaryThreadPool_GetInstance(); _ = primaryThreadPool
                primaryThreadPool.RunTask(demoGameSingleThreadStartRunnable)
            } else {
                PreLogUtil_Put(fmt.Sprint(NewStringMaker().Append("Using Pseudo Thread for Runnable under PlayN/HTML5: ").Append(StringUtil_GetInstance().ToString(runnable))), this, commonStrings.CONSTRUCTOR)
                var primaryThreadPool IThreadPool = PrimaryThreadPool_GetInstance(); _ = primaryThreadPool
                primaryThreadPool.RunTask(runnable)
            }
        }
        var thread IThread = NewThread(runnable, fmt.Sprint(runnable)); _ = thread
        this.LogUtil.PutF(fmt.Sprint(thread), this, commonStrings.CONSTRUCTOR)
        return thread
}


func (this *ThreadFactoryUtil) Start(thread IThread) {
        thread.Start()
}

