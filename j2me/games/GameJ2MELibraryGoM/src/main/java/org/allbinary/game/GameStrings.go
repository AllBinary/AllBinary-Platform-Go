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

type GameStrings struct {
    PROCESS_TICK string
    ON_GAME_INITIALIZED string
    UPDATESURFACE string
    ON_SET_DISPLAYABLE string
    UNPAUSE string
    TOGGLE_MENU string
    SET_GAME_OVER string
    GAME_STATE string
    SET_GAME_STATE string
    STOP_GAME_CANVAS_RUNNABLE_INTERFACE string
}

var GameStrings_Instance *GameStrings = NewGameStrings()

func (this *GameStrings) initFields() {
    this.PROCESS_TICK = "processTick"
    this.ON_GAME_INITIALIZED = "onGameInitialized"
    this.UPDATESURFACE = "updateSurface"
    this.ON_SET_DISPLAYABLE = "onSetDisplayable"
    this.UNPAUSE = "unPause"
    this.TOGGLE_MENU = "toggleMenu"
    this.SET_GAME_OVER = "setGameOver"
    this.GAME_STATE = "Game State: "
    this.SET_GAME_STATE = "setGameState"
    this.STOP_GAME_CANVAS_RUNNABLE_INTERFACE = "stopGameCanvasRunnableInterface"
}

func NewGameStrings() *GameStrings {
    this := &GameStrings{}
    this.initFields()
    return this
}


func GameStrings_GetInstance() *GameStrings {
        return GameStrings_Instance
}

