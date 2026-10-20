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

type EmulatorCustomRenderer struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    RenderStrings *RendererStrings
}


func (this *EmulatorCustomRenderer) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.RenderStrings = RendererStrings_GetInstance()
}

func NewEmulatorCustomRenderer() *EmulatorCustomRenderer {
    this := &EmulatorCustomRenderer{}
    this.initFields()
    return this
}


func (this *EmulatorCustomRenderer) OnSurfaceCreated(gl GL10, eglConfig *EGLConfig) {
        this.LogUtil.PutF(this.CommonStrings.START, this, this.RenderStrings.ON_SURFACE_CREATED)
        var initEmulatorFactory *InitEmulatorFactory = InitEmulatorFactory_GetInstance(); _ = initEmulatorFactory
        if !initEmulatorFactory.IsInitEmulator() {
            if tryControl1 := func() (tryControl1 int) {
                defer func() {
                    if recovered := recover(); recovered != nil {
                        e := Java_ToError(recovered)
                        _ = e
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.RenderStrings.ON_SURFACE_CREATED, e)
                    }
                }()
                var WAIT_FOR_EMULATOR string = "Waiting on Emulator Initialization"; _ = WAIT_FOR_EMULATOR
                for !initEmulatorFactory.IsInitEmulator() {
                    PreLogUtil_Put(WAIT_FOR_EMULATOR, this, this.RenderStrings.ON_SURFACE_CREATED)
                    time.Sleep(time.Duration(180) * time.Millisecond)
                }
                return
            }(); tryControl1 == 1 {
                return
            }
        }
}

