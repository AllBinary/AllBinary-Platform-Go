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

type GameCanvasInputProcessor struct {
    InputProcessor
    LogUtil *LogUtil
    AllBinaryGameCanvas IAllBinaryGameCanvas
    DownGameKeyEventHandler *DownGameKeyEventHandler
    UpGameKeyEventHandler *UpGameKeyEventHandler
    DownKeyEventHandler *DownKeyEventHandler
    SmallIntegerSingletonFactory *SmallIntegerSingletonFactory
    InputToGameKeyMapping *InputToGameKeyMapping
    NONE *GameKey
    GameKeyEventFactory *GameKeyEventFactory
}


func (this *GameCanvasInputProcessor) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    this.DownGameKeyEventHandler = DownGameKeyEventHandler_GetInstance()
    this.UpGameKeyEventHandler = UpGameKeyEventHandler_GetInstance()
    this.DownKeyEventHandler = DownKeyEventHandler_GetInstance()
    this.SmallIntegerSingletonFactory = SmallIntegerSingletonFactory_GetInstance()
    this.InputToGameKeyMapping = PlatformInputMappingFactory_GetInstance().GetPersistentInputMappingInstance().GetInputMapping()
    this.NONE = GameKeyFactory_GetInstance().NONE
    this.GameKeyEventFactory = GameKeyEventFactory_GetInstance()
}


func NewGameCanvasInputProcessor(allBinaryGameCanvas IAllBinaryGameCanvas) *GameCanvasInputProcessor {
    this := &GameCanvasInputProcessor{}
    this.InputProcessor = *NewInputProcessor()
    this.initFields()
        this.AllBinaryGameCanvas = allBinaryGameCanvas
        this.InputToGameKeyMapping.Init(this.AllBinaryGameCanvas)
    return this
}


func (this *GameCanvasInputProcessor) KeyPressedByDevice(keyCode int, deviceId int) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put("Key Event Error", this, this.GameInputStrings.ADD_KEY_EVENT, e)
                }
            }()
            var gameKey *GameKey = this.InputToGameKeyMapping.GetInstanceForCanvas(this.AllBinaryGameCanvas, keyCode); _ = gameKey
            if gameKey != this.NONE {
                var gameKeyEvent *GameKeyEvent = this.GameKeyEventFactory.GetInstanceForInput(this.AllBinaryGameCanvas, gameKey); _ = gameKeyEvent
                this.DownGameKeyEventHandler.FireEvent(gameKeyEvent)
                this.DownGameKeyEventHandler.GetInstanceForDevice(deviceId).FireEvent(gameKeyEvent)
            } else {
                this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(this.GameInputStrings.NO_KEY).Appendint(keyCode)), this, this.GameInputStrings.ADD_KEY_EVENT)
            }
            var keyCodeAsInteger int = this.SmallIntegerSingletonFactory.GetAtNoThrow(keyCode); _ = keyCodeAsInteger
            this.DownKeyEventHandler.FireEventI(keyCodeAsInteger)
            this.DownKeyEventHandler.GetInstanceForDevice(deviceId).FireEventI(keyCodeAsInteger)
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *GameCanvasInputProcessor) KeyReleasedByDevice(canvas ICanvas, keyCode int, deviceId int) {
        this.RemoveGameKeyEvent(canvas, keyCode, deviceId, false)
}


func (this *GameCanvasInputProcessor) RemoveGameKeyEvent(canvas ICanvas, keyCode int, deviceId int, repeated bool) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put("Key Event Error", this, this.GameInputStrings.REMOVE_KEY_EVENT, e)
                }
            }()
            var gameKey *GameKey = this.InputToGameKeyMapping.GetInstanceForCanvas(canvas, keyCode); _ = gameKey
            if gameKey != this.NONE {
                var gameKeyEventSourceInterface GameKeyEventSourceInterface = any(canvas).(GameKeyEventSourceInterface); _ = gameKeyEventSourceInterface
                var gameKeyEvent *GameKeyEvent = this.GameKeyEventFactory.GetInstanceForInput(gameKeyEventSourceInterface, gameKey); _ = gameKeyEvent
                this.UpGameKeyEventHandler.FireEvent(gameKeyEvent)
                this.UpGameKeyEventHandler.GetInstanceForDevice(deviceId).FireEvent(gameKeyEvent)
            } else {
                this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(this.GameInputStrings.NO_KEY).Appendint(keyCode)), this, this.GameInputStrings.REMOVE_KEY_EVENT)
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}

