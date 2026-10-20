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

type GameInputStrings struct {
    PROCESS_INPUT string
    KEY_PRESSED string
    KEY_RELEASED string
    KEY_REPEATED string
    ON_PRESS_GAME_KEY string
    ON_DOWN_GAME_KEY string
    ON_UP_GAME_KEY string
    NO_KEY string
    ADD_KEY_EVENT string
    REMOVE_KEY_EVENT string
    STRAFE_LEFT string
    STRAFE_RIGHT string
    PROCESS_MOTION_INPUT string
    UDPATE_CURRENT_TOUCH_INPUT_FACTORY string
    ENABLE_PLAYER_GAME_INPUTS string
    ENABLE_PLAYER_GAME_INPUT string
}

var GameInputStrings_Instance *GameInputStrings = NewGameInputStrings()

func (this *GameInputStrings) initFields() {
    this.PROCESS_INPUT = "processInput"
    this.KEY_PRESSED = "keyPressed"
    this.KEY_RELEASED = "keyReleased"
    this.KEY_REPEATED = "keyRepeated"
    this.ON_PRESS_GAME_KEY = "onPressGameKeyEvent"
    this.ON_DOWN_GAME_KEY = "onDownGameKeyEvent"
    this.ON_UP_GAME_KEY = "onUpGameKeyEvent"
    this.NO_KEY = "Key Code Not Mapped For Game: "
    this.ADD_KEY_EVENT = "addGameKeyEvent"
    this.REMOVE_KEY_EVENT = "removeGameKeyEvent"
    this.STRAFE_LEFT = "strafeLeft"
    this.STRAFE_RIGHT = "strafeRight"
    this.PROCESS_MOTION_INPUT = "processMotionInput"
    this.UDPATE_CURRENT_TOUCH_INPUT_FACTORY = "updateCurrentTouchInputFactory"
    this.ENABLE_PLAYER_GAME_INPUTS = "Enabling PlayerGameInputs: "
    this.ENABLE_PLAYER_GAME_INPUT = "Enabling PlayerGameInput: "
}

func NewGameInputStrings() *GameInputStrings {
    this := &GameInputStrings{}
    this.initFields()
    return this
}


func GameInputStrings_GetInstance() *GameInputStrings {
        return GameInputStrings_Instance
}

