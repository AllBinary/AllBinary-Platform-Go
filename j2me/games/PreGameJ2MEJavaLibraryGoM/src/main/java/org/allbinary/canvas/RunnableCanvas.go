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

type RunnableCanvas struct {
    MyCanvas
    NullUtil *NullUtil
    StdUtil *StdUtil
    SystemWrapper *ABSystemWrapper
    TsUtil *TsUtil
    Thread IThread
    CurrentThread IThread
    Running bool
    LoopTimeHelper ITimeDelayHelper
    RunnableCanvasRefreshHelper IProcessor
    CommonLabels *CommonLabels
    ThreadObjectUtil *ThreadObjectUtil
    SET_RUNNING string
    IS_RUNNING string
    THREAD string
    NOT_EQUAL string
    PAUSE_SLEEP string
    START_PAUSE string
    END_PAUSE string
    PROCESS_LOOP_SLEEP string
    PauseWait int64
    Notified bool
}

// Implements RunnableInterface

// The Java class RunnableCanvas has subclasses so its variables are this interface
type IRunnableCanvas interface {
    IMyCanvas
    RunnableCanvas_() *RunnableCanvas
    InitCommands(cmdListener CommandListener)
    IsMainCanvas() bool
    SetThread(thread IThread)
    SetRunning(running bool)
    IsRunning() bool
    GetLoopTimeHelperP() ITimeDelayHelper
    SetWait(wait int)
    GetWait() int
    SetCurrentThread()
    SetCurrentThreadFake()
    ShowNotify()
    StopWaiting()
    WaitOnNotify(wait int64)
    ProcessSleep()
    IsPausable() bool
    ProcessGameSleep(sleep int64)
    ProcessLoopSleep()
    IsSingleThread() bool
    Run()
    End2()
    GetType() int
}

func (this *RunnableCanvas) RunnableCanvas_() *RunnableCanvas {
    return this
}


func (this *RunnableCanvas) initFields() {
    this.NullUtil = NullUtil_GetInstance()
    this.StdUtil = StdUtil_GetInstance()
    this.SystemWrapper = ABSystemWrapper_GetInstance()
    this.TsUtil = TsUtil_GetInstance()
    this.Thread = NullThread_NULL_THREAD
    this.CurrentThread = NullThread_NULL_THREAD
    
    this.LoopTimeHelper = NewTimeDelayHelper(NullWaitGameRunnable_GetInstance().GameRunnable_().WAIT)
    this.RunnableCanvasRefreshHelper = Processor_GetInstance()
    this.CommonLabels = CommonLabels_GetInstance()
    this.ThreadObjectUtil = ThreadObjectUtil_GetInstance()
    this.SET_RUNNING = "setRunning"
    this.IS_RUNNING = "isRunning"
    this.THREAD = "Thread: "
    this.NOT_EQUAL = " != "
    this.PAUSE_SLEEP = "pause sleep"
    this.START_PAUSE = "start pause - game thread sleep at: "
    this.END_PAUSE = "end pause - game thread sleep at: "
    this.PROCESS_LOOP_SLEEP = "processLoopSleep"
    this.PauseWait = 0
    this.Notified = false
}


func NewRunnableCanvas(commandListener CommandListener, childNameList IBasicArrayList, hasParam bool) *RunnableCanvas {
    this := &RunnableCanvas{}
        this.MyCanvas = *NewMyCanvas(CommonStrings_GetInstance().UNKNOWN, childNameList)
        this.initFields()
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append("RunnableCanvas delay: ").Appendint(this.LoopTimeHelper.TimeDelayHelper_().Delay)), this, this.CommonStrings.CONSTRUCTOR)
        this.RunnableCanvasRefreshHelper = NewRunnableCanvasRefreshHelper(this)
        if commandListener != nil {
            this.InitCommands(commandListener)
        } else if hasParam {
            this.LogUtil.PutF("commandListener was null", this, "initCommands")
        }
    return this
}


func (this *RunnableCanvas) InitCommands(cmdListener CommandListener) {
        this.LogUtil.PutF(this.CommonStrings.NOT_IMPLEMENTED, this, "initCommands")
}


func (this *RunnableCanvas) IsMainCanvas() bool {
        if this.GetCustomCommandListener() != nil {
            return true
        } else {
            return false
        }
}


func (this *RunnableCanvas) SetThread(thread IThread) {
        this.Thread = thread
}


