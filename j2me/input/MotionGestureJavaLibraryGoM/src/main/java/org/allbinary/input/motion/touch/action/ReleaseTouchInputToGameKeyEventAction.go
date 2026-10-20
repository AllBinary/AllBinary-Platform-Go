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

type ReleaseTouchInputToGameKeyEventAction struct {
    GameKeyCompleteMotionGestureInputEvent
    NONE *GameKey
}

var ReleaseTouchInputToGameKeyEventAction_Instance any = NullUtil_GetInstance().NULL_OBJECT

func (this *ReleaseTouchInputToGameKeyEventAction) initFields() {
    this.NONE = GameKeyFactory_GetInstance().NONE
}


func ReleaseTouchInputToGameKeyEventAction_GetInstance() IGameKeyCompleteMotionGestureInputEvent {
        if ReleaseTouchInputToGameKeyEventAction_Instance == NullUtil_GetInstance().NULL_OBJECT {
            ReleaseTouchInputToGameKeyEventAction_Instance = NewReleaseTouchInputToGameKeyEventAction()
        }
        return any(ReleaseTouchInputToGameKeyEventAction_Instance).(IGameKeyCompleteMotionGestureInputEvent)
}


func NewReleaseTouchInputToGameKeyEventAction() *ReleaseTouchInputToGameKeyEventAction {
    this := &ReleaseTouchInputToGameKeyEventAction{}
        this.GameKeyCompleteMotionGestureInputEvent = *NewGameKeyCompleteMotionGestureInputEvent("Release Action", TouchMotionGestureFactory_GetInstance().RELEASED, PlatformInputMappingFactory_GetInstance().GetPersistentInputMappingInstance().GetInputMapping())
        this.initFields()
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.Put(commonStrings.EXCEPTION, this, commonStrings.CONSTRUCTOR, e)
                }
            }()
            this.SetGameKey(this.NONE)
            this.SetGameKeyEvent(GameKeyEventFactory_GetInstance().GetInstanceForInput(this, this.NONE))
            return
        }(); tryControl1 == 1 {
            return
        }
    return this
}


func (this *ReleaseTouchInputToGameKeyEventAction) Update() {
}

