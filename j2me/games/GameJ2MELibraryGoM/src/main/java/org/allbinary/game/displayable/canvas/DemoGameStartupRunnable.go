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

type DemoGameStartupRunnable struct {
    ABRunnable
    LogUtil *LogUtil
    DemoCanvas *DemoCanvas
}

var DemoGameStartupRunnable_TYPE int = 4

func (this *DemoGameStartupRunnable) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
}


func NewDemoGameStartupRunnable(demoCanvas *DemoCanvas) *DemoGameStartupRunnable {
    this := &DemoGameStartupRunnable{}
    this.ABRunnable = *NewABRunnable()
    this.initFields()
        this.DemoCanvas = demoCanvas
    return this
}


func (this *DemoGameStartupRunnable) Run() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.SetRunning(false)
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.Put(commonStrings.EXCEPTION, this, commonStrings.RUN, e)
                }
            }()
            this.SetRunning(true)
            this.DemoCanvas.StopGameDemo()
            this.DemoCanvas.Create()
            this.DemoCanvas.Start()
            this.SetRunning(false)
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *DemoGameStartupRunnable) GetType() int {
        return DemoGameStartupRunnable_TYPE
}

