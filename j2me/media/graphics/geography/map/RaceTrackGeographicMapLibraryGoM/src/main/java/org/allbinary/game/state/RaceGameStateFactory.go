/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2025 AllBinary 
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

type RaceGameStateFactory struct {
    START_GAME_STATE *GameState
    YELLOW_LIGHT_GAME_STATE *GameState
    RED_LIGHT_GAME_STATE *GameState
}

var RaceGameStateFactory_Instance any = NullUtil_GetInstance().NULL_OBJECT

func (this *RaceGameStateFactory) initFields() {
    this.START_GAME_STATE = GameStateFactory_GetInstance().CreateGameState("START_GAME_STATE")
    this.YELLOW_LIGHT_GAME_STATE = GameStateFactory_GetInstance().CreateGameState("YELLOW_LIGHT_GAME_STATE")
    this.RED_LIGHT_GAME_STATE = GameStateFactory_GetInstance().CreateGameState("RED_LIGHT_GAME_STATE")
}

func NewRaceGameStateFactory() *RaceGameStateFactory {
    this := &RaceGameStateFactory{}
    this.initFields()
    return this
}


func RaceGameStateFactory_GetInstance() *RaceGameStateFactory {
        if RaceGameStateFactory_Instance == NullUtil_GetInstance().NULL_OBJECT {
            RaceGameStateFactory_Instance = NewRaceGameStateFactory()
        }
        return any(RaceGameStateFactory_Instance).(*RaceGameStateFactory)
}

