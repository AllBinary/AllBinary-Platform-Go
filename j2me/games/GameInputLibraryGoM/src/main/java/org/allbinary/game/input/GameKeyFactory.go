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

type GameKeyFactory struct {
    NONE *GameKey
    UP *GameKey
    DOWN *GameKey
    LEFT *GameKey
    RIGHT *GameKey
    KEY_POUND *GameKey
    KEY_STAR *GameKey
    KEY_NUM0 *GameKey
    KEY_NUM1 *GameKey
    KEY_NUM2 *GameKey
    KEY_NUM3 *GameKey
    KEY_NUM4 *GameKey
    KEY_NUM5 *GameKey
    KEY_NUM6 *GameKey
    KEY_NUM7 *GameKey
    KEY_NUM8 *GameKey
    KEY_NUM9 *GameKey
    GAME_A *GameKey
    GAME_B *GameKey
    GAME_C *GameKey
    GAME_D *GameKey
    FIRE *GameKey
    LEVEL_DOWN *GameKey
    LEVEL_UP *GameKey
}

var GameKeyFactory_Instance any = NullUtil_GetInstance().NULL_OBJECT

func (this *GameKeyFactory) initFields() {
    this.NONE = NewGameKey(InputFactory_GetInstance().MAX - 1, "None")
    this.UP = NewGameKey(Canvas_UP, CommonPhoneStrings_GetInstance().UP)
    this.DOWN = NewGameKey(Canvas_DOWN, CommonPhoneStrings_GetInstance().DOWN)
    this.LEFT = NewGameKey(Canvas_LEFT, CommonPhoneStrings_GetInstance().LEFT)
    this.RIGHT = NewGameKey(Canvas_RIGHT, CommonPhoneStrings_GetInstance().RIGHT)
    this.KEY_POUND = NewGameKey(Canvas_KEY_POUND, "POUND")
    this.KEY_STAR = NewGameKey(Canvas_KEY_STAR, "STAR")
    this.KEY_NUM0 = NewGameKey(Canvas_KEY_NUM0, CommonPhoneStrings_GetInstance().ZERO)
    this.KEY_NUM1 = NewGameKey(Canvas_KEY_NUM1, CommonPhoneStrings_GetInstance().ONE)
    this.KEY_NUM2 = NewGameKey(Canvas_KEY_NUM2, CommonPhoneStrings_GetInstance().TWO)
    this.KEY_NUM3 = NewGameKey(Canvas_KEY_NUM3, CommonPhoneStrings_GetInstance().THREE)
    this.KEY_NUM4 = NewGameKey(Canvas_KEY_NUM4, CommonPhoneStrings_GetInstance().FOUR)
    this.KEY_NUM5 = NewGameKey(Canvas_KEY_NUM5, CommonPhoneStrings_GetInstance().FIVE)
    this.KEY_NUM6 = NewGameKey(Canvas_KEY_NUM6, CommonPhoneStrings_GetInstance().SIX)
    this.KEY_NUM7 = NewGameKey(Canvas_KEY_NUM7, CommonPhoneStrings_GetInstance().SEVEN)
    this.KEY_NUM8 = NewGameKey(Canvas_KEY_NUM8, CommonPhoneStrings_GetInstance().EIGHT)
    this.KEY_NUM9 = NewGameKey(Canvas_KEY_NUM9, CommonPhoneStrings_GetInstance().NINE)
    this.GAME_A = NewGameKey(Canvas_GAME_A, "Game A")
    this.GAME_B = NewGameKey(Canvas_GAME_B, "Game B")
    this.GAME_C = NewGameKey(Canvas_GAME_C, "Game C")
    this.GAME_D = NewGameKey(Canvas_GAME_D, "Game D")
    this.FIRE = NewGameKey(Canvas_FIRE, CommonPhoneStrings_GetInstance().FIRE)
    this.LEVEL_DOWN = NewGameKey(Canvas_KEY_NUM9 + 1, "Cheat Level Up")
    this.LEVEL_UP = NewGameKey(Canvas_KEY_NUM9 + 2, "Cheat Level Down")
}

func NewGameKeyFactory() *GameKeyFactory {
    this := &GameKeyFactory{}
    this.initFields()
    return this
}


func GameKeyFactory_GetInstance() *GameKeyFactory {
        if GameKeyFactory_Instance == NullUtil_GetInstance().NULL_OBJECT {
            GameKeyFactory_Instance = NewGameKeyFactory()
        }
        return any(GameKeyFactory_Instance).(*GameKeyFactory)
}

