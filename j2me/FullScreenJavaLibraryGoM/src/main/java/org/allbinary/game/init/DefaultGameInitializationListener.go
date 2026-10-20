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

type DefaultGameInitializationListener struct {
    LogUtil *LogUtil
    SwtJOGLProcessor *SWTJOGLProcessor
    FirstTime bool
}

// Implements GameInitializedListenerInterface


func (this *DefaultGameInitializationListener) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.SwtJOGLProcessor = SWTJOGLProcessor_GetInstance()
    this.FirstTime = true
}


func NewDefaultGameInitializationListener() *DefaultGameInitializationListener {
    this := &DefaultGameInitializationListener{}
    this.initFields()
        var gameInitializedEventHandler *GameInitializedEventHandler = GameInitializedEventHandler_GetInstance(); _ = gameInitializedEventHandler
        gameInitializedEventHandler.RemoveAllListeners()
        gameInitializedEventHandler.AddListenerInterface(any(this).(GameInitializedListenerInterface))
    return this
}


func (this *DefaultGameInitializationListener) OnEvent(eventObject IAllBinaryEventObject) {
        ForcedLogUtil_Log(EventStrings_GetInstance().PERFORMANCE_MESSAGE, this)
}


func (this *DefaultGameInitializationListener) OnGameInitialized(gameInitializedEvent *GameInitializedEvent) {
        var ON_GAME_INITIALIZED string = "onGameInitialized"; _ = ON_GAME_INITIALIZED
        var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(commonStrings.EXCEPTION, this, ON_GAME_INITIALIZED, e)
                }
            }()
            this.LogUtil.PutF(commonStrings.START, this, ON_GAME_INITIALIZED)
            for !this.SwtJOGLProcessor.IsHolderCreated() {
                this.LogUtil.PutF(commonStrings.UPDATE, this, ON_GAME_INITIALIZED)
                time.Sleep(time.Duration(20) * time.Millisecond)
            }
            FeatureResourceInitializationUtil_GetInstance().Init(gameInitializedEvent.GetLevel())
            if this.FirstTime {
                this.FirstTime = false
            } else {
                this.SwtJOGLProcessor.OnSurfaceChanged()
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}

