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

type BasicHighScoresGameInfoFactory struct {
    MULTI_PLAYER_SERVER *GameInfo
    SINGLE_PLAYER_SERVER *GameInfo
}

var BasicHighScoresGameInfoFactory_SINGLETON *BasicHighScoresGameInfoFactory = NewBasicHighScoresGameInfoFactory()

func (this *BasicHighScoresGameInfoFactory) initFields() {
    this.MULTI_PLAYER_SERVER = NewGameInfo(GameTypeFactory_GetInstance().MULTI_PLAYER, GameMode_SERVER, PlayerTypesFactory_GetInstance().PLAYER_TYPE_ONE, 0, 0)
    this.SINGLE_PLAYER_SERVER = NewGameInfo(GameTypeFactory_GetInstance().SINGLE_PLAYER, GameMode_SERVER, PlayerTypesFactory_GetInstance().PLAYER_TYPE_ONE, 0, 0)
}

func NewBasicHighScoresGameInfoFactory() *BasicHighScoresGameInfoFactory {
    this := &BasicHighScoresGameInfoFactory{}
    this.initFields()
    return this
}


func BasicHighScoresGameInfoFactory_GetInstance() *BasicHighScoresGameInfoFactory {
        return BasicHighScoresGameInfoFactory_SINGLETON
}

