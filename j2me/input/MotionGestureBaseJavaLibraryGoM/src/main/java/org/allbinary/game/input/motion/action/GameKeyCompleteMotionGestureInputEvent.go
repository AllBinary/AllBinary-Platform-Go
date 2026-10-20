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

type GameKeyCompleteMotionGestureInputEvent struct {
    CompleteMotionGestureInputEvent
    GameKeyEventFactory *GameKeyEventFactory
    SOURCE_ID int
    GameKey *GameKey
    GameKeyEvent *GameKeyEvent
    InputToGameKeyMapping *InputToGameKeyMapping
}

// Implements GameKeyEventSourceInterface

// The Java class GameKeyCompleteMotionGestureInputEvent has subclasses so its variables are this interface
type IGameKeyCompleteMotionGestureInputEvent interface {
    ICompleteMotionGestureInputEvent
    GameKeyCompleteMotionGestureInputEvent_() *GameKeyCompleteMotionGestureInputEvent
    GetSourceId() int
    Update()
    SetGameKeyEvent(gameKeyEvent *GameKeyEvent)
    GetGameKeyEvent() *GameKeyEvent
    SetGameKey(gameKey *GameKey)
    GetGameKey() *GameKey
}

func (this *GameKeyCompleteMotionGestureInputEvent) GameKeyCompleteMotionGestureInputEvent_() *GameKeyCompleteMotionGestureInputEvent {
    return this
}


func (this *GameKeyCompleteMotionGestureInputEvent) initFields() {
    this.GameKeyEventFactory = GameKeyEventFactory_GetInstance()
    this.SOURCE_ID = this.GameKeyEventFactory.MOTION_GESTURE_SOURCE_ID
    this.GameKey = GameKey_NULL_GAME_KEY
    this.GameKeyEvent = GameKeyEvent_NONE
    
}


func NewGameKeyCompleteMotionGestureInputEvent(name string, motionGestureInput *MotionGestureInput, inputToGameKeyMapping *InputToGameKeyMapping) *GameKeyCompleteMotionGestureInputEvent {
    this := &GameKeyCompleteMotionGestureInputEvent{}
        this.CompleteMotionGestureInputEvent = *NewCompleteMotionGestureInputEvent(name, motionGestureInput)
        this.initFields()
        this.InputToGameKeyMapping = inputToGameKeyMapping
        GameKeyCompleteMotionGestureInputEventFactory_GetInstance().Add(this)
        this.Update()
    return this
}


func (this *GameKeyCompleteMotionGestureInputEvent) GetSourceId() int {
        return this.SOURCE_ID
}


func (this *GameKeyCompleteMotionGestureInputEvent) Update() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.Put(commonStrings.EXCEPTION, this, commonStrings.UPDATE, e)
                }
            }()
            this.SetGameKey(this.InputToGameKeyMapping.GetInstance(this.GetMotionGestureInput().GetId()))
            this.SetGameKeyEvent(this.GameKeyEventFactory.GetInstanceForInput(this, this.GetGameKey()))
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *GameKeyCompleteMotionGestureInputEvent) SetGameKeyEvent(gameKeyEvent *GameKeyEvent) {
        this.GameKeyEvent = gameKeyEvent
}


func (this *GameKeyCompleteMotionGestureInputEvent) GetGameKeyEvent() *GameKeyEvent {
        return this.GameKeyEvent
}


func (this *GameKeyCompleteMotionGestureInputEvent) SetGameKey(gameKey *GameKey) {
        this.GameKey = gameKey
}


func (this *GameKeyCompleteMotionGestureInputEvent) GetGameKey() *GameKey {
        return this.GameKey
}

