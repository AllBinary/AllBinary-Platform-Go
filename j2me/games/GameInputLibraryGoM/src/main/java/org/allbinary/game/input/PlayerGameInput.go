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

type PlayerGameInput struct {
    GameInput
    GameInputStrings *GameInputStrings
    PlayerInputId int
}

// Implements GameKeyEventListenerInterface

// The Java class PlayerGameInput has subclasses so its variables are this interface
type IPlayerGameInput interface {
    IGameInput
    PlayerGameInput_() *PlayerGameInput
    OnPressGameKeyEvent(gameKeyEvent *GameKeyEvent)
    OnDownGameKeyEvent(gameKeyEvent *GameKeyEvent)
    OnDownKeyEvent(gameKeyEvent *GameKeyEvent)
    OnDownKey(gameKeyEvent int)
    OnUpGameKeyEvent(gameKeyEvent *GameKeyEvent)
    OnUpKeyEvent(gameKeyEvent int)
    OnEvent(eventObject IAllBinaryEventObject)
    GetPlayerInputId() int
}

func (this *PlayerGameInput) PlayerGameInput_() *PlayerGameInput {
    return this
}


func (this *PlayerGameInput) initFields() {
    this.GameInputStrings = GameInputStrings_GetInstance()
    
}


func NewPlayerGameInput(gameKeyEventList IBasicArrayList, removalGameKeyEventList IBasicArrayList, playerInputId int) *PlayerGameInput {
    this := &PlayerGameInput{}
        this.GameInput = *NewGameInput(gameKeyEventList, removalGameKeyEventList)
        this.initFields()
        this.PlayerInputId = playerInputId
    return this
}


// synchronized
func (this *PlayerGameInput) OnPressGameKeyEvent(gameKeyEvent *GameKeyEvent) {
        this.Add(gameKeyEvent)
        this.AddForRemoval(gameKeyEvent)
}


// synchronized
func (this *PlayerGameInput) OnDownGameKeyEvent(gameKeyEvent *GameKeyEvent) {
        this.Add(gameKeyEvent)
}


// synchronized
func (this *PlayerGameInput) OnDownKeyEvent(gameKeyEvent *GameKeyEvent) {
        this.Add(gameKeyEvent)
        this.AddForRemoval(gameKeyEvent)
}


// synchronized
func (this *PlayerGameInput) OnDownKey(gameKeyEvent int) {
        PreLogUtil_Put(fmt.Sprint(NewStringMaker().Append(CommonLabels_GetInstance().START_LABEL).Append(fmt.Sprint(gameKeyEvent))), this, this.GameInputStrings.ON_DOWN_GAME_KEY)
}


// synchronized
func (this *PlayerGameInput) OnUpGameKeyEvent(gameKeyEvent *GameKeyEvent) {
        this.AddForRemoval(gameKeyEvent)
}


// synchronized
func (this *PlayerGameInput) OnUpKeyEvent(gameKeyEvent int) {
}


func (this *PlayerGameInput) OnEvent(eventObject IAllBinaryEventObject) {
        ForcedLogUtil_Log(EventStrings_GetInstance().PERFORMANCE_MESSAGE, this)
}


func (this *PlayerGameInput) GetPlayerInputId() int {
        return this.PlayerInputId
}

