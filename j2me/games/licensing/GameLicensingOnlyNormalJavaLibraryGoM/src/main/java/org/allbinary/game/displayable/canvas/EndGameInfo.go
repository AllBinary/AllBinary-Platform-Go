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

type EndGameInfo struct {
    NORMAL_GAME_OVER string
    NORMAL_GAME_OVER_WIN string
    GameOverStr string
    WinStr string
}

// The Java class EndGameInfo has subclasses so its variables are this interface
type IEndGameInfo interface {
    EndGameInfo_() *EndGameInfo
    Update(gameInfo *GameInfo, myCanvas IMyCanvas)
    GetGameOverStr() string
    SetGameOverStr(gameOverStr string)
    GetWinStr() string
    SetWinStr(winStr string)
}

func (this *EndGameInfo) EndGameInfo_() *EndGameInfo {
    return this
}


func (this *EndGameInfo) initFields() {
    this.NORMAL_GAME_OVER = "Game Over"
    this.NORMAL_GAME_OVER_WIN = "You Won"
    this.GameOverStr = StringUtil_GetInstance().EMPTY_STRING
    this.WinStr = this.GameOverStr
}


func NewEndGameInfo() *EndGameInfo {
    this := &EndGameInfo{}
    this.initFields()
    return this
}


func (this *EndGameInfo) Update(gameInfo *GameInfo, myCanvas IMyCanvas) {
        this.SetGameOverStr(this.NORMAL_GAME_OVER)
        this.SetWinStr(this.NORMAL_GAME_OVER_WIN)
}


func (this *EndGameInfo) GetGameOverStr() string {
        return this.GameOverStr
}


func (this *EndGameInfo) SetGameOverStr(gameOverStr string) {
        this.GameOverStr = gameOverStr
}


func (this *EndGameInfo) GetWinStr() string {
        return this.WinStr
}


func (this *EndGameInfo) SetWinStr(winStr string) {
        this.WinStr = winStr
}

