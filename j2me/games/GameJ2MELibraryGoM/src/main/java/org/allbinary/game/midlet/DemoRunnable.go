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

type DemoRunnable struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    DemoGameMidlet IDemoGameMidlet
    StartDemoGameMidletEvent *DemoGameMidletEvent
}

// Implements Runnable


func (this *DemoRunnable) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
    
}


func NewDemoRunnable(demoGameMidlet IDemoGameMidlet) *DemoRunnable {
    this := &DemoRunnable{}
    this.initFields()
        this.DemoGameMidlet = demoGameMidlet
        this.StartDemoGameMidletEvent = NewDemoGameMidletEvent(this.DemoGameMidlet, DemoGameMidletStateFactory_GetInstance().START_DEMO)
    return this
}


func (this *DemoRunnable) Run() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.RUN, e)
                }
            }()
            this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(CommonLabels_GetInstance().START_LABEL).Append("GameCanvasRunnableInterface")), this, this.CommonStrings.RUN)
            this.DemoGameMidlet.CommandAction(MyCommandsFactory_GetInstance().SET_DISPLAYABLE, ProgressCanvasFactory_GetInstance())
            this.DemoGameMidlet.SetGameCanvasRunnableInterface(this.DemoGameMidlet.CreateDemoGameCanvasRunnableInterface())
            this.DemoGameMidlet.DemoSetup()
            DemoGameMidletEventHandler_GetInstance().FireEvent(this.StartDemoGameMidletEvent)
            this.DemoGameMidlet.StartGameCanvasRunnableInterface()
            this.DemoGameMidlet.PostDemoSetup()
            this.LogUtil.PutF(this.CommonStrings.END_RUNNABLE, this, this.CommonStrings.RUN)
            return
        }(); tryControl1 == 1 {
            return
        }
}