func (this *RunnableCanvas) SetRunning(running bool) {
        this.Running = running
        if !this.Running {
            this.Thread = NullThread_NULL_THREAD
            { // synchronized (this)
                this.ThreadObjectUtil.NotifyObject(this)
            }
        }
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(this.IS_RUNNING).Appendboolean(this.Running)), this, this.SET_RUNNING)
}


// synchronized
func (this *RunnableCanvas) IsRunning() bool {
        if this.Thread == this.CurrentThread {
            return this.Running
        } else {
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append(this.THREAD)
            if this.Thread != nil {
                stringBuffer.Append(fmt.Sprint(this.Thread))
            }
            stringBuffer.Append(this.NOT_EQUAL)
            if this.CurrentThread != nil {
                stringBuffer.Append(StringUtil_GetInstance().ToString(this.CurrentThread))
            }
            this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, this.IS_RUNNING)
            return false
        }
}


func (this *RunnableCanvas) GetLoopTimeHelperP() ITimeDelayHelper {
        return this.LoopTimeHelper
}


func (this *RunnableCanvas) SetWait(wait int) {
        this.LoopTimeHelper.TimeDelayHelper_().Delay = wait
        this.PauseWait = int64(wait) * 3
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append("setWait - delay: ").Appendint(this.LoopTimeHelper.TimeDelayHelper_().Delay)), this, this.CommonStrings.CONSTRUCTOR)
}


func (this *RunnableCanvas) GetWait() int {
        return this.LoopTimeHelper.TimeDelayHelper_().Delay
}


func (this *RunnableCanvas) SetCurrentThread() {
        this.CurrentThread = Thread_CurrentThread()
}


func (this *RunnableCanvas) SetCurrentThreadFake() {
        this.CurrentThread = this.Thread
}


func (this *RunnableCanvas) ShowNotify() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "showNotify", e)
                }
            }()
            this.StopWaiting()
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func (this *RunnableCanvas) StopWaiting() {
        this.Notified = true
        { // synchronized (this)
            this.ThreadObjectUtil.NotifyObject(this)
        }
        this.Repaint()
}


// throws Exception
// synchronized
func (this *RunnableCanvas) WaitOnNotify(wait int64) {
        if !this.Notified {
            if wait > 0 {
                this.ThreadObjectUtil.WaitObjectFor(this, wait)
            } else {
                this.ThreadObjectUtil.WaitObject(this)
            }
        }
}


// throws Exception
func (this *RunnableCanvas) ProcessSleep() {
        time.Sleep(time.Duration(this.PauseWait) * time.Millisecond)
}


func (this *RunnableCanvas) IsPausable() bool {
        return false
}


// throws Exception
func (this *RunnableCanvas) ProcessGameSleep(sleep int64) {
        time.Sleep(time.Duration(sleep) * time.Millisecond)
}


// throws Exception
func (this *RunnableCanvas) ProcessLoopSleep() {
        this.RunnableCanvasRefreshHelper.Process()
        if this.IsPaused() {
            var stringMaker *StringMaker = NewStringMaker(); _ = stringMaker
            if this.IsRunning() && !this.IsSingleThread() {
                stringMaker.Delete(0, stringMaker.Length())
                var currentTimeMillis int64 = this.SystemWrapper.CurrentTimeMillis(); _ = currentTimeMillis
                this.LogUtil.PutF(fmt.Sprint(stringMaker.Append(this.START_PAUSE).Appendlong(currentTimeMillis).Append(this.PAUSE_SLEEP).Appendlong(this.PauseWait)), this, this.PROCESS_LOOP_SLEEP)
                for this.IsPaused() && this.IsRunning() && !this.IsSingleThread() {
                    this.ProcessSleep()
                    if !this.IsPausable() {
                        stringMaker.Delete(0, stringMaker.Length())
                        this.LogUtil.PutF(fmt.Sprint(stringMaker.Append(this.END_PAUSE).Appendlong(currentTimeMillis)), this, this.PROCESS_LOOP_SLEEP)
                        return
                    }
                }
            }
        }
        var elapsedTime int64 = this.LoopTimeHelper.GetElapsedTNT(); _ = elapsedTime
        var wait int64 = int64(this.LoopTimeHelper.TimeDelayHelper_().Delay); _ = wait
        if elapsedTime > wait {
            elapsedTime = wait
        } else {
            this.ProcessGameSleep(wait - elapsedTime)
        }
}


func (this *RunnableCanvas) IsSingleThread() bool {
        return false
}


func (this *RunnableCanvas) Run() {
        this.SetCurrentThread()
}


// throws Exception
func (this *RunnableCanvas) End2() {
}


func (this *RunnableCanvas) GetType() int {
        return -1
}

