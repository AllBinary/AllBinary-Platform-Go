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

type GameStateFactory struct {
    Index int
    NO_GAME_STATE *GameState
    PLAYING_GAME_STATE *GameState
    SHOW_END_RESULT_GAME_STATE *GameState
    SHOW_HIGH_SCORE_GAME_STATE *GameState
}

var GameStateFactory_Instance any = NullUtil_GetInstance().NULL_OBJECT

func (this *GameStateFactory) initFields() {
    this.Index = 0
    this.NO_GAME_STATE = this.CreateGameState("NO_GAME_STATE")
    this.PLAYING_GAME_STATE = this.CreateGameState("PLAYING_GAME_STATE")
    this.SHOW_END_RESULT_GAME_STATE = this.CreateGameState("SHOW_END_RESULT_GAME_STATE")
    this.SHOW_HIGH_SCORE_GAME_STATE = this.CreateGameState("SHOW_HIGH_SCORE_GAME_STATE")
}


func GameStateFactory_GetInstance() *GameStateFactory {
        if GameStateFactory_Instance == NullUtil_GetInstance().NULL_OBJECT {
            GameStateFactory_Instance = NewGameStateFactory()
        }
        return any(GameStateFactory_Instance).(*GameStateFactory)
}


func NewGameStateFactory() *GameStateFactory {
    this := &GameStateFactory{}
    this.initFields()
    return this
}


func (this *GameStateFactory) CreateGameState(name string) *GameState {
        return NewGameState(name, Java_PostIncrement(&this.Index))
}

