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

type GameInfoData struct {
    GAME_TYPE string
    GAME_MODE string
    PLAYER_TYPE string
    HIGHEST_LEVEL string
    CURRENT_LEVEL string
    SOFTWARE_INFORMATION string
}

var GameInfoData_Instance *GameInfoData = NewGameInfoData()

func (this *GameInfoData) initFields() {
    this.GAME_TYPE = "GAME_TYPE"
    this.GAME_MODE = "GAME_MODE"
    this.PLAYER_TYPE = "PLAYER_TYPE"
    this.HIGHEST_LEVEL = "HIGHEST_LEVEL"
    this.CURRENT_LEVEL = "CURRENT_LEVEL"
    this.SOFTWARE_INFORMATION = "SOFTWARE_INFORMATION"
}


func GameInfoData_GetInstance() *GameInfoData {
        return GameInfoData_Instance
}


func NewGameInfoData() *GameInfoData {
    this := &GameInfoData{}
    this.initFields()
    return this
}

