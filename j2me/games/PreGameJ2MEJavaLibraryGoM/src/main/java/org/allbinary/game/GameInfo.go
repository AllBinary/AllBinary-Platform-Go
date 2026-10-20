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

type GameInfo struct {
    LogUtil *LogUtil
    GameInfoData *GameInfoData
    GameType *GameType
    GameMode *GameMode
    PlayerType *PlayerType
    HighestLevel int
    CurrentLevel int
    NEW_LEVEL string
    SET_CURRENT_LEVEL string
    NEXT_GAME_LEVEL string
}

var GameInfo_NONE *GameInfo = NewGameInfo(GameTypeFactory_GetInstance().NONE, GameMode_NONE, PlayerTypesFactory_GetInstance().PLAYER_TYPE_ONE, -1, -1)
var GameInfo_LEVEL_NAME string = "LEVEL"

func (this *GameInfo) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.GameInfoData = GameInfoData_GetInstance()
    
    
    
    
    
    this.NEW_LEVEL = "New Level: "
    this.SET_CURRENT_LEVEL = "setCurrentLevel"
    this.NEXT_GAME_LEVEL = "nextGameLevel"
}


func NewGameInfo(gameType *GameType, gameMode *GameMode, playerType *PlayerType, highestLevel int, currentLevel int) *GameInfo {
    this := &GameInfo{}
    this.initFields()
        this.GameType = gameType
        this.GameMode = gameMode
        this.SetHighestLevel(highestLevel)
        this.CurrentLevel = currentLevel
        this.PlayerType = playerType
    return this
}


func (this *GameInfo) GetGameType() *GameType {
        return this.GameType
}


func (this *GameInfo) GetGameMode() *GameMode {
        return this.GameMode
}


func (this *GameInfo) GetPlayerType() *PlayerType {
        return this.PlayerType
}


func (this *GameInfo) GetHighestLevel() int {
        return this.HighestLevel
}


func (this *GameInfo) SetHighestLevel(highestLevel int) {
        this.HighestLevel = highestLevel
}


func (this *GameInfo) GetCurrentLevel() int {
        return this.CurrentLevel
}


func (this *GameInfo) SetCurrentLevel(currentLevel int) {
        this.CurrentLevel = currentLevel
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(this.NEW_LEVEL).Appendint(this.GetCurrentLevel())), this, this.SET_CURRENT_LEVEL)
}


func (this *GameInfo) NextGameLevel() {
        if this.GetCurrentLevel() < this.GetHighestLevel() {
            this.CurrentLevel++
        }
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(CommonLabels_GetInstance().NEXT).Append(CommonSeps_GetInstance().SPACE).Append(this.NEW_LEVEL).Appendint(this.GetCurrentLevel())), this, this.NEXT_GAME_LEVEL)
}


func (this *GameInfo) PreviousGameLevel() {
        if this.GetCurrentLevel() > 1 {
            this.CurrentLevel--
        }
}


func (this *GameInfo) IsLastLevel() bool {
        if this.GetCurrentLevel() >= this.GetHighestLevel() {
            return true
        } else {
            return false
        }
}


func (this *GameInfo) ToHashtable() *ABHashtable[any, any] {
        var hashtable *ABHashtable[any, any] = StdUtil_GetInstance().CreateHashtable(); _ = hashtable
        hashtable.Put(this.GameInfoData.GAME_TYPE, fmt.Sprint(this.GetGameType()))
        hashtable.Put(this.GameInfoData.GAME_MODE, fmt.Sprint(this.GameMode))
        hashtable.Put(this.GameInfoData.PLAYER_TYPE, fmt.Sprint(this.PlayerType))
        hashtable.Put(this.GameInfoData.HIGHEST_LEVEL, fmt.Sprint(this.GetHighestLevel()))
        hashtable.Put(this.GameInfoData.CURRENT_LEVEL, fmt.Sprint(this.CurrentLevel))
        return hashtable
}


func (this *GameInfo) ToString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        var commonSeps *CommonSeps = CommonSeps_GetInstance(); _ = commonSeps
        stringBuffer.Append(this.GameInfoData.GAME_TYPE)
        stringBuffer.Append(commonSeps.EQUALS)
        stringBuffer.Append(fmt.Sprint(this.GetGameType()))
        stringBuffer.Append(commonSeps.COMMA_SEP)
        stringBuffer.Append(this.GameInfoData.GAME_MODE)
        stringBuffer.Append(commonSeps.EQUALS)
        stringBuffer.Append(fmt.Sprint(this.GameMode))
        stringBuffer.Append(commonSeps.COMMA_SEP)
        stringBuffer.Append(this.GameInfoData.PLAYER_TYPE)
        stringBuffer.Append(commonSeps.EQUALS)
        stringBuffer.Append(StringUtil_GetInstance().ToString(this.PlayerType))
        stringBuffer.Append(commonSeps.COMMA_SEP)
        stringBuffer.Append(this.GameInfoData.HIGHEST_LEVEL)
        stringBuffer.Append(commonSeps.EQUALS)
        stringBuffer.Appendint(this.GetHighestLevel())
        stringBuffer.Append(commonSeps.COMMA_SEP)
        stringBuffer.Append(this.GameInfoData.CURRENT_LEVEL)
        stringBuffer.Append(commonSeps.EQUALS)
        stringBuffer.Appendint(this.CurrentLevel)
        return fmt.Sprint(stringBuffer)
}

func (this *GameInfo) String() string {
    return this.ToString()
}

