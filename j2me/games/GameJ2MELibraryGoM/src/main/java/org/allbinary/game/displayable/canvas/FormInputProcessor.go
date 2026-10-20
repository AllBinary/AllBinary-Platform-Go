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

type FormInputProcessor struct {
    InputProcessor
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    GameKeyEventFactory *GameKeyEventFactory
    DownKeyEventHandler *DownKeyEventHandler
    UpGameKeyEventHandler *UpGameKeyEventHandler
    InputFactory *InputFactory
    PlatformKeyFactory *PlatformKeyFactory
    AllBinaryGameCanvas IAllBinaryGameCanvas
}


func (this *FormInputProcessor) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.GameKeyEventFactory = GameKeyEventFactory_GetInstance()
    this.DownKeyEventHandler = DownKeyEventHandler_GetInstance()
    this.UpGameKeyEventHandler = UpGameKeyEventHandler_GetInstance()
    this.InputFactory = InputFactory_GetInstance()
    this.PlatformKeyFactory = PlatformKeyFactory_GetInstance()
    
}


func NewFormInputProcessor(allBinaryGameCanvas IAllBinaryGameCanvas) *FormInputProcessor {
    this := &FormInputProcessor{}
    this.InputProcessor = *NewInputProcessor()
    this.initFields()
        this.AllBinaryGameCanvas = allBinaryGameCanvas
    return this
}


func (this *FormInputProcessor) KeyPressedByDevice(keyCode int, deviceId int) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.GameInputStrings.ADD_KEY_EVENT, e)
                }
            }()
            PreLogUtil_Put(fmt.Sprint(NewStringMaker().Append(this.InputFactory.KEY_CODE_LABEL).Appendint(keyCode).Append(CommonSeps_GetInstance().SPACE).Append(this.InputFactory.DEVICE_ID_LABEL).Appendint(deviceId)), this, this.GameInputStrings.KEY_PRESSED)
            var input IInput = this.InputFactory.GetInstanceById(keyCode); _ = input
            var gameKeyEvent *GameKeyEvent = this.GameKeyEventFactory.GetInstanceForInput(this.AllBinaryGameCanvas, input); _ = gameKeyEvent
            this.DownKeyEventHandler.FireEventForGameKeyEvent(gameKeyEvent)
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *FormInputProcessor) KeyReleasedByDevice(canvas ICanvas, keyCode int, deviceId int) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.GameInputStrings.ADD_KEY_EVENT, e)
                }
            }()
            var input IInput = this.InputFactory.GetInstanceById(keyCode); _ = input
            if this.PlatformKeyFactory.IsEnter(input) {
                PreLogUtil_Put(fmt.Sprint(NewStringMaker().Append(this.InputFactory.KEY_CODE_LABEL).Appendint(keyCode).Append(CommonSeps_GetInstance().SPACE).Append(this.InputFactory.DEVICE_ID_LABEL).Appendint(deviceId)), this, this.GameInputStrings.KEY_RELEASED)
                var gameKey *GameKey = GameKeyFactory_GetInstance().KEY_NUM0; _ = gameKey
                var gameKeyEventSourceInterface GameKeyEventSourceInterface = any(canvas).(GameKeyEventSourceInterface); _ = gameKeyEventSourceInterface
                var gameKeyEvent *GameKeyEvent = this.GameKeyEventFactory.GetInstanceForInput(gameKeyEventSourceInterface, gameKey); _ = gameKeyEvent
                this.UpGameKeyEventHandler.FireEvent(gameKeyEvent)
                this.UpGameKeyEventHandler.GetInstanceForDevice(deviceId).FireEvent(gameKeyEvent)
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}

