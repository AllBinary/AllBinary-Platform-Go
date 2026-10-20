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

type IntermissionFactory struct {
    START_LEVEL_INTERMISSION_GAME_STATE *GameState
    WAIT_LEVEL_INTERMISSION_GAME_STATE *GameState
    SHOW_RESULTS_LEVEL_INTERMISSION_GAME_STATE *GameState
    SHOW_HIGH_SCORE_LEVEL_INTERMISSION_GAME_STATE *GameState
}

var IntermissionFactory_Instance *IntermissionFactory = NewIntermissionFactory()

func (this *IntermissionFactory) initFields() {
    this.START_LEVEL_INTERMISSION_GAME_STATE = GameStateFactory_GetInstance().CreateGameState("START_LEVEL_INTERMISSION_GAME_STATE")
    this.WAIT_LEVEL_INTERMISSION_GAME_STATE = GameStateFactory_GetInstance().CreateGameState("WAIT_LEVEL_INTERMISSION_GAME_STATE")
    this.SHOW_RESULTS_LEVEL_INTERMISSION_GAME_STATE = GameStateFactory_GetInstance().CreateGameState("SHOW_RESULTS_LEVEL_INTERMISSION_GAME_STATE")
    this.SHOW_HIGH_SCORE_LEVEL_INTERMISSION_GAME_STATE = GameStateFactory_GetInstance().CreateGameState("SHOW_HIGH_SCORE_LEVEL_INTERMISSION_GAME_STATE")
}

func NewIntermissionFactory() *IntermissionFactory {
    this := &IntermissionFactory{}
    this.initFields()
    return this
}


func IntermissionFactory_GetInstance() *IntermissionFactory {
        return IntermissionFactory_Instance
}

